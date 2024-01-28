package main

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

// Конкурентное возведение чисел в квадрат с сохранением их порядка
func OrderedSquaring(numbers []int) {
	result := make([]int, len(numbers))

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for i, number := range numbers {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		// Запуск горутины вычисления квадрата
		go func(number, index int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			// Так как горутина работает только с конкретным индексом в slice, блокировка необязательна (нет общих ресурсов)
			result[index] = number * number
		}(number, i)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()

	// Вывод результата
	writer := bufio.NewWriter(os.Stdout)
	for _, square := range result {
		writer.WriteString(strconv.Itoa(square))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
	writer.Flush()
}

// Конкурентное возведение чисел в квадрат без сохранения их порядка с использованием канала
func UnorderedSquaringUsingChannel(numbers []int) {
	writer := bufio.NewWriter(os.Stdout)
	// Создание канала для передачи полученных квадратов на вывод
	squares := make(chan int, len(numbers))

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for _, number := range numbers {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		// Запуск горутины вычисления квадрата
		go func(number int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			squares <- number * number
		}(number)
	}

	// Создание канала для оповещения об окончании вывода данных
	done := make(chan struct{})
	go func() {
		// Закрытие канала, оповещающее о том, что вывод завершен
		defer close(done)
		for square := range squares {
			writer.WriteString(strconv.Itoa(square))
			writer.WriteByte(' ')
		}
	}()

	// Ожидание выполнения всех горутин
	wg.Wait()
	// Закрытие канала, оповещающее о том, что вычисление квадратов завершено
	close(squares)
	// Ожидание завершения вывода квадратов
	<-done

	writer.WriteByte('\n')
	writer.Flush()
}

// Конкурентное возведение чисел в квадрат без сохранения их порядка с использованием мьютекса
func UnorderedSquaringUsingMutex(numbers []int) {
	writer := bufio.NewWriter(os.Stdout)
	// Создание мьютекса для блокировки доступа к общему ресурсу
	mutex := new(sync.Mutex)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	for _, number := range numbers {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(number int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			// Блокировка доступа к общему ресурсу
			mutex.Lock()
			writer.WriteString(strconv.Itoa(number * number))
			writer.WriteByte(' ')
			// Разблокировка доступа к общему ресурсу
			mutex.Unlock()
		}(number)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()

	writer.WriteByte('\n')
	writer.Flush()
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	OrderedSquaring(numbers)
	UnorderedSquaringUsingChannel(numbers)
	UnorderedSquaringUsingMutex(numbers)
}
