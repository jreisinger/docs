package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	repoURL   = "https://github.com/jreisinger/homepage"
	repoPath  = "/tmp/homepage"
	validPath = regexp.MustCompile(`^/([a-zA-Z0-9/\-\.Φιλοσοφία]+)$`)
)

func main() {
	rp := os.Getenv("REPOPATH")
	if rp != "" {
		repoPath = rp
	}

	http.Handle("/static/", staticHandler())
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", handler)

	go gitPuller(repoURL, repoPath)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
