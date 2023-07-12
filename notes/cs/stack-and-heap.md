---
title: Stack and heap
date: 1970-01-01
tags: cs notes
---

Stack is a simple data structure that allows for two operations:

* push - add an item to the top of the stack
* pop - remove an item from the top of the stack

Functions form *call stacks*:

```go
func main() {
    fmt.Println(f1())
}

func f1() int {
    return f2()
}

func f2() int {
    return 1
}
```

```
                    +----+
                    | f2 | return
                    +----+
          +----+    +----+        +----+
          | f1 | f2 | f1 |        | f1 | return
          +----+    +----+        +----+
+----+    +----+    +----+        +----+        +----+
|main| f1 |main|    |main|        |main|        |main|
+----+    +----+    +----+        +----+        +----+
```

* (`Println` is also a function but we ignore it for simplicity)
* each time we *call* a function, we *push* it onto the call stack
* each time we *return* from a function, we *pop* it off of the stack
* each function gets some memory allocated on the stack - a stack frame
* if too much memory is consumed (e.g. when making too many recursive calls) you get *stack overflow* error
* when you call a fuction from another function, the calling function is paused in a partially completed state
Both are memory regions.

Stack (temporary to a function)

* stores temporary variables created by functions
* LIFO data structure with push/pop operations
* all vars are popped off when function exits
* very fast
* limited in size

Heap (global)

* not as tightly managed by CPU
* in C you have to manage it yourself via `malloc()`, `calloc()` or `realloc()`
* if you fail you get a memory leak
* slower access because pointers are used
* limited only by physical memory

Go example:

```go
// x.go
package main

import (
    "fmt"
    "runtime"
    "time"
)

var global *int     // "global" is the name :-)

func f() {
    var x int       // heap-allocated because ...
    x = 1 
    global = &x     // ... escapes from f()
}

func g() {
    y := new(int)   // allocated on the stack
    *y = 1 
}

func main() {
    start := time.Now()
    for {
        // Run once per second
        if time.Since(start) > time.Second {
            var r runtime.MemStats
            runtime.ReadMemStats(&r)
            fmt.Printf("Heap size %d\n", r.HeapAlloc)
            fmt.Printf("Stack size %d\n", r.StackInuse)
            fmt.Printf("NumGC %d\n", r.NumGC)
            start = time.Now()
        }   

        f() 
        g()
    }   
}
```

```sh
go run -gcflags="-m" x.go
```

More

* https://youtu.be/ZMZpH4yT7M0
* https://youtu.be/sTFJtxJXkaY
* [Garbage collection and garbage reduction](https://www.safaribooksonline.com/videos/intermediate-go-programming/9781491944073/9781491944073-video234746)
* https://www.gopl.io/, ch. 2.3
