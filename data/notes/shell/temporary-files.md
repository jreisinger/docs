Working with temporary files in a shell script:

    # Create temporary files and store their names in variables ...
    TMPFILE1=$(mktemp /tmp/$0.XXXXX)
    TMPFILE2=$(mktemp /tmp/$0.XXXXX)
    # Make sure they get removed even if the script exits abnormally.
    trap "exit 1"                     HUP INT PIPE QUIT TERM
    trap "rm -rf $TMPFILE1 $TMPFILE2" EXIT
    # ... or alternatively
    #trap "rm -rf $TMPFILE1 $TMPFILE2; exit 1" HUP INT PIPE QUIT TERM

Working with a temporary directory in a shell script:

```
#!/bin/bash

# Create a temporary directory and store its name in a variable ...
TMPDIR=$(mktemp -d)
# Make sure it gets removed even if the script exits abnormally.
trap "exit 1"         HUP INT PIPE QUIT TERM
trap "rm -rf $TMPDIR" EXIT
```
