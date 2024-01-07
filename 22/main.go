package main

import (
	"fmt"
	"math/big"
)

func main() {
	// a2^21
	a := big.NewInt(2097152)
	// 2^22
	b := big.NewInt(4194304)

	fmt.Printf("a = %d, b = %d\n", a, b)

	// Умножаем
	mult := big.NewInt(0)
	mult.Mul(a, b)
	fmt.Println("a * b = ", mult)

	// Делим
	quot := big.NewFloat(0)
	quot.Quo(big.NewFloat(0).SetInt(a), big.NewFloat(0).SetInt(b))
	fmt.Println("a / b = ", quot)

	// Складываем
	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Println("a + b = ", sum)

	// Вычитаем
	diff := big.NewInt(0)
	diff.Sub(a, b)
	fmt.Println("a - b = ", diff)
}
