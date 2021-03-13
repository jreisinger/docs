[spew](https://github.com/davecgh/go-spew) is something like Perl's [Data::Dumper](https://perldoc.perl.org/Data/Dumper.html)

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
	//spew.Config.DisableCapacities = true
	spew.Dump(employees)
}
```

```
$ go run dilbert.go
([]main.Employee) (len=3 cap=4) {
 (main.Employee) {
  ID: (int) 0,
  Name: (string) (len=7) "Dilbert",
  Position: (string) "",
  DoB: (time.Time) 0001-01-01 00:00:00 +0000 UTC,
  Manager: (string) ""
 },
 (main.Employee) {
  ID: (int) 0,
  Name: (string) (len=6) "Wallie",
  Position: (string) "",
  DoB: (time.Time) 0001-01-01 00:00:00 +0000 UTC,
  Manager: (string) ""
 },
 (main.Employee) {
  ID: (int) 0,
  Name: (string) (len=4) "Boss",
  Position: (string) "",
  DoB: (time.Time) 0001-01-01 00:00:00 +0000 UTC,
  Manager: (string) ""
 }
}
```
