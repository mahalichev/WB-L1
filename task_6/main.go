package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка горутины с помощью отправки сигнала в канал
func OffUsingChannelSignal(duration time.Duration) {
	fmt.Print("Off by channel signal: ")
	// Создание сигнала для уведомления о необходимости завершить работу горутины
	off := make(chan struct{})

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()

		// Создание тикера, посылающего сигналы в канал каждую 1 с
		ticker := time.NewTicker(1 * time.Second)
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			// Получение сигнала о необходимости завершения горутины
			case <-off:
				fmt.Print("off signal was recieved. ")
				close(off)
				return
			// Вывод в stdout при каждом тике
			case <-ticker.C:
				fmt.Print(". ")
			}
		}
	}()

	time.Sleep(duration)
	// Запись данных в канал для сигнала о необходимости завершения горутины
	off <- struct{}{}
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

// Остановка горутины с помощью закрытия канала
func OffUsingChannelClose(seconds int) {
	fmt.Print("Off by channel close: ")
	// Создание канала для передачи данных
	data := make(chan string, 3)

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()
		// Получение данных из канала до момента его закрытия
		for value := range data {
			fmt.Print(value)
		}
		fmt.Print("off signal was recieved. ")
	}()

	for i := 0; i < seconds; i++ {
		time.Sleep(1 * time.Second)
		// Отправка данных в канал
		data <- ". "
	}
	// Закрытие канала, уведомляющее об окончании передачи в него данных
	close(data)
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

// Остановка горутины с помощью завершения контекста (вызова cancel)
func OffUsingContextWithCancel(duration time.Duration) {
	fmt.Print("Off by context cancel: ")
	// Создание контекста для управления работой горутин
	ctx, cancel := context.WithCancel(context.Background())

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()
		// Создание тикера, посылающего сигналы в канал каждую 1 с
		ticker := time.NewTicker(1 * time.Second)
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			// Получение сигнала о завершении работы контекста
			case <-ctx.Done():
				fmt.Print("off signal was recieved. ")
				return
			// Вывод в stdout при каждом тике
			case <-ticker.C:
				fmt.Print(". ")
			}
		}
	}()

	time.Sleep(duration)
	// Завершение контекста
	cancel()
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

// Остановка горутины с помощью завершения контекста (достижения таймаута)
func OffUsingContextWithTimeout(duration time.Duration) {
	fmt.Print("Off by context timeout: ")
	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)

	// Создание контекста для управления работой горутин (посылает сигнал через seconds секунд с момента создания контекста)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// Завершение контекста
	defer cancel()
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()
		// Создание тикера, посылающего сигналы в канал каждую 1 с
		ticker := time.NewTicker(1 * time.Second)
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			// Получение сигнала о завершении работы контекста
			case <-ctx.Done():
				fmt.Print("off signal was recieved. ")
				return
			// Вывод в stdout при каждом тике
			case <-ticker.C:
				fmt.Print(". ")
			}
		}
	}()
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

// Остановка горутины с помощью time.After
func OffUsingTimeAfter(duration time.Duration) {
	fmt.Print("Off by timeout: ")
	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()
		// Создание таймера, который отправит сигнал в канал через seconds секунд, и тикера, посылающего сигналы в канал каждую 1 с
		timer, ticker := time.After(duration), time.NewTicker(1*time.Second)
		// Остановка тикера
		defer ticker.Stop()
		for {
			select {
			case <-timer:
				fmt.Print("off signal was recieved. ")
				return
			case <-ticker.C:
				fmt.Print(". ")
			}
		}
	}()
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

// Остановка горутины с помощью флага
func OffUsingFlag(duration time.Duration) {
	fmt.Print("Off by flag: ")
	off := false

	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)
	// Увеличение счётчика горутин на 1
	wg.Add(1)
	go func() {
		// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
		defer wg.Done()
		for {
			if !off {
				fmt.Print(". ")
				time.Sleep(1 * time.Second)
			} else {
				fmt.Print("off signal was recieved. ")
				return
			}
		}
	}()
	time.Sleep(duration)
	// Установка значения true флагу завершения
	off = true
	// Ожидание выполнения всех горутин
	wg.Wait()
	fmt.Println("End.")
}

func main() {
	OffUsingChannelSignal(3 * time.Second)
	OffUsingChannelClose(3)
	OffUsingContextWithCancel(3 * time.Second)
	OffUsingContextWithTimeout(3 * time.Second)
	OffUsingTimeAfter(3 * time.Second)
	OffUsingFlag(3 * time.Second)
}
