# User-defined types

## Python class vs Go struct

```python
class Person:

    def __init__(self, name="", age=0):
        self.name = name
        self.age = age

    def Introduce(self):
        print("Hi, my name is {} and I'm {}.".format(self.name, self.age))

p = Person()
p.name = "John"
p.age = 30
p.hair = "black" # ok in Python, compile time error in Go
p.Introduce()
```

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, my name is %s and I'm %d.\n", p.Name, p.Age)
}

func main() {
	var p Person
	p.Name = "John"
	p.Age = 30
	p.Introduce()
}
```
