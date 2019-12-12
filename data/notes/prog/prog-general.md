Computer science (CS)

* studies what can be computed
* uses techniques of design, analysis and experimentation

Computer

* universal machine for information storage and manipulation
* can carry out any instructions described in sufficient detail ([program](../terms/program))
* computers are to CS what telescopes are to astronomy

Hardware (functional view)

* CPU -- simple arithmetic (adding two numbers) and logical (testing two numbers
    are equal) operations
* Main memory - RAM (fast but volatile) -- stores information ([data](../terms/data) and programs) that CPU acts on
* Secondary memory (HD, CD, DVD) -- stores more permanent information
* Input/Output devices -- for entering/displaying information
* running a program: HD --> RAM --> CPU (fetch, decode, execute - with blazing
    speed!)

Programming languages

* formal notation for writing programs
* have precise syntax (form) and semantics (meaning)
* machine language: low-level, created by the CPU designers, understood by computer hardware (in binary notation - sequences of 0s and 1s):
    * load the number from memory location 2001 into the CPU
    * load the number from memory location 2002 into the CPU
    * add the two numbers in the CPU
    * store the result into location 2003
* high-level language: human-oriented, must be compiled or interpreted, more portable
* compiled: faster
    * compiler - complex program that translates source code into machine code
* interpreted: more flexible (interactive) development
    * interpreter - a program that simulates a computer that understands a high-level language
* traditional systems languages (C++, Ada, Java): large-scale programming, emphasis on structure and discipline
* scripting ("agile") languages: designed for writing small/medium-scaled programs easy

Programming styles

* object oriented - objects are the main focus; we tell objects to do things (by calling their methods)
* imperative - functions are the primary focus; we pass them objects to work with

(Python) program

* sequence of commands (statements) telling computer what to do
* statement = identifiers + expressions
* identifiers - names, begin with an underscore or letter followed by letter,
    digit or underscore
* expressions - fragments of program that produce data
    * literals
    * variables
    * operators

(Perl) program

* statements are made up of expressions
* expressions are made up of terms and operators
* term (terminal) - leaf node of a syntax tree; "a thing that functions grammatically as an operand for the operators in an expression"

Function

* kind of a subprogram
* why:
    * to reduce code duplication
    * to make program more modular (better structured)
    * to make a program more self-documenting (you don't need that many comments)
    * to facilitate unit testing (test each function separately)
* see [rball](../code/cs_intro/rball) for separating concerns using functions
    * abstraction (concerns separation) -- determining important
        characteristics and ignoring other details

Programming hints

* TIMTOWTDI
    * don't rush to code up the first idea that pops into your head
    * after you have an algorithm strive for clarity, simplicity, efficiency, scalability and elegance (poems of logic)
* ask yourself how would you solve the problem - the straightforward solution is often (not always) simple, clear, and efficient enough
* generality is good - consideration of a more general problem can lead to a better solution for some special case (max of "n" numbers instead of "three")
    * if the general program is as easy to write as the special one go for the general one as it is more likely to be useful in other situations (maximum utility from the effort)
* don't reinvent the wheel - a lot of very smart programmers have designed countless good algorithms and programs
    * good for learning though
* Just because you CAN do something a particular way doesn't mean that you SHOULD do it that way. Perl is designed to give you several ways to do anything, so consider picking the most readable one. See [perlstyle](http://perldoc.perl.org/perlstyle.html) for more.


Programming

1. Problem analysis
2. Program specification (what to do)
3. Design (how to do it - algorithm)
    * Top-down
    * Prototyping
    * Spiral
4. Implementation (algorithm gets translated into programming language)
5. Testing and debugging
6. Maintenance (most programs are never finished)

Why test

* find out problems sooner
* force us to write code in smaller chunks
* psychologically easier to modify the code
* more confidence answering questions about our code

Sources:

* Python Programming: An Introduction to Computer Science, 2nd
* Intermediate Perl, 2nd
