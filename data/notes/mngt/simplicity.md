![image](https://user-images.githubusercontent.com/1047259/175539790-bcdfb1f1-560c-4204-904c-8def089a6416.png)

Simplicity, so important.

Complexity is a problem because complex systems are hard to *understand* and thus hard to

* use
* modify (fix bugs, add features)
* operate

Complexity is caused by

* inexperience of the system designer/implementer
* constant pressure to add features quickly

"Complexity is multiplicative" (Rob Pike) - fixing a problem by making one part of the system more complex slowly but surely adds complexity to other parts.

To fight complexity

* read and think about, observe and produce good and bad designs ("Err and err and err again but less and less and less.")
* say no to modifications that are not really needed (you need some balance here otherwise you won't learn)
* review and simplify whenever possible (software is a very tractable medium)
* at the beginning of a project reduce an idea to its essence, producing the most simple design possible
* have discipline over the lifetime of a project to distinguish good changes from bad ones
  * good changes don't break conceptual integrity of the design
  * bad changes trade simplicity for its shallow cousin, convenience
* hide (encapsulate) necessary complexity using modules with simple interfaces
