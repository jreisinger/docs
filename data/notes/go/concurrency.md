Concurrent programming

* a way to structure software
* a composition of several independent computations

Why?

* the world is a complex system of interacting, independently behaving activities
* we want to model and interact with this world
* the number of processors in a computer grows every year the speed of processors not

Go enables two styles of concurrent programming:

1. Communicating sequential processes (CSP) - values are passed between autonomous activities (goroutines) but variables are for the most part confined to a single activity.
2. Shared memory multithreading - more traditional model.

Reasoning about concurrent programs is inherently more difficult than about sequential ones.

# Goroutines

A *goroutine* is a concurrently executing activity. When a program starts, its only goroutine is the one that calls the `main` function - the *main goroutine*. You can create new goroutines with the `go` keyword.

```
f()     // call f(); wait for it to return
go f()  // create a new goroutine that calls f(); don't wait
```

```
// fib-spinner.go -- calculate fibonacci number using 
// a slow algorithm and show a spinner while calculating
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n := 45
	fibN := fib(n) // slow algorithm
	fmt.Printf("\rfib(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
```

A goroutine is stopped by returning from `main` or by exiting the program.

# Goroutines and channels

A *channel* is a way for gouroutines to communicate with each other and *synchronize* their execution.

When `pinger` or `ponger` attempts to send a message on the channel, it will
wait until `printer` is ready to receive the message (blocking):

```go
// ping-pong.go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go pinger(ch)
    go ponger(ch)
    go printer(ch)

    // Wait for Enter to exit.
    var input string
    fmt.Scanln(&input)
}

func pinger(ch chan string) { for { ch <- "ping" } }
func ponger(ch chan string) { for { ch <- "pong" } }

func printer(ch chan string) {
    for {
        fmt.Println(<-ch)
        time.Sleep(time.Second * 1)
    }
}
```

# Select

* `select` statement is like a switch but for channels (i.e. it's not for expressions but for communications)
* first all channels are evaluated
* blocks until one communication can proceed, which then does
* if multiple can proceed, select chooses pseudo-randomly
* the default case, if present, executes immediately if no channel is ready

```go
// select.go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        for {
            ch1<- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func() {
        for {
            ch2<- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()

    go func() {
        for {
            select {
            case msg1 := <-ch1:
                fmt.Println(msg1)
            case msg2 := <-ch2:
                fmt.Println(msg2)
            default:
                fmt.Println("nothing ready")
                time.Sleep(time.Second * 1)
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
```

# Timing out a goroutine

```
// timer.go
package main

import (
    "fmt"
    "time"
)

// emit emits words on wordCh for 3 seconds.
func emit(wordCh chan string) {
    defer close(wordCh) // close channel when return-ing

    words := []string{"The", "quick", "brown", "fox"}

    i := 0                              // index
    t := time.NewTimer(3 * time.Second) // function's timer

    for {
        select { // select not switch :-)

        case wordCh <- words[i]: // someone reads from channel
            i += 1
            if i == len(words) {
                i = 0
            }

        case <-t.C: // timer goes off
            return
        }

    }
}

func main() {
    wordCh := make(chan string)

    go emit(wordCh)

    // range over the channel until closed
    for word := range wordCh {
        fmt.Printf("%s ", word)
    }
}
```

# Scalable work system using goroutines, channels and interfaces

```
// work.go
// A common use case for Go is to take a stream of jobs of work and perform them
// automatically scaling up and down as work becomes available. See the video 
// Intermediate Go programming - Building a scalable work system.
// ./work < urls.txt
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type task interface {
	process()
	output()
}

type factory interface {
	create(line string) task
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	// Create 1000 goroutines to process the tasks.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.output()
	}
}

type HTTPTask struct {
	url string
	ok  bool
}

func (h *HTTPTask) process() {
	resp, err := http.Get(h.url)
	if err != nil {
		h.ok = false
		return
	}
	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false
}

func (h *HTTPTask) output() {
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type Factory struct {
}

func (f *Factory) create(line string) task {
	h := &HTTPTask{}
	h.url = line
	return h
}

func main() {
	f := &Factory{}
	run(f)
}
```

# Sources

* Donovan, Kernighan: The Go Programming Language (2015), ch.8
* Caleb Doxsey: Introducing Go (2016)
* John Graham-Cumming: Go Programming Basics (2017)
* John Graham-Cumming: Intermediate Go Programming (2015) - Building a scalable work system

# See also

* https://github.com/jreisinger/go-concurrency-patterns
* https://golang.org/doc/effective_go.html#concurrency
