package main

import (
	"log"
	"net/http"
	"os"
)

var (
	repoURL  = "https://github.com/jreisinger/homepage"
	repoPath = "/tmp/homepage"
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

	go GitPuller(repoURL, repoPath)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
