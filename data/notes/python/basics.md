*2016-05-13*

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

>>> x = "Do not meddle    in \the affairs of wizards"
>>> x.replace("\t", "t")
'Do not meddle    in the affairs of wizards'

>>> import re
>>> tabs = re.compile(r"[\t ]+")
>>> tabs.sub(" ", x)
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
#!/usr/bin/python

import os
import sys

def walk(dirname):
"""Walk through a directory and print names of files"""
    for name in os.listdir(dirname):
        path = os.path.join(dirname, name)
        if os.path.isfile(path):
            print path
        elif os.path.isdir(path):
            walk(path)

walk(sys.argv[1])
```

# Threading

Not really a Python's strength due to GIL (see Golang).

```python
#!/usr/bin/env python3

import threading
import time

class aThread(threading.Thread):
    def __init__(self, num, val):
        threading.Thread.__init__(self)
        self.threadNum=num
        self.loopCount=val

    def run(self):
        print("Starting to run thread: ", self.threadNum)
        myfunc(self.threadNum, self.loopCount)

def myfunc(num, val):
    count=0
    while count < val:
        print(num, " : ", val*count)
        count=count+1
        #time.sleep(1)

t1=aThread(1, 15)
t2=aThread(2, 20)
t3=aThread(3, 30)

t1.start()
t2.start()
t3.start()

threads = []
threads.append(t1)
threads.append(t2)
threads.append(t3)

# wait for all threads to complete by entering them
for t in threads:
    t.join()
```

# Classes and instances

```python
#!/usr/bin/env python

class me:
    def __init__(self, foo):
        self.myvar = foo

    def getval(self):
        return self.myvar

    def setval(self, bar):
        self.myvar = bar

x = me('bla')
y = x.getval()
z = me('baz')
print(y)
x.setval('ble')
print x.getval()
print z.getval()
```

# Sources

* The Quick Python Book
* Mastering Python (Safari video)
* [Think Python](https://greenteapress.com/wp/think-python-2e/)
