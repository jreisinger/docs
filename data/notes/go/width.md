* the term comes from the gc compiler
* number of bytes of storage an instance of a type occupies
* every value in a Go [program](https://go.dev/play/p/TstWfFXXntT) has a type

```go

func main() {
	var b bool
	var i int
	var s string
	var a [3]string
	var S struct {
		a uint16
		b uint32
	}
	var S2 struct {
		a struct{}
		b struct{}
	}

	fmt.Println(unsafe.Sizeof(b))  // 1
	fmt.Println(unsafe.Sizeof(i))  // 8
	fmt.Println(unsafe.Sizeof(s))  // 16
	fmt.Println(unsafe.Sizeof(a))  // 3x16
	fmt.Println(unsafe.Sizeof(S))  // 2+4+<padding>
	fmt.Println(unsafe.Sizeof(S2)) // 0
}
```

Source: https://dave.cheney.net/2014/03/25/the-empty-struct
