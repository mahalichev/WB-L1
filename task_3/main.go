package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Получение суммы квадратов с использованием Mutex
func SumOfSquaresUsingMutex(numbers []int) int {
	result := 0
	// Создание мьютекса для блокировки доступа к общему ресурсу
	mutex := new(sync.Mutex)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for _, number := range numbers {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(num int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			// Блокировка доступа к общему ресурсу
			mutex.Lock()
			result += num * num
			// Разблокировка доступа к общему ресурсу
			mutex.Unlock()
		}(number)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()
	return result
}

// Получение суммы квадратов с использованием канала
func SumOfSquaresUsingChannel(numbers []int) int {
	result := 0
	n := len(numbers)
	// Создание канала для передачи данных между горутинами
	squares := make(chan int, n)

	for _, number := range numbers {
		go func(num int) {
			// Передача квадрата числа в канал
			squares <- num * num
		}(number)
	}

	for i := 0; i < n; i++ {
		// Получение квадрата числа из канала
		result += <-squares
	}

	// Закрытие канала
	close(squares)
	return result
}

// Получение суммы квадратов с использованием пакета atomic
func SumOfSquaresUsingAtomic(numbers []int) int {
	result := int64(0)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for _, number := range numbers {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(num int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			// Метод atomic.AddInt64() гарантирует, что операция сложения над переменной будет выполнена одной операцией для процессора.
			// Это означает, что операцию нельзя прервать на середине работы
			atomic.AddInt64(&result, int64(num*num))
		}(number)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()
	return int(result)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(SumOfSquaresUsingMutex(numbers))
	fmt.Println(SumOfSquaresUsingChannel(numbers))
	fmt.Println(SumOfSquaresUsingAtomic(numbers))
}
