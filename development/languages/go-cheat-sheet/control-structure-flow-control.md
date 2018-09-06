# CONTROL STRUCTURES / FLOW CONTROL

A control structure analyzes variables and chooses a direction
in which to go based on given parameters.
`for` Loops and `if`, `then`, `else` statements are perfect examples.

## LOOPS

Go only has for loops.

For loop basic format,

```go
for i:=0; i < 8; i++ {
    do something
}
```

For loop being a while loop,

```go
i :=0
for i < 8 {
    do something
    i++
}
```

## RANGE

When we're not really sure how much in our array.

myarray := [3]{3,4,5}

```
for i := range myarray {
    avg += myarray[i]
}
```

## CONDITIONAL STATEMENTS /  DECISION MAKING

Make decisions using boolean relational operators.

* `if`
* `if...else`
* `nested if`
* `switch`
* `defer`
* `select`

### IF

```go
a, b := 2,2

if a == b {
    do something
}
```

### IF ELSE

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

### NESTED IF

Put if in an if.

### SWTICH

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

### DEFER

### SELECT

