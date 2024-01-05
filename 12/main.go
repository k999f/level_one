package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range str {
		set[v] = struct{}{}
	}

	for k := range set {
		fmt.Println(k)
	}
}
