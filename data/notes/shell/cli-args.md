```
#!/bin/bash

echo "Running program [$0] with [$#] arguments:"
for arg in "$@"; do
    echo "[$arg]"
done
```

```
$ ./x.sh hi -v --help -n=1
Running program [./x.sh] with [4] arguments:
[hi]
[-v]
[--help]
[-n=1]
```
