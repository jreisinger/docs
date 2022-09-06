A.k.a. bisection search.

```python3
def binary_search(list, item):
    """ binary_search returns the position (index) of 
        item in a SORTED list or None if not found.
    """

    # Which part of the list to search. At the 
    # beginning it's the entire list.
    low = 0
    high = len(list) - 1

    while low <= high:          # While window is open
        mid = int((low+high)/2) # try middle element.
        guess = list[mid]
        if guess == item:       # we've found the item
            return mid
        if guess > item:        # guess was too high
            high = mid - 1
        else:                   # guess was too low
            low = mid + 1
    
    return None                 # the item not found
```

More

* [MIT OpenCourseWare](https://www.youtube.com/watch?v=SE4P7IVCunE&list=PLUl4u3cNGP63WbdFxL8giv4yhgdMGaZNA&index=11) (video)
* My [Perl](https://github.com/jreisinger/algorithms-with-perl/blob/master/binary-search) and [Go](https://github.com/jreisinger/go/blob/master/binary-search.go) implementation
