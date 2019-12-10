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
