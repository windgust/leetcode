package main

import "fmt"

func lemonadeChange(bills []int) bool {
	if len(bills) == 0 || bills[0] != 5 {
		return false
	}

	keys := [3]int{20, 10, 5}
	vals := make(map[int]int, 3)

	for _, bill := range bills {
		vals[bill]++
		for _, key := range keys {
			for vals[key] > 0 && bill > key {
				bill -= key
				vals[key]--
			}
		}
		if bill != 5 {
			return false
		}
	}
	return true
}

func main() {
	bills := []int{5,5,10, 10, 20}
	fmt.Println(lemonadeChange(bills))
}
