package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) getValue() int {
	return c.value
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	fmt.Println("Start counter value: ", c.getValue())

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End counter value: ", c.getValue())
}
