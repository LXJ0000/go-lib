package heap

type binHeap []int

func (h binHeap) Len() int           { return len(h) }
func (h binHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h binHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *binHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *binHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *binHeap) Top() interface{} {
	return (*h)[0]
}
