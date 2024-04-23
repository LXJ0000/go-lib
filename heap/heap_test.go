package heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{9, 8, 7, 6, 1, 2, 3, 4}
	h := &binHeap{}
	for _, num := range nums {
		heap.Push(h, num)
	}
	for h.Len() > 0 {
		t.Log(h.Top())
		heap.Pop(h)
	}
}
