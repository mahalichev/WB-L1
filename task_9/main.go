package main

import (
	"bufio"
	"os"
	"strconv"
)

// Первый этап - запись в первый канал чисел из slice
func FirstStage(in []int, out chan<- int) {
	// Закрытие канала, оповещающее о том, что вычисления завершены
	defer close(out)
	for _, number := range in {
		// Передача числа из slice в канал
		out <- number
	}
}

// Второй этап - запись во второй канал квадратов чисел из первого канала
func SecondStage(in <-chan int, out chan<- int) {
	// Закрытие канала, оповещающее о том, что вычисления завершены
	defer close(out)
	for number := range in {
		// Передача квадрата числа из канала in в канал out
		out <- number * number
	}
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	slice := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создание каналов
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)

	// Запуск конвеера
	go FirstStage(slice, chan1)
	go SecondStage(chan1, chan2)

	// Вывод результата работы конвеера
	for number := range chan2 {
		writer.WriteString(strconv.Itoa(number))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
	writer.Flush()
}
