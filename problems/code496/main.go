package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
	index := make(map[int]int)

	for i, num := range nums2 {
		index[num] = i
	}

	for i := range nums1 {
		for j := index[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				break
			}
		}
		if len(res) != i + 1 {
			res = append(res, -1)
		}
	}
	return res
}

func main() {
	num1 := []int{3,1,5,7,9,2,6}
	num2 := []int{1,2,3,5,6,7,9,11}
	fmt.Println(nextGreaterElement(num1, num2))
}
