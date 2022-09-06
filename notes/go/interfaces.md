```go
// Stringer.go shows the use of Stringer interface.
package main

import "fmt"

// Pair is a sample custom type.
type Pair struct {
	x, y string
}

// String method implements the Stringer interface
// (https://golang.org/pkg/fmt/#Stringer)
func (p *Pair) String() string {
	return fmt.Sprintf("x: %s, y: %s", p.x, p.y)
}

func main() {
	p := &Pair{x: "XXX", y: "YYY"}
	fmt.Println(p) // String method used here
}
```

See https://research.swtch.com/interfaces for more.
