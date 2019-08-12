package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	if mh[i] < mh[j] {
		return false
	}
	return true
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	var maxHeap MaxHeap = MaxHeap(stones)
	var x, y int
	heap.Init(&maxHeap)
	for maxHeap.Len() > 1 {
		x = heap.Pop(&maxHeap).(int)
 		y = heap.Pop(&maxHeap).(int)
		if x == y {
			continue
		}
 		heap.Push(&maxHeap, x - y)
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(&maxHeap).(int)
}

func main() {
	stones := []int{2, 2}
	fmt.Println(lastStoneWeight(stones))
}
