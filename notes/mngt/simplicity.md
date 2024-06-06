Simplicity, [so important](https://user-images.githubusercontent.com/1047259/175539790-bcdfb1f1-560c-4204-904c-8def089a6416.png).

Complexity is a problem because complex systems are hard to *understand and reason about* and therefore hard to:

* secure
* operate
* modify (fix bugs, add features, refactor)
* use (if API/UI is complicated)

Complexity is caused by

* constant pressure to add features quickly
* inexperience of the system designer/implementer
* bad mindset with this reasoning
  * we're not beginners
  * real life is complicated
  * if it was hard to write, it should be hard to read
  * if it was easy, anyone could do it

To fight complexity

* say no to production stuff that is not really needed ([no code](https://github.com/kelseyhightower/nocode) means no complexity :-)
* observe, read and think about good and bad designs
* produce good and bad designs (Piet Hein says "Err and err and err again but less and less and less.")
* at the beginning of a project reduce an idea to its essence, producing the most simple design possible
* have discipline over the lifetime of a project to distinguish good changes from bad ones
  * good changes don't break conceptual integrity of the design
  * bad changes trade simplicity for its shallow cousin, convenience
* review and simplify whenever possible (software is a very tractable medium)
* hide (encapsulate) necessary complexity using modules with simple interfaces
* strive for modularity with loose coupling to be able to make changes to parts of the system in isolation

Fred Brooks in [No Silver Bullet](https://en.wikipedia.org/wiki/No_Silver_Bullet) distinguishes between accidental complexity and essential complexity. Essential complexity is caused by the problem to be solved, and nothing can remove it; if users want a program to do 30 different things, then those 30 things are essential and the program must do those 30 different things.

Rob Pike says complexity is multiplicative: fixing a problem by making one part of the system more complex slowly but surely adds complexity to other parts.

More

* https://sre.google/sre-book/simplicity/
