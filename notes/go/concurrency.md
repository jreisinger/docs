Reviewed: 2023-06-02

Concurrent programming

* a way to structure software
* a composition of several independent computations
* concurrency != parallelism
* inherently more difficult than sequential

Why?

* the world is a complex system of interacting, independently behaving pieces
* we want to model and interact with this world
* the number of processors in a computer grows faster than their speed

Go enables two styles of concurrent programming:

1. Communicating sequential processes (CSP) - see below
2. Shared memory multithreading - more traditional model

Goroutines

* concurrently executing activities
* when a program starts, its only goroutine is the one that calls the `main` function
* you can create new goroutines with the `go` keyword

```
f()     // call f(); wait for it to return
go f()  // create a new goroutine that calls f(); don't wait
```

* a goroutine is stopped by returning from `main` or by exiting the program.

Channels - a way for gouroutines to

* communicate with each other
* synchronize their execution

Select

* `select` statement is like a switch but for channels (i.e. it's not for expressions but for communications)
* first all channels are evaluated
* blocks until one communication can proceed, which then does
* if multiple can proceed, select chooses pseudo-randomly
* the default case, if present, executes immediately if no channel is ready

More

* https://github.com/jreisinger/gokatas
* The Go Programming Language (2015)
* Go In Practice (2016)
