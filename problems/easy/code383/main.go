package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[int32]int)
	for _, letter := range magazine {
		letters[letter]++
	}
	for _, letter := range ransomNote {
		if letters[letter] == 0 {
			return false
		}
		letters[letter]--
	}
	return true
}
