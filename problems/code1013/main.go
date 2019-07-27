package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	var sum int
	for i := range A {
		sum += A[i]
	}

	if sum%3 != 0 {
		return false
	}

	mid := sum / 3
	var s, cnt int
	for i := range A {
		s += A[i]
		if s == mid {
			cnt++
			s = 0
		}
	}
	return s == 0 && cnt == 3
}

func main() {
	A := []int{3,3,6,5,-2,2,5,1,-9,4}
	fmt.Println(canThreePartsEqualSum(A))
}
