The most important property of a program is whether is accomplishes the intention of its user. -- C.A.R. Hoare (1969)

Code should be clear and simple—straightforward logic, natural expression, conventional language use, meaningful names, neat formatting, helpful comments—and it should avoid clever tricks and unusual constructions. -- Kernighan, Pike (The Practice of Programming, 1999)

It is not possible to create a perfect program the first time. The insight necessary to find the right solution comes only with a combination of thought and experience; pure introspection will not produce a good system, nor will pure hacking. Reactions from users count heavily here; a **cycle** of prototyping, experiment, user feedback, and further refinement is most effective. -- Kernighan, Pike (The Practice of Programming, 1999)

# Go

No idea went into Go until it had been simplified to its essence and then had clear benefits that justified the complexity being added. -- Russ Cox

Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

And in both [prose and code] I do a great deal of rereading and rewriting. The first draft, even if correct, is rarely ‘right’. -- [Rob Pike](https://www.red-gate.com/simple-talk/opinion/geek-of-the-week/rob-pike-geek-of-the-week/)

Don't do object oriented design but data oriented design.

Go prefers convention over configuration because configuration is limiting.

Let's not group by what one [is](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_1.go) but by what one can [do](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go). In Go there are no [classes](https://github.com/jreisinger?tab=repositories&q=animal) and subclasses but Interfaces and composition.

Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

# Perl

You need a certain amount of complexity to do any particular job. ... In short, simplicity is often the enemy of success. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html) (1999)

# Unix

All the philosophy really boils down to one iron law, the hallowed ’KISS principle’ of master engineers everywhere:

![image](https://user-images.githubusercontent.com/1047259/149891682-0e0e2633-2ea7-4862-ae9d-106f99493bb8.png)

See "The Art of UNIX Programming (esr, 2003)" for more.

# The Cathedral and the Bazaar

Main points from [ch2](http://www.catb.org/~esr/writings/cathedral-bazaar/cathedral-bazaar/ar01s02.html) of The Cathedral and the Bazaar.

1. Every good work of software starts by scratching a developer's personal itch.
2. Good programmers know what to write. Great ones know what to rewrite (and reuse).
3. "Plan to throw one away, you will anyhow." (Fred Brooks, The Mythical Man-Month, Chapter 11)
The point isn't merely that you should expect first attempt to be wrong, it's that starting over with the right idea is usually more effective than trying to salvage a mess.
5. If you have the righ attitude, interesting problems will find you.
6. When you lose interest in a program, your last duty to it is to hand it off to a competent successor.
