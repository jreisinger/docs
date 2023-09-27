package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close() // execute when surrounding function (main) returns

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handle(conn) // handle connections concurrently
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
