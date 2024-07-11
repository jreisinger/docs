Reviewed: 2023-04-07

```go
package main

import (
	"fmt"
	"regexp"
)

var (
	ws   = regexp.MustCompile(`\s+`)
	num  = regexp.MustCompile(`[0-9]+`)
	perc = regexp.MustCompile(`([0-9]+)(%)`) // capture groups
)

func main() {
	s := "battery charged to 98%"

	words := ws.Split(s, -1)
	fmt.Printf("%d\n", len(words)) // 4

	matched := num.MatchString(s)
	fmt.Printf("%t\n", matched) // true

	match := num.FindString(s)
	fmt.Printf("%s\n", match) // 98

	matches := perc.FindStringSubmatch(s)
	fmt.Printf("%v\n", matches) // [98% 98 %]
}
```

* https://yourbasic.org/golang/regexp-cheat-sheet/
* https://gobyexample.com/regular-expressions
