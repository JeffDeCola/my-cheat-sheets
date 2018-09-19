package main

import (
	"fmt"
	"math"
)

// Describe2Der describes 2D shapes
type Describe2Der interface {
	area() float64
	perim() float64
}

// Describe3Der describes 3D shapes
type Describe3Der interface {
	volume() float64
	surface() float64
}

// Circle description
type Circle struct {
	radius float64
}

// Rectangle description
type Rectangle struct {
	width  float64
	height float64
}

// Triangle description
type Triangle struct {
	a float64
	b float64
	c float64
}

// Cylinder description
type Cylinder struct {
	radius float64
	height float64
}

// Circle area
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference (will call it perim for this example)
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle area
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Rectangle perimeter
func (r Rectangle) perim() float64 {
	return (r.width + r.height) * 2
}

// Triangle area
func (t Triangle) area() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.a + t.b + t.c) / 2)
	return math.Sqrt(s * (s - t.a) * (s - t.a) * (s - t.a))
}

// Triangle perimeter
func (t Triangle) perim() float64 {
	return t.a + t.b + t.c
}

// Cylinder volume
func (c Cylinder) volume() float64 {
	return math.Pi * math.Pow(c.radius, 2) * c.height
}

// Cylinder surface area
func (c Cylinder) surface() float64 {
	return (2 * math.Pi * c.radius * c.height) + (2 * math.Pi * math.Pow(c.radius, 2))
}

func describe2D(d Describe2Der) {
	fmt.Printf("2D - area is %10.3f, circumference is %10.3f\n",
		d.area(), d.perim())
}

func describe3D(d Describe3Der) {
	fmt.Printf("3D - vol  is %10.3f, surface area is  %10.3f\n",
		d.volume(), d.surface())
}

func main() {

	// Declare and assign to struct instance
	circle1 := Circle{5}
	rectangle1 := Rectangle{5, 3}
	triangle1 := Triangle{4, 5, 6}
	cylinder1 := Cylinder{5, 3}

	// The `circle` and `rectangle` struct types both
	// implement the `Describe2Der` interface so we can use
	// instances of these structs as arguments to `desscribe2D`.
	describe2D(circle1)
	describe2D(rectangle1)
	describe2D(triangle1)
	describe3D(cylinder1)

}
