# Go for cybersecurity tools

After getting a basic idea of what TLS is in the previous post, let's write a tool the will help us finding out the TLS version of a server. The first idea might be to range over the arguments that should be TCP addresses. For each address we'll get and print the TLS version.

We might start with a pseudo-code like:

```
for all IP addresses supplied as CLI arguments:
    get the the TLS version
```

This look simple enough to implement:

```go
package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	for _, addr := range os.Args[1:] {
		ver, err := getTLSVersion(addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tlsver: %v", err)
			continue
		}
		fmt.Printf("%s\t%s\n", ver, addr)
	}
}
```

Now we need to write the GetTLSVersion function. Fortunately there's the standard library package [tls](https://pkg.go.dev/crypto/tls):

```go
func getTLSVersion(addr string) (string, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return tls.VersionName(conn.ConnectionState().Version), nil
}
```

Let's run the program (saved under filename `tlsver.go`):

```sh
$ go run tlsver.go example.com:443 example.net:443 wall.org:443
1.3	example.com:443
1.3	example.net:443
1.2	wall.org:443
```

As you can see, Larry is a bit behind :-).

See the whole program at https://github.com/jreisinger/docs/blog/gosec/tlsver/1.

## Adding concurrency

The code above works fine. But imagine that we want to check the TLS version of a thousand hosts. Connecting to the hosts one after another might take some time. Concurrency means organizing the program in a way that multiple processes can execute independently. Even at the same time if you have multiple processors (which you most certainly do nowadays). Go has an excellent support for doing this.

Basically we want to run the getTLSVersion function and forget about. Then run the next one and forget about it. And so on. Obviously we need to feed some input (function parameters) into the function and collect its output (return values). We use `in` and `out` channels for this. The channels are like typed shell pipes (`$ ls | wc -l`). The type in our case is `host` - it holds all the necessary input and output data. When we fire up goroutines we usually don't want to allow for an unlimited number of them because we might exhaust computing resources (like open sockets or file descriptors). So we run only 30 goroutines. We also want to know when the goroutines are done. For this we use the `WaitGroup`, which is kind of a concurrency-safe counter.

```go
type host struct {
	addr     string
	tlsVer   string
	insecure bool
	err      error
}

in := make(chan host)
out := make(chan host)
var wg sync.WaitGroup

for i := 0; i < 30; i++ {
	wg.Add(1)
	go func() {
		for h := range in {
			h.tlsVer, h.err = getTLSVersion(h.addr, h.insecure)
			out <- h
		}
		wg.Done()
	}()
}
```

We want to get the list of TCP addresses from command line arguments or from standard input. We don't want to wait (block) on the input so we run it in a goroutine as well. When there's no more input we close the `in` channel and decrease the `WaitGroup` counter.

```go
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
```

And we just print the output coming from the `out` channel. We close the channel when the `WaitGroup` counter is zero, i.e. all goroutines are done.

```go
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
```

Now we can quickly check many hosts:

```sh
$ cat ~/Downloads/top1000domains.txt | tlsver -insecure -concurrency 10
TLS 1.3	facebook.com:443
TLS 1.3	youtube.com:443
TLS 1.3	google.co.in:443
TLS 1.3	twitter.com:443
TLS 1.2	live.com:443
TLS 1.3	wikipedia.org:443
TLS 1.2	bing.com:443
TLS 1.2	amazon.com:443
TLS 1.2	msn.com:443
TLS 1.2	linkedin.com:443
<...SNIP...>
```

See the whole program at https://github.com/jreisinger/docs/blog/gosec/tlsver/2.

# Tips for designing programs

Design iteratively. No one designs a program top to bottom in a linear, systematic fashion.

Try out alternatives. Good design involves a lot of trial and error. When you look at someone's code, it's finished work, not the process they went through to get there.

Keep it simple. Don't design in extra complexity until it is really needed.

Solve one problem at a time, don't be overwhelmed by everything.

# More

* https://eli.thegreenplace.net/2021/go-socket-servers-with-tls/
* https://eli.thegreenplace.net/2021/go-https-servers-with-tls/
* https://github.com/lizrice/secure-connections