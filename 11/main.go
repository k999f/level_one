package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 5, 1, 3, 10}
	arr2 := []int{3, 5, 2}
	arr1Map := make(map[int]struct{})
	intersec := make(map[int]struct{})

	// Создаем map со всеми значениями arr1 в качестве ключей
	for _, e := range arr1 {
		arr1Map[e] = struct{}{}
	}

	// Проверяем, имеются ли в arr1Map ключи, равные значениям элементов из arr2
	for _, e := range arr2 {
		_, ok := arr1Map[e]
		// Добавляем в intersec, если ключ имеется
		if ok {
			intersec[e] = struct{}{}
		}
	}

	for k := range intersec {
		fmt.Println(k)
	}
}
