# DECLARATION / ASSIGNMENT

Declaration is picking a name for a data type and optionally assigning
that data type a value.

## VARIABLES

Variables always has a single type and may be assigned.

The basic format is,

```
var name type
var name type = assignment
```

For example,

```go
var greeting string = "hello, world"
var pi float32 = 3.14
```

## SCOPE RULES

* Local - Inside Function
* Global - Outside Function
* Formal - In definition of function

## CONSTANTS (LITERALS)

Constants refer to fixed values that the program may not alter during its execution.

For example, a string constant could be `"hello"`.

The const keyword, 

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

Shortcuts can not be used outside function.

## GROUPING VARIABLES

```go
var x int = 5
var y int = 8
```

Group can also infer a type,

```go
var (
    x=5
    y=8
)
```

But even better, put it on one line,

```go
var x, y = 5, 8
```

But even, even better, lets use shorthand
(again, only inside a function),

```go
x, y := 5, 8
```
