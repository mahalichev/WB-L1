package main

import "fmt"

// Создание множества из slice строк
func MakeSet(slice []string) []string {
	// Создание map множества из строк
	setMap := make(map[string]struct{})
	for _, value := range slice {
		setMap[value] = struct{}{}
	}

	// Перенос элементов множества в slice
	result := make([]string, 0, len(setMap))
	for key := range setMap {
		result = append(result, key)
	}
	return result
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(slice))
}
