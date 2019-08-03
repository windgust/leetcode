package main

import "fmt"

func isMonotonic(A []int) bool {
	inc, dec := true, true
	for i := 0; i < len(A)-1 && (inc || dec); i++ {
		inc = inc && A[i] >= A[i+1]
		dec = dec && A[i] <= A[i+1]
	}
	return inc || dec

	//if len(A) <= 1 {
	//	return true
	//}
	//var dist int
	//for i := 0; i < len(A) - 1; i++ {
	//	if A[i] == A[i+1] {
	//		continue
	//	}
	//	if dist == 0 {
	//		if A[i] < A[i+1] {
	//			dist = -1
	//		}else {
	//			dist = 1
	//		}
	//	}
	//	if dist ^ (A[i] - A[i+1]) < 0 {
	//		return false
	//	}
	//}
	//return true
}

func main() {
	A := []int{1, 2, 4, 7, 6, 6}
	fmt.Println(isMonotonic(A))
}
