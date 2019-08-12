package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	index := make(map[string]int)
	for i := range logs {
		index[logs[i]] = i
	}

	sort.SliceStable(logs, func(i, j int) bool {
		strs1 := strings.SplitN(logs[i], " ", 2)
		strs2 := strings.SplitN(logs[j], " ", 2)

		isLetter1, isLetter2 := isLetter(strs1[1][0]), isLetter(strs2[1][0])
		if !isLetter1 && !isLetter2 {
			if index[logs[i]] < index[logs[j]] {
				return true
			}
			return false
		}
		if !isLetter1 {
			return false
		}
		if !isLetter2 {
			return true
		}

		if strs1[1] < strs2[1] {
			return true
		}
		if strs1[1] > strs2[1] {
			return false
		}
		if strs1[0] < strs2[0] {
			return true
		}
		return false
	})
	return logs
}

func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func main() {
	logs := []string{"j mo", "5 m w", "g 07", "o 2 0", "t q h"}
	fmt.Println(reorderLogFiles(logs))
}
