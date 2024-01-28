package main

import "fmt"

// Установка бита в 0
func SetBitTo0(number int64, position int) int64 {
	// Побитовым сдвигом влево единицы на position и применением побитового "НЕ" создается маска
	// Применяя побитовое "И" выполняется установка position бита числа в ноль
	return number & ^(1 << position)
}

// Установка бита в 1
func SetBitTo1(number int64, position int) int64 {
	// Побитовым сдвигом влево единицы на position создается маска
	// Применяя побитовое "ИЛИ" выполняется установка position бита числа в единицу
	return number | (1 << position)
}

func main() {
	number := int64(4096)
	position := 8
	fmt.Printf("set to 1: %d %d -> %d \n", number, position, SetBitTo1(number, position))
	number = int64(256)
	fmt.Printf("set to 0: %d %d -> %d \n", number, position, SetBitTo0(number, position))
}
