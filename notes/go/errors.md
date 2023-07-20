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
- if failure has only one possible cause, it's of `bool` type
```
value, ok := cache.Lookup(key)
if !ok {
    // cache[key] does not exist
}
```
- otherwise it's of `error` type
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
- the price, and the point, is more attention paid to error handling logic

Error-handling strategies

- the caller has to check for errors and take appropriate action

1) *propagate* (and enrich) the error

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

2) *retry* if the error is transient
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

3) *stop* the program gracefully (generally from the main package)
```
// (In function main.)

// Interactive tool logging.
log.SetPrefix("wait: ") // command name
log.SetFlags(0)         // no timestamp

if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}
```

4) log the error and *continue* (with reduced functionality)
```
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled", err)
}
```

5) in rare cases we can safely *ignore* en error entirely
```
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
    return fmt.Errorf("failed to create temo dir: %v", err)
}

// ...use temp dir...

os.RemoveAll(dir) // ignore error; $TMPDIR is cleaned periodically
```

Errors after Go 1.13

- usually we compare an error to nil to see if an operation failed
- sometimes we compare to a known *sentinel* value to see if a specific error occurred
```
var ErrNotFound = errors.New("not found")

if err == ErrNotFound {
    // something wasn't found
}
```
- `error` is a built-in interface type that's easy to implement
- we can use the [type assertion](https://go.dev/tour/methods/15) (or type switch) to view an error value as a more specific type
```
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

if e, ok := err.(*NotFoundError); ok {
    // e.Name wasn't found
}
```
- an error often contains another lower-level error
```
type QueryError struct {
    Query string
    Err   error // lower-level error
}
```
- to look at (unwrap) the lower-level error before Go 1.13 we use the [type assertion](https://go.dev/tour/methods/15)
```
if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
    // query failed because of a permission problem
}
```
- Go 1.13 added three new functions in the `errors` package (`Unwrap`, `Is` and `As`), and a new formatting verb for `fmt.Errorf` (`%w`)

The Unwrap method

- an error which contains another may implement `Unwrap` method returning the underlying error
- if e1.Unwrap returns e2, then we say that e1 wraps e2, and that you can unwrap e1 to get e2
- the result of unwrapping an error may have itself an Unwrap method -> error chain
```
func (e *QueryError) Unwrap() error { return e.Err }
```
- the `errors` package contains Unwrap function but it's usually better to use Is or As
- because Is and As examine the entire error chain in a single call

Examining errors with Is and As

- Is compares an error to a value
```
if errors.Is(err, ErrNotFound( {
    // something was not found
}
```
- As tests whether an error is a specific type
```
var e *QueryError
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}
```
- in the simplest case Is behaves like a comparison to a sentinel error, and As behaves like a type assertion
- but when operating on wrapped errors, these functions consider all the errors in a chain
```
if errors.Is(err, ErrPermission) {
    // err, or some error that it wraps, is a permission problem
}
```

Wrapping errors with %w

- if %w is present, the error returned by fmt.Errorf will have an Unwrap method returning the argument of %w, which must be an error
- in all other ways, %w is identical to %v
```
if err != nil {
    // Return en error which unwraps to err.
    return fmt.Errorf("decompress %v: %w", name, err)
}
```
- wrapping an error with %w makes it available to Is and As
```
err := fmt.Errorf("access denied: %w", ErrPermission)
// ...
if errors.Is(err, ErrPermission ...
```

Sources:

* The Go Programming Language (2016, Go 1.5)
* https://go.dev/blog/go1.13-errors
