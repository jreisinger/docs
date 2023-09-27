package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	certFile := flag.String("cert", "cert.pem", "trusted CA certificate")
	flag.Parse()

	data, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(data); !ok {
		log.Fatalf("unable to parse certificate from %s", *certFile)
	}

	config := &tls.Config{RootCAs: certPool}
	conn, err := tls.Dial("tcp", "localhost:4430", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "hello from client")
	if err != nil {
		log.Fatalf("client write error: %v", err)
	}

	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("client read: %s\n", buf[:n])
}
