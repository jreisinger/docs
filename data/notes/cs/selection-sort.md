* your computer's memory is like a giant set of drawers
* when you want to store multiple elements, use an array or a linked list

arrays

* all elements are stored right next to each other
* allow for fast reads

linked lists

* elements are distributed all over the memory, and one element stores the address of the next one
* allow for fast inserts and deletes

This is an implementation in Go of the selection sort from ch2 of Grokking Algorithms ([here](https://github.com/jonatasbaldin/grokking-algorithms-golang/blob/master/ch2/selection_sort_test.go) is a similar piece):

```
// selection_sort_test.go
package main

import "testing"

func SelectionSort(s []int) []int {
	var sorted []int
	for range s {
		i := findIndexOfSmallest(s)
		sorted = append(sorted, s[i])
		// remove the item just appended
		s = append(s[:i], s[i+1:]...)
	}
	return sorted
}

func findIndexOfSmallest(s []int) int {
	var smallest int
	var smallestIdx int
	for i, v := range s {
		if i == 0 {
			smallest = v
			smallestIdx = i
			continue
		}
		if v < smallest {
			smallest = v
			smallestIdx = i
		}
	}
	return smallestIdx
}

// --- tests ---

func TestSelectionSort(t *testing.T) {
	type testpair struct {
		s      []int
		sorted []int
	}

	tp := []testpair{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{0}, []int{0}},
		{[]int{-1}, []int{-1}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{-1, 1, -1}, []int{-1, -1, 1}},
		{[]int{100, -100}, []int{-100, 100}},
		{[]int{42, 0, -1, 100, 13}, []int{-1, 0, 13, 42, 100}},
	}

	for _, p := range tp {
		sorted := SelectionSort(p.s)
		if !equals(p.sorted, sorted) {
			t.Errorf("selectionSort: wanted %v, got %v",
				p.sorted, sorted)
		}
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
```
