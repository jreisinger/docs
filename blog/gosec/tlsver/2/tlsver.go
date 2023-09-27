package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	insecure := flag.Bool("insecure", false, "don't verify server certificate")
	concurrency := flag.Int("concurrency", 30, "maximum number goroutines")
	flag.Parse()

	type host struct {
		addr     string
		tlsVer   string
		insecure bool
		err      error
	}

	in := make(chan host)
	out := make(chan host)
	var wg sync.WaitGroup

	go func() {
		if len(flag.Args()) > 0 {
			for _, addr := range flag.Args() {
				in <- host{addr: addr, insecure: *insecure}
			}
		} else {
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				in <- host{addr: s.Text(), insecure: *insecure}
			}
		}
		close(in)
		wg.Done()
	}()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			for h := range in {
				h.tlsVer, h.err = getTLSVersion(h.addr, h.insecure)
				out <- h
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for h := range out {
		if h.err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v", h.addr, h.err)
		}
		fmt.Printf("%s\t%s\n", h.tlsVer, h.addr)
	}
}

func getTLSVersion(addr string, insecureSkipVerify bool) (string, error) {
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}
	conn, err := tls.DialWithDialer(&dialer, "tcp", addr, &tls.Config{InsecureSkipVerify: insecureSkipVerify})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return tls.VersionName(conn.ConnectionState().Version), nil
}
