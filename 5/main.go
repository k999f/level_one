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
	timer := time.NewTimer(time.Duration(N) * time.Second)

	go startWriter(ch)

	for {
		select {
		case m := <-ch:
			fmt.Println("Received message: ", m)
		case <-timer.C:
			fmt.Println("Time over")
			close(ch)
			return
		}
	}
}
