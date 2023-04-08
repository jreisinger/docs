Reviewed: 2023-04-07

```go
package main

import (
	"fmt"
	"regexp"
)

var (
	ws  = regexp.MustCompile(`\s+`)
	num = regexp.MustCompile(`[0-9]+`)
)

func main() {
	s := "The quick brown fox."

	words := ws.Split(s, -1)
	fmt.Printf("%d\n", len(words)) // 4

	matched := num.MatchString(s)
	fmt.Printf("%t\n", matched) // false
}
```

* https://yourbasic.org/golang/regexp-cheat-sheet/
* https://gobyexample.com/regular-expressions
