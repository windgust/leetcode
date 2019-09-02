package main

import "fmt"

func firstUniqChar(s string) int {
	var letters [26]int

	for _, ch := range s {
		letters[ch-'a']++
	}

	for i, ch := range s {
		if letters[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcodeleetcode"))
}
