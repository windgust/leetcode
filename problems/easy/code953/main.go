package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	index := make(map[byte]int)
	for i := range order {
		index[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		var ii, jj int
		for ii < len(words[i]) && jj < len(words[i+1]) {
			if index[words[i][ii]] == index[words[i+1][jj]] {
				ii++
				jj++
				continue
			}
			if index[words[i][ii]] > index[words[i+1][jj]] {
				return false
			}
			break
		}
		if ii < len(words[i]) && jj == len(words[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"wov","woa"}
	order := "worldabcefghijkmnpqstuvxyz"

	fmt.Println(isAlienSorted(words, order))
}
