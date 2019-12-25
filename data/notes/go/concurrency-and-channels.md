A *goroutine* is a function capable of running concurrently with other functions. Create a gouroutine with the `go` keyword. 

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

`select` statement is like a switch but for channels. `select` picks the first channel that is ready a receives from it. If more than one of the channels are ready, then it randomly picks which one to receive from. The default case happens immediately if none of the channels are ready.

```go
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

Taken from "Introducing Go". See also [fetchall.go](https://github.com/jreisinger/go/blob/master/http/fetchall.go) and [fetchall2.go](https://github.com/jreisinger/go/blob/master/http/fetchall2.go).
