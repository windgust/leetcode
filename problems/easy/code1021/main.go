package main

import "fmt"

func removeOuterParentheses(S string) string {
	var result string
	var cnt, length int
	for i, c := range S {
		if string(c) == "(" {
			cnt++
		}else {
			cnt--
		}
		length++
		if cnt == 0 {
			if length > 2 {
				result += S[i - length + 2 : i]
			}
			length = 0
		}
	}
	return result
}

func main() {
	str := ""
	fmt.Println(removeOuterParentheses(str))
}
