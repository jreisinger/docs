## Computer Science (CS)

Fundamental question for CS is *what can be computed*?

## Hardware

Computers are to CS what telescopes are to astronomy (Dijkstra).

Functional view of a computer:

```
                +------------+
                |  +-----+   |    +---------+
                |  | CPU |   |--->| Output  |
                |  +-----+   |    | Devices |
+---------+     |    | ^     |    +---------+
| Input   |---->|    v |     |
| Devices |     | +--------+ |    +-----------+
+---------+     | | Main   | |--->| Secondary |
                | | Memory | |<---| Memory    |
                | +--------+ |    +-----------+
                +------------+
                
```

* Memory - stores programs and data (Main is fast but volatile, Secondary - HD, CD, DVD, ...)
* CPU - simple arithmetic operations (e.g. adding 2 numbers) and logical operations (e.g. testing 2 numbers are equal)

Running a program:

1. Instructions that comprise the program are copied (loaded) from secondary to main memory.
2. CPU starts executing the program, i.e. it does this cycle: fetch, decode, execute

## Programming languages (PL)

* notations for expressing computations in an exact and unambiguous way
* every structure in a PL has a precise form (its *syntax*) and a precise meaning (its *semantics*)

Sources

* [Grokking Algorithms](https://learning.oreilly.com/library/view/grokking-algorithms-an/9781617292231/) (2016)
* Python Programming: An Introduction to Computer Science (2010)
