package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jreisinger/homepage/util"
)

const repoURL = "https://github.com/jreisinger/homepage"
const repoPath = "/tmp/homepage"

func main() {
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

	// serve the rest
	http.HandleFunc("/", handler)

	// regularly update the local repo from the upstream repo
	go util.GitPuller(repoURL, repoPath)

	// start a webserver
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func removeTralingSlash(s string) string {
	suffix := "/"
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// handle requests
func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:] // remove leading /

	if urlPath == "" {
		http.Redirect(w, r, "/home", http.StatusFound)
	}

	p, err := util.RenderPage(repoURL, repoPath, urlPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	t, err := template.New("page.html").Funcs(template.FuncMap{"removeTrailingSlash": removeTralingSlash}).ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
