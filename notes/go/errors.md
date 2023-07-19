What are errors

- some functions can't fail, ex. strings.Contains or strconv.FormatBool
- other functions can't fail if their preconditions are met, ex. time.Date
- many functions can fail because they depend on factors beyond programmer's control (like I/O)
- errors are a way to understand why a function failed

Go philosophy

- errors are an important part of a package's API or an application's user interface
- failure is just one of several expected behaviors

Errors as values

- functions that can fail indicate this via the last returned value
- if failure has only one possible cause, it's a of `bool` type
```
value, ok := cache.Lookup(key)
if !ok {
    // cache[key] does not exist
}
```
- otherwise it's of `error` type, which is an interface
```
if err != nil {
    // something went wrong
}
```
- exceptions used in other languages often lead to incomprehensible stack traces
- Go uses ordinary values to indicate errors to encourage self-contained grep-able messages like
```
genesis: crashed: no parachute: G-switch failed: bad relay orientation
```
- the price (and the point) is more attention paid to error handling logic

Error-handling strategies

- the caller has to check for errors and take appropriate action

1) propagate the error

```
resp, err := http.Get(url)
if err != nil {
    return nil, err
}

// ...

doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
    // prepend additional info (that the callee doesn't have)
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

2) retry if the error is transient
```
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // success
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries)) // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

3) stop the program gracefully (generally from the main package)
```
// (In function main.)

// Interactive tool logging.
log.SetPrefix("wait: ") // command name
log.SetFlags(0)         // no timestamp

if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}
```

4) log the error and continue (with reduced functionality)
```
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled", err)
}
```

5) in rare cases we can safely ignore en error entirely
```
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
    return fmt.Errorf("failed to create temo dir: %v", err)
}

// ...use temp dir...

os.RemoveAll(dir) // ignore error; $TMPDIR is cleaned periodically
```

Errors after Go 1.13

Sources:

* The Go Programming Language (2016, Go 1.5)
* https://go.dev/blog/go1.13-errors
