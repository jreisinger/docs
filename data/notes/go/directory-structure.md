# The new way (Go modules)

```
# Let's use Go modules for dependency management.
cd ~/github
mkdir my-project
go mod init github.com/jreisinger/my-project

# Code away!
vim main.go
```

* https://golang.org/doc/code.html
* https://blog.golang.org/using-go-modules

# The old way (`GOPATH`, `src`, ...)

*NOTE*: See [How to Write Go Code](https://golang.org/doc/code.html) for official tutorial.

Go has a strong opinion about directory structure.

```bash
## Set location of your workspace (It's used by the `go` tool)
$ export GOPATH=`pwd`
$ echo $GOPATH
/Users/reisinge/temp/go
```

A single file program:

```plain
## Create a program
$ tree
.
└── src
    └── hello           # this directory name is the program name
        └── main.go     # package main

2 directories, 1 file

## Install the program
$ go install hello      # compile and install all *.go files in $GOPATH/src/hello 
$ tree
.
├── bin
│   └── hello           # hello binary that has been built and installed
└── src
    └── hello
        └── main.go

3 directories, 2 files

## Run the program
$ bin/hello
Hello, world!
```

A multiple file (and multiple package) program:

```plain
## Create a program
$ tree
.
└── src
    ├── hello
    │   └── main.go         # package main imports package shuffler
    └── shuffler            # this directory name is the package name
        └── shuffle.go      # package shuffler

3 directories, 2 files

## Install the program
$ go install hello          # Compile and install all *.go files in 
                            # $GOPATH/src/hello and all *.go files in the
                            # imported packages.
$ tree
.
├── bin
│   └── hello
├── pkg
│   └── darwin_amd64        # my OS and CPU architecture
│       └── shuffler.a      # object file of the imported package
└── src
    ├── hello
    │   └── main.go
    └── shuffler
        └── shuffle.go

6 directories, 4 files

## Run the program
$ bin/hello
[3 5 1 1 4 9]
```

The actual source code of the multifile program:

`src/hello/main.go`

```go
package main

import (
	"fmt"
	"shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	i := intSlice{3,1,4,1,5,9}
	shuffler.Shuffle(i)
	fmt.Println(i)
}
```

`src/shuffler/shuffle.go`

```go
package shuffler

import (
	"math/rand"
)

type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(s Shuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}
```
