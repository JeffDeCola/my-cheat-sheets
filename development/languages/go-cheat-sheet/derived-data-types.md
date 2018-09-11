# DERIVED DATA TYPES

I like to think of derived data types as
special variables built on the basic three data types
(Boolean, Numeric and String).

So you still need to declare type and assign value.

The data types in go,

* Boolean (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Numeric (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* String (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Derived (This Section)
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function Type
  * Interface Type
  * Channel Type

## ARRAY

Really just an array of a variable; a data structure.

The basic verbose format is,

```
var name = [number]type{value, value....}
```

Here is the syntax,

```go

// DECLARE TYPE
var a [2]float32{}

// ASSIGN VALUE
a[1] = 1.1
a[2] = 2.0

// DECLARE & ASSIGN
var a = [2]float32{1.1, 2.0}                    // Verbose
a := [2]float32{1.1, 2.0}                       // Array Shortcut Assignment
```

Here is printing an array,

```go
testscores := [3]float64{78.3, 98.9, 85.4}
fmt.Printf("a is %v\n", testscores)
fmt.Printf("a is %v\n", testscores[1])
fmt.Printf("a is %v\n", testscores[1:2])
fmt.Printf("a is %v\n", testscores[:3])
fmt.Printf("a is %v\n", testscores[2:])
```

The output is,

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

## SLICE (ARRAY SUB TYPE)

Making an initial array of n length long out of a
total capacity.

The basic verbose format is,

```
var name = []type{value, value....}
```

Here is the syntax,

```go
// DECLARE TYPE
var a []float64

// ASSIGN VALUE / ADD TO SLICE
a = append(a, 5.7)

// DECLARE & ASSIGN
var a = []float32{1.1, 2.0}                    // Verbose
a := []float32{3.4, 4.5}                       // Array Shortcut Assignment

// ADD TO SLICE
a := append(a, 5.7)                            // Append to different slice
```

## MAP (key:value)

Really key/value pairs, like a database.

```go
// DECLARE TYPES
var a = make(map[string]int)
a := make(map[string]int)

// ASSIGN KEY:VALUE
a["Jill"] = 23
a["Bob"] = 34
a["Mark"] = 28

// DECLARE & ASSIGN
var a = map[string]int{                        // Verbose
    "Jill": 23,
    "Bob":  34,
    "Mark": 28,
}
a := map[string]int{                           // Array Shortcut Assignment
    "Jill": 23,
    "Bob":  34,
    "Mark": 28,
}
```

## STRUCT

Elements of different types and start with capital letter.
Anything with a capital letter is exported.

```go
// CREATE STRUCT TYPE
type Rect struct {
    w, h float32
}

// DECLARE TYPE
var r1 Rect

// ASSIGN VALUE
r1.w = 6.1
r1.h = 5.0

// DECLARE & ASSIGN
var r1 Rect = Rect{6.1, 5.0}                    // Verbose
var r1 = Rect{6.1, 5.0}                         // Type Inference
r1 := Rect{6.1, 5.0}                            // Shortcut Assignment
```

## POINTER

A pointer is just a variable that holds the address (of memory)
of a value.

`*` is the `"contents of"`
`&` is `"address of"`

You can,

```go
// CREATE A POINTER TYPE AND ASSIGN
a := new(int)                                   // Create int pointer type
*a = 9                                          // "contents of a is 9"

// ASSIGN A POINTER TO A TYPE
a := 5                                          // If we have a var int 5
b := &a                                         // b is the "address of" a
// a == *b (both are 5)                         // "contents of" b is a

// ASSIGN A POINTER TO A STRUCT
b := &r1                                        // From struct Rect above
r1.w = 6.1                                      // I feel it should be *r1.w
r1.h = 5.0                                      // I feel it should be *r1.h
```

![IMAGE - go pointers - IMAGE](../../../docs/pics/go-pointers.jpg)

Cool uses for pointers are,

* Allocate space for a variable.
* Pass by reference to a function to change parameters value outside function.

## FUNCTION / CLOSURE TYPE

When used in a function, acts just like a type.
So I can use the variables int the function it lives.

```go

a, b := 3, 9

add := func() int {
    return a + b
}
```

## INTERFACE TYPE

## CHANNEL TYPE
