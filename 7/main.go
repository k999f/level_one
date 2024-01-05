package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]string)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("Goroutine %d wrote key %d\n", i, i)
			data[i] = fmt.Sprintf("value_%d", i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(data)
}