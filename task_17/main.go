package main

import (
	"cmp"
	"fmt"
)

// Возвращение индекса элемента, найденного с помощью бинарного поиска. Если элемент не был найден, функция возвращает -1.
// Использование дженериков: работает с любым типом данных, поддерживающим операции < <= >= > ==
func BinarySearch[T ~[]E, E cmp.Ordered](slice T, target E) int {
	l, r := 0, len(slice)-1
	// Итерирование пока l левее или равен r
	for l <= r {
		// Индекс центрального элемента в интервале [l;r]
		m := (l + r) / 2
		// Если элемент равный target найден
		if slice[m] == target {
			return m
			// Если элемент с индексом m меньше target
		} else if slice[m] < target {
			// Все значения в интервале [l, m] меньше target
			l = m + 1
			// Если элемент с индексом m больше target
		} else {
			// Все значения в интервале [m, r] больше target
			r = m - 1
		}
	}
	// Элемент со значением target не найден
	return -1
}

func main() {
	a := []int{-10, 1, 5, 7, 8, 20}
	fmt.Println(BinarySearch(a, 5))
	fmt.Println(BinarySearch(a, -10))
	fmt.Println(BinarySearch(a, 20))
	fmt.Println(BinarySearch(a, -20))
	fmt.Println(BinarySearch(a, 30))
}
