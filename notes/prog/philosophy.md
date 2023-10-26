The most important property of a program is whether is accomplishes the intention of its user. -- C.A.R. Hoare (1969)

## Good programs (Knuth, 1974)

Senses in which a program can be "good":

1. it works correctly
2. it's easy to change
3. it interacts gracefully with users
4. it uses computer's resources efficiently

Both 1. and 2. are achieved when the program is easy to read and understand by a person who knows the appropriate language.

Ad 4. - but premature optimization is the root of all evil!

Source: [Computer programming as an art](https://dl.acm.org/doi/epdf/10.1145/361604.361612)

## Complexity and redundancy (Wall, 1988)

You need a certain amount of complexity to do any particular job. ... In short, simplicity is often the enemy of success. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html) (1999)

Many of those parts were redundant. But that redundancy was absolutely necessary to achieve the goal of putting someone on the moon in 1969. So if some of those rocket parts had the job of being redundant, then each of those parts still had to do their part. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html) (1999)

## Programming principles (Raymond, 1999)

1. Every good work of software starts by scratching a developer's personal itch.
2. Good programmers know what to write. Great ones know what to rewrite (and reuse).
3. "Plan to throw one away, you will anyhow." (Fred Brooks, The Mythical Man-Month, Chapter 11)
4. If you have the right attitude, interesting problems will find you.
5. When you lose interest in a program, your last duty to it is to hand it off to a competent successor.

Source: The Cathedral and the Bazaar - [The mail must get through](http://www.catb.org/~esr/writings/cathedral-bazaar/cathedral-bazaar/ar01s02.html)

## Thought, experience and user feedback (Kernighan, Pike, 1999)

It is not possible to create a perfect program the first time. The insight necessary to find the right solution comes only with a combination of thought and experience; pure introspection will not produce a good system, nor will pure hacking. Reactions from users count heavily here; a **cycle** of prototyping, experiment, user feedback, and further refinement is most effective.

Source: The Practice of Programming

## The UNIX philosophy (Raymond, 2003)

All the [Unix] philosophy really boils down to one iron law, the hallowed [posvätný] ’KISS principle’ of master engineers everywhere:

![](https://user-images.githubusercontent.com/1047259/149891682-0e0e2633-2ea7-4862-ae9d-106f99493bb8.png)

Source: The Art of UNIX Programming

## Beauty (Kernighan, Pike 2007) 

Beatiful code is

* simple - clear and easy to understand
* compact - just enough code to do the job
* general - solving a broad class of problems in uniform way

Source: [A Regular Expression Matcher](https://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html)

## Go and simplicity (Griesemer, Pike, Thompson, 2009)

No idea went into Go until it had been simplified to its essence and then had clear benefits that justified the complexity being added. -- Russ Cox
