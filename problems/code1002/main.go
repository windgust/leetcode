package main

import (
	"fmt"
	"math"
)

func commonChars(A []string) []string {
	tmp := make([][26]int, len(A), len(A))
	reslut := make([]string, 0, 0)
	for i := 0; i < len(A); i++ {
		for _, c := range A[i] {
			tmp[i][c - 'a']++
		}
	}
	for i := 0; i < 26; i++  {
		min := math.MaxInt32
		for _, t := range tmp {
			if t[i] < min {
				min = t[i]
			}
		}
		if min != 0 && min != math.MaxInt32{
			for min != 0 {
				reslut = append(reslut, string('a' + i))
				min--
			}
		}
	}
	return reslut
}

func main() {
	fmt.Println(commonChars([]string{"aaaa", "aba", "aca"}))
}
