# METHODS

_A method is a function with a special receiver argument._

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

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#overview)
* [BASIC FORMAT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#basic-format)
* [PASSING PARAMETERS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#passing-parameters)
  * [PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#passing-struct-to-method-by-value-copy---struct-not-changed)
  * [PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#passing-struct-to-method-by-reference-pointer---struct-changed)
* [WHY USE METHODS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#why-use-methods)
* [EXAMPLE - SHAPES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md#example---shapes)

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

## WHY USE METHODS

You could certainly use a function in place of a method or visa
versa, so why use them?

* Implement an interface.  You you one function name to accept many
  different data types. Refer to
  [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/interfaces.md)
  cheat sheet

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
