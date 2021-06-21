Go's slice type provides a convenient and efficient means of working with sequences of typed data.

Slice type is an abstraction built on top of array type.

# Arrays

* fixed size
* the size is part of the type: `[4]int` and `[5]int` are distinct types
* don't need to be initialized explicitly

```
var a [4]int // a[0] == 0, a[1] == 0, ...
```

* literal:

```
b := [2]string{"hello", "world"} // or [...]string{"hello", "world"}
```

# Slices

* literal:

```
b := []string{"hello", "world"}
```

* you need to initiliaze a slice

```
var s []byte
fmt.Println(s[0])   // panic: runtime error: index out of range [0] with length 0
s = make([]byte, 5) // s == []byte{0, 0, 0, 0, 0}
```

* you can also make a slice by slicing an existing slice or array

```
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l' , 'a'}, sharing the same storage as b

x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
```

# Slice internals

A slice is a descriptor of an array segment.

Slicing does not copy data. It creates a new slice that points to the original array.

```
s := make([]byte, 5)


[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr ---------------> |  x   |  x   |  x   |  x   |  x   |
|     |                | byte | byte | byte | byte | byte |
|-----|                +----------------------------------+
| len |                   0      1      2      3      4         
|  5  |                                                    
|-----|                                                    
| cap |                                                    
|  5  |                                                    
+-----+      
```

```
s = s[2:4]


[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr -----------------|------|----->|  x   |  x   |      |
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
s = s[:cap(s)]


[]byte                 [5]byte                             
+-----+                +----------------------------------+
| ptr -----------------|------|----->|  x   |  x   |  x   |
|     |                | byte | byte | byte | byte | byte |
|-----|                +----------------------------------+
| len |                    0      1      2      3     4       
|  3  |                                                    
|-----|                                                    
| cap |                                                    
|  3  |                                                    
+-----+      
```

# Source

https://blog.golang.org/slices-intro
