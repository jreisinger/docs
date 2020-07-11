package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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

	// grep i.e. search inside MD files in data/
	http.HandleFunc("/grep", handleGrep)

	// serve the rest
	http.HandleFunc("/", handleRest)

	// regularly update the local repo from the upstream repo
	go util.GitPuller(repoURL, repoPath)

	// start a webserver
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func handleGrep(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("regexp")
	rx, err := regexp.Compile(pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var foundFiles []string
	err = filepath.Walk(repoPath+"/data", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := util.GrepFile(path, rx)
			if err != nil {
				return err
			}
			if content != "" {
				// Trim .md suffix and .*data/ prefix from file path.
				path := strings.TrimSuffix(path, ".md")
				rx, err := regexp.Compile(`.*data/`)
				if err != nil {
					return err
				}
				path = rx.ReplaceAllString(path, "")

				foundFiles = append(foundFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &util.Page{
		Title: "grep",
		IsDir: true,
		Files: foundFiles,
	}

	t, err := template.New("page.html").
		Funcs(template.FuncMap{"removeTrailingSlash": util.RemoveTralingSlash}).
		ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handle requests
func handleRest(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:] // remove leading "/"

	if urlPath == "" {
		http.Redirect(w, r, "/about", http.StatusFound)
	}

	p, err := util.RenderPage(repoURL, repoPath, urlPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	t, err := template.New("page.html").
		Funcs(template.FuncMap{"removeTrailingSlash": util.RemoveTralingSlash}).
		ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
