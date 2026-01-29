Q: Why is `log.LstdFlags` an int?

A: Because it's part of a bit flag system also called a bitmask pattern:

Q: What is a bitmask pattern?

A:

```go
conts (
	Ldate 				= 1 << iota     // 0001 (1)
	Ltime								// 0010 (2) 
	Lmicrosecond						// 0100 (4)
	Llongfile							// 1000 (8)
	// ...
	LstdFlags			= Ldate | Ltime // 0011 (3)
)
```

`<<` is the bitwise left shift operator: shifts bits to the left, discarding the leftmost bits and appending zeros to the right:

```
1 << 0 // 0001 -> 0001 = 1
1 << 1 // 0001 -> 0010 = 2
1 << 2 // 0001 -> 0100 = 4
1 << 3 // 0001 -> 1000 = 8
```

`|` is the bitwise OR operator: if either bit is 1, the result is 1:

```
    0001 (Ldate =  1)
|   0010 (Ltime =  2)
--------
    0011 (Result = 3)
```
