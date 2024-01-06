package main

import "fmt"

func binarySearch(num []int, targetNum int) (int, bool) {
	left, right := 0, len(num)-1

	for left <= right {
		mid := left + (right-left)/2
		if num[mid] == targetNum {
			return mid, true
		} else if num[mid] < targetNum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, false
}

func main() {
	num := []int{-3, 0, 1, 2, 4, 5, 6, 7, 11, 15, 16, 18}
	targetNum := 11
	i, found := binarySearch(num, targetNum)
	if found {
		fmt.Printf("Index of number %d: %d\n", targetNum, i)
	} else {
		fmt.Printf("Number %d not found\n", targetNum)
	}
}
