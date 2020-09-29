# Reading input in "streaming" mode (`bufio.Scanner`)

`bufio` package helps make input and output efficient and convenient. Its `Scanner` type reads input and breaks it into lines or words. `bufio` is good for "streaming" mode where input is read and broken into lines on the fly.

Read from STDIN (dup1.go):

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
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

Read from STDIN or files (dup2.go):

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
        fmt.Fprintln(os.Stderr, "dup2 reading input:", err)
    }
}
```

Test `countLines` using [testfile](https://github.com/jreisinger/testfile):

```
// dup2_test.go
package main

import (
    "os"
    "testing"

    "github.com/jreisinger/testfile"
)

func TestCountLines(t *testing.T) {
    counts := make(map[string]int)

    tf := testfile.New("line1\nline2\nline2")
    defer tf.Remove()

    file, err := os.Open(tf.Name)
    if err != nil {
        t.Errorf("cant' open test file %s: %v", tf.Name, err)
    }   
    defer file.Close()

    countLines(file, counts)
    if counts["line1"] != 1 { 
        t.Errorf("count for 'line1' should be 1 is %d", counts["line1"])
    }   
    if counts["line2"] != 2 { 
        t.Errorf("count for 'line2' should be 2 is %d", counts["line2"])
    }   
}
```

# Reading input in "slurp" mode (`ioutil.ReadFile`)

dup3.go:

```
// Dup3 reads only named files, not the standard input, 
// since ReadFile requires a file name argument.
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
