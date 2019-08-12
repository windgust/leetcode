package main

func detectCapitalUse(word string) bool {
	var lowerCnt, upperCnt int
	for i := 0; i < len(word); i++ {
		if isUpper(word[i]) {
			upperCnt++
		}
		if isLower(word[i]) {
			lowerCnt++
		}
	}
	if upperCnt == len(word) || lowerCnt == len(word) {
		return true
	}
	if upperCnt == 1 && isUpper(word[0]) {
		return true
	}
	return false
}

func isUpper(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isLower(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}
