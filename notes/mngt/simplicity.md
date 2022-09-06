Simplicity, [so important](https://user-images.githubusercontent.com/1047259/175539790-bcdfb1f1-560c-4204-904c-8def089a6416.png).

Complexity is a problem because complex systems are hard to:

* understand and reason about
* modify (fix bugs, add features)
* operate

"Complexity is multiplicative" (Rob Pike): fixing a problem by making one part of the system more complex slowly but surely adds complexity to other parts.

Complexity is caused by

* constant pressure to add features quickly
* inexperience of the system designer/implementer

To fight complexity

* say no to production stuff that is not really needed
* observe, read and think about good and bad designs
* produce good and bad designs, "Err and err and err again but less and less and less." (Piet Hein)
* at the beginning of a project reduce an idea to its essence, producing the most simple design possible
* have discipline over the lifetime of a project to distinguish good changes from bad ones
  * good changes don't break conceptual integrity of the design
  * bad changes trade simplicity for its shallow cousin, convenience
* review and simplify whenever possible (software is a very tractable medium)
* hide (encapsulate) necessary complexity using modules with simple interfaces

NOTE: Fred Brooks in [No Silver Bullet](https://en.wikipedia.org/wiki/No_Silver_Bullet) distinguishes between accidental complexity and essential complexity. I think essential complexity is what I mean by necessary complexity.
