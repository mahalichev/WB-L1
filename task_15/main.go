package main

import (
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("€", size)
}

var justString string

// Проблема 1: срез v[:100] возвращает первые 100 байтов строки, из-за чего могут возникнуть проблемы, если в строке были символы
// размером более 1 байта. Проблема 2: при использовании justString = v[:100] глобальная переменная justString продолжает ссылаться
// массив в строке v, поэтому вся строка останется в памяти (утечка памяти). Решение: создание slice рун из строки v позволит решить обе
// проблемы - justString сможет хранить первые 100 символов и не будет ссылаться на строку v.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}
