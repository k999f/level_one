package main

import (
	"fmt"
	"reflect"
)

func getType1(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func getType2(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel")
	}
}

func getType3(i interface{}) {
	fmt.Printf("%T\n", i)
}

func main() {
	// Первый способ - через TypeOf
	getType1(123)
	getType1("test")
	getType1(true)
	getType1(make(chan int))
	fmt.Println()

	// Второй способ - через type switch
	getType2(123)
	getType2("test")
	getType2(true)
	getType2(make(chan int))
	fmt.Println()

	// Третий способ - через Printf
	getType3(123)
	getType3("test")
	getType3(true)
	getType3(make(chan int))
}
