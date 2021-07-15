Golang has a unique approach to object oriented programming (OOP). I am not sure if it can even be called OOP.

# Structs

`struct` is a type that contains named fields.

```
// structs.go
package main

import (
    "fmt"
    "math"
)

// Let's define a new type called Circle that's based on struct type.
type Circle struct {
    x, y, r float64
}

// circleArea is a function using the Circle type.
func circleArea(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

func main() {
    // Initialize a variable of type Circle.
    // (usually used with & so var can by changed by methods)
    c := Circle{1, 2, 3}

    // Access a struct field.
    c.r = 5

    // Use the struct as a function argument.
    fmt.Println(circleArea(&c))
}
```

# Methods

```
// methods.go
package main

import (
    "fmt"
)

type Person struct {
    Name string
}

// Method in Go is just a function with a receiver (old OO lingo).
func (p *Person) talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person // embedded type (or anonymous field)
    Model  string
}

func main() {
    // Object instantiation - way 1
    p := &Person{Name: "John"}

    // Object instantiation - way 2
    a := new(Android) // implicit & (pointer)
    a.Name = "R2D2"

    // A method call
    p.talk()
    a.talk()
}
```

# Interfaces

Create a few small structs that have what you want, add in methods that you need, and as you build your program, useful interfaces will tend to emerge. There's no need to have them all figured out ahead of time.

```
// shuffler.go
package main

import (
    "fmt"
    "math/rand"
)

//// Interface and its function.

// Interface defines a method set (functions' prototypes). That's the list of
// methods a type must have in order to implement the interface.
type shuffler interface {
    Len() int
    Swap(i, j int)
}

// Interface type is used as an argument to a fuction here.
func shuffle(s shuffler) {
    for i := 0; i < s.Len(); i++ {
        j := rand.Intn(s.Len() - 1) // why not only s.Len()?
        s.Swap(i, j)
    }
}

//// Type intSlice satisfies shuffler interface.

type intSlice []int

func (is intSlice) Len() int {
    return len(is)
}

func (is intSlice) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

//// Main.

func main() {
    is := intSlice{1, 2, 3, 4, 5, 6}
    shuffle(is)
    fmt.Println(is) // [2 4 1 6 3 5]
}
```

# Sources

* Caleb Doxsey: Introducing Go (2016)
* John Graham-Cumming: Go Programming Basics (2017)
