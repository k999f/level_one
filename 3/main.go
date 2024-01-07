package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	sum1 := 0
	sum2 := 0

	// Первый вариант
	wg := sync.WaitGroup{}
	// Запускаем горутины, каждая из которых находит значение квадрата числа из numbers и прибавляет его к sum1
	for _, n := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sum1 += n * n
		}(n)
	}
	wg.Wait()
	fmt.Println(sum1)

	// Второй вариант
	c := make(chan int)
	// Запускаем горутины, каждая из которых находит значение квадрата числа из numbers и отправляет его в канал
	for _, n := range numbers {
		go func(n int) {
			c <- n * n
		}(n)
	}
	// Получаем значение из канала и прибавляем его к sum2
	for _, _ = range numbers {
		res := <-c
		sum2 += res
	}
	close(c)
	fmt.Println(sum2)
}
