# Client

This is an HTTP client implemented using socket-level programming:

```go
// Usage: go run telnet.go
package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // NOTE: ignoring errors by storing them into _
    conn, _ := net.Dial("tcp", "golang.org:80") // Connect over TCP
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") // Send string over the connection
    status, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print(status) // Print the first response line
}
```

To add a timeout you can use the `Dialer` structure (I've also added error
checking + reading from command line arguments):

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
    "time"
)

func main() {
    addr := os.Args[1] // e.g. "reisinge.net:80"

    d := net.Dialer{Timeout: 2 * time.Second}
    conn, err := d.Dial("tcp", addr)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Printf("%s\n", result)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err.Error())
        os.Exit(1)
    }
}
```

It's normal to have a lot of error checking in network programming because lot
of things can go wrong (e.g. syntax error in the address, service not running,
hardware failing).

# Server

Concurrent TCP server that prints (echoes) what it receives:

```go
// Usage: go run tcp_server.go
package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    // listen on a port
    ln, err := net.Listen("tcp", "127.0.0.1:9999")
    if err != nil {
        log.Fatalln(err)
    }

    for {
        // accept a connection
        c, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        // handle the connection
        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn) {
    remoteAddr := c.RemoteAddr().String()
    log.Println("Client connected from", remoteAddr)

    // echo received messages
    scanner := bufio.NewScanner(c)
    for {
        ok := scanner.Scan()
        if !ok {
            break
        }
        fmt.Println(scanner.Text())
    }

    log.Println("Client at", remoteAddr, "disconnected")
}
```

# Sources

* [Network programming with
Go](https://www.apress.com/gp/book/9781484226919): Ch 3. Socket-level programming
* [Introducing Go: Ch. 8.](https://learning.oreilly.com/library/view/introducing-go/9781491941997/ch08.html)
