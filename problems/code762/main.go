package main

import (
	"fmt"
	"math"
	"math/bits"
)

func countPrimeSetBits(L int, R int) int {
	var res int
	for i := L; i <= R; i++ {
		res += isPrime(bits.OnesCount(uint(i)))
	}
	return res
}

func isPrime(x int) int {
	if x == 1 {
		return 0
	}
	if x == 2 || x == 3 {
		return 1
	}
	if x%6 != 1 && x%6 != 5 {
		return 0
	}

	tmp := math.Sqrt(float64(x))
	for i := 5; i <= int(tmp); i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	fmt.Println(countPrimeSetBits(1, 3))
}
