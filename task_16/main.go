package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Распределение значений slice: значения элементов левой части меньше значения опорного, правой - больше.
// Использование дженериков: работает с любым типом данных, поддерживающим операции < <= >= > ==
func partition[T ~[]E, E cmp.Ordered](slice T) int {
	n := len(slice)
	// Опорный элемент (последний элемент из slice)
	pivot := slice[n-1]
	// Индекс первого элемента, который >= pivot
	i := 0
	for j := 0; j < n; j++ {
		// Если очередной элемент меньше опорного
		if slice[j] < pivot {
			// Меняются местами с первым элементом >= pivot
			slice[i], slice[j] = slice[j], slice[i]
			// Перенос указателя на следующий элемент >= pivot
			i++
		}
	}
	// Меняются местами первый элемент >= pivot и pivot
	slice[i], slice[n-1] = slice[n-1], slice[i]
	// Возврат индекса опорного элемента
	return i
}

// Рекурсивный запуск быстрой сортировки.
// Использование дженериков: работает с любым типом данных, поддерживающим операции < <= >= > ==
func QSort[T ~[]E, E cmp.Ordered](slice T) {
	if len(slice) > 1 {
		p := partition(slice)
		QSort(slice[:p])
		QSort(slice[p+1:])
	}
}

func main() {
	// Использование встроенного пакета
	slice := []int{5, 2, 4, 5, 1, 7, 6, -10, -3, 7, 8}
	slices.Sort(slice)
	fmt.Println(slice)

	slice = []int{5, 2, 4, 5, 1, 7, 6, -10, -3, 7, 8}
	QSort(slice)
	fmt.Println(slice)
}
