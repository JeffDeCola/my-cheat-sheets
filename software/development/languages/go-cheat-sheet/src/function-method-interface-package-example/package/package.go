package main

import (
	"fmt"

	"github.com/JeffDeCola/test/package/shapes"
)

func main() {

	// Declare and assign to struct
	// I like doing this extra step, but its up to you
	circle1 := shapes.Circle{Radius: 5}
	rectangle1 := shapes.Rectangle{Width: 5, Height: 3}
	// triangle1a := Triangle{4, 5, 6}
	// cylinder1a := Cylinder{5, 3}

	// Lets declair an interface type and assign the proper struct
	var circle1er shapes.Describe2Der = circle1
	var rectangle1er shapes.Describe2Der = rectangle1
	var triangle1er shapes.Describe2Der = shapes.Triangle{A: 3, B: 4, C: 5}
	var cylinder1er shapes.Describe3Der = shapes.Cylinder{Radius: 5, Height: 3}

	// Now to get shape properties,
	// Just use the interface and method name
	// Where interface already has the values of the methods
	areaCircle1 := circle1er.Area()
	circCircle1 := circle1er.Perim()
	areaRectangle1 := rectangle1er.Area()
	perimRectangle1 := rectangle1er.Perim()
	areaTriangle1 := triangle1er.Area()
	perimTriangle1 := triangle1er.Perim()
	volumeCylinder1 := cylinder1er.Volume()
	surfaceCylinder1 := cylinder1er.Surface()

	fmt.Println(areaCircle1, circCircle1)
	fmt.Println(areaRectangle1, perimRectangle1)
	fmt.Println(areaTriangle1, perimTriangle1)
	fmt.Println(volumeCylinder1, surfaceCylinder1)

	// Note how there is no way to get the values of the Triangle and Cylinder
	fmt.Printf("Circle1    (radius %.2f)                 area is %10.3f, circumference is %10.3f\n",
		circle1.Radius, areaCircle1, circCircle1)
	fmt.Printf("Rectangle1 (width %.2f, height %.2f)     area is %10.3f, perimeter is     %10.3f\n",
		rectangle1.Width, rectangle1.Height, areaRectangle1, perimRectangle1)
	fmt.Printf("Triangle1  (a %.2f, b %.2f, c %.2f)      area is %10.3f, perimeter is     %10.3f\n",
		3.0, 4.0, 5.0, areaTriangle1, perimTriangle1)
	fmt.Printf("Cylinder1  (radius %.2f, height %.2f)    vol  is %10.3f, surface area is  %10.3f\n",
		5.0, 3.0, volumeCylinder1, surfaceCylinder1)

}
