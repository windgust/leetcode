package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	aCandy := make(map[int]struct{})

	for i := range A {
		sumA += A[i]
		aCandy[A[i]] = struct{}{}
	}
	for i := range B {
		sumB += B[i]
	}

	diff := (sumA - sumB) / 2
	for i := range B {
		if _, ok := aCandy[B[i]+diff]; ok {
			return []int{B[i] + diff, B[i]}
		}
	}
	return nil
}

func main() {
	A := []int{1,2,5}
	B := []int{2,4}
	fmt.Println(fairCandySwap(A, B))
}
