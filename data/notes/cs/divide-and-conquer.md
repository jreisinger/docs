Divide & Conquer (D&C)

* relies on recursion
* not an algorithm but a *technique* for solving problems 
* quicksort is a sorting algorithm that uses D&C

To solve a problem using D&C:

1. Figure out the base case. This should be the simplest possible case.
2. Divide or reduce your problem until it becomes the base case.

Tip: If you're writing a recursive function involving an array, the base case is often an empty array or an array with one element. If you're stuck, try that first.

```python
# Sum a list of numbers using a loop.
def sumLoop(l):
    total = 0
    for e in l:
        total += e
    return total

# Sum a list of numbers using Divide & Conquer technique.
def sumDC(l):
    if l == []:
        return 0
    return l[0] + sumDC(l[1:])
```
