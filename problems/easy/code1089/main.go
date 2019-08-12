package main

import "fmt"

func duplicateZeros(arr []int) {
	var i, j int
	lenA := len(arr)
	tmp := make([]int, lenA, lenA)
	for i < lenA && j < lenA{
		tmp[j] = arr[i]
		if arr[i] == 0 && j+1 < lenA {
			j++
			tmp[j] = 0
		}
		j++
		i++
	}
	copy(arr, tmp)
}

func main() {
	arr := []int{1,2,3}
	duplicateZeros(arr)
	fmt.Println(arr)
}
