package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jreisinger/homepage/util"
)

var repoURL = "https://github.com/jreisinger/homepage"
var repoPath = "/tmp/homepage"

func main() {
	rp := os.Getenv("REPOPATH")
	if rp != "" {
		repoPath = rp
	}
	// serve static files
	// (https://www.alexedwards.net/blog/serving-static-sites-with-go)
	//
	// styles
	fileServer := http.FileServer(http.Dir(repoPath + "/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//
	// favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, repoPath+"/static/favicon.ico")
	})

	http.HandleFunc("/search", HandleSearch)

	http.HandleFunc("/", HandleRest)

	// regularly update the local repo from the upstream repo
	go util.GitPuller(repoURL, repoPath)

	// start a webserver
	log.Fatal(http.ListenAndServe(":5001", nil))
}
