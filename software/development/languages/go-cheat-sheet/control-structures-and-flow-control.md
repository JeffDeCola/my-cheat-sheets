# CONTROL STRUCTURES & FLOW CONTROL

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A control structure analyzes variables and chooses a direction
in which to go based on given parameters.
`for` Loops and `if`, `then`, `else` statements are perfect examples._

Table of Contents

* [LOOPS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#loops)
  * [FOR LOOP](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#for-loop)
  * [WHILE LOOP](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#while-loop)
  * [INFINITE LOOP](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#infinite-loop)
  * [RANGE LOOP](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#range-loop)
  * [BREAK / CONTINUE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#break--continue)
* [CONDITIONAL STATEMENTS /  DECISION MAKING](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#conditional-statements---decision-making)
  * [IF, IF ELSE, NESTED IF](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#if-if-else-nested-if)
  * [SWITCH (CASE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#switch-case)
  * [DEFER](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#defer)
  * [SELECT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-and-flow-control.md#select)

Documentation and Reference

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-cheat-sheet)
  main page

## LOOPS

Go only has for loops with the basic format,

```go
for init; condition; post {stuff}        // For loop
for condition {stuff}                    // While loop
for {stuff}                              // Infinite loop
for value:= range type {stuff}           // Range loop
```

### FOR LOOP

As an example,

```go
for i := 0; i < 8; i++ {
    fmt.Printf("count is %v\n", i)
}
```

### WHILE LOOP

```go
i :=0
for i < 8 {
    fmt.Printf("count is %v\n", i)
    i++
}
```

### INFINITE LOOP

```go
for {
    do something forever
}
```

### RANGE LOOP

When we're not really sure how much in our slice,

```go
myslice := []int{3, 4, 5}
sum := 0

for i := range myslice {
    fmt.Printf("Adding %v to sum %v\n", myslice[i], sum)
    sum += myslice[i]
}
fmt.Printf("Final Sum is sum %v\n", sum)
```

### BREAK / CONTINUE

Breaks out of the loop, and continue just starts
the loop at the start.

```go
// Print even numbers from 1-100
x := 0
for {
    x++
    if x%2 == 1 {
        continue
    }
    if x >= 100 {
        break
    }
    fmt.Println(x)
}
```

## CONDITIONAL STATEMENTS /  DECISION MAKING

Make decisions using boolean relational operators.

* `if`, `if / else`, `nested if`
* `switch (case)`
* `defer`
* `select`

### IF, IF ELSE, NESTED IF

Basic syntax is,

```go
a, b := 2,2

if a == b {
    do something
}
```

if else,

```go
a, b := 4, 3

if a == b {
    fmt.Println("equal")
} else if a > b {
    fmt.Println("higher")
} else {
    fmt.Println("Lower")
}
```

Nested if is just putting an if in an if.

### SWITCH (CASE)

Makes `else if` easier to read,

```go
switch {
case (a == b):
    fmt.Println("equal")
case (a > b):
    fmt.Println("higher")
default:
    fmt.Println("Lower")
}
```

Using `switch name`,

```go
switch a {
case == b:
    fmt.Println("equal")
default:
    fmt.Println("not equal")
}
```

Using `fallthrough`,

```go
switch {
case (0 == a%2):
    fmt.Println("a is an even number")
    fallthrough
case (a == b):
    fmt.Println("equal")
case (a > b):
    fmt.Println("higher")
default:
    fmt.Println("lower")
}
```

### DEFER

A defer statement defers the execution of a function until the surrounding
function returns.

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

Useful if opening a file, you can put the cleanup first, before you do stuff.

### SELECT

Go’s select lets you wait on multiple channel operations.

Refer to
[go routines and channels](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/goroutines-and-channels.md).
