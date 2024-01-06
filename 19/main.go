package main

import "fmt"

func reverse(str string) string {
	strRune := []rune(str)
	var reverseStr []rune
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
