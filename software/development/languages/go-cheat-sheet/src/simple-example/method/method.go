package main

import (
	"fmt"
)

// Rect - CREATE STRUCT TYPE
type rect struct {
	w, h float32
}

// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED

// Return area of rectangle
func (r Rect) area() float32 {
	return r.w * r.h
}

// Return scaled area of rectangle
func (r Rect) scaleArea(s int) float32 {
	return (r.w * r.h * float32(s))
}

// PASSING STRUCT BY REFERENCE (POINTER) -  STRUCT CHANGED

// Scale the struct by 2
func (r *Rect) scaleAreasByTwo() {
	s := 2
	r.w = ((r.w * r.h * float32(s)) / r.h)
	r.h = ((r.w * r.h * float32(s)) / r.w)
}

// Scale the struct by s of rectangle
func (r *Rect) scaleAreaByS(s float64) {
	r.w = ((r.w * r.h * float32(s)) / r.h)
	r.h = ((r.w * r.h * float32(s)) / r.w)
}

// Scale the struct by s and return area    // 1 in, 1 return
func (r *Rect) scaleAreaReturn(s float64) float32 {
	r.w = ((r.w * r.h * float32(s)) / r.h)
	r.h = ((r.w * r.h * float32(s)) / r.w)
	return r.w * r.h
}

func main() {

	// DECLARE & ASSIGN
	r1 := Rect{2.0, 3.0} // Shortcut Assignment
	fmt.Println("The Rect struct:         ", r1)

	// PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
	fmt.Println("PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED")
	fmt.Println("   area():               ", r1.area())
	fmt.Println("   scaleArea(3):         ", r1.scaleArea(3))
	fmt.Println("   area() not changed:   ", r1.area())

	// PASSING STRUCT BY REFERENCE (POINTER) -  STRUCT CHANGED
	fmt.Println("PASSING STRUCT BY REFERENCE (POINTER) -  STRUCT CHANGED")
	r1.scaleAreasByTwo()
	fmt.Println("    scaleAreasByTwo():    ", r1.area())
	r1.scaleAreaByS(3)
	fmt.Println("    scaleAreaByS(3):      ", r1.area())
	fmt.Println("    scaleAreaReturn(4):   ", r1.scaleAreaReturn(4)) // It has a return
	fmt.Println("    area():               ", r1.area())
}
