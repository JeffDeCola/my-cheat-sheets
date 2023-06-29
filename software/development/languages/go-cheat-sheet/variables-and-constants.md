# VARIABLES AND CONSTANTS

_Variables and constants._

tl;dr,

```go
// VARIABLE

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

// CONSTANT / LITERAL

    // DECLARE & ASSIGN (INITIALIZE)
    const a float32 = 3.14                          // Must have Assignment
    const a = 22                                    // Type Inference

// GROUPING VARIABLES

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

Table of Contents

* [DECLARE TYPE AND ASSIGN VALUE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#declare-type-and-assign-value)
* [VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#variable)
* [CONSTANT / LITERAL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#constant--literal)
* [IOTA](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#iota)
* [SCOPE RULES (Universe, Package, File, Block)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#scope-rules-universe-package-file-block)
* [TYPE INFERENCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#type-inference)
* [SHORTHAND ASSIGNMENT (PREFERRED METHOD)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#shorthand-assignment-preferred-method)
* [GROUPING VARIABLES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md#grouping-variables)

## DECLARE TYPE AND ASSIGN VALUE

* Declaration is `picking a name for a data type`.
* Assignment is `assigning a value`.

Initializing is both declaring and assigning.

## VARIABLE

A variable always has a single type and may be assigned.

The basic verbose format is,

```go
var name type = value
```

`var` gives the type it's "zero" value.

Here are the various syntax methods,

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

## CONSTANT / LITERAL

Constants (or Literal) refer to fixed values that the
program may not alter during its execution.
It will not change.

The basic verbose formats are,

```go
const name type = value
const name = value
```

Here is the syntax (type and untyped),

```go
const a float32 = 3.14                          // Must have Assignment
const a = 22                                    // Type Inference
```

Escape Sequences (really used in formatting are considered constants).
See [formating types](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/formating-types.md).

## IOTA

Really means just a little bit.

```go
const (
    c = iota
    d = iota
    e = iota
)
```

## SCOPE RULES (Universe, Package, File, Block)

* Universe
* Package - Access across files.
* File - File level (e.g. import "fmt").
* Block - Inside curly braces.

Order of declaration matters.

Always keep a nice tight scope.

## TYPE INFERENCE

The compile will infer a type (if none is giving)
based on your assignment value.

```go
var age = 42                                    // Will infer an int
var pi = 3.14                                   // Will infer a float64
var cNum = 3 + 5i                               // Will infer a complex128
```

## SHORTHAND ASSIGNMENT (PREFERRED METHOD)

Also, you may use the shorthand `:=`
the both infers a type and assigns a value,

```go
x := 42                                         // Shorthand Assignment
```

And to beat the dead horse, this is the same as,

```go
var x int = 42;                                 // Verbose
var x = 42;                                     // Type Inference
```

Shorthands can not be used outside a function.

## GROUPING VARIABLES

The basic verbose format is,

```go
var name1, name 2 ... type = value1, value2 ...
```

Here is the syntax,

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
