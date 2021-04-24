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

Types implemented with pointers:

* slices
* maps
* functions
* channels
* interfaces
