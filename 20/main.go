package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	var sb strings.Builder
	// Разделяем исходную строку на слайс строк по пробелу
	words := strings.Split(str, " ")
	// Проходимся в цикле с конца words и добавляем каждую строку и пробел после нее к итоговой строке
	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		sb.WriteString(" ")
	}
	// Возвращаем итоговую строку с удаленным справа лишним пробелом
	return strings.TrimRight(sb.String(), " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println("Original string: ", str)

	reverseStr := reverseWords(str)
	fmt.Println("Reversed string: ", reverseStr)
}
