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
func areaCircle(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

// Circle circumference
func circCircle(radius float64) float64 {
	return 2 * math.Pi * radius
}

// Rectangle area
func areaRectangle(width, height float64) float64 {
	return width * height
}

// Rectangle perimeter
func perimRectangle(width, height float64) float64 {
	return (width + height) * 2
}

// Triangle area
func areaTriangle(a, b, c float64) float64 {
	// Heron's Formula to get area from 3 sides
	s := ((a + b + c) / 2)
	return math.Sqrt(s * (s - a) * (s - a) * (s - a))
}

// Triangle perimeter
func perimTriangle(a, b, c float64) float64 {
	return a + b + c
}

// Cylinder volume
func volCylinder(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}

// Cylinder surface area
func surfaceCylinder(radius, height float64) float64 {
	return (2 * math.Pi * radius * height) + (2 * math.Pi * math.Pow(radius, 2))
}

func main() {

	// Declare and assign
	circle1 := Circle{5}
	rectangle1 := Rectangle{5, 3}
	triangle1 := Triangle{4, 5, 6}
	cylinder1 := Cylinder{5, 3}

	// Get shape properties
	areaCircle1 := areaCircle(circle1.radius)
	circCircle1 := circCircle(circle1.radius)
	areaRectangle1 := areaRectangle(rectangle1.width, rectangle1.height)
	perimRectangle1 := perimRectangle(rectangle1.width, rectangle1.height)
	areaTriangle1 := areaTriangle(triangle1.a, triangle1.b, triangle1.c)
	perimTriangle1 := perimTriangle(triangle1.a, triangle1.b, triangle1.c)
	volumeCylinder1 := volCylinder(cylinder1.radius, cylinder1.height)
	surfaceCylinder1 := surfaceCylinder(cylinder1.radius, cylinder1.height)

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
