The most important property of a program is whether is accomplishes the intention of its user. -- C.A.R. Hoare (1969)

# Wall: Perl (1988)

You need a certain amount of complexity to do any particular job. ... In short, simplicity is often the enemy of success. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html) (1999)

# Raymond (esr): The Cathedral and the Bazaar (1999)

Main points from [ch2](http://www.catb.org/~esr/writings/cathedral-bazaar/cathedral-bazaar/ar01s02.html) of The Cathedral and the Bazaar.

1. Every good work of software starts by scratching a developer's personal itch.
2. Good programmers know what to write. Great ones know what to rewrite (and reuse).
3. "Plan to throw one away, you will anyhow." (Fred Brooks, The Mythical Man-Month, Chapter 11)
The point isn't merely that you should expect first attempt to be wrong, it's that starting over with the right idea is usually more effective than trying to salvage a mess.
5. If you have the right attitude, interesting problems will find you.
6. When you lose interest in a program, your last duty to it is to hand it off to a competent successor.

# Kernighan, Pike: The Practice of Programming (1999)

Code should be clear and simple—straightforward logic, natural expression, conventional language use, meaningful names, neat formatting, helpful comments—and it should avoid clever tricks and unusual constructions. -- Kernighan, Pike

It is not possible to create a perfect program the first time. The insight necessary to find the right solution comes only with a combination of thought and experience; pure introspection will not produce a good system, nor will pure hacking. Reactions from users count heavily here; a **cycle** of prototyping, experiment, user feedback, and further refinement is most effective.

# Raymond (esr): The Art of UNIX Programming (2003)

All the [Unix] philosophy really boils down to one iron law, the hallowed ’KISS principle’ of master engineers everywhere:

![image](https://user-images.githubusercontent.com/1047259/149891682-0e0e2633-2ea7-4862-ae9d-106f99493bb8.png)

# Griesemer, Pike, Thompson: Go (2009)

No idea went into Go until it had been simplified to its essence and then had clear benefits that justified the complexity being added. -- Russ Cox

Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

And in both [prose and code] I do a great deal of rereading and rewriting. The first draft, even if correct, is rarely ‘right’. -- [Rob Pike](https://www.red-gate.com/simple-talk/opinion/geek-of-the-week/rob-pike-geek-of-the-week/)

Let's not group by what one [is](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_1.go) but by what one can [do](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go). In Go there are no [classes](https://github.com/jreisinger?tab=repositories&q=animal) and subclasses but Interfaces and composition. -- Bill Kennedy

# Rob Pike's 5 Rules of Programming

1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is.
2. Measure. Don't tune for speed until you've measured, and even then don't unless one part of the code overwhelms the rest.
3. Fancy algorithms are slow when n is small, and n is usually small. Fancy algorithms have big constants. Until you know that n is frequently going to be big, don't get fancy. (Even if n does get big, user rule 2 first.)
4. Fancy algorithms are buggier than simple ones, and they're much harder to implement. Use simple algorithms as well as simple data structures.
5. Data dominates. If you've chosen the right data structures and organized thigs well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming.

Pike's rules 1 and 2 restate Tony Hoare's famous maxim "Premature optimization is the root of all evil." Ken Thompson rephrased Pike's rules 3 and 4 as "When in doubt, use brute force." Rules 3 and 4 are instances of the design philosophy KISS. Rule 5 was previously stated by Fred Brook in the Mythical Man-Month. Rule 5 is often shortened to "write stupid code that uses smart objects".

Source: https://users.ece.utexas.edu/~adnan/pike.html
