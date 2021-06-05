> Types are life. -- Bill Kennedy

Go is statically typed - variables always have specific type and type cannot change during the program run time.

[Types](https://play.golang.org/p/24H4L7Gofrz) help us reason about what our program is doing and help us catch many errors.

Types classify values (data) into groups and determine:

* what is the memory size allocated for the value
* what does the value represent (e.g. byte `11111111` represents decimal number 255 if the type is `int`)
* what are intrinsic operations of that representation (e.g. arithmetic operations for numbers, indexing, `append` and `range` for slices)

Types of Go types:

* [basic](https://tour.golang.org/basics/11): `bool`, `string`, numeric types, `byte`, `rune`
* reference types: slice, map, pointer, `chan`, `func`
* built-in types: all the above
* user-defined types: structs

See also [Zero values](https://tour.golang.org/basics/12), [Type conversions](https://tour.golang.org/basics/13) and [Type inference](https://tour.golang.org/basics/14).
