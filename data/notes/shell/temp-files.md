Working with temporary files and directories in a shell script:

```
#!/bin/bash

# Create temporary file and directory and store their names in variables ...
TMPFILE=$(mktemp /tmp/"$0".XXXXX)
TMPDIR=$(mktemp -d)
# Make sure we clean up even if the script exits abnormally.
trap 'exit 1'                       HUP INT PIPE QUIT TERM
trap 'rm -rf "$TMPFILE" "$TMPDIR"'  EXIT
```

My [SO answer](https://stackoverflow.com/a/53063602).
