package main

import "fmt"

func titleToNumber(s string) int {
	var res int
	k := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i] - 'A' + 1) * k
		k *= 26
	}
	return res
}

func main() {
	s := "B"
	fmt.Println(titleToNumber(s))
}
