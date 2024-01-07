package main

import "fmt"

func binarySearch(num []int, targetNum int) (int, bool) {
	// Левая и правая граница поиска
	left := 0
	right := len(num) - 1

	for left <= right {
		// Находим индекс центрального элемента
		mid := left + (right-left)/2
		if num[mid] == targetNum {
			// Если центральный элемент - искомый, то возвращаем его
			return mid, true
		} else if num[mid] < targetNum {
			// Если центральный элемент меньше искомого, то продолжаем поиск в правой части (сдвигаем левую границу на 1 элемент правее центрального)
			left = mid + 1
		} else {
			// Если центральный элемент больше искомого, то продолжаем поиск в левой части (сдвигаем правую границу на 1 элемент левее центрального)
			right = mid - 1
		}
	}

	// Возвращаем false, если элемент не найден
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
