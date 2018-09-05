# DATA TYPES

Types or data types are a classification of data,
that tells the compiler how to use that data.

Tells the computer how much space it oocupies and how the bit
pattern is stored.

## BOOLEAN

* `true`
* `false`

Boolean uses logical operators
* `&&`, `||`, `!`_

Relational Operators returns a boolean (true or false)
*  `==`, `!=`, `<`, `>`, `>=`, `<=`_

## NUMERIC

Integers, floating point and complex numbers.

### INTEGER 

Based on bit size and sign.

* Signed,
  * `int8`
  * `int16`
  * `int32`
  * `int64`

* Unsigned,
  * `uint8` (`byte`)
  * `uint16`
  * `uint32` (`rune`)
  * `uint64`

* Machine dependent types (Depends on system),
  * `int` Most popular and inferred.
  * `uint` 
  * `uintptr`

Ranges of some integer types,

```txt
`uint16` is 2^16-1 (0 to 65,535).
`int16` is 2^15-1 (-32,768 to 32,767).
`uint64` is 2^65-1 (0 to 18,446,744,073,709,551,615).
`int64` is 2^64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).
```

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types,
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

### COMPLEX COMPLEX

* Complex number types,
  * `complex64`
  * `complex128`

## STRING

`string` types are arrays of bytes or runes.

A back-tick string can contain newlines.

A double quotes can contain special characters.

Since they are just arrays, you can index into a string.

## DERIVED

They include:

* Pointer Types 
* Array Types
* Slice Types
* Struct Types
* Union Types 
* Map Types
* Function Types (Own Section)
* Interface Types (Own Section)
* Channel Types (Own Section)

### POINTER

### ARRAY

Really just an array of a variable, a data structure.

The basic format is,

```txt
name = [number]type{assignment, assignment....}
```

For example,

```go
testscores = [3]float64{78.3, 98.9, 85.4}
```

All arrays are zero based so `testscores[0]` is 78.3.

### SLICE (ARRAY SUB TYPE)

### STRUCT

Elements of differnet types.

### UNION


### MAP

???

### FUNCTION

See own Section here

### INTERFACE

See own Section here

### CHANNEL

See own Section here

