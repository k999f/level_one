package main

import "fmt"

func bitToZero(num int64, i uint) int64 {
	return num &^ (1 << i)
}

func bitToOne(num int64, i uint) int64 {
	return num | (1 << i)
}

func main() {
	var num int64 = 999
	var i uint = 4

	fmt.Printf("Number: %d (%b)\n", num, num)

	num = bitToOne(num, i)
	fmt.Printf("Number after setting bit %d to one: %d (%b)\n", i, num, num)

	num = bitToZero(num, i)
	fmt.Printf("Number after setting bit %d to zero: %d (%b)\n", i, num, num)
}
