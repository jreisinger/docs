# Solving problems with software

Solve one problem at a time, don't be overwhelmed by everything.

Steps:

1. Identify concrete tasks that will solve the problem.
2. Write a little code implementing a concrete task.
3. Write tests for your code and deploy ASAP.
4. Repeat 1. - 3. until the problem is solved.
5. Refactor for change by [decoupling](https://github.com/ardanlabs/gotraining/tree/master/topics/go/design/composition/decoupling).
6. Refactor for readibility by reviewing mental model, naming and cleverness.

Design programs as layers of API:

1. Primitive layer - do one thing well, not exported, unit tested
2. Low layer - builds on primitive layer, may be exported, unit tested (might replace primitive layer tests)
3. High layer - for ease of use, exported, unit/integration tested 

When am I done:

* 70 - 80% test coverage
* ask what can change, from technical and business perspective, and refactor the code to be able to handle that change

Source: Bill Kennedy's [Ultimate Go Programming](https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/) video course.
