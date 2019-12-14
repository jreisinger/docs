package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/jreisinger/homepage/util"
)

const repoURL = "https://github.com/jreisinger/homepage"
const repoPath = "/tmp/homepage"

type Page struct {
	Title   string
	Body    template.HTML
	RepoURL string
	IsDir   bool
	Files   []string
	UrlPath string
}

func renderPage(urlPath string) (*Page, error) {
	dataPath := repoPath + "/data"

	// map URL path to filesystem path (without .md)
	filePath := dataPath + "/" + urlPath[1:] // remove leading /

	title := path.Base(filePath)

	if util.IsDir(filePath) {
		// if file is a dir return list of files within dir
		files := util.ListFiles(filePath)
		return &Page{Title: title, Files: files, RepoURL: repoURL, UrlPath: urlPath, IsDir: true}, nil
	} else if util.IsFile(filePath + ".md") {
		// if file is a file return file contents
		body := util.MdToHtml(filePath + ".md")
		return &Page{Title: title, Body: body, RepoURL: repoURL, UrlPath: urlPath}, nil
	}

	// FIXME: if filePath does not map to en existing directory nor file return error
	return nil, nil
}

// handle all requests
func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if urlPath == "/" {
		http.Redirect(w, r, "/about", http.StatusFound)
	}

	p, err := renderPage(urlPath)
	util.Check(err)

	t, err := template.ParseFiles("template/page.html")
	util.Check(err)

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	go util.GitPuller(repoURL, repoPath)
	log.Fatal(http.ListenAndServe(":5001", nil))
}
