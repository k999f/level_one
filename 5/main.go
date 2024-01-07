package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startWriter(ch chan int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	N := 7
	ch := make(chan int)
	// Создаем таймер длительностью N секунд
	timer := time.NewTimer(time.Duration(N) * time.Second)

	// Начинаем оправку данных в канал
	go startWriter(ch)

	for {
		select {
		case m := <-ch:
			fmt.Println("Received message: ", m)
		case <-timer.C:
			// Закрываем канал и делаем return при получении сообщения о том, что таймер истек
			fmt.Println("Time over")
			close(ch)
			return
		}
	}
}
