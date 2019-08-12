package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var sum int
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}

func main() {
	prices := []int{2, 1, 2}
	fmt.Println(maxProfit(prices))
}
