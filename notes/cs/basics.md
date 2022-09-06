# Computer Science

* fundamental question for CS is *what can be computed*?
* it uses techniques of design, analysis, and experimentation
* computers are to CS what telescopes are to astronomy (Dijkstra?)

# Computer

* computer is a universal information-processing machine
* computer can carry out any process that can be described in sufficient details

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

* Memory - stores programs and data that CPU acts on
  * Main is fast but volatile
  * Secondary is permanent: HD, CD, DVD, ...
* CPU - simple arithmetic (e.g. adding 2 numbers) and logical (e.g. testing 2 
  numbers are equal) operations
* I/O Devices - for entering/displaying data

Running a program (HD -> RAM -> CPU):

1. Instructions that comprise the program are copied (loaded) from secondary to
   main memory.
2. CPU starts executing the program, i.e. it does this cycle (really, really
   fast): fetch, decode, execute

# Programming languages

* formal notations for expressing computations in an exact and unambiguous way
* every structure in a PL has a precise form (its syntax) and a precise meaning 
  (its semantics)

CPU can only understand a very **low-level**, machine, language. Machine
language is created by the designers of the particular CPU. Adding two numbers
looks like (instructions and data are really represented in binary notation -
sequences of 0s and 1s):

```
load the number from memory location 2020 into the CPU
load the number from memory location 2021 into the CPU
add the two numbers in the CPU
store the result into location 2023
```

In **high-level**, human-oriented, language (like Python) it looks like:

```
c = a + b
```

## By translation mechanism

There are two ways to translate a high-level language to machine language: to
compile it or to interpret it.

Compiled

```
Source +---> Compiler +---> Machine
Code                        Code
                              +
                              |
                              v
             Inputs +---->  Running +---> Outputs
                            Program
```

* Compiler - a complex program; no longer needed after a program is translated

Interpreted

```
            +--------------+
Source      |              |
Code   +--> |              |
            | Interpreter  | +--> Outputs
            |              |
Inputs +--> |              |
            +--------------+
```

* Interpreter - program that simulates a computer that understands a high-level
  language
* more flexible (interactive) development

## By program scale

Systems languages (Ada, C++, Java)

* for large-scale programming
* emphasis on structure and discipline

Scripting languages (Bash, Perl, Python)

* for writing small/medium-scaled programs easy
* "agile" languages 

## By programming paradigm

Object oriented

* objects are the main focus
* we tell objects to do things by calling their methods

Imperative

* functions are the primary focus
* we pass them objects to work with

Source: Python Programming: An Introduction to Computer Science (2010)
