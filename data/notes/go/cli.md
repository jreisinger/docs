CLI tools template

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
