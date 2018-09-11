# METHODS

Methods, unlike functions which stand on their own,
are attached to data. 

A method is a function with a special receiver argument.

That's great definitions, but I like to think of methods as
`doing something with (pass value) or to (pass reference) structs data`.

## BASIC FORMATS

The basic format is,

```txt
func (receiver) name(parameter list optional) return type optional {
    stuff
}
```

## PASSING PARAMETERS

As with functions, you can still pass the struct
by value or reference.

Giving this struct,

```go
// Rect is a rectangle
type Rect struct {
	w, h float32
}
```

### PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED

Passes a copy of the struct's value and gets something back (return).

Will not change the values of the struct that was passed.

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

### PASSING STRUCT BY REFERENCE (POINTER) - STRUCT CHANGED

Passes the reference (pointer) of the struct so we can change
the values of the struct itself (return not necessary),

NOTE: WITH STRUCT POINTERS, YOU DO NOT NEED TO HAVE
SYNTAX *r.w or *r.h.  I'm not sure I like this since
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

