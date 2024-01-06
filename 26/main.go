package main

import (
	"fmt"
	"strings"
)

func charactersUnique(str string) bool {
	strLower := strings.ToLower(str)
	strMap := make(map[rune]bool)

	for _, c := range strLower {
		if strMap[c] {
			return false
		}
		strMap[c] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	fmt.Printf("String: %s, result: %t\n", str1, charactersUnique(str1))

	str2 := "abCdefAaf"
	fmt.Printf("String: %s, result: %t\n", str2, charactersUnique(str2))

	str3 := "aabcd"
	fmt.Printf("String: %s, result: %t\n", str3, charactersUnique(str3))
}
