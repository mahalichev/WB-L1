package main

import (
	"fmt"
	"slices"
)

// Реализация собственной функции переворачивания строки
func OwnReverse(original string) string {
	// Преобразование входной строки в slice рун
	result := []rune(original)

	// До середины slice поменять местами каждый элемент с начала slice с соответствующим элементом с конца slice
	n := len(result)
	for i := 0; i < n/2; i++ {
		result[i], result[n-i-1] = result[n-i-1], result[i]
	}
	return string(result)
}

// Реализация функции переворачивания строки с использованием пакета slices
func SlicesReverse(original string) string {
	// Преобразование входной строки в slice рун
	result := []rune(original)

	// Использование slices.Reverse() для slice рун
	slices.Reverse(result)
	return string(result)
}

func main() {
	input := "🢆123Тест🤡"
	fmt.Println(OwnReverse(input))
	fmt.Println(SlicesReverse(input))
}
