package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

// searchHandler handles requests for /search. It searches paths and contents of
// files in "data" folder.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("regexp")

	var searchOnlyPath bool

	if strings.HasPrefix(pattern, "path:") {
		searchOnlyPath = true
		pattern = strings.TrimPrefix(pattern, "path:")
	}

	rx, err := regexp.Compile("(?i)" + pattern) // make pattern case insensitive
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundFiles []string

	err = filepath.Walk(repoPath+"/data", func(path string, info os.FileInfo, err error) error {
		matchedFilePath := GrepFilePath(path, rx)
		matchedFileContent := false

		if !searchOnlyPath && !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			html := MdToHtml(data)
			matchedFileContent = GrepFileContent(string(html), rx)
		}

		if matchedFilePath || matchedFileContent {
			urlPath, err := FilesystemToURL(path)
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

	p := &Page{
		Title: "search",
		IsDir: true,
		Files: foundFiles,
	}

	renderTemplate(w, "page", p)
}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)$")

func getUrlPath(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid URL path")
	}
	return m[1], nil
}

// handler handles all requests not handled by other handlers.
func handler(w http.ResponseWriter, r *http.Request) {
	// urlPath := r.URL.Path[1:] // remove leading "/"
	urlPath, err := getUrlPath(w, r)
	if err != nil {
		return
	}

	if urlPath == "" {
		http.Redirect(w, r, "/about", http.StatusFound)
	}

	p, err := loadPage(repoURL, repoPath, urlPath)
	if err != nil {
		if err == ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	renderTemplate(w, "page", p)
}

// templates caches available HTML templates.
var templates = template.Must(
	template.New("page").
		Funcs(template.FuncMap{"removeTrailingSlash": RemoveTralingSlash}).
		ParseFiles(filepath.Join("tmpl", "page.html")),
)

// renderTemplate fills in tmpl template with p data and writes it to w.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// faviconHandler serves favicon.ico from "static" folder.
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(repoPath+"static"+"favicon.ico"))
}

// staticHandler serves files from "static" folder: CSS styles and pictures.
func staticHandler() http.Handler {
	fileServer := http.FileServer(http.Dir(repoPath + "/static/"))
	return http.StripPrefix("/static/", fileServer)
}
