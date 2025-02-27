> Types are life. -- Bill Kennedy

Go is statically typed. It means that variables always have specific type and the type cannot change during the program run time.

Data types help us reason about what our program is doing and help us catch many errors.

The type of a variable or expression defines the characteristics of values it may take on such as

* their [size](https://go.dev/play/p/TstWfFXXntT)
* their internal representation (e.g. the type `int` is represented as 32 bits)
* the instrinsic [operations](https://tour.golang.org/methods/16) that can be performed on them
* the methods associated with them

Go's types categories

* [basic types](https://play.golang.org/p/z5uVUJsKxBw): numbers (integers, floats, complex), strings, booleans
* aggregate types: arrays, structs (user-defined types)
* [reference types](https://play.golang.org/p/NZ0VhQ_pwYR): pointers, slices, maps, functions, channels
* interface types

More

* [Basic types](https://tour.golang.org/basics/11)
* [Zero values](https://tour.golang.org/basics/12)
* [Type conversions](https://tour.golang.org/basics/13)
* [Type inference](https://tour.golang.org/basics/14).
* https://dave.cheney.net/2014/03/25/the-empty-struct
