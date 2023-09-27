package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	certFile := flag.String("cert", "cert.pem", "certificate file")
	keyFile := flag.String("key", "key.pem", "private key file")
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "localhost:4430", config)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
