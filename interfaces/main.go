package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

//isoceles triangle
type triangle struct {
	side float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (t triangle) area() float64 {
	return .5 * t.side * t.side
}

func (t triangle) perimeter() float64 {
	return 3 * t.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("Object: %+v\n", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}

func main() {
	r := rectangle{width: 2, height: 3}
	t := triangle{side: 9}
	c := circle{radius: 5}

	measure(r)
	measure(t)
	measure(c)
}
