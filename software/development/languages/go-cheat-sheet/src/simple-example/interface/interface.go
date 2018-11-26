/*
- Changed method names from
    - areaRec to area
    - scaleAreaRec scaleArea
    - scaleAreaRecRef scaleAreaRef
- Added triangle struct type
- Added methods for triangle to match rectangle: area, scaleArea, scaleAreaRef
- Added describer interface
- Using the interface in main()
*/

package main

import (
	"fmt"
	"math"
)

// Describer describes 2D shapes
type Describer interface {
	area() float64
	scaleArea(float64) float64
	scaleAreaRef(float64)
}

// Rect - Rectangle
type rect struct {
	w, h float64
}

// Tri - Triangle
type tri struct {
	a, b, c float64
}

// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
// area - Return area of Rectangle
func (r rect) area() float64 {
	return r.w * r.h
}

// area - Return area of Rectangle
func (t tri) area() float64 {
	return t.a * t.b * t.c
}

// scaleAreaRec - Return scaled area of Rectangle
func (r rect) scaleArea(s float64) float64 {
	return r.w * r.h * s
}

// scaleArea - Return scaled area of Triangle
func (t tri) scaleArea(s float64) float64 {
	return t.a * t.b * t.c * s
}

// PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED
// scaleAreaRecRef - scaled sides of Rectangle
func (r *rect) scaleAreaRef(s float64) {
	r.w = r.w * math.Sqrt(s)
	r.h = r.h * math.Sqrt(s)
}

func (t *tri) scaleAreaRef(s float64) {
	t.a = t.a * math.Sqrt(s)
	t.b = t.b * math.Sqrt(s)
	t.c = t.c * math.Sqrt(s)
}

func main() {

	// DECLARE & ASSIGN (INITIALIZE)
	r := rect{3.0, 4.0}
	t := tri{3.0, 4.0, 5.0}
	s := 3.0

	// Initialize the interface -  Declare and assign the proper struct
	var r1 Describer = &r
	var t1 Describer = &t

	// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
	fmt.Println("PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED")
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)                  is %6.2f\n", r.w, r.h, r1.area())
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f)          * %5.2f is %6.2f\n", r.w, r.h, s, r1.scaleArea(s))
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)                  is %6.2f <- STRUCT NOT CHANGED\n", r.w, r.h, r1.area())

	fmt.Printf("   The area of the triangle         (%6.2f * %6.2f * %6.2f)         is %6.2f\n", t.a, t.b, t.c, t1.area())
	fmt.Printf("   The scaled area of the triangle  (%6.2f * %6.2f * %6.2f) * %5.2f is %6.2f\n", t.a, t.b, t.c, s, t1.scaleArea(s))
	fmt.Printf("   The area of the triangle         (%6.2f * %6.2f * %6.2f)         is %6.2f <- STRUCT NOT CHANGED\n\n", t.a, t.b, t.c, t1.area())

	// PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED
	fmt.Println(`PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED`)
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)                  is %6.2f\n", r.w, r.h, r1.area())
	fmt.Printf("   Scale by %v\n", s)
	r.scaleAreaRef(s)
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f)                  is %6.2f <- STRUCT CHANGED\n", r.w, r.h, r1.area())

	fmt.Printf("   The area of the triangle         (%6.2f * %6.2f * %6.2f)         is %6.2f\n", t.a, t.b, t.c, t1.area())
	fmt.Printf("   Scale by %v\n", s)
	t.scaleAreaRef(s)
	fmt.Printf("   The scaled area of the triangle  (%6.2f * %6.2f * %6.2f) * %5.2f is %6.2f <- STRUCT CHANGED\n", t.a, t.b, t.c, s, t1.area())
}
