package main

import (
	"fmt"
	"strings"
)

func charactersUnique(str string) bool {
	// Преобразуем строку к нижнему регистру
	strLower := strings.ToLower(str)
	// Создаем map для символов строки
	strMap := make(map[rune]bool)

	for _, c := range strLower {
		// Если значение по ключу - true (следовательно, символ встречается уже не первый раз), то возвращаем false
		if strMap[c] {
			return false
		}
		// Если символ встерчается первый раз, то создаем ключ и значение по ключу делаем true
		strMap[c] = true
	}

	// Возвращаем true, если цикл закончился (следовательно, не были найдены повторяющиеся символы)
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
