package main

import (
	"fmt"
	"strings"
)

var i = 1
func letterCasePermutation(S string) []string {
	result := make([]string,0)
	calueResult("", S, &result)
	return result
}

func calueResult(start, end string, result *[]string) {
	if end == "" {
		*result = append(*result, start)
		return
	}
	endS := end[:1]
	calueResult(start+strings.ToLower(endS), end[1:], result)
	if endS < "0" || endS > "9" {
		calueResult(start+strings.ToUpper(endS), end[1:], result)
	}
}

func main() {
	S := "0"
	fmt.Println(letterCasePermutation(S))
}
