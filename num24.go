package main

import (
	"fmt"
	"math"
)

// структура с полями x, y
type Point struct {
	x, y float64
}

// констуктор структуры
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// нахождение дистанции между точками по теореме пифагора
func dist(a, b Point) float64 {
	xDist := a.x - b.x
	yDist := a.y - b.y
	return math.Sqrt(xDist*xDist + yDist*yDist)
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(4, 6)
	fmt.Println(dist(a, b))
}
