* the term comes from the gc compiler
* number of bytes of storage an instance of a type occupies
* every value in a Go [program](https://go.dev/play/p/i_F_ZLnRzV1) has a type

```go
var b bool
var i int
var s string
var a [3]string
var S struct {
    a uint16
    b uint32
}

fmt.Println(unsafe.Sizeof(b)) // 1
fmt.Println(unsafe.Sizeof(i)) // 8
fmt.Println(unsafe.Sizeof(s)) // 16
fmt.Println(unsafe.Sizeof(a)) // 3x16
fmt.Println(unsafe.Sizeof(S)) // 2+4+<padding>
```

Source: https://dave.cheney.net/2014/03/25/the-empty-struct
