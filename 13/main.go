package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	fmt.Printf("Before switching: num1 = %d, num2 = %d\n", num1, num2)

	num1, num2 = num2, num1
	fmt.Printf("After switching: num1 = %d, num2 = %d", num1, num2)
}
