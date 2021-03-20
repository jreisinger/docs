package util

import (
	"bufio"
	"math/rand"
	"net/http"
	"strings"
)

type term struct {
	name        string
	explanation string
}

func fetchTerms(url string) ([]term, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var terms []term

	s := bufio.NewScanner(resp.Body)
	t := term{}
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "##") {
			terms = append(terms, t)
			t = term{}
			t.name = s.Text()
		} else {
			t.explanation += "\n" + s.Text()
		}
	}
	if s.Err() != nil {
		return nil, s.Err()
	}

	return terms, nil
}

// RandTerm returns a random term in MarkDown format.
func RandTerm() ([]byte, error) {
	terms, err := fetchTerms("https://raw.githubusercontent.com/jreisinger/homepage/master/data/notes/terms.md")
	if err != nil {
		return nil, err
	}
	t := terms[rand.Intn(len(terms))]
	md := "\n" + t.name + "\n" + t.explanation
	return []byte(md), nil
}
