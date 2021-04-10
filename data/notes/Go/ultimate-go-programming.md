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

Don't do object oriented design but data oriented design.

> Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

Go prefers convention over configuration because configuration is limiting. Let's [group](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go) by what one can do not by what one is. Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

We are writing code for today, designing and architecting for tomorrow.

Layers of APIs:

* High level - ease of use, exported, unit/integration tested (might replace tests below)
* Low level - maybe exported, unit tested (might replace tests below)
* Primitive level - do one thing well, not exported, unit tested

When am I done:

* 70 - 80% test coverage
* ask what can change, from technical and business perspective, and refactor the code to be able to handle that change

Hints:

* solve one problem at a time, don't be overwhelmed by everything

Steps:

1. Identify problems to solve
2. Write a little code
3. Write tests
4. Refactor for change
5. Refactor for simplicity, consistency, readibility

---

* https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/
* https://github.com/ardanlabs/gotraining-studyguide
