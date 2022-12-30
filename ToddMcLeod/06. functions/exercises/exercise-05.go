package main

import (
	"fmt"
	"math"
)

type square struct {
	l int
	w int
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	c := circle{
		radius: 3.00,
	}

	fmt.Println(c.area())

	s := square{
		l: 4,
		w: 4,
	}

	fmt.Println(s.area())
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return float64(s.l * s.w)
}

func area(s shape) float64 {
	return s.area()
}
