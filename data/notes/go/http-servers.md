# Basic servers

## with default handler

```
package main

import (
        "fmt"
        "log"
        "net/http"
)

// handler function
func hello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
        // register the handler function for the given pattern "/hello"
        http.HandleFunc("/hello", hello)
        // nil handler means DefaultServeMux is used
        log.Fatal(http.ListenAndServe(":8080", nil))
}
```

```
> curl localhost:8080/hello?name=dude
Hello dude
```

## with custom handler - simple router

```
package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

// ServeHTTP method satisfies the Handler interface
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a")
	case "/b":
		fmt.Fprint(w, "Executing /b")
	case "/c":
		fmt.Fprint(w, "Executing /c")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}
```

* see https://github.com/jreisinger/util for a bit more sophisticated router

## with custom handler - simple middleware

```
package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.RequestURI()
	remAddr := r.RemoteAddr
	log.Printf("start handling request for %s from %s\n", uri, remAddr)
	l.Inner.ServeHTTP(w, r)
	log.Printf("done handling request for %s from %s\n", uri, remAddr)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f} // wrapper around f
	http.ListenAndServe(":8000", &l)
}
```

# Reponse handling

## 200 OK

These two handler functions are equivalent:

```
func ok1(w http.ResponseWriter, r *http.Request) {
}

func ok2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
}
```

both generate line of information with protocol and status code plus two headers:

```
< HTTP/1.1 200 OK
< Date: Sat, 16 May 2020 15:17:55 GMT
< Content-Length: 0
< 
```

## 500 Internal Server Error

```
func err1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError) // 500
}
```

generates only line of information (Go docs calls it response header):

```
< HTTP/1.1 500 Internal Server Error
< Date: Sat, 16 May 2020 15:25:28 GMT
< Content-Length: 0
< 
```

```
func err2(w http.ResponseWriter, r *http.Request) {
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

# The `http.Handler` Inteface

* foundational element of `net/http`

```
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

`ListenAndServe`

* function that runs forever until it fails (always with a non-nil error)
* requires an instance of the `Handler` interface to which all requests should be dispatched

Super simple e-commerce site:

```
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
```

# Sources

* https://learning.oreilly.com/library/view/black-hat-go
* The Go Programming Language (ch. 7.7)
