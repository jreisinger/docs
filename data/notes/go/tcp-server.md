```
// tcp-clock.go -- TCP clock server that sends current time each second
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // blocks until a connection is made
		if err != nil {
			log.Println(err) // e.g., connection aborted
			continue
		}
		// Handle one connection at a time. Just add
		// "go" to handle concurrent connections.
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err) // e.g., client disconnect
			return
		}
		time.Sleep(1 * time.Second)
	}
}
```

Source

* The Go Programming Language (2015), ch. 8.2
