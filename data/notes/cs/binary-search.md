```python3
def binary_search(list, item):
    """ binary_search returns the position of 
        the item in the list or None if not found.
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

[Perl](https://github.com/jreisinger/algorithms-with-perl/blob/master/binary-search) ang [Go](https://github.com/jreisinger/go/blob/master/binary-search.go) of binary search.
