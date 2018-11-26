package main

import (
	"fmt"
	"math"
)

// PASSING VARIABLES BY VALUE (COPY) - VARIABLES NOT CHANGED
// areaRec - Return area of Rectangle
func areaRec(w, h float64) float64 {
	return w * h
}

// scaleAreaRec - Return scaled area of Rectangle
func scaleAreaRec(w, h, s float64) float64 {
	return (w * h) * s
}

// PASSING VARIABLES BY "REFERENCE" (POINTER) -  VARIABLES CHANGED
// scaleAreaRecRef - scaled sides of Rectangle
func scaleAreaRecRef(w, h, s *float64) {
	*w = *w * math.Sqrt(*s)
	*h = *h * math.Sqrt(*s)
}

func main() {

	// DECLARE & ASSIGN (INITIALIZE)
	w := 3.0
	h := 4.0
	s := 3.0

	// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
	fmt.Println("PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED")
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f\n", w, h, areaRec(w, h))
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f) * %5.2f is %6.2f\n", w, h, s, scaleAreaRec(w, h, s))
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f <- VARIABLES NOT CHANGED\n", w, h, areaRec(w, h))

	// PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED
	fmt.Println(`PASSING STRUCT BY "REFERENCE" (POINTER) -  STRUCT CHANGED`)
	fmt.Printf("   The area of the rectangle        (%6.2f * %6.2f)         is %6.2f\n", w, h, areaRec(w, h))
	fmt.Printf("   Scale by %v\n", s)
	scaleAreaRecRef(&w, &h, &s)
	fmt.Printf("   The scaled area of the rectangle (%6.2f * %6.2f)         is %6.2f <- VARIABLES CHANGED\n", w, h, areaRec(w, h))

}
