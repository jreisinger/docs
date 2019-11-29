# Go

* [Links to explore](#links-to-explore)
* [My Go stuff on the Web](#my-go-stuff-on-the-web)
* [Reading input in "streaming" mode (`bufio.Scanner`)](#reading-input-in-"streaming"-mode-`bufio.scanner`)
* [Reading input in "slurp" mode (`ioutil.ReadFile`)](#reading-input-in-"slurp"-mode-`ioutil.readfile`)
* [Regular expressions](#regular-expressions)
* [CLI tools template](#cli-tools-template)
* [Slices and `append`](#slices-and-`append`)
* [Interfaces](#interfaces)
* [Goroutines, channels and `select`](#goroutines,-channels-and-`select`)
* [Modules](#modules)

## Links to explore

* [wiki](https://github.com/golang/go/wiki)
* [asciigraph](https://github.com/guptarohit/asciigraph)

## My Go stuff on the Web

*  [notes](https://jreisinger.github.io/notes/tags/go/)
*  [repos](https://github.com/jreisinger?utf8=%E2%9C%93&tab=repositories&q=&type=&language=go), [gists](https://gist.github.com/search?utf8=%E2%9C%93&q=user%3Ajreisinger+language%3Ago)
*  PerlMonks: [Camel vs. Gopher](https://perlmonks.org/?node_id=1226977), [Does Go steal from Perl? :-)](https://perlmonks.org/?node_id=1219775)

## Reading input in "streaming" mode (`bufio.Scanner`)

Read from STDIN:

```
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t%d\n", line, n)
        }
    }
}
```

Read from STDIN or files:

```
// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    if err := input.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
```

## Reading input in "slurp" mode (`ioutil.ReadFile`)

```
// dup3 reads only named files, not the standard input, since ReadFile requires
// a file name argument.
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

Under the covers, `bufio.Scanner`, `ioutil.ReadFile`, and `ioutil.WriteFile` use the `Read` and `Write` methods of `*os.File`, but it's rare that most programmers need to access those lower-level routines directly.

## Regular expressions

* https://yourbasic.org/golang/regexp-cheat-sheet/

```
// wc counts the number of words in a string
package main

import (
    "fmt"
    "regexp"
)

func main() {
    s := "The quick brown fox. Or was it a blue fox?"
    ws := regexp.MustCompile(`\s+`)
    words := ws.Split(s, -1) 
    fmt.Printf("%d\n", len(words)) // 10
}
```

## CLI tools template

```
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "PATTERN", "FILE")
        os.Exit(1)
    }

    //pattern := os.Args[1]
    //file := os.Args[2]

    // Set up CLI tool style of logging.
    log.SetPrefix(os.Args[0] + ": ")
    log.SetFlags(0)
}
```

## Slices and `append`

```
// Create a slice add and then append some elements.
package main

import (
    "fmt"
)

// printSlice prints a slice of bytes with some useful info.
func printSlice(s []string) {
    fmt.Printf("0th_elem_addr=%p len=%d  cap=%d  %v\n",
        s, len(s), cap(s), s)
}

func main() {
    s := make([]string, 2, 3)
    printSlice(s) // 0th_elem_addr=0xc000064180 len=2  cap=3  [ ]
    s[0] = "the"
    s[1] = "quick"
    printSlice(s) // 0th_elem_addr=0xc000064180 len=2  cap=3  [the quick]
    s = append(s, "brown")
    printSlice(s) // 0th_elem_addr=0xc000064180 len=3  cap=3  [the quick brown]
    s = append(s, "fox")
    printSlice(s) // 0th_elem_addr=0xc00004a0c0 len=4  cap=6  [the quick brown fox]
}
```

## Interfaces

```
package main

import (
    "fmt"
    "math/rand"
)

type shuffler interface {
    Len() int
    Swap(i, j int)
}

func shuffle(s shuffler) {
    for i := 0; i < s.Len(); i++ {
        j := rand.Intn(s.Len() - 1) // why not only s.Len()?
        s.Swap(i, j)
    }
}

type intSlice []int

func (is intSlice) Len() int {
    return len(is)
}

func (is intSlice) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

func main() {
    is := intSlice{1, 2, 3, 5, 6}
    shuffle(is)
    fmt.Println(is) // [2 6 1 3 5]
}
```

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

## Modules

A module is a collection of related Go packages that are versioned together as a single unit.

Relationship between repositories, modules and packages:

* A *repository* contains one or more Go modules
* Each *module* contains one or more Go packages
* Each *package* consists of one or more Go source files in a single directory

More 

* https://github.com/golang/go/wiki/Modules
