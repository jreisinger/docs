## Terminology

* mechanics - how things work
* semantics - how things behave

## Philosophy

1. Make it correct.
2. Make it clear.
3. Make it fast.

In that order.

> Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code. -- Brian Kernighan

## Types

Types are life. Types tell you

* what is the [memory size](https://play.golang.org/p/s1M9pNjIIKp) allocated for the value
* what does the value represent (e. g. this byte `00001010` represents number 10 if the type is `int`)

A [struct](https://play.golang.org/p/Q90vDc_T77X) is a user-defined type that contains named fields:

```go
// Declaration of anonymous (literal) struct type + initialization.
e := struct {
	flag    bool
	counter int
}{
	flag: true,
	counter: 10,
}
```

## Design

Go prefers convention over configuration because configuration is limiting. Let's [group](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go) by what one can do not by what one is. Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

---

* https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/
* https://github.com/ardanlabs/gotraining-studyguide
