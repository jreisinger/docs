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
	if s.Err() != nil {
		return nil, s.Err()
	}

	return terms, nil
}

func randTerm() ([]byte, error) {
	terms, err := fetchTerms("https://raw.githubusercontent.com/jreisinger/homepage/master/data/notes/terms.md")
	if err != nil {
		return nil, err
	}
	count := len(terms)
	i := rand.Intn(count)
	t := terms[i]
	md := fmt.Sprintf("\n\nTerm %d/%d\n\n", i+1, count) + t.name + "\n" + t.explanation
	return []byte(md), nil
}
