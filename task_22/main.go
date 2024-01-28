package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// Инициализация переменных типа *big.Int для работы с большими числами числами
	num1, num2, result := new(big.Int), new(big.Int), new(big.Int)
	if _, ok := num1.SetString("98765432123456789", 10); !ok {
		log.Fatalf("an error occurred during the number initialization process")
	}

	if _, ok := num2.SetString("12345678987654321", 10); !ok {
		log.Fatalf("an error occurred during the number initialization process")
	}

	// Выполнение операций
	fmt.Println("Add:", result.Add(num1, num2).String())
	fmt.Println("Sub:", result.Sub(num1, num2).String())
	fmt.Println("Mul:", result.Mul(num1, num2).String())
	fmt.Println("Div:", result.Div(num1, num2).String())
}
