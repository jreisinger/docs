Divide & Conquer (D&C)

* relies on recursion
* not an algorithm but a *technique* for solving problems 
* quicksort is a sorting algorithm that uses D&C

To solve a problem using D&C:

1. Figure out the base case. This should be the simplest possible case.
2. Divide or reduce your problem until it becomes the base case.

Tip: If you're writing a recursive function involving an array, the base case is often an empty array or an array with one element. If you're stuck, try that first.

Python:

```
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

Go (with tests):

```
// dc_test.go - sum a list of integers using D&C technique and test it
// Usage: go test -v
package main

import "testing"

func TestSumDC(t *testing.T) {
	type testpair struct {
		ints []int
		sum  int
	}
	testpairs := []testpair{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1, 3, 2}, 6},
	}
	for _, tp := range testpairs {
		got := SumDC(tp.ints)
		if got != tp.sum {
			t.Errorf("got %v want %v", got, tp.sum)
		}
	}
}

func SumDC(ints []int) int {
	if len(ints) == 0 { // base case
		return 0
	}
	return ints[0] + SumDC(ints[1:])
}
```
