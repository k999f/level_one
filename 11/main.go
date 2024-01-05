package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 5, 1, 3, 10}
	arr2 := []int{3, 5, 2}
	arr1Map := make(map[int]struct{})
	intersec := make(map[int]struct{})

	for _, e := range arr1 {
		arr1Map[e] = struct{}{}
	}

	for _, e := range arr2 {
		_, ok := arr1Map[e]
		if ok {
			intersec[e] = struct{}{}
		}
	}

	for k := range intersec {
		fmt.Println(k)
	}
}
