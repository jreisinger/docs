To understand recursion, you must first understand recursion.

Recursion
* should be used when it makes the solution clearer
* there is no performance benefit to using recursion; in fact loops are sometimes better for performance

Every recursive function should have two parts:

* base case - the function doesn't call itself
* recursive case - the function calls itself

Wrong:

```python
def countdown(i):
    # no base case -> infinite loop
    print(i)
    countdown(i-1)
```

Right:

```
def countdown(i):
    print(i)
    if i <= 0:    # base case
        return
    else:         # recursive case
        countdown(i-1)
```

The call stack with recursions (from "Grokking Algorithms")

```
def fact(x):
    if x == 1:
        return 1
    else:
        return x * fact(x-1)
```

<img src="https://user-images.githubusercontent.com/1047259/132297124-1d1f431e-22c9-423e-8651-ec9365580d02.png" style="max-width:100%;height:auto;"> 
