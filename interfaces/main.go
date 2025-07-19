package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("area is %.2f\n", s.area())
	fmt.Printf("perimeter is %.2f", s.perimeter())
}

func main() {
	c := circle{
		radius: 5}

	info(c)

}
