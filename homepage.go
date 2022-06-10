package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	repoURL   = "https://github.com/jreisinger/homepage"
	repoPath  = "/tmp/homepage"
	validPath = regexp.MustCompile(`^/([a-zA-Z0-9/\-\.Φιλοσοφία]+)$`)
)

func main() {
	http.Handle("/static/", staticHandler())
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", handler)

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		log.Printf("cloning %s to %s", repoURL, repoPath)
		gitClone(repoURL, repoPath)
	}
	go func() {
		for {
			gitPull(repoPath)
			time.Sleep(time.Second * 2)
		}
	}()

	addr := ":5001"
	log.Printf("started HTTP server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
