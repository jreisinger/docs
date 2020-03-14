See also https://github.com/jonatasbaldin/grokking-algorithms-golang/blob/master/ch2/selection_sort_test.go

```
// selection_sort_test.go
package main

import "testing"

func findIndexOfSmallest(s []int) int {
	var smallest int
	var smallestIdx int
	for i, v := range s {
		if i == 0 {
			smallest = v
			smallestIdx = i
			continue
		}
		if s[i] < smallest {
			smallest = s[i]
			smallestIdx = i
		}
	}
	return smallestIdx
}

func selectionSort(s []int) []int {
	var sorted []int
	for range s {
		i := findIndexOfSmallest(s)
		sorted = append(sorted, s[i])
		// remove the item just appended
		s = append(s[:i], s[i+1:]...)
	}
	return sorted
}

func TestFindIndexOfSmallest(t *testing.T) {
	type testpair struct {
		s []int
		i int
	}

	tp := []testpair{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{-1}, 0},
		{[]int{1000}, 0},
		{[]int{-1000}, 0},
		{[]int{0, 1}, 0},
		{[]int{-1, 1}, 0},
		{[]int{1, -1}, 1},
		{[]int{1, 1}, 0},
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 100}, 0},
		{[]int{100, 1, 100}, 1},
		{[]int{1, 100, 100}, 0},
		{[]int{1, 2, 3}, 0},
		{[]int{3, 2, 1}, 2},
		{[]int{42, 0, -1, 100, 13}, 2},
	}

	for _, p := range tp {
		i := findIndexOfSmallest(p.s)
		if p.i != i {
			t.Errorf("findIndexOfSmallest: from %v wanted %d, got %d", p.s, p.i, i)
		}
	}
}

func TestSlectionSort(t *testing.T) {
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
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 1}, []int{1, 1, 2}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{-1, 1, -1}, []int{-1, -1, 1}},
		{[]int{100, -100}, []int{-100, 100}},
	}

	for _, p := range tp {
		sorted := selectionSort(p.s)
		if !egual(p.sorted, sorted) {
			t.Errorf("selectionSort: wanted %v, got %v", p.sorted, sorted)
		}
	}
}

// --- helper functions ---

func egual(a, b []int) bool {
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
