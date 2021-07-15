A [struct](https://play.golang.org/p/Av0NOh_cu_K) is a user-defined type that contains named fields:

```go
// Declaration of anonymous (literal) struct type + initialization.
e := struct {
	flag    bool
	counter int
}{
	flag: true,
	counter: 10,
}
```