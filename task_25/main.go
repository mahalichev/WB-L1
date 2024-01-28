package main

import (
	"fmt"
	"time"
)

// Использование каналов (time.After)
func SleepUsingTimeAfter(duration time.Duration) {
	// Функция time.After() передает значение в канал по истечению указанного времени
	<-time.After(duration)
}

// Сравнивание времени в цикле for (больше нагрузка на процессор)
func SleepUsingFor(duration time.Duration) {
	for now := time.Now(); ; {
		// Сравнивание времени прошедшего с момента вызова функции с необходимой длительностью паузы
		if time.Since(now) >= duration {
			return
		}
	}
}

func main() {
	timeDuration := time.Duration(3 * time.Second)
	fmt.Println("Start SleepUsingTimeAfter")
	now := time.Now()
	SleepUsingTimeAfter(timeDuration)
	fmt.Printf("Stopped after %s\n", time.Since(now))

	fmt.Println("Start SleepUsingFor")
	now = time.Now()
	SleepUsingFor(timeDuration)
	fmt.Printf("Stopped after %s\n", time.Since(now))
}
