package shapes

import (
	"math"
)

// Describe2Der describes 2D shapes
type Describe2Der interface {
	Area() float64
	Perim() float64
}

// Describe3Der describes 3D shapes
type Describe3Der interface {
	Volume() float64
	Surface() float64
}

// Circle description
type Circle struct {
	Radius float64
}

// Rectangle description
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle description
type Triangle struct {
	A float64
	B float64
	C float64
}

// Cylinder description
type Cylinder struct {
	Radius float64
	Height float64
}

// Area is Circle area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perim is Circle circumference (will call it perim for this example)
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

// Area is Rectangle area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perim is Rectangle perimeter
func (r Rectangle) Perim() float64 {
	return (r.Width + r.Height) * 2
}

// Area is Triangle area
func (t Triangle) Area() float64 {
	// Heron's Formula to get area from 3 sides
	s := ((t.A + t.B + t.C) / 2)
	return math.Sqrt(s * (s - t.A) * (s - t.A) * (s - t.A))
}

// Perim is Triangle perimeter
func (t Triangle) Perim() float64 {
	return t.A + t.B + t.C
}

// Volume is Cylinder volume
func (c Cylinder) Volume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height
}

// Surface is Cylinder surface area
func (c Cylinder) Surface() float64 {
	return (2 * math.Pi * c.Radius * c.Height) + (2 * math.Pi * math.Pow(c.Radius, 2))
}
