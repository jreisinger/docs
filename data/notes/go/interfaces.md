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
