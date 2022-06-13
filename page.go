package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
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

var (
	ErrorNotFound = errors.New(`<a href="/search">search</a>, and you will find`)
)

// loadPage maps urlPath fo file or directory from the local repo "data" folder
// and returns it as a page. Thus it implicitly limits URL paths to files or
// directories that exist in the repo "data" folder.
func loadPage(urlPath string) (*Page, error) {
	// map URL path to filesystem path (without .md)
	dataPath := path.Join(repoLocalPath, "/data")
	filePath := path.Join(dataPath, urlPath)

	if isDir(filePath) { // if file is a dir return list of files within dir
		files := listFiles(filePath, false)
		if strings.HasSuffix(filePath, "blog") {
			files = listFiles(filePath, true)
		}
		return &Page{Title: path.Base(filePath), Files: files, RepoURL: repoRemoteUrl, UrlPath: urlPath, IsDir: true}, nil
	} else if isFile(filePath + ".md") { // if file is a file return file contents
		lastModified, err := lastModified(repoLocalPath, filePath+".md")
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadFile(filePath + ".md")
		if err != nil {
			return nil, err
		}

		body := mdToHtml(data)

		return &Page{Title: path.Base(filePath), Body: body, RepoURL: repoRemoteUrl, UrlPath: urlPath, LastModified: lastModified}, nil
	}

	err := ErrorNotFound
	return nil, err
}

// mdToHtml converts markdown to HTML.
func mdToHtml(md []byte) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs ^ parser.MathJax

	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(md, p, nil)
	return template.HTML(html)
}

// listFiles returns a sorted list of files in a directory.
func listFiles(filePath string, reverseSort bool) []string {
	dir, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // -1 means return all entries
	if err != nil {
		log.Fatal(err)
	}
	var files []string
	for _, fi := range fileInfos {
		basename := fi.Name()

		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		if fi.IsDir() {
			name += "/"
		}

		files = append(files, name)
	}

	// Sort list of files.
	if reverseSort {
		sort.Sort(sort.Reverse(sort.StringSlice(files)))
	} else {
		sort.Sort(sort.StringSlice(files))
	}

	return files
}

// isDir returns true is filePath is a directory.
func isDir(filePath string) bool {
	fi, err := os.Lstat(filePath)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

// isFile returns true is filePath is a regular file.
func isFile(filePath string) bool {
	fi, err := os.Lstat(filePath)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

// lastModified returns when a file from a repo was last modified according to git.
func lastModified(repoPath, filename string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%ci", filename)
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, output)
	}
	return fmt.Sprintf("%s", output), nil
}
