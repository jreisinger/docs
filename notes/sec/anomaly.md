# The anomaly of cheap complexity

(This is my summary of Andrew Appel's [sunmmary](https://freedom-to-tinker.com/2022/08/03/the-anomaly-of-cheap-complexity/) of Thomas Dullien's [talk](http://rule11.tech/papers/2018-complexitysecuritysec-dullien.pdf).)

> How does one design an electric motor? Would you attach a bathtub to it, simply because one was available? Would a bouquet of flowers help? A heap of rocks? No, you would use just those elements necessary to its purpose and make it no larger than needed -- and you would incorporate safety factors. Function controls design.

-- Prof. Bernardo de la Paz in The Moon Is A Harsh Mistress (Robert A. Heinlein)

Why are computer systems so insecure?

The reason is that they have so many complex layers. Why there are so many layers, and why those layers are so complex-even for what should be a simple thing like counting up votes?

It's because complexity is cheap. For most of human history, a more complex device was more expensive to build than a simpler device. This is not the case in modern computing. It is often more cost-effective to take a very complicated device or software, and make it simulate simplicity, than to make a simpler one. 

You need a machine that does something. Complex general-purpose CPUs are cheap (economies of scale and Moore's law). ARM Cortex-M0 CPUs cost pennies. Software specializes a CPU that could do anything (universal computation) to become a device that does something.

A (huge and complex) general-purpose operating system is free, but a simpler, custom-designed, perhaps more secure OS would be very expensive to build.

