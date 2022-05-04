Difficulties in software engineering:

* essential - inherent in the nature of the software
* accidental - attend software production today but are not inherent

The hard parts of building software are

* specification
* design
* testing of the conceptual construct

not the labor of representing it and testing the fidelity of the representation. We still make syntax errors but they are fuzz compared to the conceptual errors in most systems.

If this is true; building software will be always hard. There is inherently **no silver bullet**.

Inherent properties of the irreducible essence of modern software systems:

Complexity

* for their size, software entities are more complex than any other human construct, because no two parts are alike (at least above the statement level)
* thus scaling-up a software entity is not merely a repetition of the same elements in larger size
* in most cases the elements interact in some nonlinear fashion -> the complexity of the whole increases much more than linearly
* from the complexity come difficulties in: communication among team members (delays, cost overrun, product flaws), understanding all possible states of the program (unreliability, security trapdoors), invoking the functions (hard to use), extending programs to new functions (side effects), management problems

Conformity

Einstein repeatedly argued that there must be simplified explanations of nature, because God is not capricious or arbitrary. No such faith comforts the software engineer. Much of the complexity he must master is arbitrary, forced by the many human institutions and systems to which his interfaces must conform. These differ from interface to interface, and from time to time, not because of necessity but only because they were designed by different people, rather than by God.

Changeability

* software is constantly subject to pressure for change

Invisibility

* geometric abstractions are powerful tools
* but software is invisible and unvisualizable
* this lack not only impedes the process of design within one mind, it severely hinders communication among minds

Past breakthroughs solved accidental difficulties

* high-level languages - they eliminate the complexity of mapping the conceptual constructs (operations, data types, sequences and communication) to concrete machine program (concerned with bits, registers, conditions, branches, channels, disks and such)
* time-sharing - preserves immediacy (in contrast to slow turnaround of batch programming)
* unified programming environments (like Unix) - solve the difficulties of using programs together, by providing integrated libraries, unified file formats, and pipes and filters

---

Source: Mythical Man-Month (1995), ch 16. No Silver Bullet—Essence and Accident in Software Engineering
