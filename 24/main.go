package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}

}

func (p1 Point) Distance(p2 Point) float64 {
	distance := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return float64(distance)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := Point{4, 6}
	fmt.Printf("Расстояние между двумя точками: %.1f", p1.Distance(p2))
}
