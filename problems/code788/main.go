package main

import "fmt"

func rotatedDigits(N int) int {
	var res int

	for i := 1; i <= N; i++ {
		var valid bool
		k := i
		for k > 0 {
			x := k % 10
			k /= 10
			if x == 3 || x == 4 || x == 7 {
				valid = false
				break
			}

			if x == 0 || x == 1 || x == 8 {
				continue
			}
			valid = true
		}
		if valid {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(rotatedDigits(857))
}
