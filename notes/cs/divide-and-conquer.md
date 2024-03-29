Divide and Conquer (DaC)

* not an algorithm but a *technique* for solving problems 
* quicksort is a sorting algorithm that uses DaC
* relies on recursion

To solve a problem using DaC:

1. Figure out the base case. This should be the simplest possible case.
2. Divide or reduce your problem until it becomes the base case.

TIP: If you're writing a recursive function involving an array, the base case is often an empty array or an array with one element. If you're stuck, try that first.

```python
# Sum a list of numbers using a loop.
def sumLoop(l):
    sum = 0
    for e in l:
        sum += e
    return sum

# Sum a list of numbers using DaC.
def sumDaC(l):
    if l == []:
        return 0
    return l[0] + sumDaC(l[1:])
```
