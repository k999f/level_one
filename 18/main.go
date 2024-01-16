package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Структура, содержащая значение счетчика и мьютекс
type Counter1 struct {
	value int
	mu    sync.Mutex
}

func (c *Counter1) Increment() {
	// Блокируем мьютекс
	c.mu.Lock()
	// Увеличиваем счетчик
	c.value++
	// Отключаем блокировку мьютекса
	c.mu.Unlock()
}

func (c *Counter1) getValue() int {
	return c.value
}

// Структура, содержащая значение счетчика
type Counter2 struct {
	value int32
}

// Метод для атомарной инкрементации счетчика
func (c *Counter2) Increment() {
	atomic.AddInt32(&c.value, 1)
}

func (c *Counter2) getValue() int32 {
	return c.value
}

func main() {
	// Первый способ - с использованием мьютекса
	с1 := Counter1{}
	wg1 := sync.WaitGroup{}
	fmt.Println("Start counter 1 value: ", с1.getValue())

	// Запускаем 5 горутин, каждая из которых вызывает метод инкрементирования Counter1
	for i := 0; i < 50; i++ {
		wg1.Add(1)
		go func() {
			с1.Increment()
			wg1.Done()
		}()
	}

	wg1.Wait()
	fmt.Println("End counter 1 value: ", с1.getValue())

	// Второй способ - с использованием атомиков
	с2 := Counter2{}
	wg2 := sync.WaitGroup{}
	fmt.Println("Start counter 2 value: ", с2.getValue())

	// Запускаем 5 горутин, каждая из которых вызывает метод инкрементирования Counter2
	for i := 0; i < 50; i++ {
		wg2.Add(1)
		go func() {
			с2.Increment()
			wg2.Done()
		}()
	}

	wg2.Wait()
	fmt.Println("End counter 2 value: ", с2.getValue())
}
