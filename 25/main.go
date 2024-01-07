package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	// Приостанавливаем время работы программы на d
	<-time.After(d)
}

func main() {
	fmt.Println("Program started")
	sleep(3 * time.Second)
	fmt.Println("Program stopped")
}
