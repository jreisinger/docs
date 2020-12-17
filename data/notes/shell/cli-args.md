```
#!/bin/bash
# cli-args.sh -- demo of how to handle CLI arguments in bash

echo "Running program [$0] with [$#] arguments:"
for arg in "$@"; do
    echo "[$arg]"
done
```

```
$ ./cli-args.sh hi -v -n=1 --help
Running program [./cli-args.sh] with [4] arguments:
[hi]
[-v]
[-n=1]
[--help]
```
