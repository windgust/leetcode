package main

import "fmt"

func numberOfLines(widths []int, S string) []int {
	lines, plus := 1, 0
	for _, ch := range S {
		if widths[ch-'a']+plus <= 100 {
			plus += widths[ch-'a']
			continue
		}
		lines++
		plus = widths[ch-'a']
	}
	return []int{lines, plus}
}

func main() {
	width := []int{4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	S := "bbbcccdddaaa"
	fmt.Println(numberOfLines(width, S))
}
