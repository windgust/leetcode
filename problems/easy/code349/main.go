package main

import (
	"fmt"
	"unsafe"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	n1 := make(map[int]struct{})

	for i := range nums1 {
		n1[nums1[i]] = struct{}{}
	}

	for i := range nums2 {
		if _, ok := n1[nums2[i]]; !ok {
			continue
		}
		res = append(res, nums2[i])
		delete(n1, nums2[i])
	}

	return res
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}

	fmt.Println(intersection(nums1,nums2))
	fmt.Println(unsafe.Sizeof(struct {}{}))
}
