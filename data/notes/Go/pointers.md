```go
var x int32 = 10  // 0000 1010
var y bool = true // 0000 0001
pointerX := &x
pointerY := &y
var pointerZ *string // pointerZ == nil

Value    |  0 |  0 |  0 | 10 |  1 |  0 |  0 |  0 |  1 |  0 |  0 |  0 |  5 |  0 |  0 |  0 |  0 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Address  |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 | 19 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 |
---------|----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+----+
Variable | x                 | y  | pointerX          | pointerY          | pointerZ          |
```

* pointers are [really](https://play.golang.org/p/kgslsZufai2) 8 bytes
* the point is they are always the same size no matter what type they point to
* `nil` is an untyped identifier that represents a lack of value

Types implemented with pointers:

* slices
* maps
* functions
* channels
* interfaces
