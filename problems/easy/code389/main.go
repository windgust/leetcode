package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var res byte
	for i := range s {
		res ^= s[i]
	}
	for i := range t {
		res ^= t[i]
	}
	return res
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(string(findTheDifference(s, t)))
}
