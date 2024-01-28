package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Реализация воркера
func Worker(ctx context.Context, wg *sync.WaitGroup, id int, data <-chan int) {
	// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
	defer wg.Done()
	for {
		select {
		// Получение сигнала о завершении работы контекста
		case <-ctx.Done():
			return
		// Получение данных из канала
		case number := <-data:
			fmt.Printf("Worker %d - %d\n", id, number)
		}
	}
}

// Реализация генератора данных
func Producer(ctx context.Context, wg *sync.WaitGroup, data chan<- int) {
	// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
	defer wg.Done()

	// Создание тикера, посылающего сигналы в канал каждые 200 мс
	ticker, i := time.NewTicker(200*time.Millisecond), 0
	// Остановка тикера
	defer ticker.Stop()
	for {
		select {
		// Получение сигнала о завершении работы контекста
		case <-ctx.Done():
			return
		// Генерация данных и запись их в канал
		case <-ticker.C:
			data <- i
			i++
		}
	}
}

func main() {
	// Получение количества воркеров
	workersSize := 0
	fmt.Print("Number of workers = ")
	fmt.Scan(&workersSize)
	if workersSize < 1 {
		return
	}

	// Создание канала для передачи данных
	data := make(chan int, workersSize)
	defer close(data)

	// Создание контекста для управления работой горутин
	ctx, cancel := context.WithCancel(context.Background())

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)

	for i := 0; i < workersSize; i++ {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go Worker(ctx, wg, i+1, data)
	}

	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go Producer(ctx, wg, data)

	// Создание канала для передачи сигналов ОС
	c := make(chan os.Signal, 1)
	defer close(c)
	// Регистрация уведомления канала c о нажатии сочетания ctrl-c
	signal.Notify(c, os.Interrupt)
	// Ожидание нажатия пользователем сочетания ctrl-c
	<-c

	// Оповещение горутин о завершении работы программы и ожидание выполнения всех горутин
	cancel()
	wg.Wait()

	// В данном случае при нажатии ctrl-c программа не закроется принудительно и горутины смогут корректно завершить свою работу
}
