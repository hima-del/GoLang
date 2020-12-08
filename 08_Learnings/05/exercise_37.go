package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	sq := square{
		length: 10,
	}
	ci := circle{
		radius: 5.67,
	}
	info(sq)
	info(ci)
}
