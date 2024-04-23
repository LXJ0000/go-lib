package queue

import "github.com/LXJ0000/go-lib"

type PriorityQueue[T lib.Ordered] struct {
	data    []T
	size    int
	maxHeap bool
}

func NewPriorityQueue[T lib.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data:    []T{0},
		size:    0,
		maxHeap: false,
	}
}

func WithMaxHeap[T lib.Ordered](q *PriorityQueue[T]) *PriorityQueue[T] {
	q.maxHeap = true
	return q
}

func (h *PriorityQueue[T]) down(u int) {
	t, n := u, h.size

	check := func(a, b T) bool {
		return a > b
	}
	if h.maxHeap {
		check = func(a, b T) bool {
			return a < b
		}
	}

	if u*2 <= n && check(h.data[t], h.data[u*2]) {
		t = u * 2
	}

	if u*2+1 <= n && check(h.data[t], h.data[u*2+1]) {
		t = u*2 + 1
	}
	if t != u {
		h.data[u], h.data[t] = h.data[t], h.data[u]
		h.down(t)
	}
}

func (h *PriorityQueue[T]) up(u int) {
	check := func(a, b T) bool {
		return a < b
	}
	if h.maxHeap {
		check = func(a, b T) bool {
			return a > b
		}
	}

	for u/2 != 0 && check(h.data[u], h.data[u/2]) {
		h.data[u], h.data[u/2] = h.data[u/2], h.data[u]
		u >>= 1
	}
}

func (h *PriorityQueue[T]) Push(x T) {
	h.data = append(h.data, x)
	h.size++
	h.up(h.size)
}

func (h *PriorityQueue[T]) Pop() T {
	move := h.data[1]
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.size--
	h.down(1)
	return move
}

func (h *PriorityQueue[T]) Front() T {
	return h.data[1]
}

func (h *PriorityQueue[T]) Size() int {
	return h.size
}
func (h *PriorityQueue[T]) Empty() bool {
	return h.size == 0
}
