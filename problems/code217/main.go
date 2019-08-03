package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}
