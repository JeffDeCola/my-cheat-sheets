# VARIABLES AND CONSTANTS

* [DECLARE TYPE AND ASSIGN VALUE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#declare-type-and-assign-value)
* [VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#variable)
* [CONSTANT / LITERAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#constant--literal)
* [IOTA](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#iota)
* [SCOPE RULES (Universe, Package, File, Block)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#scope-rules-universe-package-file-block)
* [TYPE INFERENCE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#type-inference)
* [SHORTHAND ASSIGNMENT (PREFERRED METHOD)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#shorthand-assignment-preferred-method)
* [GROUPING VARIABLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#grouping-variables)

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
var a string

// ASSIGN VALUE
a = "happy"

// DECLARE & ASSIGN (INITIALIZE)
var a int32 = 22                                // Verbose
var a = 22                                      // Type Inference
a := 32                                         // Shorthand Assignment (Preferred)
```

The preferred method is the shorthand.

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
```
