package main

import "fmt"

type originalType struct{}

// Исходный тип
func (o *originalType) OriginalTypePrint() {
	fmt.Println("Original type print")
}

// Интерфейс требуемого типа
type desiredType interface {
	DesiredTypePrint()
}

// Адаптер для исходного типа
type originalTypeAdapter struct {
	originalType
}

func (a *originalTypeAdapter) DesiredTypePrint() {
	a.OriginalTypePrint()
}

// Функция, с аргументом  требуемого типа
func funcForDesiredType(d desiredType) {
	d.DesiredTypePrint()
}

func main() {
	orig := originalType{}
	origAdapted := originalTypeAdapter{orig}
	funcForDesiredType(&origAdapted)
}
