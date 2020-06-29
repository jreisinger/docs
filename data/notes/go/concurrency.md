Concurrent programming structures the program as a composition of several independent activities. It's important because:

* web servers handle thousands of clients at once
* GUIs render animations while simultaneously performing computations and network requests
* even traditional batch programs (read input data, compute, write output) hide I/O latency and exploit multiple processors

The number of processors in a computer grows every year. The speed of processors not so much.

Go enables two styles of concurrent programming:

1. CSP (communicating sequential processes) - values are passed between autonomous activities (goroutines) but variables are for the most part confined to a single activity.
2. Shared memory multithreading - more traditional model.

Reasoning about concurrent programs is inherently more difficult than about sequential ones.

# Goroutines

A *goroutine* is a concurrently executing activity. When a program starts, its only goroutine is the one that calls the `main` function - the *main goroutine*. You can create new goroutines with the `go` keyword.

```
f()     // call f(); wait for it to return
go f()  // create a new goroutine that calls f(); don't wait
```

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

`select` statement is like a switch but for channels. `select` picks the first channel that is ready a receives from it. If more than one of the channels are ready, then it randomly picks which one to receive from. The default case happens immediately if none of the channels is ready.

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

# Sources

* Caleb Doxsey: Introducing Go (2016)
* John Graham-Cumming: Go Programming Basics (2017)

# See also

* [fetchall.go](https://github.com/jreisinger/go/blob/master/http/fetchall.go)
* [fetchall2.go](https://github.com/jreisinger/go/blob/master/http/fetchall2.go)
