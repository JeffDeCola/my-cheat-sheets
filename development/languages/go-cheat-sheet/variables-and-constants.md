# VARIABLES AND CONSTANTS

## DECLARE TYPE AND ASSIGN VALUE

* Declaration is picking a `name for a data type`.
* Assignment is `assigning a value`.

## VARIABLE

A variable always has a single type and may be assigned.

The basic verbose format is,

```
var name type = value
```

Here is the syntax,

```go
// DECLARE TYPE
var a string

// ASSIGN VALUE
a = "happy"

// DECLARE & ASSIGN
var a int32 = 22                                // Verbose
var a = 22                                      // Type Inference
a := 32                                         // Shortcut Assignment
```

## CONSTANT / LITERAL

Constants (or Literal) refer to fixed values that the
program may not alter during its execution.

The basic verbose format is,

```
const name type = value
```

Here is the syntax,

```go
const a float32 = 3.14                          // Must have Assignment
const a = 22                                    // Type Inference
```

Escape Sequences (really used in formatting are considered constants).
See [formating types](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/formating-types.md).

## SCOPE RULES

* Local _- Inside function._
* Global _- Outside function._
* Formal _- In definition of function._

## TYPE INFERENCE

As seen above, the compile will infer a type (if none is giving)
based on your assignment value.

```go
var age = 42                                    // Will infer an int
var pi = 3.14                                   // Will infer a float64
var cNum = 3 + 5i                               // Will infer a complex128
```

## SHORTCUT ASSIGNMENT

Also, as seen above, you may use the shortcut `:=`
the both infers a type and assigns a value,

```go
x := 42                                         // Shortcut Assignment
```

And to beat the dead horse, this is the same as,

```go
var x int = 42;                                 // Verbose
var x = 42;                                     // Type Inference
```

Shortcuts can not be used outside a function.

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

// DECLARE & ASSIGN
var a, b string = "hello a", "hello b"          // Verbose
var a, b = "hello a", "hello b"                 // Type Inference
var (                                           // Parenthesis
    a = "hello a"
    b = "hello b"
)
a, b := "hello a", "hello b"                    // Group Shortcut Assignment
```
