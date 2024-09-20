package main

import "fmt"

type shape interface{
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sidelength float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sidelength: 10}

	printArea(t)
	printArea(s)
}

func (ti triangle) getArea() float64 {
	return 0.5 * ti.base * ti.height
}

func (sq square) getArea() float64 {
	return sq.sidelength * sq.sidelength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}