package main

import "fmt"

func divisorGame(N int) bool {
	if (N & 1) == 1  {
		return false
	}
	return true
}

func main() {
	fmt.Println(divisorGame(4))
}
