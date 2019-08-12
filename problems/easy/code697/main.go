package main

import (
	"fmt"
	"math"
)

type section struct {
	Left  int
	Right int
	Cnt   int
}

func newSection(idx int) *section {
	return &section{
		Left:  idx,
		Right: idx,
		Cnt:   0,
	}
}

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxCnt int
	maxLen := math.MaxInt32
	m := make(map[int]*section)
	for i, num := range nums {
		sec, exist := m[num]
		if !exist {
			sec = newSection(i)
			m[num] = sec
		}
		sec.Right = i
		sec.Cnt++
		if maxCnt == sec.Cnt {
			maxLen = min(maxLen, sec.Right-sec.Left+1)
		}
		if maxCnt < sec.Cnt {
			maxCnt = sec.Cnt
			maxLen = sec.Right - sec.Left + 1
		}
	}
	return maxLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}
	fmt.Println(findShortestSubArray(nums))
}
