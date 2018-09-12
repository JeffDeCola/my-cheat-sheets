# INTERFACES

Interfaces are verbs, they do something.

Syntactic way to have multiple structure do the same thing differently.
Implementing the same verb in a different way.

It makes your code cleaner.

## BASIC FORMAT

End interface names with `er`.

```go
type name interface {
    methodName()
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

You would make a method for each thing you wanted
to do with the data (e.g. area, volume, circumference, etc...)

```go
// Return circle area
func (c Circle) areaCircle() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

// Return circle circumference
func (c Circle) circCircle() float64 {
    return 2 * math.Pi * c.radius
}

// Return cylinder volume
func (c Cylinder) volCylinder() float64 {
    return math.Pi * math.Pow(c.radius, 2) * c.height
}

// Return cylinder circumference
func (c Cylinder) circCylinder() float64 {
    return 2 * math.Pi * c.radius
}
```

And use them as,

```go
    myCircle := Circle{5}
    myCylinder := Cylinder{5, 3}

    areaCircle := myCircle.areaCircle()
    circCircle := myCircle.circCircle()
    volCylinder := myCylinder.volCylinder()
    circCylinder := myCylinder.circCylinder()
```

### WITH INTERFACE

But we can make the code neater if we make an interface,

```go
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

And use them as,

```go
myCircle := Circle{5}
myCylinder := Cylinder{5, 3}

areaCircle, circCircle := myCircle.describe()
volCylinder, circCylinder := myCylinder.describe()
```
