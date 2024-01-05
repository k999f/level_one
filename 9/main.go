package main

import (
	"fmt"
	"sync"
)

func main() {
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)
	doubleCh := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for n := range ch {
			doubleCh <- 2 * n
		}
		close(doubleCh)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for n := range doubleCh {
			fmt.Println(n)
		}
		wg.Done()
	}()

	for _, n := range num {
		ch <- n
	}
	close(ch)
	wg.Wait()
}
