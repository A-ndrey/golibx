package queue

type Interface[T any] interface {
	Peek() T
	Pop() T
	Push(T)
	Size() int
}

type Queue[T any] struct {
	first int
	last  int
	s     []T
	empty bool
}

func New[T any]() *Queue[T] {
	q := Queue[T]{
		empty: true,
	}

	return &q
}

func (q *Queue[T]) Peek() T {
	if q.empty {
		var t T
		return t
	}

	return q.s[q.first]
}

func (q *Queue[T]) Pop() T {
	var t T
	if q.empty {
		return t
	}

	t = q.s[q.first]

	q.first++
	q.first = q.first % len(q.s)

	if q.first == q.last {
		q.empty = true
	}

	return t
}

func (q *Queue[T]) Push(t T) {
	if q.needGrow() {
		q.grow()
	}

	q.s[q.last] = t
	q.last = (q.last + 1) % len(q.s)

	q.empty = false
}

func (q *Queue[T]) Size() int {
	if q.empty {
		return 0
	}

	diff := q.last - q.first
	if diff > 0 {
		return diff
	}

	return diff + len(q.s)
}

func (q *Queue[T]) needGrow() bool {
	return q.first == q.last
}

func (q *Queue[T]) grow() {
	size := len(q.s)
	if size == 0 {
		size = 4
		q.s = make([]T, size)
		return
	}

	size = int(float64(size) * 1.75)

	newS := make([]T, size)
	for i := 0; i < len(q.s); i++ {
		newS[i] = q.s[q.first]
		q.first = (q.first + 1) % len(q.s)
	}

	q.last = len(q.s)
	q.first = 0
	q.s = newS
}
