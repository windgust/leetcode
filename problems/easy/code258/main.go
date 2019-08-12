package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	res := num % 9
	if res == 0 {
		return 9
	}
	return res
}

func main() {
	fmt.Println(addDigits(38))
}
