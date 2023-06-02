Reviewed: 2023-06-02

# Package `net/http`

* standard library package for implementing HTTP servers (and clients)

## `Handler` interface

* foundational element of `net/http`

```
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

## `ListenAndServe`

* function that runs forever until it fails (always with a non-nil error)
* requires an instance of the `Handler` interface to which all requests should be dispatched

```
// Shop - a super simple e-shop showing price of shoes and socks
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
```

* the server will list all the inventory for every request, regardless of URL

A more realistic server triggers different behaviours based on the path component of the URL:

```
// Shop v2
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	// called like: /price?item=socks
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}
```

* `WriteHeader` must be called before anything is written to `w`
* equivalently you could use the `http.Error` utility function:

```
msg := fmt.Sprintf("no such page: %s\n", r.URL)
http.Error(w, msg, http.StatusNotFound) // 404
```

## `ServeMux` struct

* it's convenient to define logic for each URL in a separate function or method
* related URLs (e.g. `/images/*.png`) might need similar logic
* enters `ServeMux`, a request multiplexer, to simplify the association between URLs and handlers
* a `ServeMux` aggregates a collection of `http.Handler`s into a single `http.Handler`

```
// Shop v3
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```

* `db.list` is a function (or a method) that implements handler-like behavior but since it does not satisfy `http.Handler` interface it can't be passed directly to `mux.Handle`
* the expression `http.HandlerFunc(db.list)` is a *conversion*, not a function call, since `http.HandlerFunc` is a type
* because registering a handler this way is so common, `ServeMux` has a convenience method `HandleFunc`

```
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
```

* also for convenience `net/http` provides a global `ServeMux` instance called `DefaultServeMux` and package level functions `http.Handle` and `http.HandleFunc`
* to use `DefaultServeMux` pass `nil` to `ListenAndServe`

```
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```

NOTE: the web server invokes each handler in a new goroutine, so handlers must take precautions such as locking when accessing variables that other goroutines, including other requests to the same handler, may be accessing.

## `WriteHeader` method of `ResponseWriter` interface

These two handler functions are equivalent:

```
func ok1(w http.ResponseWriter, r *http.Request) {
}

func ok2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
}
```

both generate line of information with protocol and status code (Go doc calls this line response header) plus two headers:

```
< HTTP/1.1 200 OK
< Date: Sat, 16 May 2020 15:17:55 GMT
< Content-Length: 0
< 
```

## `Error` function

```
func err(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 something's wrong with the server", http.StatusInternalServerError)
}
```

generates also response body:

```
< HTTP/1.1 500 Internal Server Error
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sat, 16 May 2020 15:29:09 GMT
< Content-Length: 38
< 
500 something's wrong with the server
```

# Sources

* The Go Programming Language (2015)
* Black Hat Go (2020)
