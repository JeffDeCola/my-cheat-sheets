# INTERFACES (SET OF METHOD SIGNATURES)

tl;dr,

```go
// CREATE INTERFACE TYPE
// I like to end interface names with er
type namer interface {
    methodName1()
    methodName2()
    ...
}
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#overview)
* [BASIC FORMAT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#basic-format)
* [MAKING YOUR CODE CLEANER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#making-your-code-cleaner)
  * [WITHOUT INTERFACE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#without-interface)
  * [WITH INTERFACE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#with-interface)
* [EXAMPLE - SHAPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md#example---shapes)

## OVERVIEW

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
    methodName1()
    methodName2()
    ...
}
```

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

This example will calculate the **area** and **perimeter**
of three shapes: **circles**, **rectangles** and **triangles**.

the example is located in `my-go-examples` repo done three different ways,

* [shapes-using-functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes-using-functions)
* [shapes-using-methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/shapes-using-methods)
* [shapes-using-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-using-interfaces)

You will see that using an interface you can go from this,

```go
c1Area := circleArea(c1)
c1Perimeter := circlePerimeter(c1)
r1Area := rectangleArea(r1)
r1Perimeter := rectanglePerimeter(r1)
t1Area := triangleArea(t1)
t1Perimeter := trianglePerimeter(t1)
```

To this,

```go
// This is the power of interfaces.
c1Area = area(c1)
c1Perimeter = perimeter(c1)
r1Area = area(r1)
r1Perimeter = perimeter(r1)
t1Area = area(t1)
t1Perimeter = perimeter(t1)
```
