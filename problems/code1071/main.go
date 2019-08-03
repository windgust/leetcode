package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	k := len(str2)
	for k > 0 {
		gcd := str2[0:k]
		if isGcd(str2, gcd) && isGcd(str1, gcd) {
			return gcd
		}
		k--
		for k > 0 && len(str2)%k != 0 {
			k--
		}
	}
	return ""

}

func isGcd(str1, gcd string) bool {
	if len(str1)%len(gcd) != 0 {
		return false
	}
	k := len(gcd)
	for i := 0; i <= len(str1)-k; i += k {
		if str1[i:i+k] != gcd {
			return false
		}
	}
	return true
}

func main() {
	str1 := "ABCABCADC"
	str2 := "ABC"

	fmt.Println(gcdOfStrings(str1, str2))
}
