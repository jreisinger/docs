## Terminology

* mechanics - how things work
* semantics - how things behave

## Philosophy

1. Make it correct.
2. Make it clear.
3. Make it fast.

In that order.

> Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code. -- Brian Kernighan

Engineering is not about just hacking the code. It's about evaluating the costs and benefits.

## Types

Types are life. [Types](https://play.golang.org/p/Rdoskvi8e_c) tell you

* what is the memory size allocated for the value
* what does the value represent (e. g. byte `11111111` represents number 255 if the type is `int`)

A [struct](https://play.golang.org/p/jFQMm91N0nQ) is a user-defined type that contains named fields:

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

## Functions

Arguments are passed by [value](https://play.golang.org/p/LoFTsgS3BUQ). A goroutine can access only data within its active frame (a slice of the stack) so the argument needs to be copied upon a function call. This is a mechanism for isolating memory mutations. It also means that if you need to mutate data you need to use a pointer.

## Design

Don't do object oriented design but data oriented design.

> Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

Go prefers convention over configuration because configuration is limiting. Let's group by what one can do not by what one is. In Go there are [no classes](https://github.com/jreisinger?tab=repositories&q=animal) and subclasses but [Interfaces](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go) and composition.

Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

We are writing code for today, designing and architecting for tomorrow.

Solve one problem at a time, don't be overwhelmed by everything.

Layers of APIs:

* [High](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L149-#L166) level - ease of use, exported, unit/integration tested (might replace tests below)
* [Low](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L104-#L147) level - builds on primitive layer, maybe exported, unit tested (might replace tests below)
* [Primitive](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L66-#L102) level - do one thing well, not exported, unit tested

When am I done:

* 70 - 80% test coverage
* ask what can change, from technical and business perspective, and refactor the code to be able to handle that change

Steps:

1. Identify problems to solve
2. Write a little [code](https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/decoupling/example1/example1.go) at primitive/low/high level (start at primitive level and deploy ASAP)
3. Write tests at primitive/low/high level
4. [Refactor](https://github.com/ardanlabs/gotraining/tree/master/topics/go/design/composition/decoupling) for change by decoupling
5. Refactor for readibility (mental model, names, semantics, cleverness)

---

* https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/
* https://github.com/ardanlabs/gotraining-studyguide
* https://github.com/ardanlabs/gotraining
