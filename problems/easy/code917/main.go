package main

import (
	"fmt"
	"unsafe"
)

func reverseOnlyLetters(S string) string {
	i, j := 0, len(S)-1
	bs := []byte(S)
	for i < j {
		for i < j && !isLetter(&bs[i]) {
			i++
		}
		for i < j && !isLetter(&bs[j]) {
			j--
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return *(*string)(unsafe.Pointer(&bs))
}

func isLetter(ch *byte) bool {
	return *ch >= 'a' && *ch <= 'z' || *ch >= 'A' && *ch <= 'Z'
}

func main() {
	S := "v1a"
	fmt.Println(reverseOnlyLetters(S))
}
