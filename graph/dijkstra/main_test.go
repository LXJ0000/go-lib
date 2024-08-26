package main

import (
	"testing"

	"github.com/LXJ0000/go-lib/queue"
)

func TestDijkstra(t *testing.T) {
	graph := [][][2]int{
		{},
		{{2, 2}, {4, 3}},
		{{1, 3}},
		{},
	}
	n := len(graph) - 1
	st := make([]bool, n+1)
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = 0x3f3f3f3f
	}

	queue := queue.NewPriorityQueue(func(first, second [2]int) bool {
		return first[0] < second[0]
	})
	queue.Push([2]int{0, 1})
	dist[1] = 0

	for queue.Size() != 0 {
		top := queue.Pop()
		node, distance := top[1], top[0]
		if st[node] {
			continue
		}
		st[node] = true
		for _, tmp := range graph[node] {
			w, v := tmp[0], tmp[1]
			if dist[v] > distance+w {
				dist[v] = distance + w
				queue.Push([2]int{dist[v], v})
			}
		}
	}
	t.Error(dist[1:])
}
