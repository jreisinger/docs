# Python

## Finding duplicate lines

*A.k.a. templates for working with files and STDIN*

```python
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
```
