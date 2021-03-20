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

// RandQuote returns a random quote in MarkDown format.
func RandQuote() ([]byte, error) {
	url := "https://raw.githubusercontent.com/jreisinger/quotes/master/quotes.txt"
	quotes, err := fetchQuotes(url)
	if err != nil {
		return nil, err
	}
	i := rand.Intn(len(quotes))
	q := quotes[i]
	stats := fmt.Sprintf("(quote %d/[%d](%s))", i+1, len(quotes), url)
	md := "\n\n" + "> " + q.what + " --- " + q.author + "\n\n" + stats + "\n"
	return []byte(md), nil
}
