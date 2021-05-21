```go
func describe(s []int) {
	fmt.Printf("Size\t%d\nType\t%s\nValue\t%v\nLen\t%d\nCap\t%d\n---\n",
		unsafe.Sizeof(s), reflect.TypeOf(s), s, len(s), cap(s))
}
```
