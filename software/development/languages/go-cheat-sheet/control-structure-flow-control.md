# CONTROL STRUCTURES / FLOW CONTROL

A control structure analyzes variables and chooses a direction
in which to go based on given parameters.
`for` Loops and `if`, `then`, `else` statements are perfect examples.

Control flow / Flow Control is basically,

* Sequential
* Loops
* Conditionals

Table of Contents,

* [LOOPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#loops)
  * [FOR LOOP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#for-loop)
  * [WHILE LOOP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#while-loop)
  * [INFINITE LOOP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#infinite-loop)
  * [RANGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#range)
  * [BREAK / CONTINUE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#break--continue)
* [CONDITIONAL STATEMENTS /  DECISION MAKING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#conditional-statements---decision-making)
  * [IF, IF ELSE, NESTED IF](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#if-if-else-nested-if)
  * [SWITCH (CASE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#switch-case)
  * [DEFER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#defer)
  * [SELECT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#select)

## LOOPS

Go only has for loops.

For loop basic format,

```go
for init; condition; post {}        // normal
for condition {}                    // while
for {}                              // infinite
```

### FOR LOOP

As an example,

```go
for i:=0; i < 8; i++ {
    do something
}
```

### WHILE LOOP

For loop being a while loop,

```go
i :=0
for i < 8 {
    do something
    i++
}
```

### INFINITE LOOP

```go
for {
    do something
}
```

### RANGE

When we're not really sure how much in our array.

```go
myarray := [3]{3,4,5}

for i := range myarray {
    avg += myarray[i]
}
```

### BREAK / CONTINUE

Breaks, breaks out of the loop, and continue just starts
the loop at the start.

```go
// Print even numbers from 1-100
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

Using fallthrough,

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


