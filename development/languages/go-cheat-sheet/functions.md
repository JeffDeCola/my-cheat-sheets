# FUNCTIONS

Functions stand on their own, a black box.

## BASIC FORMATS

The basic format is,

```txt
func name(parameter list optional) return type optional {
    stuff
}
```

As an example,

```go
func add(a, b int) int {
    return a + b
}
```

Variadic parameter list is when you don't know how many parameters are being passed,

```go
func add(name ...int) int {
    ....
    return sum
}
```

Multiple returns,

```go
func swap(a, b int) (int, int) {
    x, y := b, a
    return x, y
}
```

Named returns,

```go
func swap(a, b int) (x int, y int) {
    x, y = b, a
    return
}
```
## PASSING PARAMETERS

![IMAGE - go function passing by reference and value - IMAGE](../../../docs/pics/go-function-passing-by-reference-and-value.jpg)

### PASSING PARAMETERS TO FUNCTION BY VALUE (COPY) - PARAMETER NOT CHANGED

Passes a copy of the parameter's value and gets something back (return).

Will not change the value of the parameter that was passed.

```go
a := 33
a = negateValue(a)

func negateValue(i int) int {
    return i * -1
}
```

### PASSING PARAMETERS TO FUNCTION BY REFERENCE (POINTER) - PARAMETER CHANGED

Passes the reference (pointer) of the parameter so we can change
the value of the parameter itself (return not necessary),

```go
a := 33
negateReference(&a)

func negateReference(i *int) {
    *i = *i * -1
}
```
