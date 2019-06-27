package main

import "fmt"

func swap(s []int, i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func diStringMatch(S string) []int {
	if len(S) == 0 {
		return []int{}
	}
	reslut := make([]int, 0, len(S) + 1)
	min, max := 0, len(S)
	for _, s := range S {
		if s == 'I' {
			reslut = append(reslut, min)
			min++
		}else {
			reslut = append(reslut, max)
			max--
		}
	}
	reslut = append(reslut, min)
	return reslut
}

func main() {
	fmt.Println(diStringMatch("D"))
}
