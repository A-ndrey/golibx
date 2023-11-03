package slicex

import "testing"

func TestFilterInPlace(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	even := func(v int) bool {
		return v%2 == 0
	}

	res := FilterInPlace(src, even)

	if len(res) != 2 || res[0] != 2 || res[1] != 4 || src[0] != res[0] || src[1] != res[1] {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	even := func(v int) bool {
		return v%2 == 0
	}

	res := Filter(src, even)

	if len(res) != 2 || res[0] != 2 || res[1] != 4 || src[0] == res[0] || src[1] == res[1] {
		t.Fail()
	}
}
