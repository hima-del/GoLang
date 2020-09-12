package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length  int
	breadth int
}

type circle struct {
	radius float64
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rectangle1 := rectangle{
		length:  4,
		breadth: 10,
	}
	fmt.Printf("area of rectangle is %d\n", rectangle1.area())
	circle1 := circle{
		radius: 2,
	}
	fmt.Printf("area of circle is %f", circle1.area())
}
