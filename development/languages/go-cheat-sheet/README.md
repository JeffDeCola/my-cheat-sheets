# GO CHEAT SHEET

```
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`go` _is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems._

This is a very abbreviated cheat-sheet highlighting the main
syntax of go. I have lots go examples in my appropriately named repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## LET'S GO

The cheat sheet is broken up into the following sections,

* [INSTALL & CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/install-and-configure.md)
  * Install
  * Configure
    * Linux
    * Bash on Ubuntu on Windows
    * Windows

* [BASIC CONCEPTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/basic-concepts.md)
  * Basic Structure of go
  * Basic Syntax
  * Packages
  * _run_
  * _build_
  * _install_

* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md)
  * Boolean
  * Numeric
  * String

* [TYPE CONVERSION / TYPE CASTING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/type-conversion-type-casting.md)

* [VARIABLES & CONSTANTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/variables-and-constants.md)
  * Declare Type and Assign Value
  * Variable
  * Constant / Literal
  * Scope Rules (Local, Global, Formal)
  * Type Inference
  * Shortcut Assignment
  * Grouping Variables

* [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/derived-data-types.md)
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md))
  * Interface (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md))

* [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md)
  * Basic Format
  * Passing Parameters by Value (_Copy_) - - PARAMETER NOT CHANGED
  * Passing Parameters by Reference (_Pointer_) - PARAMETER CHANGED

* [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/methods.md)
  * Basic Format
  * Passing Struct by Value (_Copy_) - STRUCT NOT CHANGED
  * Passing Struct by Reference (_Pointer_) - STRUCT CHANGED

* [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md)
  * Basic Format

* [CONCURRENCY / CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md)
  * _goroutines_

* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/operators.md)
  * Arithmetic (_Math_)
  * Relational (_Compare_)
  * Logical (_Boolean_)
  * Bitwise (_Bits_)
  * Assignment
  * Miscellaneous

* [CONTROL STRUCTURES / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/control-structure-flow-control.md)
  * Loops (_for loop_)
  * Range
  * Conditional Statements / Decision Making (_if_/_else_, _switch_, _defer_, _select_)

* [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/error-handling.md)

* [FORMATING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/formating-types.md)
  * Format Specifiers
  * Escape Sequences

## GO SYNTAX OVERVIEW

```go

// DATA TYPES

    // BOOLEAN
        // true, false

    // NUMERIC
        // singed - int8, int16, int32, int64
        // unsigned - uint8 (byte), uint16, uint32 (rune) uint64
        // machine - int, uint, uintptr
        // float - float32, float64
        // complex - complex64, complex128

    // STRING
        // string

// TYPE CONVERSION /  TYPE CASTING

    a:= 33
    nowAFloat = float32(a)                          // int to float

// VARIABLE

    // DECLARE TYPE
    var a string

    // ASSIGN VALUE
    a = "happy"

    // DECLARE & ASSIGN
    var a int32 = 22                                // Verbose
    var a = 22                                      // Type Inference
    a := 32                                         // Shortcut Assignment

// CONSTANT / LITERAL

    const a float32 = 3.14                          // Must have Assignment
    const a = 22                                    // Type Inference

// GROUPING VARIABLES

    // DECLARE TYPE
    var a, b string

    // ASSIGN VALUE
    a = "hello a"
    b = "hello b"

    // DECLARE & ASSIGN
    var a, b string = "hello a", "hello b"          // Verbose
    var a, b = "hello a", "hello b"                 // Type Inference
    var (                                           // Parenthesis
        a = "hello a"
        b = "hello b"
    )
    a, b := "hello a", "hello b"                    // Group Shortcut Assignment

// ARRAY

    // DECLARE TYPE
    var a [2]float32{}

    // ASSIGN VALUE
    a[1] = 1.1
    a[2] = 2.0

    // DECLARE & ASSIGN
    var a = [2]float32{1.1, 2.0}                    // Verbose
    a := [2]float32{1.1, 2.0}                       // Array Shortcut Assignment

// SLICE

    // DECLARE TYPE
    var a []float64

    // ASSIGN VALUE / ADD TO SLICE
    a = append(a, 5.7)

    // DECLARE & ASSIGN
    var a = []float32{1.1, 2.0}                    // Verbose
    a := []float32{3.4, 4.5}                       // Array Shortcut Assignment

    // ADD TO SLICE
    a := append(a, 5.7)                            // Append to different slice

// MAP

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

// STRUCT

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

// POINTER

    // CREATE A POINTER TYPE AND ASSIGN
    a := new(int)                                   // Create int pointer type
    *a = 9                                          // "contents of a is 9"

    // ASSIGN A POINTER TO A TYPE
    a := 5                                          // If we have a var int 5
    b := &a                                         // b is the "address of" a
    // a == *b (both are 5)                         // "contents of" b is a
    
    // ASSIGN A POINTER TO A STRUCT
    b := &r1                                        // From struct Rect above
    r1.w = 6.1                                      // I wish it was *r1.w
    r1.h = 5.0                                      // I wish it was *r1.h

// FUNCTION

    // PASSING PARAMETERS BY VALUE (COPY) - PARAMETER NOT CHANGED
        func name(a int) {                          // 1 in - You would never d
        func name(a, b int) int32 {                 // 2 in, 1 return
        func name(name ...int) int {                // Variadic in, 1 return
        func name(a int, b string) (x int32) {      // 2 in, 1 NAMED return
        func name(a, b int) (x int, y string) {     // 2 in, 2 NAMED return

    // PASSING PARAMETERS BY REFERENCE (POINTER) - PARAMETER CHANGED
        func name (a *int) {                        // In is a pointer, 0 return
        func name (a *int) float32 {                // In is a pointer, 1 return

// METHOD

    // PASSING STRUCT BY VALUE (COPY) - STRUCT NOT CHANGED
        // Return area
       func (r Rect) area() float32 {               // 0 in, 1 Return
            return r.w * r.h
        }
        // Return scaled area
        func (r Rect) scaleArea(s int) float32 {    // 1 in, 1 Return
            return (r.w * r.h * float32(s))
        }

    // PASSING STRUCT BY REFERENCE (POINTER) -  STRUCT CHANGED
        // Scale the struct by 2                    // 0 in, 0 return
        func (r *Rect) scaleByTwo() {               
            r.w = r.w * 2.0                         // I wish it was *r1.w
            r.h = r.h * 2.0                         // I wish it was *r1.h
        }
        // Scale the struct by s                    // 1 in, 0 return
        func (r *Rect) scaleStruct(s float64) {
            r.w = r.w * float32(s)
            r.h = r.h * float32(s)
        }
        // Scale the struct by s and return area    // 1 in, 1 return
        func (r *Rect) scaleStructArea(s float64) float32 { 
            r.w = r.w * float32(s)
            r.h = r.h * float32(s)
            return r.w * r.h
        }

// INTERFACE

    ????

// CHANNEL

    ????

// OPERATORS

    // ARITHMETIC (MATH)
        // +, i, *, /, %,  ++, --

    // RELATIONAL (COMPARE)
        // ==, !=, >, <, >=, <=

    // LOGICAL (BOOLEAN)
        // &&, ||, ! 

    // BITWISE (BITS)
        // &, |, ^, <<, >>

    // ASSIGNMENT
        // =, +=, -=, *=, /=, %=, <=, >>=, &=, ^=, |= 

    // MISCELLANEOUS
        // &, *

// CONTROL STRUCTURE / FLOW CONTROL

    // LOOP

    // IF


// ERROR HANDLING


// FORMAT SPECIFIERS

    // VALUE OF DEFAULT FORMAT
        // %v
            // %t                                   // boolean
            // %d                                   // signed
            // %d                                   // unsigned
            // %g                                   // float
            // %g                                   // complex
            // %s                                   // string
            // %p                                   // pointer
            // %p                                   // channel
        // %#v                                      // Rep of the value
        // %T                                       // Rep of the type of the value
        // %%	                                    // a literal percent sign

    // BOOLEAN
        // %t
    
    // NUMERIC
        // %b, %c, %d, %o, %q, %x, %V, %U           // signed, unsigned, machine
        // %b, %e, %E, %f, %F, %g, %G               // float, complex
        
    // STRING
        // $s, %q, %x, %X

    // SLICE
        // %p

    // POINTER
        // %p

// ESCAPE SEQUENCES

    // \n                                           // newline.
    // \?                                           // The ? character.
    // \b                                           // backspace.
    // \"                                           // The " character.
```

## REFERENCES / DOCUMENTATION

* [golang.org](http://golang.org) _- Home base for everything._
* [tutorialspoint.com](https://www.tutorialspoint.com/go/go_data_types.htm)
  _- A great summary of syntax._
* [golang.org docs](https://golang.org/doc/) _- A good collection of docs._
* [A tour of go](https://tour.golang.org/welcome/1) _- A good place to start._
* [Effective go](https://golang.org/doc/effective_go.html)
  _- A must read to create great things._
* [Plugin for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
  _- I use
  [visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)
  ._
