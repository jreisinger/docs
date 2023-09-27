package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	for _, addr := range os.Args[1:] {
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
