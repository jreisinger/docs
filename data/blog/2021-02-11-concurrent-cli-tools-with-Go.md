# Concurrent CLI tools with Go

Imagine you have a list of 100 URLs and you want to check whether they are OK
(i.e. they return 200 HTTP status code). Well, easy enough you might think - 
I'll run `curl` with some fancy options in a for loop:

```
$ time for url in $(cat urls.txt); do curl -s -o /dev/null -w "%{http_code}" -L $url; echo " $url"; done
200 https://golang.org/doc
200 https://perl.org
404 https://perl.org/python

<...SNIP...>

real	0m37.991s
user	0m3.491s
sys     0m1.243s
```

Wait a second! Almost 40 seconds? What if I needed to check 10,000 URLs!

Let's try another approach. Go is famous for its easy concurrency by the
virtue of goroutines and channels. It's true that it is much easier (and cheaper) to write
concurrent programs in Go than in many other languages. But writing a
concurrent program is still more difficult than a sequential one. So you
might find useful the [work](https://github.com/jreisinger/work) package that
abstracts the concurrency code away. To use it you just need to implement
Factory and Task interfaces.

To implement the Factory interface you need to create a data structure (an
empty `struct` in this case) with the Generate method.

```
type factory struct{}

func (f *factory) Generate(line string) work.Task {
	t := &task{URL: line}
	return t
}
```

The Generate method takes a line from STDIN and creates a task from it. Task
is another interface. It requires Process method that will process the task
and Print method to print the results.

```
type task struct {
	URL    string
	Status bool
}

func (t *task) Process() {
	resp, err := http.Get(t.URL)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		t.Status = true
	}
}

func (t *task) Print() {
	status := map[bool]string{
		true:  "OK",
		false: "NOTOK",
	}
	fmt.Printf("%-5s %s\n", status[t.Status], t.URL)
}
```

Now you run the factory with some number of workers:

```
func main() {
	w := flag.Int("w", 100, "number of concurrent workers")
	flag.Parse()

	f := &factory{}
	work.Run(f, *w)
}
```

Let's see if we can check those 100 URLs faster:

```
$ go build
$ time ./urlchecker < urls.txt
NOTOK https://nonexistent.net
OK    https://reisinge.net/notes/go/basics
OK    https://golang.org/doc

<...SNIP...>

real	0m1.641s
user	0m0.438s
sys     0m0.222s
```

Less than two seconds. That's not bad :-). See
[examples](https://github.com/jreisinger/work/tree/main/examples) for the
full program described above and more.
