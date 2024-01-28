package main

import "fmt"

func main() {
	// Одновременное присваивание переменным значения
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

	// Использование сложения и вычитания
	a, b = 1, 2
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)

	// Использование умножения и деления (возможен неправильный результат при переполнении)
	a, b = 1, 2
	a *= b
	b = a / b
	a /= b
	fmt.Println(a, b)

	// Использование XOR
	a, b = 1, 2
	a ^= b
	b = a ^ b
	a ^= b
	fmt.Println(a, b)
}
