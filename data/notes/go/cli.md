CLI tools template

```
package main

import (
	"log"
	"os"
)

func main() {
	// Set up CLI tool style of logging.
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0) // no timestamp

	if len(os.Args) != 3 {
		// write to stderr and call os.Exit(1)
		log.Fatalf("usage: %s %s %s", os.Args[0], "PATTERN", "FILE")
	}

	//pattern := os.Args[1]
	//file := os.Args[2]
}
```

[Run](https://play.golang.org/p/omvP2uhNVQX).
