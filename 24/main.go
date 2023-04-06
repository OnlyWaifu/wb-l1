package main

import (
	"fmt"
	"math"
)

// Point представляет собой координатную точку
type Point struct {
	x, y float64
}

// X приёмник для x
func (p Point) X() float64 {
	return p.x
}

// Y приёмник для y
func (p Point) Y() float64 {
	return p.y
}

// New функция-конструктор для типа Point
func New(x, y float64) *Point {
	return &Point{x, y}
}

// Distance возвращает расстояние между двумя точками
func (p Point) Distance(that *Point) float64 {
	// Реализация математической формулы
	x, y := p.x-that.x, p.y-that.y
	return math.Sqrt(x*x + y*y)
}

func main() {
	p1 := New(2.0, 16.0)
	p2 := New(-11.0, 7.0)

	fmt.Println(p1.Distance(p2))
}
