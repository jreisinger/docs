package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ikeikeikeike/go-sitemap-generator/stm"
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

	err = filepath.Walk(repoLocalPath+"/data", func(path string, info os.FileInfo, err error) error {
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

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path)
	if err != nil {
		handleError(w, r, err)
		return
	}
	renderTemplate(w, "about", p)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path)
	if err != nil {
		handleError(w, r, err)
		return
	}
	renderTemplate(w, "notes", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path)
	if err != nil {
		handleError(w, r, err)
		return
	}
	renderTemplate(w, "blog", p)
}

// pageHandler handles all requests not handled by other handlers.
func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/about", http.StatusFound)
		return
	}

	p, err := loadPage(r.URL.Path)
	if err != nil {
		handleError(w, r, err)
		return
	}
	renderTemplate(w, "page", p)
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		if err == ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, err)
	}
}

// templates caches available HTML templates.
var templates = template.Must(
	template.New("page").
		Funcs(template.FuncMap{"removeTrailingSlash": removeTralingSlash, "removeLeadingSlash": removeLeadingSlash}).
		ParseGlob(filepath.Join("tmpl", "*.html")),
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
	icon := filepath.Join(repoLocalPath, "static", "favicon.ico")
	http.ServeFile(w, r, icon)
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("http://reisinge.net")

	sm.Create()
	sm.Add(stm.URL{"loc": "/notes", "changefreq": "daily"})
	sm.Add(stm.URL{"loc": "/blog", "changefreq": "monthly"})
	sm.Add(stm.URL{"loc": "/about", "changefreq": "monthly"})

	w.Write(sm.XMLContent())
}

// staticHandler serves files from "static" folder: CSS styles and pictures.
func staticHandler() http.Handler {
	fileServer := http.FileServer(http.Dir(repoLocalPath + "/static/"))
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

// RemoveLeadingSlash removes slash ("/") if it's the first character in a
// string.
func removeLeadingSlash(s string) string {
	prefix := "/"
	if strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	return s
}
