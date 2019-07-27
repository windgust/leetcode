package main

import "fmt"

func canWinNim(n int) bool {
	if uint8(n<<6) == 0 {
		return false
	}
	return true
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i, canWinNim(i))
	}
}
