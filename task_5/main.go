package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Вывод данных из канала
func Consumer(wg *sync.WaitGroup, data <-chan int) {
	// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
	defer wg.Done()
	for i := range data {
		fmt.Printf("%d ", i)
	}
}

// Использование контекста для отсчёта времени
func TimeoutUsingContext(duration time.Duration) {
	// Создание канала для передачи данных
	data := make(chan int, 5)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 2
	wg.Add(2)

	// Создание контекста для управления работой горутин (посылает сигнал через seconds секунд с момента создания контекста)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// Завершение контекста
	defer cancel()

	now := time.Now()
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()

		// Создание тикера, посылающего сигналы в канал каждые 500 мс
		ticker, i := time.NewTicker(500*time.Millisecond), 0
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			// Получение сигнала о завершении работы контекста
			case <-ctx.Done():
				// Закрытие канала, уведомляющее об окончании передачи в него данных
				close(data)
				return
			// Генерация данных и запись их в канал
			case <-ticker.C:
				data <- i
				i++
			}
		}
	}()
	go Consumer(wg, data)

	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Printf("\nfunc stopped after %s", time.Since(now))
}

// Использование time.After для отсчёта времени
func TimeoutUsingTimeAfter(duration time.Duration) {
	// Создание канала для передачи данных
	data := make(chan int, 5)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 2
	wg.Add(2)

	now := time.Now()
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()

		// Создание таймера, который отправит сигнал в канал через seconds секунд, и тикера, посылающего сигналы в канал каждые 500 мс
		timer, ticker, i := time.After(duration), time.NewTicker(500*time.Millisecond), 0
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			case <-timer:
				// Закрытие канала, уведомляющее об окончании передачи в него данных
				close(data)
				return
			// Генерация данных и запись их в канал
			case <-ticker.C:
				data <- i
				i++
			}
		}
	}()
	go Consumer(wg, data)

	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Printf("\nfunc stopped after %s", time.Since(now))
}

func main() {
	fmt.Println("Using context:")
	TimeoutUsingContext(5 * time.Second)

	fmt.Println("\nUsing time.After:")
	TimeoutUsingTimeAfter(5 * time.Second)
}
