package main

import (
	"fmt"
	"math"
)

type Square struct {
	width float64
}

func (s *Square) Area() float64 {
	return s.width * s.width
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Traingle struct {
	width  float64
	height float64
}

func (t *Traingle) Area() float64 {
	return t.height * t.width / 2
}

func Area(shape interface{}) float64 {
	switch shape := shape.(type) {
	case Square:
		return shape.width * shape.width
	case Circle:
		return shape.radius * shape.radius * math.Pi
	case Traingle:
		return shape.width * shape.height / 2
	}
	return 0
}

type Shape interface {
	Area() float64
}

func Area2(s Shape) float64 {
	return s.Area()
}

func main() {
	c := Circle{3}
	Area(c)
	fmt.Println(Area(c))
	Area2(&c)
	fmt.Println(Area2(&c))
}
