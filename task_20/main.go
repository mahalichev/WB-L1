package main

import (
	"fmt"
	"slices"
	"strings"
)

// Реализация собственной функции переворачивания слов в строке
func ReverseWords(original string) string {
	// Разбиение строки на slice слов (слово - подстрока, разделённая пробельными символами)
	fields := strings.Fields(original)

	// До середины slice поменять местами каждый элемент с начала slice с соответствующим элементом с конца slice
	n := len(fields)
	for i := 0; i < n/2; i++ {
		fields[i], fields[n-i-1] = fields[n-i-1], fields[i]
	}
	// Объединение подстрок из slice с использованием разделителя " "
	return strings.Join(fields, " ")
}

// Реализация функции переворачивания слов в строке с использованием пакета slices
func SlicesReverseWords(original string) string {
	// Разбиение строки на slice слов (слово - подстрока, разделённая пробельными символами)
	fields := strings.Fields(original)

	// Использование slices.Reverse() для slice рун
	slices.Reverse(fields)
	return strings.Join(fields, " ")
}

func main() {
	text := "snow dog sun"
	fmt.Println(ReverseWords(text))
	fmt.Println(SlicesReverseWords(text))
}
