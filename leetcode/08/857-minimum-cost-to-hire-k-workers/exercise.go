package exercise

import (
	"container/heap"
	"sort"
)

type qualityHeap struct {
	pool []int
	p    int
}

func (h *qualityHeap) Len() int           { return h.p }
func (h *qualityHeap) Less(i, j int) bool { return h.pool[i] > h.pool[j] }
func (h *qualityHeap) Swap(i, j int)      { h.pool[i], h.pool[j] = h.pool[j], h.pool[i] }
func (h *qualityHeap) Push(x interface{}) {
	h.pool[h.p] = x.(int)
	h.p++
}
func (h *qualityHeap) Pop() interface{} {
	h.p--
	return h.pool[h.p]
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	indices, ratios := make([]int, n), make([]float64, n)
	for i := 0; i < n; i++ {
		indices[i], ratios[i] = i, float64(wage[i])/float64(quality[i])
	}
	sort.Slice(indices, func(i, j int) bool {
		return ratios[indices[i]] < ratios[indices[j]]
	})

	pq := &qualityHeap{make([]int, k), 0}
	sum := 0
	i := 0
	for ; i < k; i++ {
		heap.Push(pq, quality[indices[i]])
		sum += quality[indices[i]]
	}
	min := float64(sum) * ratios[indices[i-1]]
	for ; i < n; i++ {
		if quality[indices[i]] > pq.pool[0] {
			continue
		}
		sum += quality[indices[i]] - heap.Pop(pq).(int)
		heap.Push(pq, quality[indices[i]])
		if res := float64(sum) * ratios[indices[i]]; res < min {
			min = res
		}
	}

	return min
}
