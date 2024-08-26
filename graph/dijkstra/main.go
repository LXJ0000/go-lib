package main

import (
	"container/heap"
	"fmt"
)

func Dijkstra(graph [][][2]int) int {
	n := len(graph) - 1
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = 0x3f3f3f3f
	}
	st := make([]bool, n+1)
	h := &minHeap{}

	heap.Push(h, [2]int{0, 1})
	dist[1] = 0

	for h.Len() != 0 {
		// t := h.Top().([2]int)
		t := heap.Pop(h).([2]int)
		node := t[1]
		dis := t[0]
		if st[node] {
			continue
		}
		st[node] = true
		for _, tmp := range graph[node] {
			w, v := tmp[0], tmp[1]
			if dis+w < dist[v] {
				dist[v] = dis + w
				heap.Push(h, [2]int{dist[v], v})
			}
		}
	}
	return dist[n]
}

func main() {
	// var n, m int
	// var graph [][][2]int
	// fmt.Scan(&n, &m)
	// for i := 0; i <= n; i++ {
	// 	graph = append(graph, [][2]int{})
	// }
	// for i := 0; i < m; i++ {
	// 	var u, v, w int
	// 	fmt.Scan(&u, &v, &w)
	// 	graph[u] = append(graph[u], [2]int{w, v})
	// }
	// fmt.Println(graph[0])
	// fmt.Println(graph[1])
	// fmt.Println(graph[2])
	// fmt.Println(graph[3])

	// []
	// [[2 2] [4 3]]
	// [[1 3]]
	// []
	graph := [][][2]int{
		{},
		{{2, 2}, {4, 3}},
		{{1, 3}},
		{},
	}
	fmt.Println(Dijkstra(graph))
}

type minHeap [][2]int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *minHeap) Top() interface{} {
	return (*h)[0]
}
