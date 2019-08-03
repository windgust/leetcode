package main

import "fmt"

func countBinarySubstrings(s string) int {
	var cnt int
	var zeroAndOneCnt [2]int
	prev := s[0]
	for i := range s {
		if s[i] == prev {
			zeroAndOneCnt[s[i]-'0']++
		}
		if s[i] != prev {
			zeroAndOneCnt[s[i]-'0'] = 1
		}
		if s[i] == '0' && zeroAndOneCnt[0] <= zeroAndOneCnt[1] ||
			s[i] == '1' && zeroAndOneCnt[1] <= zeroAndOneCnt[0] {
			cnt++
		}
		prev = s[i]
	}
	return cnt
}

func main() {
	s := "01100"
	fmt.Println(countBinarySubstrings(s))
}
