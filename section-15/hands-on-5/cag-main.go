package main

import (
	"fmt"
	"math"
	"strings"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (s square) perimeter() float64 {
	return 2*s.length + 2*s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

type shape interface {
	area() float64
	perimeter() float64
}

func info(s shape) {

	stype := strings.Split(fmt.Sprintf("%T", s), ".")[1]
	fmt.Printf("Area of %s is %6.3f\n", stype, s.area())
	fmt.Printf("Perimeter of %s is %6.3f\n", stype, s.perimeter())
}

func main() {

	s1 := square{
		length: 2.5,
		width:  3.5,
	}

	c1 := circle{
		radius: 2.75,
	}

	fmt.Printf("Square dims are %v, area is %6.3f and perim is %6.3f\n", s1, s1.area(), s1.perimeter())
	fmt.Printf("Circle dims are %v, area is %6.3f and perim is %6.3f\n", c1, c1.area(), c1.perimeter())
	fmt.Println()
	info(s1)
	info(c1)

}
