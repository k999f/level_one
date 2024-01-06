package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 5
	fmt.Println("Original slice: ", sl)

	sl = append(sl[:i], sl[i+1:]...)

	fmt.Println("New slice: ", sl)
}
