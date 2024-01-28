package main

import (
	"fmt"
)

// Определение пересечения двух множеств.
// Использование дженериков: работает с любым сравниваемым типом данных (к которому применима операция ==).
// Результатом работы функции будет иметь тот же тип данных, что и переданные аргументы
func SetsIntersection[T ~[]E, E comparable](set1, set2 T) T {
	// Создание map из первого множества
	setMap := make(map[E]struct{}, len(set1))
	for _, value := range set1 {
		setMap[value] = struct{}{}
	}

	// Проверка наличия элементов второго множества в map
	result := T{}
	for _, value := range set2 {
		if _, ok := setMap[value]; ok {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	set1 := []int{-2, 1, -7, -1, 8}
	set2 := []int{-7, -2, 3, 1, 10}
	fmt.Println(SetsIntersection(set1, set2))
}
