Reviewed: 2023-03-14

Go has pointers 👉. A pointer is a value denoting (or pointing to) the memory address of a variable.

```go
func main() {
	a := 200
	b := &a			// &a is a pointer	
	fmt.Printf("%v\n", a) 	// 200
	fmt.Printf("%v\n", b) 	// 0xc000018030
}
```

<img width="287" alt="image" src="https://user-images.githubusercontent.com/1047259/171110311-7456b542-da42-4a50-8e28-380ea01e7abc.png">

# How do they work

You can think of computer memory (RAM) as a sequence of boxes. Each box is labeled with a number. These numbers increment sequentially (1, 2, 3 ...). These numbers are called memory **addresses**. Memory address denotes a piece of storage that can contain a **value**. A **variable** is a convenient, alphanumeric pseudonym for a memory address.

```go
var x int32 = 10  // 4 bytes at address 1, holds value 10
var y bool = true // 1 byte at address 5, holds value 1
px := &x          // 4 bytes at address 6, holds value 1 (memory address of x)
py := &y          // 4 bytes at address 10, holds value 5 (memory address of y)
var pz *string    // 4 bytes at address 14, holds no value (nil)

Value    |  0 |  0 |  0 | 10 |  1 |  0 |  0 |  0 |  1 |  0 |  0 |  0 |  5 |  0 |  0 |  0 |  0 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Address  |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |  9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Variable | x                 | y  | px                | py                | pz                |
```

Pointers are [always](https://go.dev/play/p/t638QHuE21E) only a machine word ~~in size~~ wide (usually 32 or 64 bits) no matter what type they point to.

# How to work with them

The type `*T` is a pointer to a `T` value. Its zero value is `nil`. `nil` is an untyped identifier (in the universe block) that represents a lack of value for pointer types.

```go
var p *int  // The * here means that p holds a pointer to int.
i := 42
p = &i      // The & operator generates a pointer to its operand.
*p = 21     // The * operator denotes the pointer's underlying value.
```

The operation on the last line is known as "dereferencing" or "indirecting". Before dereferencing a pointer you must make sure it's not nil.

```go
var p *int
fmt.Println(p == nil) // true
fmt.Println(*p)       // panics
```

## `new` and `&`

The built-in `new` function creates a pointer to zero value of the given type.

```go
var x = new(int)
fmt.Println(x == nil) // false
fmt.Println(*x)       // 0
```

`new` is rarely used though because you can take address of a struct literal.

```go
x := &Foo{}
```

You can't use `&` before primitive literals (numbers, booleans and strings) or a constant.

```go
type person struct {
    firstName  string
    lastName   *string
}

p := person{
    firstName:  "Pat",
    lastName:   &"Peterson",  // This line won't compile
}
```

Use a helper function to turn a constant value into a pointer.

```go
func stringp(s string) *string {
    return &s
}

p := person{
    firstName: "Pat",
    lastName: stringp("Peterson"), // This works
}
```

## Reference and non-reference types

Types implemented with references: pointers, slices, maps, functions, channels, interfaces

Non-reference types: primitives (numbers, booleans and strings), structs, arrays

## Method receivers

There are two reasons to use a [pointer receiver](https://go.dev/tour/methods/4):

1. so that the method can modify the value the receiver points to
2. avoid copying the value on each method call (this can be more efficient if the receiver is a large struct, for example)

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

```go
type person struct {
	firstName string
	lastName  string
}

// If we didn't use pointer receiver (*person) this method would be ineffective.
// A value receiver (person) would operate on a copy of the original value.
func (p *person) rename(lastName string) {
	p.lastName = lastName
}

func (p *person) name() string {
	return p.firstName + " " + p.lastName
}
```

As a convenience, [Go interprets](https://go.dev/tour/methods/6) the statement `john.rename("Smith")` as `(&john).Rename("Smith")` since the rename method has a pointer receiver.

```go
func main() {
	john := person{firstName: "John", lastName: "Doe"}
	john.rename("Smith")
	fmt.Println(john.name())
}
```

## Notes on mutability

* mutability means changing data in place
* mutability brings flexibility and sometimes performance (you don't have to copy the data but the garbage collector might work more)
* immutable types are safer from bugs and easier to understand
* you should use pointers as last resort
* Go developers use pointers to indicate that a function parameter is mutable

# Sources and more

* https://tour.golang.org/moretypes/1
* https://yourbasic.org/golang/pointers-explained
* https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
