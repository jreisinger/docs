package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	rand.Seed(time.Now().UTC().UnixNano())
	rawurl := "https://raw.githubusercontent.com/jreisinger/terms/main/terms.md"
	url := "https://github.com/jreisinger/terms/blob/main/terms.md"
	terms, err := fetchTerms(rawurl)
	if err != nil {
		return nil, err
	}
	i := rand.Intn(len(terms))
	t := terms[i]
	stats := fmt.Sprintf("[term](%s) %d/%d", url, i+1, len(terms))
	md := t.name + "\n" + t.explanation + "\n" + stats
	return []byte(md), nil
}
