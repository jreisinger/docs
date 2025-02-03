Created: 2016-05-13

All Python *objects* are instances of one class or another.

# Variables

Python variables are more like labels that reference various objects (integers, strings, class instances, ...):

```python
x = "2"     # x is string
x = int(x)  # x is now integer
```

* types are associated with objects, not variables
* variables are not checked for type at compile time
* variables don't have to be declared
* variables can have any built-in data type, user-defined object, function, or module assigned to them

# Built-in data types

Built-in data types can be manipulated with operators, built-in functions, library functions, or a data type's own methods.

## numbers

```python
>>> type(1)
<class 'int'>
>>> type(1.1)
<class 'float'>
```

## strings

Strings are immutable.

Raw strings (`r""`) do not interpret escape sequences.

```python
>>> print(r"hello\nworld")
hello\nworld
>>> print("hello\nworld")
hello
world

>>> x = "Do not meddle in \the affairs of wizards"
>>> x.replace("\t", "t")
'Do not meddle in the affairs of wizards'

>>> import re
>>> tabs = re.compile(r"[\t]+")
>>> tabs.sub("t", x)
'Do not meddle in he affairs of wizards'
```

## lists

```python
[]
[1]
[1, "two", (3, 4), ["a", "b"]]
```

* built-in functions: `len()`, `max()`, `min()`
* operators: `in`, `+`, `*` (create new list, original unchanged)
* statements: `del`
* methods: `append()`, `count()`, `extend()`, `index()`, `insert()`, `pop()`, `remove()`, `reverse()`, `sort()`

## tuples

immutable lists

```python
(0, 1, 2)
(1, ) # one element tuple
```

## dictionaries


```python
{ 1: "one", 2: "two" }
```

* associative array functionality implemented using hash tables
* keys must be of an immutable type (numbers, strings, tuples)

## sets

```python
>>> x = set([1, 2, 3])
>>> 1 in x
True
>>> 4 in x
False
```

* an unordered collection of objects
* when membership and uniqueness in the set are important
* like dictionary keys without values

## file objects

```python
f = open("myfile", "r")
line1 = f.readline()
f.close()
```

# What is Truth

* `False` (boolean constant), `0`, `None` (Python nil value) and empty values (ex. `[ ]`, `""`)
    are taken as `False`
* `True` and everything else are considered `True`

# Exceptions

* special objects to manage **errors** that arise during execution
* whenever en error occurs Python creates an exception object
* if you handle the exception (by `try-except-finally-else` compound statement) the program continues to run
* any uncaught exception will cause the program to exit and show a *traceback* describing the exception that was raised

```python
#!/usr/bin/env python3
# Argument parsing and error handling

import argparse

parser = argparse.ArgumentParser(description="This program's description")
parser.add_argument('-f', type=str, help='Name of file to open', required=True)

cmdargs = parser.parse_args()
f = cmdargs.f

try:
    fh = open(f)
    line = fh.readline()
except Exception as e:  # generic exception;
                        # more specific exceptions could be caught before
    print("There was an error:", e)
    exit(1) # like tchrist said
else:
    print("1st line from", f, "is", line, end="")
    fh.close()
```

# Calling to System

## `os` module

```python
import os
os.getcwd()
os.getenv('PATH')
os.system('ls -la')
```

## `subprocess` module

```python
from subprocess import call
call(['ls', '-la'])
```

# Filenames and Paths

```python
#!/usr/bin/env python3

import os
import sys

def walk(dirname):
    """Walk dirname recursively and print filenames
    """
    for name in os.listdir(dirname):
        path = os.path.join(dirname, name)
        if os.path.isfile(path):
            print(path)
        elif os.path.isdir(path):
            walk(path)

if __name__ == "__main__":
    walk(sys.argv[1])

```

# Threading

Not really a Python's strength due to GIL (see Golang).

```python
#!/usr/bin/env python3

import threading
import time

# Define a function for the thread to execute
def print_numbers(thread_name, delay):
    for i in range(5):
        time.sleep(delay)
        print(f"{thread_name}: {i}")

# Create two threads
thread1 = threading.Thread(target=print_numbers, args=("Thread-1", 1))
thread2 = threading.Thread(target=print_numbers, args=("Thread-2", 2))

# Start the threads
thread1.start()
thread2.start()

# Wait for both threads to complete
for thread in [thread1, thread2]:
    thread.join()

print("Exiting Main Thread")
```

# Classes and methods

```python
#!/usr/bin/env python3

class Time:
    """ Represents the time of day.
        attributes: hours (int), minutes (int), seconds (int)
    """

    # __init__ method gets invoked when an object is instantiated
    def __init__(self, hours=0, minutes=0, seconds=0):
        self.hours = hours
        self.minutes = minutes
        self.seconds = seconds

    # __str__ method gets invoked when you print an object
    def __str__(self):
        return "{:02}:{:02}:{:02}".format(self.hours, self.minutes, self.seconds)

    def add_seconds(self, seconds_to_add):
        total_seconds = self.hours * 3600 + self.minutes * 60 + self.seconds + seconds_to_add

        self.minutes, self.seconds = divmod(total_seconds, 60)
        self.hours, self.minutes = divmod(self.minutes, 60)

time = Time()
print("initial time\t{}".format(time))

for seconds in [15, 3600, 7245]:
    time.add_seconds(seconds)
    print("+ {:>4} seconds\t{}".format(seconds, time))
```

# Sources

* Copilot (AI)
* The Quick Python Book
* Mastering Python (Safari video)
* [Think Python](https://greenteapress.com/wp/think-python-2e/)
