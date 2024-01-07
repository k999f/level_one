package main

import (
	"fmt"
)

func quickSort(num []int) []int {
	// Если длина слайса < 2, то возвращаем его как есть
	if len(num) < 2 {
		return num
	}

	left := []int{}
	right := []int{}
	result := []int{}

	// Берем первый элемент как опорный
	pivot := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] <= pivot {
			// Если текущий элемент меньше опорного, то добавляем его в левую часть
			left = append(left, num[i])
		} else {
			// Если текущий элемент больше опорного, то добавляем его в правую часть
			right = append(right, num[i])
		}
	}

	// Рекурсивно сортируем левую и правую часть
	left = quickSort(left)
	right = quickSort(right)

	// Соединяем левую часть, опорный элемент и правую часть
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
