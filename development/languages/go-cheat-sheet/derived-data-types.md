# DERIVED DATA TYPES

The derived data types are built on the basic three data types
(Boolean, Numeric and String).

The Data Types in go,

* Boolean (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Numeric (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* String (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Derived
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md))
  * Interface (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md))

## ARRAY

Really just an array of a variable; a data structure.

The basic format is,

```
name = [number]type{assignment, assignment....}
```

Here is the syntax,

```go

// DECLARE
var a [2]float32

// ASSIGN
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

## SLICE (ARRAY SUB TYPE) (make)

Making an initial array of n length long out of a
total capacity,

```go
// DECLARE
var a []float64

// ASSIGN
a = append(a, 5.7)

// DECLARE & ASSIGN
var a = []float32{1.1, 2.0}                    // Verbose
a = append(a, 5.7)                             // Append to same slice
a := []float32{3.4, 4.5}                       // Array Shortcut Assignment
b := append(a, 5.7)                            // Append to different slice
```

## MAP (key:value)

Really key/value pairs, like a database.

```go
// DECLARE
var a = make(map[string]int)
a := make(map[string]int)

// ASSIGN
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

Elements of different types and start
with capital letter.  Because
anything with a capital letter is exported.

```go
// CREATE
type Rect struct {
    w, h float32
}

// DECLARE
var r1 Rect

// ASSIGN
r1.w = 6.1
r1.h = 5.0

// DECLARE & ASSIGN
var r Rect = Rect{6.1, 5.0}                     // Verbose
var r = Rect{6.1, 5.0}                          // Type Inference
r := Rect{6.1, 5.0}                             // Shortcut Assignment
```

## POINTER

A pointer is just a variable that holds the address (of memory)
of a value.

`*` is the "contents of"
`&` is "address of"

You can,

```go
// CREATE A POINTER TYPE AND ASSIGN
a := new(int)                                   // Create int pointer type
*a = 9                                          // "Contents of a is 9"

// ASSIGN A POINTER TO A TYPE
a := 5                                          // If we have a var int 5
b := &a                                         // b is the "address of" a
// a == *b (both are 5)                         // "Contents of" b is a
```

![IMAGE - go pointers - IMAGE](../../../docs/pics/go-pointers.jpg)

Cool uses for pointers are,

* Allocate space for a variable.
* Pass by reference to a function to change parameters value outside function.
