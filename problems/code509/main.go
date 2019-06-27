package main

import "fmt"

func fib(N int) int {
	first, second := 0, 1
	if N == 0 {
		return first
	}
	if N == 1 {
		return second
	}

	for i := 2; i <= N; i++  {
		first, second = second, first + second
	}
	return second
}

func main() {
	fmt.Println(fib(2))
}
