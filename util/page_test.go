package util

import "testing"

func TestTitleCasePathComponents(t *testing.T) {
	type testpair struct {
		in  string
		out []string
	}
	testpairs := []testpair{
		{"./data/notes/go/oo", []string{"./Data/notes/go/oo", "./data/Notes/go/oo", "./data/notes/Go/oo", "./data/notes/go/Oo"}},
	}
	for _, tp := range testpairs {
		got := titleCasePathComponents(tp.in)
		for i := range got {
			if got[i] != tp.out[i] {
				t.Fatalf("got %s but wanted %s", got[i], tp.out[i])
			}
		}
	}
}
