package main

import "fmt"

func repeatedNTimes(A []int) int {
	m := make(map[int]struct{})
	for _, a := range A {
		_, ok := m[a]
		if ok {
			return a
		}
		m[a] = struct{}{}
	}
	return 0
}

func main() {
	arr1 := []int{5,1,5,2,5,3,5,4}
	fmt.Println(repeatedNTimes(arr1))
}
