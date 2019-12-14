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
	http.HandleFunc("/", handler)
	go util.GitPuller(repoURL, repoPath)
	log.Fatal(http.ListenAndServe(":5001", nil))
}

// handle all requests
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
