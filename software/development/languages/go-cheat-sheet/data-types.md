# DATA TYPES

_Types or data types are a classification of data,
that tells the compiler how to use that data.
It also tells the computer how much space the data type
occupies and how the bit pattern is stored._

tl;dr,

```go
// BOOLEAN
    // true, false

// NUMERIC
    // singed - int8, int16, int32 (rune), int64
    // unsigned - uint8 (byte), uint16, uint32, uint64
    // machine - int, uint, uintptr
    // float - float32, float64
    // complex - complex64, complex128

// STRING
    // string (immutable array of bytes (or runes))
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#overview)
* [BOOLEAN](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#boolean)
* [NUMERIC](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#numeric)
  * [INTEGER](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#integer)
  * [FLOATING POINT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#floating-point)
  * [COMPLEX NUMBERS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#complex-numbers)
* [STRING](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#string)

## OVERVIEW

The data types in go are,

* **BOOLEAN**
* **NUMERIC**
* **STRING**
* Derived (see this
  [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/derived-data-types.md))
  * **ARRAY** (Data Structure) (_new_)
  * **SLICE** (Data Structure, Reference Type) (_make_)
  * **MAP** (Data Structure, Reference Type) (_make_)
  * **STRUCT** (Data Structure)
  * **POINTER**
  * **FUNCTION AS A TYPE**
  * **INTERFACE** (see this
    [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md))
  * **CHANNEL** (Reference Type) (_make_) (see this
    [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md))

## BOOLEAN

* `true`
* `false`

Boolean uses [logical operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#logical-boolean).

* `&&`, `||`, `!`

[Relational operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#relational-compare)
(compare) returns a boolean (true or false).

* `==`, `!=`, `<`, `>`, `>=`, `<=`

## NUMERIC

Integers, floating point and complex numbers.

### INTEGER

Based on bit size and sign.

* Signed
  * `int8`
  * `int16`
  * `int32` (`rune`)
  * `int64`

* Unsigned
  * `uint8` (`byte`)
  * `uint16`
  * `uint32`
  * `uint64`

* Machine Dependent Types (Depends on your machine)
  * `int` _(The most popular and inferred)_
  * `uint`
  * `uintptr`

Ranges of some integer types,

```txt
`uint16` is 2^16-1 (0 to 65,535).
`int16` is 2^15-1 (-32,768 to 32,767).
`uint64` is 2^65-1 (0 to 18,446,744,073,709,551,615).
`int64` is 2^64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).
```

`Rune` is nice because it's a 4 bytes which is a character
of any language in the world.  UTF-8 uses 1-4 bytes where each number corresponds
to a character.

Use single quote for runes,

```go
fmt.Println('e') // Prints 105
```

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

### COMPLEX NUMBERS

* Complex number types
  * `complex64`
  * `complex128`

## STRING

`string` types are an immutable array of bytes (or runes).
They are just numbers.

* A `back-tick` string can contain newlines and double quotes.
* A `double quotes` string can contain special characters.

Since they are just arrays, you can index into a string.

As an example `e` is the rune number 101. Hence,

```go
    a := "hello"
    fmt.Println(string(a[1])) // Prints e
    fmt.Println(a[1])         // Prints 101
    fmt.Println('e')          // Prints 101
    fmt.Println("e")          // Prints e
```
