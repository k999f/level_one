package main

import "fmt"

func reverse(str string) string {
	// Преобразуем строку к слайсу рун, чтобы работать с символами unicode
	strRune := []rune(str)
	var reverseStr []rune
	// Проходимся в цикле с конца strRune и добавляем каждый элемент в конец reverseStr
	for i := len(strRune) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, strRune[i])
	}
	return string(reverseStr)
}

func main() {
	str := "главрыба 我爱皇帝"
	fmt.Println("Original string: ", str)

	reverseStr := reverse(str)
	fmt.Println("Reversed string: ", reverseStr)
}
