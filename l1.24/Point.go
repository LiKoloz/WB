package main

import (
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(o Point) float64 {
	return math.Sqrt(math.Pow(o.x-p.x, 2) + math.Pow(o.y-p.y, 2))
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}
