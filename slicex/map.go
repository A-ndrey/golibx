package slicex

func Map[T any, R any](slice []T, f func(int, T) R) []R {
	res := make([]R, len(slice))
	for i, v := range slice {
		res[i] = f(i, v)
	}

	return res
}
