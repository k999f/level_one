package main

import (
	"fmt"
)

func quickSort(num []int) []int {
	if len(num) < 2 {
		return num
	}

	left := []int{}
	right := []int{}
	result := []int{}

	pivot := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] <= pivot {
			left = append(left, num[i])
		} else {
			right = append(right, num[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)

	return result
}

func main() {
	num := []int{5, 4, 1, 3, 8, -7, 0, 9, 19}
	fmt.Println("Unsorted: ", num)
	sortedNum := quickSort(num)
	fmt.Println("Sorted: ", sortedNum)
}
