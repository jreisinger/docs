See also `notes/go/basics`.

Consider this program that de-duplicates input lines:

```
// dedup.go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    files := os.Args[1:]

    if len(files) == 0 {
        dedup(os.Stdin)
    } else {
        for _, file := range files {
            f, err := os.Open(file)
            if err != nil {
                fmt.Fprintf(os.Stderr, "%s", err)
                continue
            }
            defer f.Close()
            dedup(f)
        }
    }
}

func dedup(f *os.File) {
    seen := make(map[string]bool) // a set of strings

    input := bufio.NewScanner(f)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
```

`bufio.NewScanner` (inside `dedup` function) takes `io.Reader` as parameter. `io.Reader` is an interface:

```
package io // import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
}
```

It means that anything that has a `Read` function with the given signature ^ is (or implements) a Reader. `bufio.NewScanner` can take `*os.File` since `*os.File` implements the given Read function:

```
package os // import "os"

type File struct {
	// Has unexported fields.
}
    File represents an open file descriptor.

func (f *File) Read(b []byte) (n int, err error)
```
