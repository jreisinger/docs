package util

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

type quote struct {
	what   string
	author string
}

func fetchQuotes(url string) ([]quote, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var quotes []quote
	for _, q := range strings.Split(string(data), "\n\n") {
		parts := strings.Split(q, " -- ")
		if len(parts) != 2 { // some quotes are without author
			quotes = append(quotes, quote{q, ""})
			continue
		}
		quotes = append(quotes, quote{parts[0], parts[1]})
	}
	return quotes, nil
}

func randQuote() ([]byte, error) {
	quotes, err := fetchQuotes("https://raw.githubusercontent.com/jreisinger/quotes/master/quotes.txt")
	if err != nil {
		return nil, err
	}
	count := len(quotes)
	i := rand.Intn(count)
	q := quotes[i]
	md := fmt.Sprintf("Quote %d/%d\n\n> ", i+1, count) + q.what + " --- " + q.author
	return []byte(md), nil
}
