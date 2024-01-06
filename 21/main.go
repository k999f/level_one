package main

import "fmt"

type originalType struct{}

func (o *originalType) OriginalTypePrint() {
	fmt.Println("Original type print")
}

type desiredType interface {
	DesiredTypePrint()
}

type originalTypeAdapter struct {
	originalType
}

func (a *originalTypeAdapter) DesiredTypePrint() {
	a.OriginalTypePrint()
}

func funcForDesiredType(d desiredType) {
	d.DesiredTypePrint()
}

func main() {
	orig := originalType{}
	origAdapted := originalTypeAdapter{orig}
	funcForDesiredType(&origAdapted)
}
