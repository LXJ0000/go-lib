package queue

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestPriorityQueue1(t *testing.T) {
	nums := []int{20, 17, 3, 12, 5, 9, 7, 8, 6, 15, 11, 4, 13, 14, 10, 16, 2, 18, 19, 1}
	q := NewPriorityQueue[int]()
	for _, num := range nums {
		q.Push(num)
	}
	res := []int{}
	for q.Size() > 0 {
		res = append(res, q.Pop())
	}
	got := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	assert.Equal(t, got, res)
}

func TestPriorityQueue2(t *testing.T) {
	nums := []int{20, 17, 3, 12, 5, 9, 7, 8, 6, 15, 11, 4, 13, 14, 10, 16, 2, 18, 19, 1}
	q := NewPriorityQueue[int]()
	q = WithMaxHeap(q)
	for _, num := range nums {
		q.Push(num)
	}
	res := []int{}
	for q.Size() > 0 {
		res = append(res, q.Pop())
	}
	got := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slices.Reverse(got)
	assert.Equal(t, got, res)
}
