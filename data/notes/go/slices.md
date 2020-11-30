This won't give you a slice

```
names := [4]string{"John", "Paul", "George", "Ringo"}
a := names[0] // a = John (string)
```

this will

```
names := [4]string{"John", "Paul", "George", "Ringo"}
a := names[0:1] // a = [John] ([]string)
```

## Slices and `append`

```
// Create a slice add and then append some elements.
package main

import (
    "fmt"
)

// printSlice prints a slice of bytes with some useful info.
func printSlice(s []string) {
    fmt.Printf("0th_elem_addr=%p len=%d  cap=%d  %v\n",
        s, len(s), cap(s), s)
}

func main() {
    s := make([]string, 2, 3)
    printSlice(s) // 0th_elem_addr=0xc000064180 len=2  cap=3  [ ]
    s[0] = "the"
    s[1] = "quick"
    printSlice(s) // 0th_elem_addr=0xc000064180 len=2  cap=3  [the quick]
    s = append(s, "brown")
    printSlice(s) // 0th_elem_addr=0xc000064180 len=3  cap=3  [the quick brown]
    s = append(s, "fox")
    printSlice(s) // 0th_elem_addr=0xc00004a0c0 len=4  cap=6  [the quick brown fox]
}
```

More

* https://tour.golang.org/moretypes/7
