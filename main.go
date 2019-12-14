package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jreisinger/homepage/util"
)

const repoURL = "https://github.com/jreisinger/homepage"
const repoPath = "/tmp/homepage"

func main() {
	// serve static files
	// (https://www.alexedwards.net/blog/serving-static-sites-with-go)
	fileServer := http.FileServer(http.Dir(repoPath + "/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// serve the rest
	http.HandleFunc("/", handler)

	// regularly update the local repo from the upstream repo
	go util.GitPuller(repoURL, repoPath)

	// start a webserver
	log.Fatal(http.ListenAndServe(":5001", nil))
}

// handle requests
func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if urlPath == "/" {
		http.Redirect(w, r, "/home", http.StatusFound)
	}

	p, err := util.RenderPage(repoURL, repoPath, urlPath)
	util.Check(err)

	t, err := template.ParseFiles("template/page.html")
	util.Check(err)

	t.Execute(w, p)
}
