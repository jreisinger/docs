> The most important property of a program is whether is accomplishes the intention of its user. -- C.A.R. Hoare (1969)

> Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code. -- Brian Kernighan

> In short, simplicity is often the enemy of success. -- [Larry Wall](https://www.oreilly.com/openbook/opensources/book/larry.html)

> So all else being equal, scientists prefer hypotheses which are simple, uniform, common-sensical and aesthetically pleasing. ... Nevertheless, ultimately the criterion of elegance is subordinate to observations.  It doesn't matter how beautiful or simple your theory is, if it gets the facts wrong. -- [Aron Wall](http://www.wall.org/~aron/blog/pillar-of-science-ii-elegent-hypotheses/)

1. Make it correct.
2. Make it simple and clear.
3. Make it fast.

In that order.

Engineering is not about just hacking the code. It's about evaluating the costs and benefits.

We are writing code for today, designing and architecting for tomorrow.

# Go

> Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming. -- Rob Pike

Don't do object oriented design but data oriented design.

Go prefers convention over configuration because configuration is limiting. Let's group by what one can do not by what one is. In Go there are [no classes](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_1.go) (or [Perl version](https://github.com/jreisinger?tab=repositories&q=animal)) and subclasses but [Interfaces](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/grouping_types_2.go) and composition.

Little copying is better than little dependency. Cost of DRY in Go is bigger than in other languages.
