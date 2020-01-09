# Types

Go is statically typed - variables always have specific type and type cannot
change during the program run time.

Types help us reason about what our program is doing and help us catch many
errors.

Types are similar to sets in mathematics. They classify data into groups and
determine:

* characteristics of data (e.g. all strings have length)
* operations that can be performed on data (e.g. `len("a string")`)
* data size (e.g. `int8`)
* how data is stored in memory

See also [Basic types](https://tour.golang.org/basics/11), [Zero values](https://tour.golang.org/basics/12) and [Type conversions](https://tour.golang.org/basics/13).

## Numbers

Computers use base-2 binary system to store and work with the numbers. So
computers count like this: 0, 1, 10, 11, 100, 101, 110, 111, ...

Integer types

* [u]int{8, 16, 32, 64}
* machine dependent: uint, **int**, uintptr
* byte - alias for uint8
* rune - alias for int32 (represents a Unicode code point)

Floating-point types

* float32 (single precision), **float64** (double precision)
* complex64, complex128
* contain decimal component (i.e. real numbers)
* their actual representation on computer is quite complicated but not needed to
    know to use them
* inexact (1.01 – 0.99 using floating-point arithmetic results in
    0.020000000000000018)
* have certain size (32 or 64 bit)
* NaN - not a number (for things like 0/0), +∞, -∞

```go
// we use .0 to tell Go it's a floating-point number
fmt.Prinln("1 + 1 =", 1.0 + 1.0)
```

## Strings

* sequences of characters used to represent text
* made up of individual bytes, usually (but not always) one for each character

String literals are created with:

* double quotes (`"Hello world"`) - cannot contain newlines and allow escape
    sequeences (e.g. `\t`, `\n`)
* backticks (`` `Hello world` ``)

Common operations on strings:

* find length: `len("Hello world")`
* access a character: `"Hello world"[1]` -> 101 instead of e because the
    character is represented by a byte (i.e. and integer)
* concatenate strings: `"Hello " + "world"` -> Go figures out what to do based
    on the type of the arguments

## Booleans

* special 1-bit integer type used to represent true and false (or on and off)
* logical operators: `&&`, `||`, `!`
* truth tables define how these operators work

# Variables

* variable - storage location, with a specific type and an associated name
* [scope](https://golang.org/ref/spec#Declarations_and_scope) - the range of places where you are allowed to use a variable ("Go is lexically scoped using block.")
* constants - variables whose values cannot be changed during program run time

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Print("Enter distance in feet: ")
    var feet float64 // one way to define var
    fmt.Scanf("%f", &feet)
    meters := feet * 0.3048 // another way to define var
    fmt.Printf("%.2f ft = %.2f m\n", feet, meters)
}
```

See also [Type inference](https://tour.golang.org/basics/14).

# Control Structures

## The for Statement

Other programming languages have various types of loops (while, until, foreach,
...). Go only has for loop that can be used in various ways, e.g.:

```go
// traditional c-style
for i := 1; i <= 10; i++ {
    fmt.Println(i)
}

i := 1              // declaration + initialization
for i <= 10 {       // condition
    fmt.Println(i)
    i++             // increment
}

// loop over array/slice
for i, value := range x {
    ...
}
```

## The if and switch Statements

If the if statement becomes too verbose use the switch statement.

```go
for i := 1; i <= 10; i++ {
    switch i { // expression after switch can be omitted
    case 1: fmt.Println("one")
    case 5: fmt.Println("five")
    case 6: fmt.Println("six")
    case 10: fmt.Println("ten")
    default: fmt.Println(i) // similar to else
    }
}
```

The value of the expressions (in this example `i`) is compared to the
expression following each `case` keyword. If they are equivalent the statements
following `:` are executed. The first one to succeed is chosen.

# More built-in types

## Arrays

Array is a numbered sequence of elements of a single type with a *fixed length*.

```go
var a1 [3]int // array of three integers
a1[0] = 10
a1[1] = 20
a1[2] = 30

// shorter syntax for creating arrays
a2 := [3]int{ 10, 20, 30, }
```

Now, you rarely see arrays used directly in Go code :-).

## Slices

Slice is a segment of an array. Like arrays, they are indexable and have a length. Unlike arrays, the *length is allowed to change*.

```go
    var s1 []float64            // []
    s2 := make([]float64, 3)    // [0 0 0]

    // define length (3) and capacity (5)
    s3 := make([]float64, 3, 5) // [0 0 0]

    // create slice from array
    a := [5]float64{1,2,3,4,5}
    s4 := a[1:3]                // [2 3]
    s5 := a[:]                  // [1 2 3 4 5]
```

Slices are always associated with some array. The are like
[references](https://tour.golang.org/moretypes/8) to arrays.

See also [Slice literals](https://tour.golang.org/moretypes/9).

### `append` operator

```go
s1 := []int{1,2,3}
s2 := append(s1, 4, 5)
```

* adds elements onto the end of a slice and creates a *new slice*
* if there's sufficient capacity, the backing array's length is incremented
* if not, a new backing array is created and all the existing elements are copied over

See also [Appending to a slice](https://tour.golang.org/moretypes/15).

### `copy` operator

```go
s1 := []int{1,2,3}
s2 := make([]int, 2)
// func copy(dst, src []Type) int
copy(s2, s1)
// s1: [1,2,3], s2: [1,2]
```

* all of the entries in `src` are copied into `dst` overwriting whatever is there
* if lengths are not the same, the smaller of the two will be used

See also [copy](https://golang.org/pkg/builtin/#copy).

## Maps

* unordered collection of key-value pairs (also called associative arrays, hash tables, or dictionaries)

```go
// x is a map of strings into ints

// WRONG: this will yield a run time error
var x map[string]int
x["key"] = 10 // panic: assignment to entry in nil map

// maps have to be initialized before they can be used
var x = make(map[string]int)
x["key"] = 10

// delete an item from a map
delete(x, "key")
```

* maps are often used as lookup tables (dictionaries)

```go
elements := map[string]string{
    "H":  "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
}

if name, ok := elements["He"]; ok {
    fmt.Printf("He is %s\n", name)
}
```

# Functions

A function (aka a procedure, or a subroutine) is an independent section of code that maps zero or more input parameters to zero or more output parameters:

```
Inputs -> [ func f(i, j int) int {} ] -> Outputs
```

* collectively, the parameters (i, j) and the return type (int) are called function's signature

Functions form call stacks:

```go
func main() {
    fmt.Println(f1())
}

func f1() int {
    return f2()
}

func f2() int {
    return 1
}
```

```
                    +----+
                    | f2 | return
                    +----+
          +----+    +----+        +----+
          | f1 | f2 | f1 |        | f1 | return
          +----+    +----+        +----+
+----+    +----+    +----+        +----+        +----+
|main| f1 |main|    |main|        |main|        |main|
+----+    +----+    +----+        +----+        +----+
```

* each time we call a function, we push it onto the call stack
* each time we return from a function, we pop it off of the stack

Return types can have names:

```go
func f2() (r int) {
    r := 1
    return
}
```

Multiple values can be returned:

```go
func f() (int, int) {
    return 4, 2
}

func main() {
    x, y := f()
}
```

Multiple values are often used to return an error value along with the result (`x, err := f()`), or a boolean to indicate success (`x, ok := f()`).

## Variadic functions

There is a special form available for the last parameter:

```go
// sump up zero or more integers
func sum(args ...int) int {     // prefix ellipsis
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(sum())
    fmt.Println(sum(1, 2))

    xs := []int{1,2,3}
    fmt.Println(sum(xs...))     // suffix ellipsis
}
```

fmt.Println can take any number of values (...) of any type (interface{}):

```go
func Println(a ...interface{}) (n int, err error)
```

## Closures

It's possible to create functions inside functions. These local functions have
access to other local variables:

```go
func main() {
    // local variable accessible by increment
    x := 0

    // local variable of type func() int
    increment := func() int {
        x++
        return x
    }

    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}
```

* a function like this together with nonlocal variables it references is known
  as closure

One way to use closure is to write a function that returns another function:

```go
func makeEvenGenerator() func() uint {
    i := uint(0) // unlike normal local variable this one persists between calls
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}
```

## Recursion

A function is able to call itself:

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
```

factorial(2):

1. Is x == 0? No (x is 2).
2. Find the factorial of x - 1.
    1. Is x == 0? No (x is 1).
    2. Find the factorial of x - 1.
        1. Is x == 0? Yes, return 1.
    3. Return 1 * 1.
3. Return 2 * 1.

## defer

defer schedules a function call to be run before a function returns. It's often used when resources need to be freed in some way, e.g.:

```go
func main() {
    f, _ := os.Open(filename)
    defer f.Close()
}
```

This has three advantages:

* you keep the closing call close to the opening call
* if a function had multiple return calls (like within an if statement) defer would call Close before any of them
* deferred functions run even if a runtime panic occurs

## panic and recover

* `panic` causes a runtime error immediately stopping the function's execution
* `recover` stops the panic and returns the value that was passed to `panic`

WRONG:

```go
func main() {
    panic("PANIC")
    str := recover() // this will never happen!
    fmt.Println(str)
}
```

CORRECT:

```go
func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
```

A panic generally indicates a programmer's error or an exceptional condition that's not easy to recover from.

See https://blog.golang.org/defer-panic-and-recover for more.

## Pointers

Normally a function's argument is copied:

```go
func zero(x int) {
    x = 0
}

func main() {
    x := 1
    zero(x)
    // x is still 1
}
```

If we want to modify the original argument one way to do it is to use a special data type known as a pointer:

```go
func zero(xPtr *int) {
    *xPtr = 0
}

func main() {
    x := 1
    zero(&x)
    // x is 0
}
```

Pointers reference a *location* in memory where a value is stored rather than the *value* itself.

* `*` represents a pointer, e.g. `*int` means a pointer to an integer value
* `*` is also used to dereference a pointer variable (`*xPtr`), i.e. to get to the value a pointer points to
* `&` operator finds the memory location (address) of a variable

Another way to get a pointer is to use the `new` built-in function:

```go
xPtr := new(int)
```

* `new` takes a type as an argument, allocates enough memory to fit a value of that type, and returns a pointer to it

Go is a garbage-collected language. It means memory is cleaned up automatically when nothing refers to it anymore.

Pointers are rarely used with Go's built-in types but are extremely useful when paired with structs.

# Structs and interfaces

At some point it would become tedious and error prone to write programs using only Go's built-in types.

## Structs

A struct is a type that contains named fields:

```go
// Circle represents, well, a circle.
type Circle struct {
    x, y, r float64
}

// Several ways to do initialization.
var c Circle
c := Circle{x: 0, y: 0, r: 5}
c := Circle{0, 0, 5}
c := new(Circle) // returns pointer
c := &Circle{0, 0, 5} // most typical

// Accessing fields.
fmt.Println(c.x)
c.r = 10
```

## Methods

Using structs with functions:

A normal function:

```go
func circleArea(c *Circle) float64 {
    // no dereferencing needed... ahaa, that's r not c!
    return math.Pi * c.r*c.r
}

c := &Circle{0, 0, 5}
fmt.Println(circleArea(c))
```

A special function - method:

```go
func (c *Circle) area() float64 { // (c *Circle) is called a receiver
    return math.Pi * c.r*c.r
}

c := Circle{0, 0, 5}  // no & needed here as Go automatically
fmt.Println(c.area()) // knows to pass a pointer to the Circle
```

## Embedded types

A struct's fields usually represent the *has-a* relationship, e.g. `Person` has a name:

```go
type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
```

We use embedded types (anonymous fields) to represent the *is-a* relationship, e.g. `Android` is a person (so it can `Talk()`):

```go
type Android struct {
    Person // embedded type
    Model string
}

a := new(Android) // you cannot do &Android{Name: "R2D2"} here
a.Name = "R2D2"
a.Talk() // could be also: a.Person.Talk()
```

## Interfaces

Interfaces are similar to structs but instead of fields they have a method set. A method set is a list of methods that a type must have in order to *implement* the interface. We can use interface types as arguments to functions. Interfaces define behaviour instead of defining types.

```go
package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int      // Any type that has this method set satisfies
	Swap(i, j int) // the shuffler interface, i.e. is a shuffler.
}

func shuffle(s shuffler) { // interface (not a concrete type) used as argument
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
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

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// Note that we can use shuffle with *any* type that satisfies the shuffler
	// interfaces. This is pretty close to duck typing and provides for a lot
	// of flexibility.

	is := intSlice{1, 2, 3, 4, 5}
	shuffle(is)

	s := stringSlice{"The", "Quick", "Brown", "Fox"}
	shuffle(s)

	fmt.Println(is, s)
}
```

Code taken from John Graham-Cumming: [Interfaces](https://learning.oreilly.com/learning-paths/learning-path-go/9781491990384/9781491913871-video191862).

# Standard Packages

* Go was designed to encourage good software engineering practices
* one of them is code reuse (DRY - Don't Repeat Yourself): functions, packages

## Input/Output

### io

`io` package consists of a few functions, but mostly interfaces used in other packages. The two main interfaces are `Reader` and `Writer`. `Reader`s support reading via the `Read` method. `Writer`s support writing via the `Write` method. Many functions in Go take Readers or Writers as arguments. E.g. the io.Copy function copies data from a Reader to a Writer:

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

### bufio and ioutil

See notes/go/io.

### bytes strings

To read/write a `[]byte` or a `string`, you can use the `Buffer` type (struct) from `bytes` package:

```go
var b bytes.Buffer
b.Write([]byte("Hello "))
fmt.Fprintf(&b, "world!\n")
b.WriteTo(os.Stdout)
```

* a `Buffer` doesn't have to be initialized
* it supports both the `Reader` and `Writer` interfaces
* you can convert it into a `[]byte` by calling `buf.Bytes()`

If you only need to read from a string, you can use the more efficient `strings.NewReader` function.

## Files and Folders

The easiest way to read the entire file into memory:

```go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("/etc/passwd")
    if err != nil {
        // handle error here
        return
    }
    str := string(bs)
    fmt.Println(str)
}
```

Create a file:

```go
package main

import (
    "os"
)

func main() {
    file, err := os.Create("test.txt") // file is os.File
    if err != nil {
        // handle error here
        return
    }
    defer file.Close()
    file.WriteString("test")
}
```

Get contents of a directory:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    dir, err := os.Open(".")
    if err != nil {
        // handle error here
        return
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1) // -1 means return all entries
    if err != nil {
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}
```

Recursively walk a folder (read the folder’s contents, all the subfolders, all the sub-subfolders, etc.):

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // func() is called for every file and folder in "."
    filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
    })
}
```

## errors

Go has a built-in type for errors (the `error` type). We can also create our own errors:

```go
package main

import "errors"

func main() {
    err := errors.New("error message")
}
```

## sort

* `sort` package contains functions for sorting arbitrary data
* there are several predefined sorting functions, e.g. for [slices of ints](https://golang.org/pkg/sort/#Ints) or floats
* you can also [sort](https://github.com/jreisinger/go/blob/master/sort.go) your own data

# Sources

* Caleb Doxsey: Introducing Go (O'Reilly, 2016)
* John Graham-Cumming: [Go Programming Basics](https://learning.oreilly.com/learning-paths/learning-path-go/9781491990384/)
* [A Tour of Go](https://tour.golang.org)
