package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func startWorker(ch chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for m := range ch {
		fmt.Println("Worker ", i, " received message: ", m)
	}
}

func startWriter(ch chan int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(1 * time.Second)
	}
}

func getWorkersNum() int {
	var input string
	var num int
	fmt.Println("Enter number of workers")
	for {
		_, err := fmt.Scan(&input)
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number")
		} else if num <= 0 {
			fmt.Println("Number of workers can't be <= 0")
		} else {
			break
		}
	}
	return num
}

func main() {
	workersNum := getWorkersNum()

	ch := make(chan int)
	wg := sync.WaitGroup{}

	go startWriter(ch)
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go startWorker(ch, i, &wg)
	}

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT)

	go func() {
		<-exit
		close(exit)
		close(ch)
	}()

	wg.Wait()
	fmt.Println("\nProgram stopped")
}
