package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	maxHeap := MaxHeap(A)
	heap.Init(&maxHeap)
	a, b, c := heap.Pop(&maxHeap).(int), heap.Pop(&maxHeap).(int), heap.Pop(&maxHeap).(int)
	if b+c > a {
		return a + b + c
	}
	for maxHeap.Len() > 0 {
		a, b, c = b, c, heap.Pop(&maxHeap).(int)
		if b+c > a {
			return a + b + c
		}
	}
	return 0
}

func main() {
	A := []int{3, 6, 2, 3}
	fmt.Println(largestPerimeter(A))
}
