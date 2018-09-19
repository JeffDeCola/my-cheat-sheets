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

`var` gives the type it's `zero` value.

Here is the various syntax methods,

```go
// DECLARE TYPE
var a string

// ASSIGN VALUE
a = "happy"

// DECLARE & ASSIGN (INITIALIZE)
var a int32 = 22                                // Verbose
var a = 22                                      // Type Inference
a := 32                                         // Shorthand Assignment
```

The preferred method is the shorthand.

## CONSTANT / LITERAL

Constants (or Literal) refer to fixed values that the
program may not alter during its execution.
It will not change.

The basic verbose format is,

```
const name type = value
const name = value
```

Here is the syntax (types and untyped),

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

## SCOPE RULES

* Universe
* Package (access across files)
* File (e.g. import "fmt")
* Block (Inside curly braces)

Order of declaration matters.

Keep a nice tight scope.

## CLOSURE

Helps us limit the scope of variables used by multiple functions.
Without closure, for two or more functions to have access ti same variable,
that variable would need to be at the package scope.

Here is an example,

```go
func main() {
    x := 0
    increment := func() int {
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
}

```

## TYPE INFERENCE

As seen above, the compile will infer a type (if none is giving)
based on your assignment value.

```go
var age = 42                                    // Will infer an int
var pi = 3.14                                   // Will infer a float64
var cNum = 3 + 5i                               // Will infer a complex128
```

## SHORTHAND ASSIGNMENT (PREFERRED METHOD)

Also, as seen above, you may use the shorthand `:=`
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
