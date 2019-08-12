package main

import "fmt"

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	return N ^ onlyOne(N)
}

func onlyOne(N int) int {
	var one int
	for N > 0  {
		N = N >> 1
		one = (one << 1) + 1
	}
	return one
}

func main() {
	fmt.Println(bitwiseComplement(10))
}