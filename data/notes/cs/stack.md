Stack is a simple data structure that allows for two operations:

* push - add an item to the top of the stack
* pop - remove an item from the top of the stack

Call stack

* all function calls go onto the call stack
* when you call a fuction from another function, the calling function is paused in a partially completed state
* each function gets some memory allocated on the stack
* if too much memory is consumed (e.g. when making too many recursive calls) you get stack overflow error
