package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

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

	// search paths and contents of files in data/
	http.HandleFunc("/search", handleSearch)

	// serve the rest
	http.HandleFunc("/", handleRest)

	// regularly update the local repo from the upstream repo
	go util.GitPuller(repoURL, repoPath)

	// start a webserver
	log.Fatal(http.ListenAndServe(":5001", nil))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("regexp")
	rx, err := regexp.Compile(pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundFiles []string

	err = filepath.Walk(repoPath+"/data", func(path string, info os.FileInfo, err error) error {
		matchFilePath := util.GrepFilePath(path, rx)
		matchFileContent := ""

		if !info.IsDir() { // i.e. it's a file
			matchFileContent, err = util.GrepFile(path, rx)
			if err != nil {
				return err
			}
		}

		if matchFilePath != "" || matchFileContent != "" {
			urlPath, err := util.FilesystemToURL(path)
			if err != nil {
				return err
			}
			foundFiles = append(foundFiles, urlPath)
		}

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &util.Page{
		Title: "search",
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
