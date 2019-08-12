package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	var res []int
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	fmt.Println(findDisappearedNumbers(nums))
}
