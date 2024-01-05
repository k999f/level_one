package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	// Первый вариант
	wg := sync.WaitGroup{}
	for _, n := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(n)
	}
	wg.Wait()

	fmt.Println()

	// Второй вариант
	c := make(chan int)
	for _, n := range numbers {
		go func(n int) {
			c <- n * n
		}(n)
	}
	for _, _ = range numbers {
		res := <-c
		fmt.Println(res)
	}
	close(c)
}
