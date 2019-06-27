package main

import "fmt"

func minDeletionSize(A []string) int {
	lenStr := len(A[0])
	reslut := 0
	for i := 0; i < lenStr; i++ {
		for j := 0; j < len(A) - 1; j++ {
			if A[j][i] > A[j + 1][i] {
				reslut++
				break
			}
		}
	}
	return reslut
}

func main() {
	fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"}))
}
