# DECLARATION & ASSIGNMENT

Declaration is picking a name for a data type and optionally assigning
that data type a value.

## VARIABLE

A variable always has a single type and may be assigned.

The basic format is,

```
var name type
var name type = assignment
```

Here is the syntax,

```go
    // DECLARE                               
    var a string

    // ASSIGN
    a = "happy"

    // DECLARE & ASSIGN
    var a int32 = 22                                // Verbose
    var a = 22                                      // Type Inference
    a := 32                                         // Shortcut Assignment
```

## SCOPE RULES

* Local _- Inside function._
* Global _- Outside function._
* Formal _- In definition of function._

## CONSTANT (LITERAL)

Constants refer to fixed values that the program may not alter during its execution.

For example, a string constant could be `"hello"`.

```go
const variable type = value;
```

Escape Sequences (really used in formatting are considered constants).
See [formating-types](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/formating-types.md).

## TYPE INFERENCE

The compile will infer a type if none is giving
based on your assignment.

```go
var age = 42         // will infer an int
var pi = 3.14        // Will infer a float64
var cNum = 3 + 5i    // will infer a complex128
```

## SHORTCUT ASSIGNMENT

You may use the shortcut `:=`,

```go
x := 42
```

This is the same as,

```go
var x int = 42;
var x = 42;
```

Shortcuts can not be used outside a function.

## GROUPING VARIABLES

```go
var a string = "hello a"
var b string = "hello b"
```

Here is the syntax,

```go
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
```
