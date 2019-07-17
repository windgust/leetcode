package main

import "fmt"

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	nums := []int{2,2,1}
	fmt.Println(singleNumber(nums))
}
