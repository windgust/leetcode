package main

import "fmt"

func moveZeroes(nums []int) {
	var k int
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		nums[k] = nums[i]
		if k != i {
			nums[i] = 0
		}
		k++
	}
}

func main() {
	nums := []int{2, 0, 1, 0, 3, 0, 12, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
