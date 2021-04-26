# How do they work

```go
var x int32 = 10 // 0000 1010
var y bool = true
pointerX := &x // address of x
pointerY := &y
var pointerZ *string // pointerZ == nil

Value    |  0 |  0 |  0 | 10 |  1 |  0 |  0 |  0 |  1 |  0 |  0 |  0 |  5 |  0 |  0 |  0 |  0 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Address  |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 | 19 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Variable | x                 | y  | pointerX          | pointerY          | pointerZ          |
```

* pointers are [really](https://play.golang.org/p/3Lz_C_sXaHv) 8 bytes
* the point is they are always the same size no matter what type they point to
* `nil` is an untyped identifier (in the universe block) that represents a lack of value

`&` - the address operator, `*` - the indirection operator or denoting a pointer type

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

`new` is rarely used because you can take address of a struct literal. You can't use `&` before primitive literals (numbers, booleans and strings) or a constant because they don’t have memory addresses; they exist only at compile time.

```go
x := &Foo{}

var y string
z := &y
```

Use a helper function to turn a constant value into a pointer.

```go
type person struct {
    FirstName  string
    MiddleName *string
    LastName   string
}

p := person{
    FirstName:  "Pat",
    MiddleName: "Perry", // This line won't compile
    LastName:   "Peterson",
}

func stringp(s string) *string {
    return &s
}

p := person{
    FirstName:  "Pat",
    MiddleName: stringp("Perry"), // This works
    LastName:   "Peterson",
}
```

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

# Mutability

* mutability means changing data in place
* immutable types are safer from bugs and easier to understand
* mutability brings flexibility and sometimes performance (you don't have to copy the data but the garbage collector might work more)
* Go developers use pointers to indicate that a function parameter is mutable
* you should use pointers as last resort

# How to use them

## Populating structs

Rather than populating a struct by passing a pointer to it into a function, have the function instantiate and return the struct.

Don't do this:

```go
func MakePerson(p *Person) error {
    p.Name = "John"
    p.Age = 41
    return nil
}
```

Do this:

```go
func MakePerson() (Person, error) {
    p := Person{
        Name: "John",
        Age: 41,
    }
    return p, nil
}
```

One exception is when a function expects an interface:

```go
f := struct {
	Name string
	Age  int
}{}

err := json.Unmarshal([]byte(`{"Name": "John", "Age": 41}`), &f)
```

## Method receiver

There are two reasons to use a pointer receiver:

1. so that the method can modify the value the receiver points to
2. avoid copying the value on each method call (this can be more efficient if the receiver is a large struct, for example)

Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

# Sources

* Learning Go, 2021
* https://tour.golang.org/methods/4
