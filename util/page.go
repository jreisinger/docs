package util

import (
	"errors"
	"html/template"
	"log"
	"path"
)

type Page struct {
	Title        string
	Body         template.HTML
	RepoURL      string
	IsDir        bool
	Files        []string
	UrlPath      string
	LastModified string
}

// RenderPage renders a file or a directory as an HTML page.
func RenderPage(repoURL string, repoPath string, urlPath string) (*Page, error) {
	dataPath := repoPath + "/data"

	// map URL path to filesystem path (without .md)
	filePath := dataPath + "/" + urlPath

	title := path.Base(filePath)

	if IsDir(filePath) {
		// if file is a dir return list of files within dir
		files := ListFiles(filePath)
		return &Page{Title: title, Files: files, RepoURL: repoURL, UrlPath: urlPath, IsDir: true}, nil
	} else if IsFile(filePath + ".md") {
		lastModified, err := LastModified(filePath + ".md")
		if err != nil {
			log.Fatal(err)
		}
		// if file is a file return file contents
		body := MdToHtml(filePath + ".md")
		return &Page{Title: title, Body: body, RepoURL: repoURL, UrlPath: urlPath, LastModified: lastModified}, nil
	}

	err := errors.New("rendering page: resource not found (not all those who wander are lost)")
	return nil, err
}
