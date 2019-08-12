package main

import "fmt"

func removeDuplicates(S string) string {
	var stack []rune
	bytes := []rune(S)
	for _, ch := range bytes {
		stack = append(stack, ch)
		for {
			if len(stack) < 2 {
				break
			}
			length := len(stack)
			if stack[length - 1] != stack[length - 2] {
				break
			}
			stack = stack[:length - 2]
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates(""))
}
