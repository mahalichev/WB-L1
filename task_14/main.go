package main

import (
	"fmt"
	"reflect"
)

// Использование переключателя типов
func GetTypeUsingSwitch(value interface{}) string {
	result := "not defined in type switch"
	switch value.(type) {
	case int:
		result = "int"
	case string:
		result = "string"
	case bool:
		result = "bool"
	case chan int:
		result = "chan int"
	case chan string:
		result = "chan int"
	case chan bool:
		result = "chan bool"
	}
	return result
}

// Использование рефлексии (получение информации о переменной во время выполнения программы)
func GetTypeUsingReflect(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func main() {
	slice := []interface{}{1, "Hello", false, make(chan int), make(chan string), make(chan bool), 1.1, make(chan float64)}
	fmt.Println("Using type switch:")
	for _, value := range slice {
		fmt.Printf("%v - %s\n", value, GetTypeUsingSwitch(value))
	}

	fmt.Println("\nUsing reflect:")
	for _, value := range slice {
		fmt.Printf("%v - %s\n", value, GetTypeUsingReflect(value))
	}
}
