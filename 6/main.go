package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Первый способ - при помощи отправки значения в канал
	ch1 := make(chan struct{})

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("Goroutine 1 stopped")
				close(ch1)
				return
			default:
				fmt.Println("Goroutine 1 is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ch1 <- struct{}{}

	// Второй способ - при помощи закрытия канала
	ch2 := make(chan int)
	go func() {
		for {
			v, ok := <-ch2
			if !ok {
				fmt.Println("Goroutine 2 stopped")
				return
			} else {
				fmt.Println("Goroutine 2 is working, received: ", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond)
		}
		close(ch2)
	}()

	time.Sleep(2 * time.Second)

	// Третий способ - при помощи контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 3 stopped")
				return
			default:
				fmt.Println("Goroutine 3 is working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// Четвертый способ - при помощи WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine 4 is working")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("Goroutine 4 stopped")
}
