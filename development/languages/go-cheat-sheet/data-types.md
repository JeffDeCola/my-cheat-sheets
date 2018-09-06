# DATA TYPES

Types or data types are a classification of data,
that tells the compiler how to use that data.

Tells the computer how much space it occupies and how the bit
pattern is stored.

The Data Types in go,

* Boolean
* Numeric
* String
* Derived
  * Array
  * Slice 
  * Map
  * Struct
  * Pointer
  * Function (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md))
  * Interface (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md)) 
  * Channel (See own[Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md)

## OVERVIEW

```go
// VARIABLES
var vA int32 = 1                                // Verbose
var vB = 2                                      // Type Inference
vC := 3                                         // Shortcut assignment

// GROUPING VARIABLES
var gA, gB string = "hello gA", "hello gB"      // Verbose
var gC, gD = "hello gC", "hello gD"             // Type Inference
var (                                           // Neater Form
	gE = "hello gE"
	gF = "hello gF"
)
gG, gH := "hello gG", "hello gH"                // Group Shortcut Assignment

// ARRAYS
aA := [2]float32{1.1, 2.0}                      // Assign

// SLICE
sA := make([]int, 5, 10)                        // Make
sA[0] = 3                                       // Assign
sA = append(sA, 5, 6)                           // Append to slice
sB := []float32{3.4, 4.5}                       // Create another slice
sB = append(sB, 5.7)                            // Append to slice

// MAP
mA := map[string]int                            // A Little database of key:value pairs
	"Jill": 23,
	"Bob":  34,
	"Mark": 28,
}

// STRUCT
type Rect struct {                              // Create your struct type
	w, h float32
}
var r1 Rect                                     // Declare type Rect
r1.w = 5.5	                                    // Assign width
r1.h = 2.2                                      // Assign height

// POINTER
????

// FUNCTION
func name (a int) {
func name (a, b int) int32 {     
func name (a int, b string) (c int32) {

// INTERFACE
????
```

## BOOLEAN

* `true`
* `false`

Boolean uses logical operators
* `&&`, `||`, `!`

Relational Operators returns a boolean (true or false)
*  `==`, `!=`, `<`, `>`, `>=`, `<=`

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

### COMPLEX NUMBERS

* Complex number types,
  * `complex64`
  * `complex128`

## STRING

`string` types are immutable array of bytes (or runes).

A back-tick string can contain newlines.

A double quotes can contain special characters.

Since they are just arrays, you can index into a string.

## DERIVED

### ARRAY

Really just an array of a variable, a data structure.

The basic format is,

```
name = [number]type{assignment, assignment....}
testscores := [3]float64{78.3, 98.9, 85.4}
```

All arrays are zero based so `testscores[0]` is 78.3.

For example,

```go
testscores := [3]float64{78.3, 98.9, 85.4}
fmt.Printf("a is %v\n", testscores)
fmt.Printf("a is %v\n", testscores[1])
fmt.Printf("a is %v\n", testscores[1:2])
fmt.Printf("a is %v\n", testscores[:3])
fmt.Printf("a is %v\n", testscores[2:])
```

And output is,

```go
a is [78.3 98.9 85.4]
a is 98.9
a is [98.9]
a is [78.3 98.9 85.4]
a is [85.4]
```

Remember a string is an array of bytes. Here is a good example,

```go
func main() {
	testString := "Happy Birthday Jeff"
	fmt.Printf("The first word of testString is %s\n", firstWord(testString))
}

func firstWord(str string) (word []byte) {
	for i := range str {
		if str[i] == ' ' {
			break
		} else {
			word = append(word, str[i])
		}
	}
	return word
}
```

### SLICE (ARRAY SUB TYPE) (make)

Making an initial array of n length long out of a
total capacity,

```go
// initial slice is 10 long out of 100
a := make([]int, 10, 100)
fmt.Printf("a is %v\n", a)

// adding to a slice
a = append(a, 20)
fmt.Printf("a is %v\n", a)
```

Assignment, rather then using make,

```go
c := []int{5, 6}
fmt.Println(c)
c = append(c, 7)
fmt.Println(c)
```

Output is,

```
[5 6]
[5 6 7]
```

### MAP (key:value)

Really key/value pairs, like a database.

```go
ages := map[string]int{
	"Jill": 23,
	"Bob":  34,
	"Mark": 28,
}

fmt.Println(ages["Jill"])
```

### STRUCT

Elements of different types and start
with capital letter.  Because
anything with a capital letter is exported.

type Rect struct {
    w,h,float64
}

### POINTER

### FUNCTION

See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md).

### INTERFACE

See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md).

### CHANNEL

See own[Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md).

