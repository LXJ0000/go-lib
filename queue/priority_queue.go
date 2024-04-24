package queue

type PriorityQueue[T interface{}] struct {
	data    []T
	size    int
	compare func(first, second T) bool
}

func NewPriorityQueue[T interface{}](compare func(first, second T) bool) *PriorityQueue[T] {
	var zero T
	return &PriorityQueue[T]{
		data:    []T{zero},
		size:    0,
		compare: compare,
	}
}

func (h *PriorityQueue[T]) down(u int) {
	t, n := u, h.size

	if u*2 <= n && h.compare(h.data[u*2], h.data[t]) {
		t = u * 2
	}

	if u*2+1 <= n && h.compare(h.data[u*2+1], h.data[t]) {
		t = u*2 + 1
	}
	if t != u {
		h.data[u], h.data[t] = h.data[t], h.data[u]
		h.down(t)
	}
}

func (h *PriorityQueue[T]) up(u int) {
	for u/2 != 0 && h.compare(h.data[u], h.data[u/2]) {
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
