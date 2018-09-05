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

Variadic Parameter List - Don't know how many parameters are being passed,

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

## PASSING PARAMETERS BY REFERENCE

## PASSING PARAMETERS BY VALUE
