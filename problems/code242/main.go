package main

import "fmt"

func isAnagram(s string, t string) bool {
	var letterCnt [26]int
	for i := range s {
		letterCnt[s[i]-'a']++
	}
	for i := range t {
		letterCnt[t[i]-'a']--
	}
	for i := range letterCnt {
		if letterCnt[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := ""
	t := "car"
	fmt.Println(isAnagram(s, t))
}
