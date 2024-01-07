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
	// Получаем количество воркеров
	workersNum := getWorkersNum()

	ch := make(chan int)
	wg := sync.WaitGroup{}

	// Запускаем в горутине постоянную запись в канал
	go startWriter(ch)
	// Запускаем воркеров в отдельных горутинах
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go startWorker(ch, i, &wg)
	}

	// Создаем канал для перехвата нажатия Ctrl+C
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT)

	// Запускаем горутину, ожидающую значение из канала exit
	go func() {
		<-exit
		// При получении сигнала закрываем каналы
		close(exit)
		close(ch)
	}()

	wg.Wait()
	fmt.Println("\nProgram stopped")
}
