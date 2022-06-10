package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

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
		matchedFilePath := grepFilePath(path, rx)
		matchedFileContent := false

		if !searchOnlyPath && !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			html := mdToHtml(data)
			matchedFileContent = grepFileContent(string(html), rx)
		}

		if matchedFilePath || matchedFileContent {
			urlPath, err := filesystemToURL(path)
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

// getUrlPath validates the URL path of the request by matching it against
// validPath regex.
func getUrlPath(r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		return "", errors.New("invalid URL path")
	}
	return m[1], nil
}

// handler handles all requests not handled by other handlers.
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/about", http.StatusFound)
		return
	}

	urlPath, err := getUrlPath(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid URL path: %q", r.URL.Path)
		return
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
		Funcs(template.FuncMap{"removeTrailingSlash": removeTralingSlash}).
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
	icon := filepath.Join(repoPath, "static", "favicon.ico")
	http.ServeFile(w, r, icon)
}

// staticHandler serves files from "static" folder: CSS styles and pictures.
func staticHandler() http.Handler {
	fileServer := http.FileServer(http.Dir(repoPath + "/static/"))
	return http.StripPrefix("/static/", fileServer)
}

// RemoveTrailingSlash removes slash ("/") if it's the last character in a
// string.
func removeTralingSlash(s string) string {
	suffix := "/"
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
