package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {
	ma, end := "ma", "a"
	strs := strings.Split(S, " ")
	for i := range strs {
		if isVowelFirst(strs[i]) {
			strs[i] += ma + end
		} else {
			strs[i] = strs[i][1:] + strs[i][:1] + ma + end
		}
		end += "a"
	}
	return strings.Join(strs, " ")
}

func isVowelFirst(str string) bool {
	if str == "" {
		return false
	}
	vowel := []string{"a", "e", "i", "o", "u"}
	for i := range vowel {
		if strings.ToLower(string(str[0])) == vowel[i] {
			return true
		}
	}
	return false
}

func isContainVowel(str string) bool {
	if str == "" {
		return false
	}
	vowel := []string{"a", "e", "i", "o", "u"}
	lowerStr := strings.ToLower(str)
	for i := range vowel {
		if strings.Contains(lowerStr, vowel[i]) {
			return true
		}
	}
	return false
}

func main() {
	S := "HZ sg L"
	fmt.Println(toGoatLatin(S))
}
