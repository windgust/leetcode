package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	matrixs := make([][]int, 0, R * C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			matrixs = append(matrixs, []int{i,j})
		}
	}
	sort.SliceStable(matrixs, func(i, j int) bool {
		t1 := abs(matrixs[i][0] - r0) + abs(matrixs[i][1] - c0)
		t2 := abs(matrixs[j][0] - r0) + abs(matrixs[j][1] - c0)
		return t1 < t2
	})
	return matrixs
}

func main() {
	fmt.Println(allCellsDistOrder(2,3,1,2))
}