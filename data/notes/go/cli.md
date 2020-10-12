A simple CLI tool:

```
// Mygrep reads from STDIN or file(s) and prints lines containing pattern.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	progName := os.Args[0]

	log.SetPrefix(progName + ": ")
	log.SetFlags(0) // no timestamp

	if !(len(os.Args) > 1) {
		fmt.Fprintf(os.Stderr, "Usage: %s PATTERN [FILE ...]\n", progName)
		os.Exit(1)
	}

	pattern := os.Args[1]
	files := os.Args[2:]
	if len(files) == 0 {
		matchLines(os.Stdin, pattern)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			matchLines(f, pattern)
		}
	}
}

func matchLines(f *os.File, pattern string) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		if re.MatchString(s.Text()) {
			fmt.Println(s.Text())
		}
	}
}
```

See also

* https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/
