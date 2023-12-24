package slicex

func Zip[T1 any, T2 any](s1 []T1, s2 []T2, f func(T1, T2)) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}

	maxLen := max(len(s1), len(s2))

	for i := 0; i < maxLen; i++ {
		f(s1[i%len(s1)], s2[i%len(s2)])
	}
}
