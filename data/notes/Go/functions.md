Arguments are passed by [value](https://play.golang.org/p/LoFTsgS3BUQ). A goroutine can access only data within its active frame (a slice of the stack) so the argument needs to be copied upon a function call. This is a mechanism for isolating memory mutations. It also means that if you need to mutate data you need to use a pointer.

Prefer functions over methods because functions are more precise.

Use methods for state or decoupling.