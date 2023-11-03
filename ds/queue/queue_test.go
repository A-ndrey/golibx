package queue

import "testing"

func TestQueue(t *testing.T) {
	checkSize := func(t *testing.T, exp, act int) {
		if exp != act {
			t.Errorf("Size: expected=%d, actual=%d", exp, act)
		}
	}

	checkElement := func(t *testing.T, exp, act int) {
		if exp != act {
			t.Errorf("Element: expected=%d, actual=%d", exp, act)
		}
	}

	t.Run("1 element, only push", func(t *testing.T) {
		q := New[int]()
		q.Push(1)
		checkSize(t, 1, q.Size())
	})

	t.Run("1 element, push and pop", func(t *testing.T) {
		q := New[int]()
		q.Push(1)
		q.Push(2)
		e := q.Pop()
		checkElement(t, 1, e)
		checkSize(t, 1, q.Size())
	})

	t.Run("4 elements, only push", func(t *testing.T) {
		q := New[int]()
		q.Push(1)
		q.Push(2)
		q.Push(1)
		q.Push(2)
		checkSize(t, 4, q.Size())
	})

	t.Run("5 elements, only push", func(t *testing.T) {
		q := New[int]()
		q.Push(1)
		q.Push(2)
		q.Push(1)
		q.Push(2)
		q.Push(2)
		checkSize(t, 5, q.Size())
	})

	t.Run("5 elements, push and pop", func(t *testing.T) {
		q := New[int]()
		q.Push(1)
		q.Push(2)
		q.Push(1)
		e := q.Pop()
		checkElement(t, 1, e)
		q.Push(2)
		q.Push(2)
		e = q.Pop()
		checkElement(t, 2, e)
		q.Push(3)
		q.Push(3)
		checkSize(t, 5, q.Size())
	})

	t.Run("0(10) elements, push and pop", func(t *testing.T) {
		q := New[int]()
		for i := 1; i <= 10; i++ {
			q.Push(i)
		}
		checkSize(t, 10, q.Size())
		for i := 1; i <= 10; i++ {
			checkElement(t, i, q.Pop())
		}
		checkSize(t, 0, q.Size())
	})

	t.Run("0 elements, only pop", func(t *testing.T) {
		q := New[int]()
		checkElement(t, 0, q.Pop())
		checkSize(t, 0, q.Size())
	})
}
