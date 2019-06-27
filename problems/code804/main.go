package main

func uniqueMorseRepresentations(words []string) int {
	letters := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	result := make(map[string]bool)
	for _, word := range words {
		var tmp string
		for _, ch := range word {
			tmp += letters[ch - 'a']
		}
		result[tmp] = true
	}
	return len(result)
}