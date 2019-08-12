package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var maxCnt, cnt int

	for i := range nums {
		if nums[i] == 0 {
			cnt = 0
			continue
		}
		if nums[i] == 1 {
			cnt++
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}

func main() {
	nums := []int{1,1,0,0,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

