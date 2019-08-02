# GO CHEAT SHEET

`go` _is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems._

This is a very abbreviated cheat-sheet highlighting the main
syntax of go. I have lots go examples in my appropriately named repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples).

* [LET'S GO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#lets-go)
  * [INSTALL & CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#install--configure)
  * [BASIC CONCEPTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#basic-concepts)
  * [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#data-types)
    * BOOLEAN
    * NUMERIC
    * STRING
  * [TYPE CONVERSION & TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#type-conversion--type-assertion)
  * [VARIABLES & CONSTANTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#variables--constants)
  * [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#derived-data-types)
    * ARRAY (Data Structure) (_new_)
    * SLICE (Data Structure, Reference Type) (_make_)
    * MAP (Data Structure, Reference Type) (_make_)
    * STRUCT (Data Structure)
    * POINTER
    * FUNCTION AS A TYPE
    * INTERFACE (see below)
    * CHANNEL (Reference Type) (_make_) (see below)
  * [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#functions-black-box)
  * [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#methods-attached-to-data)
  * [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#interfaces-set-of-method-signatures)
  * [CONCURRENCY / CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#concurrency--channels)
  * [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#operators)
  * [CONTROL STRUCTURES / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#control-structures--flow-control)
  * [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#error-handling)
  * [FORMATING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#formating-types)
  * [PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#packages)
* [GO SYNTAX OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-syntax-overview)
  * [GO DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-data-types)
  * [GO TYPE CONVERSION & TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-type-conversion--type-assertion)
  * [VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#variable)
  * [CONSTANT / LITERAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#constant--literal)
  * [GROUPING VARIABLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#grouping-variables)
  * [ARRAY (Data Structure) (_new_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#array-data-structure-new)
  * [SLICE (Data Structure, Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#slice-data-structure-reference-type-make)
  * [MAP (Data Structure, Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#map-data-structure-reference-type-make)
  * [STRUCT (Data Structure)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#struct-data-structure)
  * [POINTER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#pointer)
  * [FUNCTION AS A TYPE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#function-as-a-type)
  * [FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#function)
  * [METHOD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#method)
  * [INTERFACE (Reference Type)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#interface-reference-type)
  * [CHANNEL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#channel)
  * [GO OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-operators)
  * [CONTROL STRUCTURE / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#control-structure--flow-control)
  * [GO ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-error-handling)
  * [FORMAT SPECIFIERS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#format-specifiers)
  * [ESCAPE SEQUENCES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#escape-sequences)
* [REFERENCES / DOCUMENTATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#references--documentation)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## LET'S GO

This cheat sheet is broken up into the following sections,

### INSTALL & CONFIGURE

* [INSTALL & CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/install-and-configure.md)
  * Install
  * Configure
  * Check
  * Install Go Tools

### BASIC CONCEPTS

* [BASIC CONCEPTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/basic-concepts.md)
  * Basic Structure of go
  * Basic Syntax
  * Packages (See own [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md))
  * _go run_
  * _go build_
  * _go install_
  * The Object Side of go

### DATA TYPES

* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/data-types.md)
  * All Types
  * Boolean
  * Numeric
  * String

### TYPE CONVERSION & TYPE ASSERTION

* [TYPE CONVERSION & TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/type-conversion-and-type-assertion.md)
  * Type Conversion (Type casting)
  * Type Assertion

### VARIABLES & CONSTANTS

* [VARIABLES & CONSTANTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/variables-and-constants.md)
  * Declare Type and Assign Value
  * Variable
  * Constant / Literal
  * IOTA
  * Scope Rules (Universe, Package, File, Block)
  * Type Inference
  * Shorthand Assignment (Preferred Method)
  * Grouping Variables

### DERIVED DATA TYPES

* [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/derived-data-types.md)
  * All Types
  * Array (Data Structure) (_new_)
  * Slice (Data Structure, Reference Type) (_make_)
  * Map (Data Structure, Reference Type) (_make_)
  * Struct (Data Structure)
  * Pointer
  * Function as a Type
  * Interface (See own
    [Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (Reference Type) (_make_) (See own
    [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md))

### FUNCTIONS (BLACK BOX)

* [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/functions.md)
  * Basic Formats
  * Variadic Functions
  * Closure (func Expression & Anonymous Function)
    * Assign anonymous Function (func Literal) to a Variable
    * Return a Function to a Function - Closure
  * Passing a Function (As an Argument) to a Function - Callback
  * Passing Arguments - Go Passes by Value Only
    * Passing Arguments to Function by Value (_Copy_) - Parameter not Changed
    * Passing Arguments to Function by "Reference" (_Pointer_) - Parameter Changed
  * Recursion (function calling itself)
  * Anonymous Self Executing Function
  * Example - Shapes

### METHODS (ATTACHED TO DATA)

* [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/methods.md)
  * Overview
  * Basic Format
  * Passing Parameters
    * Passing Struct to Method by Value (_Copy_) - Struct not Changed
    * Passing Struct to Method by Reference (_Pointer_) - Struct Changed
  * Example - Shapes

### INTERFACES (SET OF METHOD SIGNATURES)

* [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md)
  * Overview
  * Basic Format
  * Making your Code Cleaner
    * Without Interface
    * With Interface
  * Example - Shapes

### CONCURRENCY / CHANNELS

* [CONCURRENCY / CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md)
  * tbd

### OPERATORS

* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md)
  * Arithmetic (_Math_)
  * Relational (_Compare_)
  * Logical (_Boolean_)
  * Bitwise (_Bits_)
  * Assignment
  * Miscellaneous

### CONTROL STRUCTURES / FLOW CONTROL

* [CONTROL STRUCTURES / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/control-structure-flow-control.md)
  * Loops
    * _for, _while_, _infinite loops_, _range_, _break/continue_
  * Conditional Statements / Decision Making
    * _if_, _if else_, _nested if_, _switch (case)_, _defer_, _select_

### ERROR HANDLING

* [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/error-handling.md)
  * tbd

### FORMATING TYPES

* [FORMATING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/formating-types.md)
  * Read Input (bufio.NewReader Package)
  * Scan Input (fmt.Scan Package)
  * Format Specifiers
  * Escape Sequences

### PACKAGES

* [PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md)
  * Go Get a Package and use it
  * Lets create a package

## GO SYNTAX OVERVIEW

### GO DATA TYPES

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

### GO TYPE CONVERSION & TYPE ASSERTION

```go
    // TYPE CONVERSION
    i := 33
    nowAFloat := float32(i) / 2.5                   // Result 13.2 (int to float)

    // TYPE ASSERTION
    var j interface{} = "jeff"
    s, ok := j.(string)
    fmt.Println(s, ok)                              // Prints jeff true
    f, ok := j.(float32)
    fmt.Println(f, ok)                              // Prints 0, false
```

### VARIABLE

```go
    // DECLARE TYPE
    var a string                                    // var name type

    // ASSIGN VALUE
    a = "happy"                                     // name = value

    // DECLARE & ASSIGN (INITIALIZE)
    var b int32 = 22                                // Verbose - var name type = value
    var c = 22                                      // Type Inference
    d := 32                                         // Shorthand Assignment (Preferred)

    // PRINT
    fmt.Println(a, b ,c ,d)                         // happy 22 22 32
```

### CONSTANT / LITERAL

```go
    const a float32 = 3.14                          // Must have Assignment
    const a = 22                                    // Type Inference
```

### GROUPING VARIABLES

```go
    // GROUP DECLARE TYPE
    var a, b string                                 // var name1, name 2 ... type

    // GROUP ASSIGN VALUE
    a = "hi a"                                      // name1 = value
    b = "hi b"                                      // name2 = value

    // GROUP DECLARE & ASSIGN (INITIALIZE)
    var c, d string = "hi c", "hi d"                // Verbose - var name1, name 2 ... type = value1, value2, ...
    var e, f = "hi e", "hi f"                       // Type Inference
    var (                                           // Parenthesis
        g = "hi g"
        h = "hi h"
    )
    i, j := "hi i", "hi j"                          // Group Shorthand Assignment

    // PRINT
    fmt.Println(a, b ,c ,d, e, f ,g, h, i, j)       // hi a hi b hi c hi d hi e hi f hi g hi h hi i hi j
```

### ARRAY (Data Structure) (_new_)

```go
    // DECLARE TYPE
    var a [2]float32                                // var name [number]type

    // ASSIGN VALUE
    a[0] = 1.1                                      // name[number] = value
    a[1] = 2.0

    // DECLARE & ASSIGN (INITIALIZE)
    var b = [2]float32{1.1, 2.0}                    // Verbose - var name = [number]type{value, value, ...}
    c := [2]float32{1.1, 2.0}                       // Array Shorthand Assignment

    // PRINT
    fmt.Println(a, b, c)                            // [1.1 2] [1.1 2] [1.1 2]
```

### SLICE (Data Structure, Reference Type) (_make_)

```go
    // DECLARE TYPE - NO SIZE
    var a []float64                                 // var name []type

    // ASSIGN VALUE - ADD LENGTH TO SLICE
    a = append(a, 5.7)                              // name = append(name, value, value, ...)

    // DECLARE TYPE - WITH SIZE (make)
    b := make([]string, 1, 25)                      // name := make([]type, length, capacity)

    // ASSIGN VALUE
    b[0] = "hello"                                  // name[index] = value

    // DECLARE & ASSIGN (INITIALIZE)
    var c = []float32{1.1, 2.0}                     // Verbose - var name = []type{value, value, ...}
    d := []float32{3.4, 4.5}                        // Array Shortcut Assignment

    // PRINT
    fmt.Println(a, b, c, d)                         // [5.7] [hello] [1.1 2] [3.4 4.5]
```

### MAP (Data Structure, Reference Type) (_make_)

```go
    // DECLARE TYPES - THIS IS A NIL MAP - DON'T DO THIS
    var a map[string]int                            // var name map[keytype]valuetype

    // DECLARE TYPES (make)
    var b = make(map[string]int)                    // var name make(map[keytype]valuetype)
    c := make(map[string]int)                       // name := make(map[keytype]valuetype)

    // ASSIGN KEY:VALUE
    b["Jill"] = 23                                  // name[key] = value
    b["Bob"] = 34
    b["Mark"] = 28
    c["Jill"], c["Bob"], c["Mark"] = 23, 34, 28

    // DECLARE & ASSIGN KEY:VALUE (INITIALIZE)
    var d = map[string]int{                         // Verbose - var name = map[keytype]valuetype {key:value, key:value, ...}
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }
    e := map[string]int{                            // Array Shortcut Assignment
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }

    // PRINT
    fmt.Println(a, b, c, d, e)                      // map[Jill:23 Bob:34 Mark:28] (For all of the maps)

    // DELETE A KEY
    delete(e,"Jill")                                // Delete key "Jill"
```

### STRUCT (Data Structure)

 ```go
    // CREATE STRUCT TYPE
    type Rect struct {                              // Define a struct
        w, h int
        }

    // DECLARE TYPE
    var r1 Rect                                     // Create a struct
    r2 := new(Rect)                                 // Returns Pointer

    // ASSIGN VALUE TO FIELDS
    r1.w = 2                                        // name.field = value
    r1.h = 4
    r2.w, r2.h = 3, 5

    // DECLARE & ASSIGN (INITIALIZE)
    var r3 Rect = Rect{2, 4}                        // Verbose (Don't use)
    var r4 = Rect{2, 4}                             // Type Inference - var name = structName{value, value, ....}
    r5 := Rect{w: 2, h: 4}                          // Shorthand Assignment
    r6 := Rect{2, 4}                                // Shorthand Assignment

    // PRINT
    fmt.Println(r1, *r2, r3, r4, r5, r6)            // {2 4} {3 5} {2 4} {2 4} {2 4} {2 4}
```

### POINTER

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
    r1.w = 6.1                                      // I wish it was *r1.w
    r1.h = 5.0                                      // I wish it was *r1.h
```

### FUNCTION AS A TYPE

```go
    // FUNCTION AS A TYPE
    a, b := 4, 4
    var add = func() int {
        return a + b
    }
    fmt.Println(add())
```

### FUNCTION

```go
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
```

### METHOD

```go
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
```

### INTERFACE (Reference Type)

```go
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
```

### CHANNEL

```go
    tbd
```

### GO OPERATORS

```go
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
```

### CONTROL STRUCTURE / FLOW CONTROL

```go
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
```

### GO ERROR HANDLING

```go
    tbd
```

### FORMAT SPECIFIERS

```go
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
```

### ESCAPE SEQUENCES

```go
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

### GO PACKAGES

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
