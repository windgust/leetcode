package main

import (
	"fmt"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	letters := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := range licensePlate {
		if licensePlate[i] < 'a' || licensePlate[i] > 'z' {
			continue
		}
		letters[licensePlate[i]]++
	}

	var res string
	for i := range words {
		tmp := make(map[byte]int)
		isValid := true
		lowerWord := strings.ToLower(words[i])
		for j := range lowerWord {
			if lowerWord[j] < 'a' || lowerWord[j] > 'z' {
				continue
			}
			tmp[lowerWord[j]]++
		}
		for k, v := range letters {
			if v > tmp[k] {
				isValid = false
				break
			}
		}
		if isValid && (res == "" || len(res) > len(words[i])) {
			res = words[i]
		}
	}
	return res
}

func main() {
	licensePlate := "AN87005"
	words := []string{"participant","individual","start","exist","above","already","easy","attack","player","important"}
	fmt.Println(shortestCompletingWord(licensePlate, words))
}
