package main

import (
	"fmt"

	"github.com/mahalichev/WB-L1/task_24/point"
)

func main() {
	point1 := point.New(2.5, 3)
	point2 := point.New(-1, 10)
	fmt.Println(point1.DistanceTo(point2))
}
