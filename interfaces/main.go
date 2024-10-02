package main

import (
	"fmt"
)

func main() {
	sq := square{sideLength: 10}
	tr := triangle{height: 10, base: 10}
	printArea(sq)
	printArea(tr)
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
