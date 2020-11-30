## Gotcha

This won't give you a slice

```
names := [4]string{"John", "Paul", "George", "Ringo"}
a := names[0] // a = John -> a string
```

this will

```
names := [4]string{"John", "Paul", "George", "Ringo"}
a := names[0:1] // a = [John] -> a slice of strings
```

## Operations on a slice, length and capacity

* length - the number of elements it contains
* capacity - the number of elements in the underlying array, counting from the first element in the slice

```
package main

import (
	"fmt"
)

func main() {
	s := make([]string, 2, 3)
	printSlice(s) // 0th_elem_addr=0xc000064180 len=2 cap=3 [ ]

	s[0] = "the"
	s[1] = "quick"
	printSlice(s) // 0th_elem_addr=0xc000064180 len=2 cap=3 [the quick]

	s = append(s, "brown")
	printSlice(s) // 0th_elem_addr=0xc000064180 len=3 cap=3 [the quick brown]

	s = append(s, "fox")
	printSlice(s) // 0th_elem_addr=0xc00004a0c0 len=4 cap=6 [the quick brown fox]

	// Slice the slice to give it zero lenth.
	s = s[:0]
	printSlice(s) // 0th_elem_addr=0xc00004a0c0 len=0 cap=6 []

	// Extend its length.
	s = s[:3]
	printSlice(s) // 0th_elem_addr=0xc00004a0c0 len=2 cap=6 [the quick brown]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // 0th_elem_addr=0xc0000521a0 len=1 cap=4 [brown]
}

// printSlice prints a slice of strings with some useful info.
func printSlice(s []string) {
	fmt.Printf("0th_elem_addr=%p len=%d cap=%d %v\n", s, len(s), cap(s), s)
}
```

More

* https://tour.golang.org/moretypes/7
