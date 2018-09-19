# DATA TYPES

Types or data types are a classification of data,
that tells the compiler how to use that data.
It also tells the computer how much space the data type
occupies and how the bit pattern is stored.

The data types in go are,

* Boolean
* Numeric
* String
* Derived (_See next [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/derived-data-types.md)_)
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function Type
  * Interface Type
  * Channel Type

## BOOLEAN

* `true`
* `false`

Boolean uses [logical operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#logical-boolean).

* `&&`, `||`, `!`

[Relational operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#relational-compare)
returns a boolean (true or false).

* `==`, `!=`, `<`, `>`, `>=`, `<=`

## NUMERIC

Integers, floating point and complex numbers.

### INTEGER

Based on bit size and sign.

* Signed,
  * `int8`
  * `int16`
  * `int32` (`rune`)
  * `int64`

* Unsigned,
  * `uint8` (`byte`)
  * `uint16`
  * `uint32`
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

Rune is nice because it's a 4 bytes which is a character
of any language in the world.  UTF-8 uses 1-4 bytes.
The number corresponds to a character.

Use single quote for runes,

```go
fmt.Println('e')
```

Prints 105.

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types,
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

### COMPLEX NUMBERS

* Complex number types,
  * `complex64`
  * `complex128`

## STRING

`string` types are immutable array of bytes (or runes).
They are just numbers.

* A back-tick string can contain newlines and double quotes.
* A double quotes string can contain special characters.

Since they are just arrays, you can index into a string.

As an example e is the rune number 101. Hence,

```go
a := "hello"
fmt.Println(string(a[1]))
fmt.Println(a[1])
fmt.Println('e')
```

Prints,

```
e
101
101
```

