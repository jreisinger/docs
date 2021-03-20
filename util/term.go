package util

import (
	"bufio"
	"fmt"
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
	terms = append(terms, t)
	if s.Err() != nil {
		return nil, s.Err()
	}

	// First item is empty. Not nice, I know ...
	return terms[1:], nil
}

// RandTerm returns a random term in MarkDown format.
func RandTerm() ([]byte, error) {
	url := "https://raw.githubusercontent.com/jreisinger/homepage/master/data/notes/terms.md"
	terms, err := fetchTerms(url)
	if err != nil {
		return nil, err
	}
	i := rand.Intn(len(terms))
	t := terms[i]
	stats := fmt.Sprintf("(term %d/[%d](%s))", i+1, len(terms), url)
	md := t.name + "\n" + t.explanation + "\n" + stats
	return []byte(md), nil
}
