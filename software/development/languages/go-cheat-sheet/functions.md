# FUNCTIONS

_Functions stand on their own, a black box._

Just a note,

* `Parameters` are what you declare in a function
* `Arguments` are passed to functions

```go
// BASIC FORMAT
    func receiver name(parameter list) (return type) {
        stuff
    }
// VARIADIC FUNCTIONS
    func name(name ...int) int {                    // Variadic in, 1 return
// PASSING ARGUMENTS BY VALUE (COPY) - ARGUMENT NOT CHANGED
    func negateValue(i int) {
        i = -i
    }
    func main() {
        a := 33
        fmt.Println(a)                              // 33
        negateValue(a)
        fmt.Println(a)                              // 33
    }
// PASSING ARGUMENTS BY "REFERENCE" (POINTER) - ARGUMENT CHANGED
    func negateValue(i *int) {
        *i = -*i
    }
    func main() {
        a := 33
        fmt.Println(a)                              // 33
        negateValue(&a)
        fmt.Println(a)                              // -33
    }
// FUNCTION AS A TYPE
    // ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE
    a, b := 3, 9
    add := func() int {                             // Anonymous func as a type (no name)
        return a + b                                // returns int
    }
    fmt.Println(add())                              // 12
    a = 9
    fmt.Println(add())                              // 18
    // CLOSURE - RETURN A FUNCTION TO A FUNCTION
    func addThis(a, b int) func() int {
        return func() int {
            return a + b
        }
    }
    func main() {
        a, b := 3, 9
        add := addThis(a, b)                        // Think of the func add like a variable
        fmt.Println(add())                          // 12
        a = 9
        fmt.Println(add())                          // 12 <- NOTE THIS
    }
// CALLBACK - PASSING A FUNCTION (AS AN ARGUMENT) TO A FUNCTION
    // Receiving a function
    func math(numbers []int, callback func(int)) {
        for _, n := range numbers {
            callback(n)
        }
    }
    func main() {
        numbers := []int{1, 2, 3, 4}
        // Passing a function (as an argument) to a function - callback
        math(numbers, func(x int) {
            fmt.Printf("called %v ", x)             // called 1 called 2 called 3 called 4
        })
    }
// RECURSION - FUNCTION CALLING ITSELF
    func factorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }
    func main() {
        var n int
        fmt.Scanf("%d\n", &n)                       // if input 4
        fmt.Println(factorial(n))                   // 24 (4*3*2*1)
    }
// ANONYMOUS SELF EXECUTING FUNCTION
    func main() {
        func() {
            fmt.Println("hi")                       // hi
        }()
    }
```

Table of Contents

* [BASIC FORMATS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#basic-formats)
* [VARIADIC FUNCTIONS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#variadic-functions)
* [PASSING ARGUMENTS - GO PASSES BY VALUE ONLY](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#passing-arguments---go-passes-by-value-only)
  * [PASSING ARGUMENTS TO FUNCTION BY VALUE (COPY) - ARGUMENT NOT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#passing-arguments-to-function-by-value-copy---argument-not-changed)
  * [PASSING ARGUMENTS TO FUNCTION BY "REFERENCE" (POINTER) - ARGUMENT CHANGED](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#passing-arguments-to-function-by-reference-pointer---argument-changed)
* [FUNCTION AS A TYPE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#function-as-a-type)
  * [ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#assign-anonymous-function-func-literal-to-a-variable)
  * [CLOSURE - RETURN A FUNCTION TO A FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#closure---return-a-function-to-a-function)
* [CALLBACK - PASSING A FUNCTION (AS AN ARGUMENT) TO A FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#callback---passing-a-function-as-an-argument-to-a-function)
* [RECURSION - FUNCTION CALLING ITSELF](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#recursion---function-calling-itself)
* [ANONYMOUS SELF EXECUTING FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#anonymous-self-executing-function)
* [EXAMPLE - SHAPES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#example---shapes)

## BASIC FORMATS

The basic format is,

```go
func receiver identifier parameters returns
```

or,

```go
func receiver name(parameter list) (return type) {
    stuff
}
```

As an example,

```go
func add(a, b int) int {
    return a + b
}
```

Multiple returns,

```go
func swap(a, b int) (int, int) {
    x, y := b, a
    return x, y
}
```

Named returns (not a good coding style),

```go
func swap(a, b int) (x int, y int) {
    x, y = b, a
    return
}
```

## VARIADIC FUNCTIONS

Variadic arguments lists when you don't know how many arguments to pass.
Variadic parameter list is when you don't know how many parameters are
being passed.

```go
func average(n ...float64) float64 {
    total := 0.0
    for _, f := range n {
        total += f
    }
    return total / float64(len(n))
}

func main() {
    data := []float64{43, 44, 55, 66, 77, 88}
    fmt.Println("Average of", data, "is", average(data...))
}
```

## PASSING ARGUMENTS - GO PASSES BY VALUE ONLY

![IMAGE - go function passing by reference and value - IMAGE](../../../../docs/pics/go-function-passing-by-reference-and-value.jpg)

### PASSING ARGUMENTS TO FUNCTION BY VALUE (COPY) - ARGUMENT NOT CHANGED

Passes a "copy" of the parameter's value and gets something back (return).
Take the word copy with a grain of salt.

Really passing a value (the argument) and assigning it to the function's parameter.

Will not change the value of the argument that was passed.  Because of scope.

```go
func negateValue(i int) {
    i = -i
}

func main() {
    a := 33
    fmt.Println(a) // 33
    negateValue(a)
    fmt.Println(a) // 33
}
```

### PASSING ARGUMENTS TO FUNCTION BY "REFERENCE" (POINTER) - ARGUMENT CHANGED

Go only passes by value, hence the quotes on "reference".

Passes the reference (pointer) (the argument) to the function parameter
so we can change the value of the argument itself (return not necessary),

```go
func negateValue(i *int) {
    *i = -*i
}

func main() {
    a := 33
    fmt.Println(a) // 33
    negateValue(&a)
    fmt.Println(a) // -33
}
```

Again you do this because you want to actually modify whatever you're passing
(“read/write” as opposed to just “read”)

## FUNCTION AS A TYPE

A closure is a function value that references variables from outside its body.
**We already did this in derived types**, but will repeat here for fun with
different examples.

Helps us limit the scope of variables used by multiple functions.
Without closure, for two or more functions to have access to the same variable,
that variable would need to be at the package scope.

There are two methods as follows.

### ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE

Here is an example of closure, assigning an anonymous function
(also called func literal) to a variable.

This is called closure because the function captures (enclose)
the surrounding environment and can use it.

```go
x:= 5

// This is also called a func literal.
increment1 := func() int {                      // Got x = 5
    x++
    return x
}

// When you call function increment1, it uses x, which is outside function.
fmt.Println(increment1())                       // 6
fmt.Println(increment1())                       // 7
fmt.Println(x)                                  // 7
fmt.Println(increment1())                       // 8
x = 30
fmt.Println(x)                                  // 30
fmt.Println(increment1())                       // 31 <- NOTE THIS
fmt.Println(x)                                  // 31
```

Just think/treat the function as a variable and closure makes sense.
People try to make this too complicated, like I did above but its just
treating the unnamed function as a variable `increment1`.  Simple.

### CLOSURE - RETURN A FUNCTION TO A FUNCTION

Same program as above, but with function outside main.

But this acts like a variable where once increment2 or increment3
is declared and assigned, x is set.

```go
func inc(i int) func() int {
    return func() int {
        i++
        return i
    }
}

// CLOSURE METHOD TWO - Return a function to a function
// This is different because now the value is bound to the variable.
// The function captures the scope it is in
// Think of it as an assigned variable, BECAUSE it is.
increment2 := inc(x)                            // Got x = 31
increment3 := inc(x)                            // Got x = 31

// When you call function increment2, it uses x, which is outside function.
fmt.Println(increment2())                       // 32
fmt.Println(increment2())                       // 33
fmt.Println(increment2())                       // 34
fmt.Println(x)                                  // 31

// increment2 has its own value of x
// Treat the function like a variable, because it is.
fmt.Println(increment3())                       // 32
fmt.Println(increment3())                       // 33
fmt.Println(increment3())                       // 34
x = 20
fmt.Println(x)                                  // 20
// Nothing to do with x because we assigned inc() to to increment3
fmt.Println(increment3())                       // 35
fmt.Println(increment3())                       // 36
```

Another example,

```go
// Is the integer even or not
func even() func(int) bool {
    return func(n int) bool {
        return n%2 == 0
    }
}

func main() {

    n := []int{1, 2, 3, 4, 5}

    // r is the returned function
    r := even()

    for _, s := range n {
        fmt.Printf("%v %v,", s, r(s))           // 1 false,2 true,3 false,4 true,5 false,
    }
}
```

That really looks like this (a func expression),

```go
func main() {

    n := []int{1, 2, 3, 4, 5}

    // r is the returned function
    // Is the integer even or not
    // This is a func expression
    r := func(n int) bool {
        return n%2 == 0
    }

    for _, s := range n {
        fmt.Printf("%v %v,", s, r(s))           // 1 false,2 true,3 false,4 true,5 false,
    }
}
```

## CALLBACK - PASSING A FUNCTION (AS AN ARGUMENT) TO A FUNCTION

Passing a function (as an argument) to a function.

Callback means call back the function you passed in.
Calling back home.

```go
// Receiving a function
func math(numbers []int, callback func(int)) {
    for _, n := range numbers {
        callback(n)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7}

    // Passing a function (as an argument) to a function - callback
    math(numbers, func(x int) {
        fmt.Println("I've been called", x)
    })
}
```

Another example is in
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/blob/master/basic-programming/callback/callback.go).

Why would you do this? Not very popular in go. You
don't want complexity in go or being too clever.

## RECURSION - FUNCTION CALLING ITSELF

This is very straightforward.

Using recursion to get a factorial,

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    fmt.Println(factorial(n))
}
```

Another example is in
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/blob/master/basic-programming/recursion/recursion.go).

## ANONYMOUS SELF EXECUTING FUNCTION

Simply a function that executes in your code.

```go
func main() {
    func() {
        fmt.Println("hi")
    }()
}
```

## EXAMPLE - SHAPES

This example will calculate the **area** and **perimeter**
of three shapes: **circles**, **rectangles** and **triangles**.

The example is located in `my-go-examples` repo done three different ways,

* [shapes-using-functions](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/functions/shapes-using-functions)
* [shapes-using-methods](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/methods/shapes-using-methods)
* [shapes-using-interfaces](https://github.com/JeffDeCola/my-go-examples/tree/master/basic-syntax/interfaces/shapes-using-interfaces)

You will see that using an interface you can go from this,

```go
c1Area := circleArea(c1)
c1Perimeter := circlePerimeter(c1)
r1Area := rectangleArea(r1)
r1Perimeter := rectanglePerimeter(r1)
t1Area := triangleArea(t1)
t1Perimeter := trianglePerimeter(t1)
```

To this,

```go
// This is the power of interfaces.
c1Area = area(c1)
c1Perimeter = perimeter(c1)
r1Area = area(r1)
r1Perimeter = perimeter(r1)
t1Area = area(t1)
t1Perimeter = perimeter(t1)
```
