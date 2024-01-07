package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func FindDistance(p1, p2 *Point) float64 {
	n1 := math.Pow(p2.x-p1.x, 2)
	n2 := math.Pow(p2.y-p1.y, 2)
	return math.Sqrt(n1 + n2)
}
func main() {
	// Создаем точки
	p1 := NewPoint(10, 20)
	p2 := NewPoint(30, 40)
	fmt.Printf("Point 1: %f, %f; point 2: %f, %f\n", p1.x, p1.y, p2.x, p2.y)
	// Ищем расстояние
	d := FindDistance(p1, p2)
	fmt.Println("Distance = ", d)
}
