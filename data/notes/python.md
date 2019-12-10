Working with files and STDIN
----------------------------

Finding duplicate lines

    #!/usr/bin/env python3
    # dup prints lines from file(s) or STDIN that appear more than once.
    import fileinput

    counts = dict()

    for line in fileinput.input():
        line = line.rstrip()
        if line not in counts:
            counts[line] = 1
            continue
        counts[line] += 1

    for line, n in counts.items():
        if n > 1:
            print(line, n)

Reading from files

    #!/usr/bin/env python3

    filename = "/usr/share/dict/words"

    # Read the entire file as a single string
    with open(filename, "rt") as f:
        data = f.read()

    # Iterate over the lines of a file
    with open(filename, "rt") as f:
        for line in f:
            print(line, end="")
