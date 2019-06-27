package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	strs := strings.Split(text, " ")
	i := 0
	reslut := make([]string, 0, 0)
	for i + 2 < len(strs) {
		if strs[i] == first && strs[i + 1] == second {
			reslut = append(reslut, strs[i + 2])
		}
		i++
	}
	return reslut
}

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student","a","good"))
}
