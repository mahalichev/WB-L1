package point

import "math"

type Point struct {
	// Идентификаторы, начинающиеся со строчной буквы, видны только внутри пакета
	x, y float64
}

// Создание новой точки
func New(x, y float64) Point {
	return Point{x: x, y: y}
}

// Вычисление расстояния между двумя точками
func (point Point) DistanceTo(pointTo Point) float64 {
	dX := point.x - pointTo.x
	dY := point.y - pointTo.y
	return math.Sqrt(dX*dX + dY*dY)
}
