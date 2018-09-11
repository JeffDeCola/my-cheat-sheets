# METHODS

Methods, unlike functions which stand on their own,
are attached to data. 

A method is a function with a special receiver argument.

That's a great definitions, but I like to think of methods as
`doing something with (pass value) or to (pass reference) a structs data`.

## BASIC FORMATS

The basic format is,

```
func (receiver) name(parameter list optional) return type optional {
    stuff
}
```

## PASSING PARAMETERS

As with functions, you can pass the struct by value or reference.

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

Get something back (return),

```go
// Return area
func (r Rect) area() float32 {
	return r.w * r.h
}

// Return scaled area
func (r Rect) scaleArea(s int) float32 {
	return (r.w * r.h * float32(s))
}
```

Where the output is,

```go
// PASS STRUCT TO METHOD BY VALUE - STRUCT NOT CHANGED
// Get something back (return).

// Pass the method some data (the struct).
fmt.Println("The area of", r1.w, "and", r1.h, "is", r1.area())

// Pass the method some data (the struct), plus parameters.
fmt.Println("The area of", r1.w, "and", r1.h, "scaled by 3 is", r1.scaleArea(3))
```

### PASSING STRUCT TO METHOD BY REFERENCE (POINTER) - STRUCT CHANGED

Passes the reference (pointer) of the struct so we can change
the values of the struct itself (return not necessary),

Work on struct itself (return optional),

NOTE: For some reason, struct pointer don't need to use the syntax
`*r.w` or `*r.h`.  I'm not sure I like this since
I like to know I'm working on the "contents of" a pointer.

```go
// Scale the struct by 2
func (r *Rect) scaleByTwo() {
	r.w = r.w * 2.0
	r.h = r.h * 2.0
}

// Scale the struct by s
func (r *Rect) scaleStruct(s float64) {
	r.w = r.w * float32(s)
	r.h = r.h * float32(s)
}

// Scale the struct by s and return the area
func (r *Rect) scaleStructArea(s float64) float32 {
	r.w = r.w * float32(s)
	r.h = r.h * float32(s)
	return r.w * r.h
}
```

Where the output is,

```go
// PASS STRUCT TO METHOD BY REFERENCE - STRUCT CHANGED
// Work on struct itself (return optional).

// Pass the method address/pointer of struct.
r1.scaleByTwo()
fmt.Println("The area of", r1.w, "and", r1.h, "is", r1.area())

// Pass the method address/pointer of struct, plus parameters.
r1.scaleStruct(2.0)
fmt.Println("The area of", r1.w, "and", r1.h, "is", r1.area())

// Pass method address/pointer of struct, plus parameters and return.
area := r1.scaleStructArea(2.0)
fmt.Println("The area of", r1.w, "and", r1.h, "is", area)
```
