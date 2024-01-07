package main

import (
	"fmt"
	"sync"
)

// Структура, содержащая значение счетчика и мьютекс
type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	// Блокируем доступ к счетчику
	c.mu.Lock()
	// Увеличиваем счетчик
	c.value++
	// Отключаем блокировку доступа к счетчику
	c.mu.Unlock()
}

func (c *Counter) getValue() int {
	return c.value
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	fmt.Println("Start counter value: ", c.getValue())

	// Запускаем 5- горутин, каждая из которых вызывает метод инкрементирования счетчика
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
