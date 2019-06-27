package main

import (
	"fmt"
	"strings"
)

func replcaeAll(localName string) string {
	k := 0
	src := []rune(localName)
	for i, v := range src {
		if string(v) != "." {
			src[k] = src[i]
			k++
		}
	}
	return string(src[0:k])
}

func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, email := range emails {
		strs := strings.Split(email, "@")
		localName, domainName := strs[0], strs[1]
		localName = strings.Split(localName, "+")[0]
		localName = replcaeAll(localName)
		fmt.Println(localName + "@" + domainName)
		set[localName + "@" + domainName] = struct{}{}
	}
	return len(set)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
