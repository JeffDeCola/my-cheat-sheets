package main

import (
	"fmt"
	"math"
)

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
func (c Circle) areaCircle() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference
func (c Circle) circCircle() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle area
func (r Rectangle) areaRectangle() float64 {
	return r.width * r.height
}

// Rectangle perimeter
func (r Rectangle) perimRectangle() float64 {
	return (r.width + r.height) * 2
}

// Triangle area
func (t Triangle) areaTriangle() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.a + t.b + t.c) / 2)
	return math.Sqrt(s * (s - t.a) * (s - t.a) * (s - t.a))
}

// Triangle perimeter
func (t Triangle) perimTriangle() float64 {
	return t.a + t.b + t.c
}

// Cylinder volume
func (c Cylinder) volCylinder() float64 {
	return math.Pi * math.Pow(c.radius, 2) * c.height
}

// Cylinder surface area
func (c Cylinder) surfaceCylinder() float64 {
	return (2 * math.Pi * c.radius * c.height) + (2 * math.Pi * math.Pow(c.radius, 2))
}

func main() {

	// Declare and assign
	circle1 := Circle{5}
	rectangle1 := Rectangle{5, 3}
	triangle1 := Triangle{4, 5, 6}
	cylinder1 := Cylinder{5, 3}

	// Get shape properties
	areaCircle1 := circle1.areaCircle()
	circCircle1 := circle1.circCircle()
	areaRectangle1 := rectangle1.areaRectangle()
	perimRectangle1 := rectangle1.perimRectangle()
	areaTriangle1 := triangle1.areaTriangle()
	perimTriangle1 := triangle1.perimTriangle()
	volumeCylinder1 := cylinder1.volCylinder()
	surfaceCylinder1 := cylinder1.surfaceCylinder()

	fmt.Println(circle1.radius, areaCircle1, circCircle1)
	fmt.Println(rectangle1.width, rectangle1.height, areaRectangle1, perimRectangle1)
	fmt.Println(triangle1.a, triangle1.b, triangle1.c, areaTriangle1, perimTriangle1)
	fmt.Println(cylinder1.radius, cylinder1.height, volumeCylinder1, surfaceCylinder1)

	fmt.Printf("Circle1    (radius %.2f)                 area is %10.3f, circumference is %10.3f\n",
		circle1.radius, areaCircle1, circCircle1)
	fmt.Printf("Rectangle1 (width %.2f, height %.2f)     area is %10.3f, perimeter is     %10.3f\n",
		rectangle1.width, rectangle1.height, areaRectangle1, perimRectangle1)
	fmt.Printf("Triangle1  (a %.2f, b %.2f, c %.2f)      area is %10.3f, perimeter is     %10.3f\n",
		triangle1.a, triangle1.b, triangle1.c, areaTriangle1, perimTriangle1)
	fmt.Printf("Cylinder1  (radius %.2f, height %.2f)    vol  is %10.3f, surface area is  %10.3f\n",
		cylinder1.radius, cylinder1.height, volumeCylinder1, surfaceCylinder1)
}
