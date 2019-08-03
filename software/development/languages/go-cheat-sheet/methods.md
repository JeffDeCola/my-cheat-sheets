# METHODS

A method is a function with a special receiver argument.

tl;dr,

```go
// BASIC FORMAT
    func (receiver) name(parameter list) return type {
// PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED
    // Return Area
    func (r Rect) area() float32 {
        return r.w * r.h
    }
    // Scale Struct & Return Area
    func (r Rect) scaleArea(s float32) float32 {
        r.w = r.w * s
        r.h = r.h * s
        return r.area()
    }
    func main() {
        r1 := Rect{2, 4}
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area())       // Rect Area: 2 x 4 = 8
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.scaleArea(3)) // Rect Area: 6 x 12 = 72
    }
// PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED
    // Return Area
    func (r Rect) area() float32 {
        return r.w * r.h
    }
    // Scale Struct
    func (r *Rect) scaleStruct(s float32) {
        r.w = r.w * s
        r.h = r.h * s
    }
    func main() {
        r1 := Rect{2, 4}
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 2 x 4 = 8
        r1.scaleStruct(3)                                          // Passed by Reference
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 6 x 12 = 72
    }
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#overview)
* [BASIC FORMAT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#basic-format)
* [PASSING PARAMETERS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#passing-parameters)
  * [PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#passing-struct-to-method-by-value-copy---struct-not-changed)
  * [PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#passing-struct-to-method-by-reference-pointer---struct-changed)
* [EXAMPLE - SHAPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md#example---shapes)

## OVERVIEW

A method is a function with a special receiver argument.

Methods (unlike functions which stand on their own)
are attached to data (struct). I like to think of methods as,

* Doing something with (pass by value) a structs data.
* Doing something to (pass by "reference") a structs data.

## BASIC FORMAT

AS we learned from function, the basic format is,

```go
func (receiver) name(parameter list) return type {
func receiver identifier parameters returns
```

## PASSING PARAMETERS

As with functions, you can pass the struct by "value" or "reference".

Giving this struct,

```go
// Rect is a rectangle
type Rect struct {
    w, h float32
}
```

### PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED

Passes a copy of the struct's value and gets something back (return).
Will not change the values of the struct that was passed.

Get get something back (return),

```go
// PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED
// Return Area
func (r Rect) area() float32 {
    return r.w * r.h
}

// Scale Struct & Return Area
func (r Rect) scaleArea(s float32) float32 {
    r.w = r.w * s
    r.h = r.h * s
    return r.area()
}

func main() {
    r1 := Rect{2, 4}
    fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area())       // Rect Area: 2 x 4 = 8
    fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.scaleArea(3)) // Rect Area: 6 x 12 = 72
}
```

### PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED

Again, go only passes by value, hence the quotes on "reference".

Passes the reference (pointer) of the struct so we can change
the values of the struct itself (return not necessary),

Work on struct itself (return optional),

```go
// PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED
// Return Area
func (r Rect) area() float32 {
    return r.w * r.h
}

// Scale Struct
func (r *Rect) scaleStruct(s float32) {
    r.w = r.w * s
    r.h = r.h * s
}

func main() {
    r1 := Rect{2, 4}
    fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 2 x 4 = 8
    r1.scaleStruct(3)                                          // Passed by Reference
    fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 6 x 12 = 72
}
```

Again you do this because you want to actually modify whatever you're passing
(the receiver) (“read/write” as opposed to just “read”)

## EXAMPLE - SHAPES

This example will be used for function, method,
interface and package sections.

Print the area, volume, circumference, perimeter, surface volume of
shapes such as circles, rectangles, triangles and cylinders.

The major thing to notice is how our methods are attached to structs.
So instead of passing info to a function, you only tell the method what
struct you are using, e.g. `rectangle1.perimRectangle()`.

```go

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
```
