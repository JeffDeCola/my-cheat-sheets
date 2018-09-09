# FUNCTIONS

Functions stand on their own, a black box.

## BASIC FORMATS

The basic format is,

```txt
func name(parameter list - optional)(return type - optional) {
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

Name your returned parameters,

```go
func swap(a, b int) (x int, y int) {
    x, y = b, a
    return
}
```
## PASSING PARAMETERS

![IMAGE - go function passing by reference and value - IMAGE](../../../docs/pics/go-function-passing-by-reference-and-value.jpg)

### PASSING PARAMETERS BY VALUE (COPY)

Passes a copy of that value.

So if I alter the value inside a function
it will not change it outside the function.


```go
a := 33
a = negateValue(a)

func negateValue(i int) int {
    return i * -1
}
```

### PASSING PARAMETERS BY REFERENCE (POINTER)

Passes the reference (pointer) instead
so we can change the type in the function.

```go
a := 33
negateReference(&a)

func negateReference(i *int) {
    *i = *i * -1
}
```