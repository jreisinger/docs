Go's slice type provides a convenient and efficient means of working with sequences of typed data.

Slice type is an abstraction built on top of array type.

# Arrays

* fixed-sized composite value
* the size is part of the type: `[4]int` and `[5]int` are distinct types
* value, not a pointer to the first array element like in C
* like a struct but with indexed rather than named fields

Don't need to be initialized explicitly:

```
var a [3]int // a[0] == 0, a[1] == 0, a[2] == 0
```

Literal:

```
b := [2]string{"hello", "world"} // or [...]string{"hello", "world"}
```

# Slices

A slice does not store any data, it just describes a section of an underlying array.

It's zero value is `nil`:

```
var s []byte
fmt.Println(s == nil)          // true
fmt.Println(s, len(s), cap(s)) // [] 0 0
fmt.Println(s[0]) // panic: runtime error: index out of range [0] with length 0
```

You initiliaze a slice by using `make`:

```
s = make([]byte, 5) // s == []byte{0, 0, 0, 0, 0}
fmt.Println(s[0])   // 0
```

or a literal:

```
b := []string{"hello", "world"}
```

You can also make a slice by slicing an existing slice or array. Slicing is done by specifying a half-open range with two indices like `b[1:4]` or `x[0:len(x)]`:

```
s := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
t := s[1:4] // []byte{'o', 'l' , 'a'}, sharing the same storage as s

a := [3]string{"Лайка", "Белка", "Стрелка"}
s := a[:] // s references the storage of a
```

# Slice internals

A slice is a descriptor of an array segment.

```
s := make([]byte, 5) // allocate array of size 5 and return slice pointing to it

[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr ---------------> | xxxx | xxxx | xxxx | xxxx | xxxx |
|     |                | byte | byte | byte | byte | byte |
|-----|                +----------------------------------+
| len |                   0      1      2      3      4         
|  5  |                                                    
|-----|                                                    
| cap |                                                    
|  5  |                                                    
+-----+      
```

Slicing does not copy data. It creates a new slice that points to the original array.

```
s = s[2:4] // slice s to length shorter than its capacity

[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr -----------------|------|----->| xxxx | xxxx |      |
|     |                | byte | byte | byte | byte | byte |
|-----|                +----------------------------------+
| len |                    0      1      2      3     4       
|  2  |                                                    
|-----|                                                    
| cap |                                                    
|  3  |                                                    
+-----+      
```

```
s = s[:cap(s)] // grow s to its capacity

[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr -----------------|------|----->| xxxx | xxxx | xxxx |
|     |                | byte | byte | byte | byte | byte |
|-----|                +----------------------------------+
| len |                    0      1      2      3     4       
|  3  |                                                    
|-----|                                                    
| cap |                                                    
|  3  |                                                    
+-----+      
```

# Growing slices

To increase the capacity of a slice one must create a new, larger slice and `copy` the contents of the original slice into it.

```
// Double the capacity of s.
t := make([]string, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
copy(t, s)                                // copy the contents of s into t
s = t                                     // assign the slice value t to s
```

The `append` function appends the elements x to the end of slice s, and grows the slice if a greater capacity is needed.

```
func append(s []T, x ...T) []T
```

```
a := make([]int, 1)    // a == []int{0}
a = append(a, 1, 2, 3) // a == []int{0, 1, 2, 3} 
```

To append one slice to another, use `...` to expand the second argument to a list of arguments:

```
a := []string{"Polkinghorne", "Lennox"}
b := []string{"Brooks", "Knuth", "Wall"}
a = append(a, b...) // equivalent to append(a, b[0], b[1], b[2])
```

# Source

* https://go.dev/tour/moretypes/7
* https://go.dev/blog/slices-intro
* https://go.dev/blog/slices
