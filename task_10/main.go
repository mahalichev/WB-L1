package main

import "fmt"

// Распределение slice температур по группам
func GroupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temperature := range temperatures {
		// Определение подгруппы, к которой относится значение температуры
		subgroupKey := int(temperature/10) * 10

		// Добавление значения температуры в slice температур соответствующей подгруппы
		groups[subgroupKey] = append(groups[subgroupKey], temperature)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(GroupTemperatures(temperatures))
}
