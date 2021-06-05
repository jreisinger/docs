# Solving problems with software

Solve one problem at a time, don't be overwhelmed by everything.

Steps:

1. Identify concrete tasks that will solve the problem.
2. Write a little [code](https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/composition/decoupling/example1/example1.go) implementing a concrete task.
3. Write tests for your code and deploy ASAP.
4. Repeat 1. - 3. until the problem is solved.
5. [Refactor](https://github.com/ardanlabs/gotraining/tree/master/topics/go/design/composition/decoupling) for change by decoupling.
6. Refactor for readibility by reviewing mental model, naming and cleverness.

Design programs as layers of API:

1. [Primitive](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L66-#L102) layer - do one thing well, not exported, unit tested
2. [Low](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L104-#L147) layer - builds on primitive layer, may be exported, unit tested (might replace primitive layer tests)
3. [High](https://github.com/ardanlabs/gotraining-studyguide/blob/master/go/design/decoupling_1.go#L149-#L166) layer - for ease of use, exported, unit/integration tested 

When am I done:

* 70 - 80% test coverage
* ask what can change, from technical and business perspective, and refactor the code to be able to handle that change

Source: Bill Kennedy's [Ultimate Go Programming](https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/) video course.
