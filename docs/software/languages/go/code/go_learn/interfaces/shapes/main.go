package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	t := triangle{30, 12}
	printArea(t)

	s := square{30}
	printArea(s)

}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
