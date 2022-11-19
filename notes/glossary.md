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
* used by `systemd`, Docker

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

## philosophy

literally, the love of wisdom; the science that seeks to understand all things by knowing their causes by natural reason

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

## IPv4 address classes

    .--------------------------------------------------------------------------------------------------------.
    | Class | First octet | Networks                  | Number of networks | Purpose                         |
    +-------+-------------+---------------------------+--------------------+---------------------------------+
    | A     | 1   - 126   | 1.0.0.0   - 126.0.0.0     | 2^7 - 2 = 126      | Unicast (large networks)        |
    | B     | 128 - 191   | 128.0.0.0 - 191.255.0.0   | 2^14 = 16,384      | Unicast (medium-sized networks) |
    | C     | 192 - 223   | 192.0.0.0 - 223.255.255.0 | 2^21 = 2,097,152   | Unicast (small networks)        |
    | D     | 224 - 239   |                           |                    | Multicast                       |
    | E     | 240 - 255   |                           |                    | Experimental                    |
    '-------+-------------+---------------------------+--------------------+---------------------------------'

## latency numbers (~2012)

```
L1 cache reference                           0.5 ns
Branch mispredict                            5   ns
L2 cache reference                           7   ns                      14x L1 cache
Mutex lock/unlock                           25   ns
Main memory reference                      100   ns                      20x L2 cache, 200x L1 cache
Compress 1K bytes with Zippy             3,000   ns        3 us
Send 1K bytes over 1 Gbps network       10,000   ns       10 us
Read 4K randomly from SSD*             150,000   ns      150 us          ~1GB/sec SSD
Read 1 MB sequentially from memory     250,000   ns      250 us
Round trip within same datacenter      500,000   ns      500 us
Read 1 MB sequentially from SSD*     1,000,000   ns    1,000 us    1 ms  ~1GB/sec SSD, 4X memory
Disk seek                           10,000,000   ns   10,000 us   10 ms  20x datacenter roundtrip
Read 1 MB sequentially from disk    20,000,000   ns   20,000 us   20 ms  80x memory, 20X SSD
Send packet CA->Netherlands->CA    150,000,000   ns  150,000 us  150 ms
```

Source: https://gist.github.com/jboner/2841832

## ldd

list dynamic dependencies

## linux bridge

a Linux kernel feature that connects two network segments (poor man's router)

## loose coupling

A system property and design strategy in which a system’s components have minimal knowledge of any other components. Two systems can be said to be loosely coupled when changes to one component generally don’t require changes to the other.

## marshalling

(or serializing, or just encoding) converting non-bytes data into bytes

A program will typically build a complex data structure to hold the current
program state. To transfer this data outside of the program's own address space
(e.g. to another application over the network) it needs to be serialized.

## media speed (~2010)

    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | What              | Size       | Sequential Read Speed| Random Read Speed | Write Speed | Cost     |
    +-------------------+------------+----------------------+-------------------+-------------+----------+
    | LTO-3 write speed |            |                      |                   |     80 MB/s | $0.06/GB |
    | LTO-4 write speed |            |                      |                   |    120 MB/s | $0.05/GB |
    | HD                | Terrabytes |             100 MB/s |            2 MB/s |             | $0.10/GB |
    | SSD               | Gigabytes  |             250 MB/s |          250 MB/s |             | $3.00/GB |
    '-------------------+------------+----------------------+-------------------+-------------+----------'

Source: ULSAH, p. 210, 204

## mutual authentication

Process between a sender and a receiver authenticating each other's identity to be sure who they are talking to. Mutual TLS (mTLS) is one well-known way to achieve mutually authenticated (and encrypted) traffic.

<img width="493" alt="image" src="https://user-images.githubusercontent.com/1047259/173361854-0449cd88-d4ec-4e03-aff7-3c06c4a9f0c1.png">

source: https://isovalent.com/blog/post/2022-05-03-servicemesh-security/

## newline

Set of characters indicating the end of a line:

* Linux, MacOS: `\n` (one character represented by two symbols)
* Windows: `\r\n` (two characters)

## nil

in Go is an untyped identifier that represents the lack of a value for pointer types (e.g. *int) and types implemented with pointers (slices, maps, functions, channels, interfaces). Unlike NULL in C, nil is not another name for 0; you can't convert it back and forth with a number.

## parallelism

* doing a lot of things at once
* you need multiple physical processors for this

## preemption

The interruption of a computer process without its cooperation in order to
perform another task.

Preemptive OS means that the rules governing which
processes receive use of the CPU and for how long are determined by the kernel
process scheduler rather than by the processes themselves.

## Private Address Space

[RFC 1918](https://tools.ietf.org/html/rfc1918)

    .----------------------------------------------------.
    | Class | Networks                      | Prefix     |
    +-------+-------------------------------+------------+
    | A     |    10.0.0.0 - 10.255.255.255  | 10/8       |
    | B     |  172.16.0.0 - 172.31.255.255  | 172.16/12  |
    | C     | 192.168.0.0 - 192.168.255.255 | 192.168/16 |
    '-------+-------------------------------+------------'

All

```
# rfc 1918
10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
# https://en.wikipedia.org/wiki/Reserved_IP_addresses
100.64.0.0/10
192.0.0.0/24
198.18.0.0/15
```

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

- "universal encoding", 32 bits to encode a character
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

## webhook

HTTP-based callback function for event-driven communication between two APIs.
E.g. you give Twilio an URL to be called with specific payload (data) when error
or warning occurs on your account.

## xor

Exclusive or, the opposite of inclusive (normal) or.

Normal or allows both possibilities as well as either:

    >>> True or True
    True
    >>> True or False
    True
    >>> False or True
    True
    >>> False or False
    False

Exclusive or expression in Python:

    >>> True != True
    False
    >>> True != False
    True
    >>> False != True
    True
    >>> False != False
    False

Example of exclusive or in English: "You can have pizza or chicken" - you probably
don't mean you can have both.

