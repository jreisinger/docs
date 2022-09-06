To run test file selectively put this at the top of a `foo_test.go`:

```
// +build manual

package foo
```

Then run it with:

```
go test ./... -tags manual
```
