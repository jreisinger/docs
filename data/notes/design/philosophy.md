 The most important property of a program is whether is accomplishes the intention of its **user**. -- C.A.R. Hoare (1969)

**Simple**, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have **clean**, simple code. -- Brian Kernighan

You need a certain amount of **complexity** to do any particular job. ... In short, simplicity is often the enemy of success. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html)

So all else being equal, scientists prefer hypotheses which are simple, uniform, common-sensical and aesthetically pleasing. ... Nevertheless, ultimately the criterion of elegance is subordinate to observations.  It doesn't matter how beautiful or simple your theory is, if it gets the facts wrong. -- [Aron Wall](http://www.wall.org/~aron/blog/pillar-of-science-ii-elegent-hypotheses/)

It is not possible to create a perfect program the first time. The insight necessary to find the right solution comes only with a combination of thought and experience; pure introspection will not produce a good system, nor will pure hacking. Reactions from users count heavily here; a **cycle** of prototyping, experiment, user feedback, and further refinement is most effective. -- The Practice of Programming

**Data** dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

# Go

Don't do object oriented design but data oriented design.

Go prefers convention over configuration because configuration is limiting.

Let's not group by what one [is](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_1.go) but by what one can [do](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go). In Go there are no [classes](https://github.com/jreisinger?tab=repositories&q=animal) and subclasses but Interfaces and composition.

Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.

# Perl

* TIMTOWTDI
    * don't rush to code up the first idea that pops into your head
    * after you have an algorithm strive for clarity, simplicity, efficiency, scalability and elegance (poems of logic)
* ask yourself how would you solve the problem - the straightforward solution is often (not always) simple, clear, and efficient enough
* generality is good - consideration of a more general problem can lead to a better solution for some special case (max of "n" numbers instead of "three")
    * if the general program is as easy to write as the special one go for the general one as it is more likely to be useful in other situations (maximum utility from the effort)
* don't reinvent the wheel - a lot of very smart programmers have designed countless good algorithms and programs
    * good for learning though
* Just because you CAN do something a particular way doesn't mean that you SHOULD do it that way. Perl is designed to give you several ways to do anything, so consider picking the most readable one. See [perlstyle](http://perldoc.perl.org/perlstyle.html) for more.
