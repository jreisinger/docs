Go has pointers 👉. A pointer is a value that points to the memory address of another value.

```go
type Student struct{ Name string }

func main() {
	s := &Student{"John"}
	fmt.Printf("%p", s) // 0xc000092050
}
```

# How do they work

A variable is a convenient, alphanumeric pseudonym for a memory address. Memory address denotes a piece of storage that can contain a value.

```go
var x int32 = 10        // 4 bytes
var y bool = true       // 1 byte
pointerX := &x          // address of x
pointerY := &y
var pointerZ *string    // pointerZ == nil

Value    |  0 |  0 |  0 | 10 |  1 |  0 |  0 |  0 |  1 |  0 |  0 |  0 |  5 |  0 |  0 |  0 |  0 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Address  |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |  9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Variable | x                 | y  | pointerX          | pointerY          | pointerZ          |
```

* pointers are [really](https://go.dev/play/p/Js27uJKClEg) 8 bytes (not 4 as in the picture)
* the point is they are always the same size no matter what type they point to
* `nil` is an untyped identifier (in the universe block) that represents a lack of value for pointer types

# How to work with them

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```go
var p *int  // The * here means that p holds a pointer to int.
i := 42
p = &i      // The & operator generates a pointer to its operand.
*p = 21     // The * operator denotes the pointer's underlying value.
```

The operation on the last line is known as "dereferencing" or "indirecting".

Before dereferencing a pointer you must make sure it's not nil.

```go
var p *int
fmt.Println(p == nil) // true
fmt.Println(*p)       // panics
```

The built-in `new` function creates a pointer to zero value of the given type.

```go
var x = new(int)
fmt.Println(x == nil) // false
fmt.Println(*x)       // 0
```

`new` is rarely used because you can take address of a struct literal.

```go
x := &Foo{}
```

You can't use `&` before primitive literals (numbers, booleans and strings) or a constant
because they don’t have memory addresses; they exist only at compile time.

```go
type person struct {
    FirstName  string
    MiddleName *string
    LastName   string
}

p := person{
    FirstName:  "Pat",
    MiddleName: &"Perry", // This line won't compile
    LastName:   "Peterson",
}
```

Use a helper function to turn a constant value into a pointer.

```go
func stringp(s string) *string {
    return &s
}

p := person{
    FirstName:  "Pat",
    MiddleName: stringp("Perry"), // This works
    LastName:   "Peterson",
}
```

## Mutability

* mutability means changing data in place
* mutability brings flexibility and sometimes performance (you don't have to copy the data but the garbage collector might work more)
* immutable types are safer from bugs and easier to understand
* you should use pointers as last resort
* Go developers use pointers to indicate that a function parameter is mutable

## Method receivers

There are two reasons to use a pointer receiver:

1. so that the method can modify the value the receiver points to
2. avoid copying the value on each method call (this can be more efficient if the receiver is a large struct, for example)

Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

# Pointer and non-pointer types

Types implemented with pointers:

* slices
* maps
* functions
* channels
* interfaces

Non-pointer types:

* primitives (numbers, booleans and strings)
* structs
* arrays

# Sources and more

* https://tour.golang.org/moretypes/1
* https://yourbasic.org/golang/pointers-explained
* https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
