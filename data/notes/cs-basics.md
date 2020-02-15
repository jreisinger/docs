## Computer Science (CS)

* fundamental question for CS is *what can be computed*?
* it uses techniques of design, analysis, and experimentation 

## Hardware

* computer is a universal information-processing machine
* computer can carry out any process that can be described in sufficient details
* computers are to CS what telescopes are to astronomy (Dijkstra)

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

CPU can only understand a very *low-level* language - machine language. Machine language is created by the designers of the particular CPU. Adding two numbers looks like:

```
load the number from memory location 2020 into the CPU
load the number from memory location 2021 into the CPU
add the two numbers int the CPU
store the result into location 2023
```

(Instructions and data are really represented in *binary* notation - sequences of 0s and 1s.)

In *high-level* (human-oriented) language - like Python - it looks like:

```
c = a + b
```

There are two ways to translate a high-level language to machine language:

1) Compiling:

```
Source +---> Compiler +---> Machine
Code                        Code
                              +
                              |
                              v
             Inputs +---->  Running +---> Outputs
                            Program
```

* compiler - a complex program

2) Interpreting:

```
            +--------------+
Source      |              |
Code   +--> |              |
            | Interpreter  | +--> Outputs
            |              |
Inputs +--> |              |
            +--------------+
```

* interpreter - program that simulates a computer that undestands a high-level language
* more flexible (interactive) development

Sources

* [Grokking Algorithms](https://learning.oreilly.com/library/view/grokking-algorithms-an/9781617292231/) (2016)
* Python Programming: An Introduction to Computer Science (2010)
