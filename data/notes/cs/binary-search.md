A.k.a. bisection search.

```python3
def binary_search(list, item):
    """ binary_search returns the position of item 
        in a sorted list or None if not found.
    """

    # Which part of the list to search. At the 
    # beginning it's the entire list.
    low = 0
    high = len(list) - 1

    while low <= high:
        mid = int((low + high) / 2)
        guess = list[mid]
        if guess == item:
            return mid
        if guess > item:    # guess was too high
            high = mid - 1
        else:               # guess was too low
            low = mid + 1
    
    return None
```

My [Perl](https://github.com/jreisinger/algorithms-with-perl/blob/master/binary-search) ang [Go](https://github.com/jreisinger/go/blob/master/binary-search.go) implementation of binary search.

More

* [MIT OpenCourseWare](https://www.youtube.com/watch?v=SE4P7IVCunE&list=PLUl4u3cNGP63WbdFxL8giv4yhgdMGaZNA&index=11) (video)
