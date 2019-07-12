# GO CHEAT SHEET

`go` _is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems._

This is a very abbreviated cheat-sheet highlighting the main
syntax of go. I have lots go examples in my appropriately named repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## LET'S GO

This cheat sheet is broken up into the following sections,

* [INSTALL & CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/install-and-configure.md)
  * Install
  * Configure
    * Linux
    * Bash on Ubuntu on Windows (Windows System for Linux)
    * Windows
  * Check
  * Install Go Tools

* [BASIC CONCEPTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/basic-concepts.md)
  * Basic Structure of go
  * Basic Syntax
  * Packages (See own [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md))
  * _run_
  * _build_
  * _install_
  * The Object Side of go

* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/data-types.md)
  * Boolean
  * Numeric
  * String

* [TYPE CONVERSION / TYPE CASTING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/type-conversion-type-casting.md)
  * Basic Format

* [VARIABLES & CONSTANTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/variables-and-constants.md)
  * Declare Type and Assign Value
  * Variable
  * Constant / Literal
  * IOTA
  * Scope Rules (Universe, Package, File, Block)
  * Type Inference
  * Shorthand Assignment (Preferred Method)
  * Grouping Variables

* [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/derived-data-types.md)
  * Array (Data structure) (_new_)
  * Slice (Data Structure) (Reference Type) (_make_)
  * Map (Data Structure) (Reference Type) (_make_)
  * Struct (Data Structure)
  * Pointer
  * Function (as a Type)
  * Interface (See own [Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (Reference Type) (_make_) (See own [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md))

* [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/functions.md)
  * Basic Formats
  * Variadic Functions
  * Closure (func expression & anonymous function)
    * Assign anonymous Function (func Literal) to a Variable
    * Return a Function to a Function - Closure
  * Passing a function (as an argument) to a Function - Callback
  * Passing Arguments - Go Passes by Value Only
    * Passing Arguments to Function by Value (_Copy_) - Parameter not Changed
    * Passing Arguments to Function by "Reference" (_Pointer_) - Parameter Changed
  * Recursion (function calling itself)
  * Anonymous Self Executing Function

* [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md)
  * Basic Format
  * Passing Parameters
  * Passing Struct to Method by Value (_Copy_) - Struct not Changed
  * Passing Struct to Method by Reference (_Pointer_) - Struct Changed

* [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md)
  * Basic Format
  * Making your code cleaner
    * Without Interface
    * With Interface

* [CONCURRENCY / CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md)
  * _goroutines_

* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md)
  * Arithmetic (_Math_)
  * Relational (_Compare_)
  * Logical (_Boolean_)
  * Bitwise (_Bits_)
  * Assignment
  * Miscellaneous

* [CONTROL STRUCTURES / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/control-structure-flow-control.md)
  * Loops (_for, while, infinite loops, range, break continue_)
  * Conditional Statements / Decision Making (_if_/_else_, _switch_, _defer_, _select_)

* [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/error-handling.md)

* [FORMATING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/formating-types.md)
  * Read Input (bufio.NewReader Package)
  * Scan Input (fmt.Scan Package)
  * Format Specifiers
  * Escape Sequences

* [PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md)

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

    // DECLARE & ASSIGN (INITIALIZE)
    var a int32 = 22                                // Verbose
    var a = 22                                      // Type Inference
    a := 32                                         // Shorthand Assignment

// CONSTANT / LITERAL

    const a float32 = 3.14                          // Must have Assignment
    const a = 22                                    // Type Inference

// GROUPING VARIABLES

    // DECLARE TYPE
    var a, b string

    // ASSIGN VALUE
    a = "hello a"
    b = "hello b"

    // DECLARE & ASSIGN (INITIALIZE)
    var a, b string = "hello a", "hello b"          // Verbose
    var a, b = "hello a", "hello b"                 // Type Inference
    var (                                           // Parenthesis
        a = "hello a"
        b = "hello b"
    )
    a, b := "hello a", "hello b"                    // Group Shorthand Assignment

// ARRAY

    // DECLARE TYPE
    var a [2]float32{}

    // ASSIGN VALUE
    a[1] = 1.1
    a[2] = 2.0

    // DECLARE & ASSIGN (INITIALIZE)
    var a = [2]float32{1.1, 2.0}                    // Verbose
    a := [2]float32{1.1, 2.0}                       // Array Shorthand Assignment

// SLICE (Reference Type) (_make_)

    // DECLARE TYPE - NO SIZE
    var a []float64

    // ASSIGN VALUE - ADD LENGTH TO SLICE
    a = append(a, 5.7)

    // DECLARE TYPE - WITH SIZE (make)
    m := make([]string, 1, 25)

    // ASSIGN VALUE
    m[0] = "hello"

    // DECLARE & ASSIGN (INITIALIZE)
    var a = []float32{1.1, 2.0}                     // Verbose
    a := []float32{3.4, 4.5}                        // Array Shortcut Assignment

    // ADD TO ANY SLICE
    a := append(a, 5.7)                             // Append to different slice

// MAP (Reference Type) (_make_)

    // DECLARE TYPES - THIS IS A NIL MAP - DON'T DO THIS
    var a map[string]int

    // DECLARE TYPES (make)
    var m1 = make(map[string]int)
    m2 := make(map[string]int)

    // ASSIGN KEY:VALUE
    m1["Jill"] = 23
    m1["Bob"] = 34
    m1["Mark"] = 28
    m2["Jill"], m2["Bob"], m2["Mark"] = 23, 34, 28

    // DECLARE & ASSIGN  (INITIALIZE)
    var m3 = map[string]int{                        // Verbose
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }
    m4 := map[string]int{                           // Array Shortcut Assignment
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }

    fmt.Println(m1, m2, m3, m4)

// STRUCT

    // CREATE STRUCT TYPE
    type Rect struct {
        w, h float32
        }

    // DECLARE TYPE
    var r1 Rect
    r2 := new(Rect)                                 // Returns Pointer

    // ASSIGN VALUE TO FIELDS
    r1.w = 6.1
    r1.h = 5.0
    r2.w, r2.h = 6.1, 6.0

    // DECLARE & ASSIGN (INITIALIZE)
    var r3 Rect = Rect{6.1, 5.0}                    // Verbose
    var r4 = Rect{6.1, 5.0}                         // Type Inference
    r5 := Rect{w: 6.1, h: 5.0}                      // Shorthand Assignment
    r6 := Rect{6.1, 5.0}                            // Shorthand Assignment

    fmt.Println(r1, *r2, r3, r4, r5, r6)

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

// FUNCTION TYPE

    // FUNCTION AS A TYPE
    a, b := 4, 4
    var add = func() int {
        return a + b
    }
    fmt.Println(add())

// FUNCTION

    // PASSING ARGUMENTS BY VALUE (COPY) - ARGUMENT NOT CHANGED
        func name(a int) {                          // 1 in - You would never d
        func name(a, b int) int32 {                 // 2 in, 1 return
        func name(name ...int) int {                // Variadic in, 1 return
        // Not a fan of Named returns
        func name(a int, b string) (x int32) {      // 2 in, 1 NAMED return
        func name(a, b int) (x int, y string) {     // 2 in, 2 NAMED return

    // PASSING ARGUMENTS BY "REFERENCE" (POINTER) - ARGUMENT CHANGED
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

// INTERFACE (Reference Type)

    // CREATE INTERFACE TYPE
    type Describer interface {
        describe()

    type Cylinder struct {
        radius float64
        height float64
    }

    // ATTACH TO A METHOD
    func (c Circle) describe() (area float64, circ float64) {
        area = math.Pi * math.Pow(c.radius, 2)
        circ = 2 * math.Pi * c.radius
        return
    }

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

    // LOOPS
        // FOR LOOP
        for inti; condition; post {
        for i:=0; i < 8; i++ {
            do something
        }
        // AS A WHILE LOOP
        i :=0
        for i < 8 {
            do something
            i++
        }

    // RANGE
    myarray := [3]{3,4,5}
    for i := range myarray {
        avg += myarray[i]
    }

    // CONDITIONAL

        // IF, IF / ELSE, NESTED IF
        if a == b {                                 // Relational / Compare operator
            fmt.Println("equal")
        } else if a > b {
            fmt.Println("higher")
        } else {
            fmt.Println("Lower")
        }

        // SWITCH (CASE)
        switch {
        case (a == b):
            fmt.Println("equal")
        case (a > b):
            fmt.Println("higher")
        default:
            fmt.Println("Lower")
        }

        // DEFER
        func main() {
            defer fmt.Println("world")
            fmt.Println("hello")
        }

        // SELECT

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
        // %%                                       // a literal percent sign

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

// PACKAGES

```

## REFERENCES / DOCUMENTATION

### SYNTAX

* [golang.org](http://golang.org)
  _- Home base for everything._
* [golang.org docs](https://golang.org/doc/)
  _- A good collection of docs._
* [golang.org spec](https://golang.org/ref/spec)
  _- I'll be honest, way to much stuff to make your head spin._

### RUNNING CODE

* [golang.org go playground](https://play.golang.org/)
  _- Lets you write, compile and share code.  Just awesome._

### TUTORIALS

* [golang.org tour of go](https://tour.golang.org/welcome/1)
  _- A good place to start._
* [golang.org effective go](https://golang.org/doc/effective_go.html)
  _- A must read to create great things._
* [tutorialspoint.com](https://www.tutorialspoint.com/go/go_data_types.htm)
  _- A great summary of syntax._
* [gobyexample.com](https://gobyexample.com/)
  _- The title says it all._
* [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)
  _- Exactly that._

### PACKAGES

* [godoc.org](https://godoc.org/)
  _- Both standard and user packages._
* [golang.org](https://golang.org/pkg/)
  _- Just official standard packages._

### OTHER STUFF

* [Go Plugin for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
  _- My cheat sheet on setting up
  [setting up visual studio code with go](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
  ._

### HELP

* [Go forum](https://forum.golangbridge.org/)
  _- Community forum._
