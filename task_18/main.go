package main

import (
	"fmt"
	"sync"
)

// Структура-счётчик
type Counter struct {
	value int
	mutex sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		// Создание мьютекса для блокировки доступа к общему ресурсу
		mutex: sync.Mutex{},
	}
}

// Безопасное увеличение счетчика в конкурентной среде
func (c *Counter) Increase() {
	// Блокировка доступа к общему ресурсу
	c.mutex.Lock()
	c.value++
	// Разблокировка общего ресурса
	c.mutex.Unlock()
}

// Получение значения счетчика
func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	// Получение количества горутин
	goroutinesCount := 0
	fmt.Print("Number of goroutines = ")
	fmt.Scan(&goroutinesCount)
	if goroutinesCount < 1 {
		return
	}

	// Создание структуры-счётчика
	counter := NewCounter()

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for i := 0; i < goroutinesCount; i++ {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func() {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for i := 0; i < 100; i++ {
				counter.Increase()
			}
		}()
	}
	// Ожидание выполнения всех горутин
	wg.Wait()

	fmt.Println(counter.GetValue())
}
