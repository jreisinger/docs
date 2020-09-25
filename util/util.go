package util

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func LastModified(repoPath, filename string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%ci", filename)
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, output)
	}
	return fmt.Sprintf("%s", output), nil
}
