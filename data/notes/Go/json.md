Networked programs need to exchange information via messages. TCP and UDP
provide a transport mechanism to do this. However, at transport level the
messages are just sequences of bytes with no structure.

A program will typically build a complex data structure to hold the current
program state. To transfer this data outside of the program's own address space
(e.g. to another application over the network) it needs to be serialized. This
process is also called marshalling or encoding.

JSON is a text-based format used for serialization by many programming
languages.

```go
// SaveJSON.go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name  Name
    Email []Email
}

type Name struct {
    Family   string
    Personal string
}

type Email struct {
    Kind    string
    Address string
}

func main() {
    person := Person{
        Name:   Name{Family: "Newmarch", Personal: "Jan"},
        Email:  []Email{
                    Email{Kind: "home", Address: "jan@newmarch.name"},
                    Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"},
                },
    }
    saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := json.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)
    outFile.Close()
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err.Error())
        os.Exit(1)
    }
}
```

```go
// LoadJSON.go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name  Name
    Email []Email
}

type Name struct {
    Family   string
    Personal string
}

type Email struct {
    Kind    string
    Address string
}

// This method gets used implicitly by Println() in the main() below.
func (p Person) String() string {
    s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s += "\n" + v.Kind + ": " + v.Address
    }
    return s
}

func main() {
    var person Person
    loadJSON("person.json", &person)
    fmt.Println("Person", person)
}

func loadJSON(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := json.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err.Error())
        os.Exit(1)
    }
}
```

More: [Network programming with Go](https://www.apress.com/gp/book/9781484226919): 4. Data Serialization
