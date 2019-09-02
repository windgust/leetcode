package main

func twoSum(numbers []int, target int) []int {
	i, j := 1, len(numbers)

	for i < j {
		s := numbers[i-1] + numbers[j-1]
		if s == target {
			break
		}
		if s < target {
			i++
		}
		if s > target {
			j--
		}
	}
	return []int{i, j}
}
