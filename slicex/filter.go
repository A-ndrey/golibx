package slicex

func FilterInPlace[T any, S ~[]T](s S, f func(T) bool) S {
	if len(s) == 0 {
		return s
	}

	var i int
	for _, v := range s {
		if f(v) {
			s[i] = v
			i++
		}
	}

	return s[:i]
}

func Filter[T any, S ~[]T](s S, f func(T) bool) S {
	res := make(S, 0, len(s))
	if len(s) == 0 {
		return res
	}

	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}
