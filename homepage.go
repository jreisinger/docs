package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	repoRemoteUrl = "https://github.com/jreisinger/homepage"
	repoLocalPath = "/tmp/homepage"
)

func main() {
	http.Handle("/static/", staticHandler())
	http.HandleFunc("/favicon.ico", faviconHandler)

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/notes", notesHandler)
	http.HandleFunc("/blog", blogHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/", pageHandler)

	if err := os.RemoveAll(repoLocalPath); err != nil {
		log.Fatal(err)
	}
	log.Printf("cloning %s to %s", repoRemoteUrl, repoLocalPath)
	if err := gitClone(repoLocalPath, repoRemoteUrl); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := gitPull(repoLocalPath); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	addr := "localhost:5001"
	log.Printf("started HTTP server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
