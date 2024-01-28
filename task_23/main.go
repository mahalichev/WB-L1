package main

import (
	"fmt"
)

func main() {
	index := 2
	// Изменение исходного slice с сохранением порядка элементов
	slice := []int{0, 1, 2, 3, 4, 5}
	n := len(slice)
	if index >= 0 && index < n {
		// Добавление к элементам slice в интервале [0:index) элементов из интервала [index+1:len(slice))
		slice = append(slice[:index], slice[index+1:]...)
		fmt.Println(slice)
	}

	// Изменение исходного slice без сохранения порядка элементов
	slice = []int{0, 1, 2, 3, 4, 5}
	n = len(slice)
	if index >= 0 && index < n {
		// Присваивание элементу по индексу значения последнего элемента slice
		slice[index] = slice[n-1]
		// Уменьшение длины slice на 1
		slice = slice[:n-1]
		fmt.Println(slice)
	}

	// Создание нового slice с сохранением порядка элементов
	slice = []int{0, 1, 2, 3, 4, 5}
	n = len(slice)
	if index >= 0 && index < n {
		// Создание нового slice размером len(slice)-1
		newSlice := make([]int, n-1)
		// Копирование значений в интервале [0, index) в новый slice
		copy(newSlice, slice[:index])
		// Копирование значений в интервале [0, index) в новый slice
		copy(newSlice[index:], slice[index+1:])
		fmt.Println(newSlice)
	}
}
