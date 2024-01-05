package main

import (
	"strings"
)

// Исходный код

// Смысл исходного кода я понимаю следующим образом: мы хотим получить строку,
// которая представляет из себя первые 100 символов другой - очень большой строки

// Проблема исходного кода: В переменную justString записываются первые 100 символов
// большой строки v => justString ссылается на v => v не будет удалена сборщиком мусора

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func createHugeString(n int) string {
// 	return strings.Repeat("a", n)
// }

// func main() {
// 	someFunc()
// }

// Исправленный код

// Чтобы решить описанную выше проблему создадим слайс байт, выполним копирование
// первых 100 символов очень большой строки в этот слайс и присвоим justString
// его значение => теперь justString ссылается на отдельную строку, содержащую только
// нужные элементы => v будет удалена сборщиком мусора и не вызовет утечку памяти

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	sl := make([]byte, 100)
	copy(sl, v[:100])
	justString = string(sl)
}

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

func main() {
	someFunc()
}
