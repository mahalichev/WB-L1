package main

import (
	"fmt"
	"time"
)

// Объявление структуры Human
type Human struct {
	Name    string
	Surname string
	Age     int
	City    string
}

// Метод представления человека
func (h Human) Introduce() {
	fmt.Printf("Hello, my name is %s %s. I am %d years old. I'm from %s\n", h.Name, h.Surname, h.Age, h.City)
}

// Метод вывода информации о Human
func (h Human) SayInfo() {
	fmt.Printf("Human info:\n    Name - %s;\n    Surname - %s;\n    Age - %d;\n    City - %s\n", h.Name, h.Surname, h.Age, h.City)
}

// Объявление структуры Action
type Action struct {
	// Встроенный тип данных Human
	Human
	Action   string
	City     string
	Duration time.Duration
}

// Метод вывода информации об Action
func (a Action) SayInfo() {
	// Так как поля City присутствуют в двух структурах, для получения города человека необходимо явно указать Human.City
	// Нет необходимости явно указывать, если требуемое поле присутствует только в одной из структур
	fmt.Printf("Action info:\n    Person - %s %s from %s, %d y.o.;\n    Action - %s for %s in %s\n", a.Name, a.Surname, a.Human.City, a.Age, a.Action, a.Duration, a.City)
}

// Метод выводит информацию о том, что делает человек
func (a Action) ActionInfo() {
	fmt.Printf("%s %s is %s for %s in %s", a.Name, a.Surname, a.Action, a.Duration, a.City)
}

func main() {
	action := Action{
		Human: Human{
			Name:    "Mike",
			Surname: "Smith",
			Age:     37,
			City:    "Hong Kong",
		},
		Action:   "Programming",
		City:     "Shanghai",
		Duration: time.Duration(8 * time.Hour),
	}

	// Так же, как в случае с полями структур, нет необходимости явно указывать структуру, если метод присутствует только
	// в одной структуре
	action.Introduce()

	// Методы SayInfo() объявлены в двух структурах
	action.Human.SayInfo()
	action.SayInfo()

	action.ActionInfo()
}
