2023-03-13

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
* each function gets some memory allocated on the stack
* if too much memory is consumed (e.g. when making too many recursive calls) you get *stack overflow* error
* when you call a fuction from another function, the calling function is paused in a partially completed state
