package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(A []int) []int {
	i, j, k := 0, len(A) - 1, len(A) - 1
	result := make([]int, len(A))
	for i <= j  {
		if abs(A[i]) > abs(A[j]) {
			result[k] = A[i] * A[i]
			i++
			k--
		}else {
			result[k] = A[j] * A[j]
			j--
			k--
		}
	}
	return result
}

func main() {
	A := []int{-7,-3,2,3,11}
	fmt.Println(sortedSquares(A))
}
