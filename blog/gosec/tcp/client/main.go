package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "Hello from client.")
	if err != nil {
		log.Fatalf("client write error: %s", err)
	}

	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("client read: %s\n", buf[:n])
}
