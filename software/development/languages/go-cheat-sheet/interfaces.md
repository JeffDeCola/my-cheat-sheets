# INTERFACES

Interfaces are verbs, they do something.

Syntactic way to have multiple structs do the same thing differently.
Implementing the same verb in a different way.

It makes your code cleaner.

From using methods like,

```go
r.areaRectangle()
t.areaTriangle()
c.areaCircle()
```

To a lot cleaner code,

```go
r.area()
t.area()
c.area()
```

## BASIC FORMAT

End interface names with `er`.

```go
// CREATE INTERFACE TYPE
type namer interface {
    methodName()
    ...
}
```

Should end in `er` so you know its an interface.

## MAKING YOUR CODE CLEANER

Giving the following structs,

```go
type Circle struct {
    radius float64
}

type Cylinder struct {
    radius float64
    height float64
}
```

### WITHOUT INTERFACE

Make a method for each property of the shape you wanted
(e.g. area, volume, circumference, etc...)

```go
// Circle area
func (c Circle) areaCircle() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

// Circle circumference
func (c Circle) circCircle() float64 {
    return 2 * math.Pi * c.radius
}

// Cylinder volume
func (c Cylinder) volCylinder() float64 {
    return math.Pi * math.Pow(c.radius, 2) * c.height
}

// Cylinder surface
func (c Cylinder) surfaceCylinder() float64 {
    return (2 * math.Pi * c.radius * c.height) + (2 * math.Pi * math.Pow(c.radius, 2))
}
```

And assign values to them as,

```go
    // Declare and assign
    circle1 := Circle{5}
    cylinder1 := Cylinder{5, 3}

    // Get properties of shapes
    areaCircle1 := circle1.areaCircle()
    circCircle1 := circle1.circCircle()
    volumeCylinder1 := cylinder1.volCylinder()
    surfaceCylinder1 := cylinder1.surfaceCylinder()

    fmt.Println(circle1.radius, areaCircle1, circCircle1)
    fmt.Println(cylinder1.radius, cylinder1.height, volumeCylinder1, surfaceCylinder1)
```

### WITH INTERFACE

But we can make the code neater if we make an interface,

```go
// CREATE INTERFACE TYPE
type Describer interface {
    describe()
}
```

And create the methods attached to the interface as,

```go
func (c Circle) describe() (area float64, circ float64) {
    area = math.Pi * math.Pow(c.radius, 2)
    circ = 2 * math.Pi * c.radius
    return
}

func (c Cylinder) describe() (volume float64, circ float64) {
    volume = math.Pi * math.Pow(c.radius, 2) * c.height
    circ = 2 * math.Pi * c.radius
    return
}
```

And assign values to them as,

```go
myCircle := Circle{5}
myCylinder := Cylinder{5, 3}

areaCircle, circCircle := myCircle.describe()
volCylinder, circCylinder := myCylinder.describe()


```

## EXAMPLE - SHAPES

This example will be used for function, method,
interface and package sections.

Print the area, volume, circumference, perimeter, surface volume of
shapes such as circles, rectangles, triangles and cylinders.

Notice there are two interfaces, one for 2D object and one
for 3D objects. We declare the interface and assign the proper struct.
e.g. `var triangle1er Describe2Der = Triangle{3, 4, 5}`

```go
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

func main() {

    // Declare and assign to struct
    // I like doing this extra step, but its up to you
    circle1 := Circle{5}
    rectangle1 := Rectangle{5, 3}
    // triangle1 := Triangle{4, 5, 6}
    // cylinder1 := Cylinder{5, 3}

    // Lets declare an interface and assign the proper struct
    var circle1er Describe2Der = circle1
    var rectangle1er Describe2Der = rectangle1
    var triangle1er Describe2Der = Triangle{3, 4, 5}
    var cylinder1er Describe3Der = Cylinder{radius: 5, height: 3}

    // Now to get shape properties,
    // Just use the interface and method name
    // Where interface has the values of the methods
    areaCircle1 := circle1er.area()
    circCircle1 := circle1er.perim()
    areaRectangle1 := rectangle1er.area()
    perimRectangle1 := rectangle1er.perim()
    areaTriangle1 := triangle1er.area()
    perimTriangle1 := triangle1er.perim()
    volumeCylinder1 := cylinder1er.volume()
    surfaceCylinder1 := cylinder1er.surface()

    fmt.Println(areaCircle1, circCircle1)
    fmt.Println(areaRectangle1, perimRectangle1)
    fmt.Println(areaTriangle1, perimTriangle1)
    fmt.Println(volumeCylinder1, surfaceCylinder1)

    // Note how there is no way to get the values of the Triangle and Cylinder
    fmt.Printf("Circle1    (radius %.2f)                 area is %10.3f, circumference is %10.3f\n",
        circle1er, areaCircle1, circCircle1)
    fmt.Printf("Rectangle1 (width %.2f, height %.2f)     area is %10.3f, perimeter is     %10.3f\n",
        rectangle1.width, rectangle1.height, areaRectangle1, perimRectangle1)
    fmt.Printf("Triangle1  (a %.2f, b %.2f, c %.2f)      area is %10.3f, perimeter is     %10.3f\n",
        3.0, 4.0, 5.0, areaTriangle1, perimTriangle1)
    fmt.Printf("Cylinder1  (radius %.2f, height %.2f)    vol  is %10.3f, surface area is  %10.3f\n",
        5.0, 3.0, volumeCylinder1, surfaceCylinder1)

    fmt.Sprintf("%v", circle1er["radius"])
}
```