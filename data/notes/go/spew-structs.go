```
// dilbert.go
package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	DoB      time.Time
	Manager  string
}

func main() {
	var employees []Employee
	for _, name := range []string{"Dilbert", "Wallie", "Boss"} {
		employees = append(employees, Employee{Name: name})
	}
	//fmt.Printf("%+v\n", employees)
	spew.Dump(employees)
}
```
