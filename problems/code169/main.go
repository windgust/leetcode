package main

import "fmt"

func majorityElement(nums []int) int {
	res, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		}
	}
	return res
}

func main() {
	nums := []int{2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))
}