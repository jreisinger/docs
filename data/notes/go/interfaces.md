```
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

bytecounter.go:

```
// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
// https://github.com/adonovan/gopl.io/blob/master/ch7/bytecounter/main.go
package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
```
