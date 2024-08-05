package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c Circle) circumference() float64 {
	return 2 * PI * c.radius
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func main() {
	circle := NewCircle(5)

	fmt.Printf(" %+v \n", circle)

	fmt.Println(circle.area())

	fmt.Println(circle.circumference())
}
