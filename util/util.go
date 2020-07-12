package util

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// MdToHtml converts markdown to HTML.
func MdToHtml(filePath string) template.HTML {
	md, err := ioutil.ReadFile(filePath)
	Check(err)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs ^ parser.MathJax

	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(md, p, nil)
	return template.HTML(html)
}

// ListFiles returns a list of files in a directory.
func ListFiles(filePath string) []string {
	dir, err := os.Open(filePath)
	Check(err)
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // -1 means return all entries
	Check(err)
	var files []string
	for _, fi := range fileInfos {
		basename := fi.Name()

		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		if fi.IsDir() {
			name += "/"
		}

		files = append(files, name)
	}

	sort.Strings(files)
	return files
}

// IsDir returns true is filePath is a directory.
func IsDir(filePath string) bool {
	fi, err := os.Lstat(filePath)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

// IsFile returns true is filePath is a regular file.
func IsFile(filePath string) bool {
	fi, err := os.Lstat(filePath)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

// Check handles error.
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// RemoveTrailingSlash removes slash ("/") if it's the last character in a
// string.
func RemoveTralingSlash(s string) string {
	suffix := "/"
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// GrepFileContent searches for pattern inside inside a file.
func GrepFileContent(content string, pattern *regexp.Regexp) bool {
	return pattern.Match([]byte(content))
}

// GrepFilePath searches for pattern inside a file pathh.
func GrepFilePath(path string, pattern *regexp.Regexp) bool {
	match := pattern.FindString(path)
	if match == "" {
		return false
	}
	return true
}

// FilesystemToURL changes filesystem path into URL path.
func FilesystemToURL(path string) (string, error) {
	// Trim .md suffix
	path = strings.TrimSuffix(path, ".md")

	// Trim .*data/ prefix
	rx, err := regexp.Compile(`.*data/`)
	if err != nil {
		return path, err
	}
	path = rx.ReplaceAllString(path, "")
	return path, nil
}
