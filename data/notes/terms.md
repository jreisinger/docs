## big data

The idea that large databases of seemingly random data about people (ex. purchasing habits, location information) are valueable.

## binary data 

Data represented in binary form rather than interpreted on higher level or converted to some other form. While most binary
data has symbolic meaning not all binary data is numeric (ex. computer instructions within processor registers).

## block size

The POSIX standard defines a block size of 512 bytes. However, this size is harder to read, so by default, the `df` and `du` output in most Linux distributions is in 1024-byte blocks.

## caching

means that data likely to be used in the future is kept "nearby"

## cgroups

* (control groups)
* an optional Linux kernel feature that allows for finer tracking of a process hierarchy
* used by `systemd`

## cloud native

Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds…. These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.

-- Cloud Native Computing Foundation, CNCF Cloud Native Definition v1.0

## concurrency

* managing a lot of things at once

## data

Qualitative or quantitative facts about things. Information and then knowledge
are derived from data.

## directory permissions

You can list the contents of a directory if it's readable (`r`), but you can only access a file in a directory if the directory is executable (`x`). Typical (absolute) permission modes for directories: `755`, `700`, `711`.

## filesystem

A form of database. It supplies the structure to transform a simple block device into the sophisticated hierachy of files and subdirectories that users can understand.

## first-class functions

In a language with first-class functions, functions can be assigned to variables, and then called later using those variables.

Functions can also be passed as arguments when calling other functions.

## fstab fields

```
<...> 0 1
```

* backup information for use by the dump command - should be always 0
* the filesystem (FS) integrity test order - 1 for the root FS, 2 for any other FS on a hard disk, 0 to disable the bootup check (ex. CD-ROM, swap, proc)

## index

In a typical database, it's an internal structure that is used to *increase the speed of data retrieval*. It's a copy of selected data that can be searched very efficiently, which might also include a file-level disk block number or even a direct connection to the entire set of data it was copied from.

## instantiate

create an object from the class, ex.

```python
class Eyes(object):

    def __init__(self, color='brown'):
        self.color = color

    def print_color(self):
        print("Eyes are", self.color)

eyes_obj = Eyes(color='blue')  # instantiation
```

## ldd

list dynamic dependencies

## linux bridge

a Linux kernel feature that connects two network segments (poor man's router)

## loose coupling

a system property and design strategy in which a system’s components have minimal knowledge of any other components. Two systems can be said to be loosely coupled when changes to one component generally don’t require changes to the other

## marshalling

(or serializing, or just encoding) converting non-bytes data into bytes

A program will typically build a complex data structure to hold the current
program state. To transfer this data outside of the program's own address space
(e.g. to another application over the network) it needs to be serialized.

## newline

Set of characters indicating the end of a line:

* Linux, MacOS: `\n` (one character represented by two symbols)
* Windows: `\r\n` (two characters)

## parallelism

* doing a lot of things at once
* you need multiple physical processors for this

## preemption

The interruption of a computer process without its cooperation in order to
perform another task.

Preemptive OS means that the rules governing which
processes receive use of the CPU and for how long are determined by the kernel
process scheduler (rather than by the processes themselves).

## program

Instructions (or commands, or statements) for a computer to execute. Cooking
recipe or direction instructions are kinds of programs.

Executable file residing on disk in a directory. It is read into memory and is
executed by the kernel as a result of one of the seven `exec` functions.

## servlets

Java programs that run on the server on top of an application server platform.

## stateless protocol

A communications protocol that treats each request as an independent
transaction that is unrelated to any previous request (ex. HTTP).

## unicode

- "universal encoding"
- 32 bits to encode a character
- problem: wasted space since 8 bits are often enough for enconding a char
- solution: "Unicode Transformation Format 8 bits" (UTF-8) - encodes most
  common characters using 8 bits, and then "escapes" into larger numbers when
  needed

DBES - decode bytes encode strings:

```
>>> raw_bytes = b'\xe6\x96\x80'
>>> raw_bytes.decode()   # DB
'斀'
>>> utf_string = '斀'
>>> utf_string.encode()  # ES
b'\xe6\x96\x80'
```

## variable

(in Go) is a piece of storage containing a value

## xor

Exclusive or - the opposite of inclusive (the normal) or, which allows both
possibilities as well as either:

    >>> True or True
    True
    >>> True or False
    True
    >>> False or True
    True
    >>> False or False
    False

Example of exlusive or in English: "You can have pizza or chicken" - you probably
don't mean you can have both. Exclusive or expression in Python:

    >>> True != True
    False
    >>> True != False
    True
    >>> False != True
    True
    >>> False != False
    False
