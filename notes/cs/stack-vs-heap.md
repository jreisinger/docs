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

* https://youtu.be/sTFJtxJXkaY
* [Garbage collection and garbage reduction](https://www.safaribooksonline.com/videos/intermediate-go-programming/9781491944073/9781491944073-video234746)
* https://www.gopl.io/, ch. 2.3
