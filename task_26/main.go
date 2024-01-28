package main

import (
	"fmt"
	"strings"
)

// Проверка уникальности символов в строке
func IsAllUnique(str string) bool {
	// Создание map множества символов
	symbolSet := make(map[rune]struct{})
	// Преобразование строки в нижний регистр и итерация по символам
	for _, symbol := range strings.ToLower(str) {
		// Если символ уже находится в множестве
		if _, ok := symbolSet[symbol]; ok {
			return false
		}
		// Добавление символа в множество
		symbolSet[symbol] = struct{}{}
	}
	return true
}

func main() {
	text := "abcd"
	fmt.Println(IsAllUnique(text))
	text = "abCdefAaf"
	fmt.Println(IsAllUnique(text))
	text = "abCdefAaf"
	fmt.Println(IsAllUnique(text))
}
