# Go

## My Go stuff on the Web

*  [repos](https://github.com/jreisinger?utf8=%E2%9C%93&tab=repositories&q=&type=&language=go)
*  [gists](https://gist.github.com/search?utf8=%E2%9C%93&q=user%3Ajreisinger+language%3Ago)
*  [notes](https://jreisinger.github.io/notes/tags/go/)
*  PerlMonks: [Camel vs. Gopher](https://perlmonks.org/?node_id=1226977) [Does Go steal from Perl? :-)](https://perlmonks.org/?node_id=1219775)

## Finding duplicate lines

*A.k.a. templates for working with files and STDIN*

### Reading input in "streaming" mode (`bufio.Scanner`)

Read from STDIN:

```go
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

## Useful functions

```
// printSlice prints a slice of bytes with some useful info.
func printSlice(s []byte) {
    fmt.Printf("0th_elem_addr=%p len=%d  cap=%d  %v\n",
                              s, len(s), cap(s), s)
}
```
