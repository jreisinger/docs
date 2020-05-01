## Simple server

```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
    // register the handler function (hello)
    // for the given pattern ("/hello")
    http.HandleFunc("/hello", hello)
    // nil handler means DefaultServeMux is used
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

```
> curl localhost:8080/hello?name=dude
Hello dude
```

## Simple router

```
package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

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

## Simple middleware

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

## More

* https://learning.oreilly.com/library/view/black-hat-go
