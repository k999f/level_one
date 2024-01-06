package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	var sb strings.Builder
	words := strings.Split(str, " ")
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		sb.WriteString(" ")
	}
	return strings.TrimRight(sb.String(), " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println("Original string: ", str)

	reverseStr := reverseWords(str)
	fmt.Println("Reversed string: ", reverseStr)
}
