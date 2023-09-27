package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
)

func main() {
	insecure := flag.Bool("insecure", false, "don't verify server certificate")
	flag.Parse()

	for _, addr := range flag.Args() {
		ver, err := getTLSVersion(addr, *insecure)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tlsver: %v\n", err)
			continue
		}
		fmt.Printf("%s\t%s\n", ver, addr)
	}
}

func getTLSVersion(addr string, insecureSkipVerify bool) (string, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: insecureSkipVerify})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return tls.VersionName(conn.ConnectionState().Version), nil
}
