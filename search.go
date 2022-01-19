package main

import (
	"regexp"
	"strings"
)

func What(pattern string) bool {
	var searchOnlyPath bool
	if strings.HasPrefix(pattern, "path:") {
		searchOnlyPath = true
		pattern = strings.TrimPrefix(pattern, "path:")
	}
	return searchOnlyPath
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
