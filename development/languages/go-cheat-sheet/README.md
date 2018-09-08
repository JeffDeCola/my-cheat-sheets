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

* [DECLARATION & ASSIGNMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/declaration-and-assignment.md)
  * Overview
  * Variables
  * Scope Rules (Local, Global, Formal)
  * Constants (Literals)
  * Type Inference
  * Shortcut Assignment
  * Grouping Variables

* [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md)
  * Array
  * Slice (_make_)
  * Map (_key:value_)
  * Struct
  * Pointer (_new_)
  * Function (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md))
  * Interface (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md))

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

* [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md)
  * Basic Format
  * Passing Parameters by Reference
  * Passing Parameters by Value

* [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/methods.md)
  * Basic Format

* [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md)
  * Basic Format

* [CONCURRENCY / CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md)
  * _goroutines_

* [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/error-handling.md)

* [FORMATING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/formating-types.md)
  * Format Specifiers
  * Escape Sequences

## GO SYNTAX OVERVIEW

```go
// VARIABLES

    // DECLARE                                    
    var a string

    // ASSIGN
    a = "happy"

    // DECLARE & ASSIGN
    var a int32 = 22                                // Verbose
    var a = 22                                      // Type Inference
    a := 32                                         // Shortcut Assignment

// CONSTANT

    const a float32 = 3.14                          // Must have assignment

// GROUPING VARIABLES

    // DECLARE
    var a, b string

    // ASSIGN
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

// ARRAYS
    
    // DECLARE
    var a [2]float32

    // ASSIGN
    a[1] = 1.1
    a[2] = 2.0

    // DECLARE & ASSIGN
	var a = [2]float32{1.1, 2.0}                    // Verbose
    a := [2]float32{1.1, 2.0}                       // Array Shortcut Assignment

// SLICE
    
    // DECLARE
    var a []float64

    // ASSIGN
    a = append(a, 5.7) 
    
    // DECLARE & ASSIGN
    var a = []float32{1.1, 2.0}                    // Verbose
    a = append(a, 5.7)                             // Append to same slice
    a := []float32{3.4, 4.5}                       // Array Shortcut Assignment
    b := append(a, 5.7)                            // Append to different slice

// MAP

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

// STRUCT

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


// POINTER
    
    // DECLARE
    ??

    // ASSIGN
    ??

    // DECLARE & ASSIGN
    ??

// FUNCTION
    
    func name (a int) {
    func name (a, b int) int32 {
    func name (a int, b string) (c int32) {

// INTERFACE
    
    ????

// CHANNEL
    
    ????
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
