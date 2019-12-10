## Goroutines, channels and `select`

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
