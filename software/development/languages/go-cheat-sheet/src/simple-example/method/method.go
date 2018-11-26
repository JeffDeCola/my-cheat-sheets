/*
- Added struct type rect
- Added receiver to functions (methods)
*/

package main

import (
	"fmt"
	"math"
)

// Rect - Rectangle
type rect struct {
	w, h float64
}

// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
// areaRec - Return area of Rectangle
func (r rect) areaRec() float64 {
	return r.w * r.h
}

// scaleAreaRec - Return scaled area of Rectangle
func (r rect) scaleAreaRec(s float64) float64 {
	return r.w * r.h * s
}

// PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED
// scaleAreaRecRef - scaled sides of Rectangle
func (r *rect) scaleAreaRecRef(s float64) {
	r.w = r.w * math.Sqrt(s)
	r.h = r.h * math.Sqrt(s)
}

func main() {

	// DECLARE & ASSIGN (INITIALIZE)
	r := rect{3.0, 4.0}
	s := 3.0

	// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
	fmt.Println("PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED")
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f\n", r.w, r.h, r.areaRec())
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f) * %5.2f is %6.2f\n", r.w, r.h, s, r.scaleAreaRec(s))
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f <- STRUCT NOT CHANGED\n", r.w, r.h, r.areaRec())

	// PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED
	fmt.Println(`PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED`)
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f\n", r.w, r.h, r.areaRec())
	fmt.Printf("   Scale by %v\n", s)
	r.scaleAreaRecRef(s)
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f)         is %6.2f <- STRUCT CHANGED\n", r.w, r.h, r.areaRec())

}
