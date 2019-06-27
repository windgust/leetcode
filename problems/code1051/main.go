package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	var reslut int
	tmp := make([]int, len(heights), len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	for i, h := range heights {
		if tmp[i] != h{
			reslut++
		}
	}
	return reslut
}

func main() {
	fmt.Println(heightChecker([]int{1,2,1,2,1,1,1,2,1}))
	//                              1,1,1,1,1,1,2,2,2
}
