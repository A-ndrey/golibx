package slicex

import (
	"strconv"
	"testing"
)

func TestZip(t *testing.T) {
	testData := []struct {
		s1 []string
		s2 []int
		result [] string
	} {
		{
			s1: []string{"a", "b", "c"},
			s2: []int{1, 2, 3},
			result: []string{"a1", "b2", "c3"},
		},
		{
			s1: []string{"a", "b", "c"},
			s2: []int{1, 2},
			result: []string{"a1", "b2", "c1"},
		},
		{
			s1: []string{"a", "b"},
			s2: []int{1, 2, 3},
			result: []string{"a1", "b2", "a3"},
		},
	}

	for _, td := range testData {
		res := make([]string, 0, len(td.result))
		Zip[string, int](td.s1, td.s2, func(s string, i int) {
			res = append(res, s + strconv.Itoa(i))
		})
		if len(res) != len(td.result) {
			t.Fatalf("Different length")
		}
		for i := range res {
			if res[i] != td.result[i] {
				t.Fatalf("Diff on [%d] positioin", i)
			}
		}
	}
}