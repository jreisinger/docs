## Bill's Terminology

* mechanics - how things work
* semantics - how things behave

## Philosophy

> Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code. -- Brian Kernighan

1. Make it correct.
2. Make it clear (simple, straightforward).
3. Make it fast.

In that order.

Engineering is not about just hacking the code. It's about evaluating the costs and benefits.

We are writing code for today, designing and architecting for tomorrow.

> Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

Don't do object oriented design but data oriented design.

Go prefers convention over configuration because configuration is limiting. Let's group by what one can do not by what one is. In Go there are [no classes](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_1.go) (or [Perl version](https://github.com/jreisinger?tab=repositories&q=animal)) and subclasses but [Interfaces](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go) and composition.

Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

## Types

Types are life. [Types](https://play.golang.org/p/24H4L7Gofrz) tell you

* what is the memory size allocated for the value
* what does the value represent (e. g. byte `11111111` represents number 255 if the type is `int`)

A [struct](https://play.golang.org/p/Av0NOh_cu_K) is a user-defined type that contains named fields:

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

Prefer functions over methods because functions are more precise.

Use methods for state or decoupling.

## Solving problems with software

Solve one problem at a time, don't be overwhelmed by everything.

Steps:

1. Identify concrete tasks that will solve the problem.
2. Write a little [code](https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/decoupling/example1/example1.go) for a concrete task.
3. Start at primitive level and deploy ASAP.
4. Write tests for your code.
5. Repeat 1. - 4. until the problem is solved.
6. [Refactor](https://github.com/ardanlabs/gotraining/tree/master/topics/go/design/composition/decoupling) for change by decoupling.
7. Refactor for readibility (mental model, names, cleverness).

Design programs as layers of API:

* [Primitive](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L66-#L102) layer - do one thing well, not exported, unit tested
* [Low](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L104-#L147) layer - builds on primitive layer, maybe exported, unit tested (might replace primitive layer tests)
* [High](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L149-#L166) layer - ease of use, exported, unit/integration tested 

When am I done:

* 70 - 80% test coverage
* ask what can change, from technical and business perspective, and refactor the code to be able to handle that change

---

* https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/
* https://github.com/ardanlabs/gotraining-studyguide
* https://github.com/ardanlabs/gotraining
