package main

import (
	"regexp"
	"strings"
)

// grepFileContent searches for pattern inside inside a file.
func grepFileContent(content string, pattern *regexp.Regexp) bool {
	return pattern.Match([]byte(content))
}

// grepFilePath searches for pattern inside a file pathh.
func grepFilePath(path string, pattern *regexp.Regexp) bool {
	match := pattern.FindString(path)
	if match == "" {
		return false
	}
	return true
}

// filesystemToURL changes filesystem path into URL path.
func filesystemToURL(path string) (string, error) {
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
