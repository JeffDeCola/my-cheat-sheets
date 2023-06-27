# GO CHEAT SHEET

_Go is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems._

**LET'S GO (LINKS)**

* [INSTALL & CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md)
* [BASIC CONCEPTS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md)
  * SYNTAX, RUN, BUILD, INSTALL
* [EXECUTABLE (YOUR CODE & GO RUNTIME)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/executable-your-code-and-go-runtime.md)
* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md)
  * BOOLEAN
  * NUMERIC
  * STRING
* [TYPE CONVERSION & TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/type-conversion-and-type-assertion.md)
* [VARIABLES & CONSTANTS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/variables-and-constants.md)
* [DERIVED DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/derived-data-types.md)
  * ARRAY (Data Structure)
  * SLICE (Data Structure, Reference Type) (_make_)
  * MAP (Data Structure, Reference Type) (_make_)
  * STRUCT (Data Structure)
  * POINTER (_new_)
  * FUNCTION AS A TYPE
  * INTERFACE (see below)
  * CHANNEL (Reference Type) (_make_) (see below)
* [FUNCTIONS (BLACK BOX)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md)
  * PASSING ARGUMENTS (VALUE (COPY) AND REFERENCE (POINTER))
  * CALL BACK
  * RECURSION
* [METHODS (ATTACHED TO DATA)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/methods.md)
  * PASSING PARAMETERS (VALUE (COPY) AND REFERENCE (POINTER))
* [INTERFACES (SET OF METHOD SIGNATURES)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/interfaces.md)
* [GOROUTINES & CHANNELS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/goroutines-and-channels.md)
* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/operators.md)
* [CONTROL STRUCTURES / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/control-structures-flow-control.md)
  * LOOPS
  * CONDITIONAL STATEMENTS
* [ERROR HANDLING & LOGGING](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/error-handling-and-logging.md)
* [FORMATTING TYPES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/formatting-types.md)
* [PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/packages.md)
* [CGO (CALLING C WITH GO)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/cgo.md)

**REFERENCES/DOCUMENTATION (LINKS)**

* MY REPOS
  * [my-go-examples](https://github.com/JeffDeCola/my-go-examples)
  * [my-go-packages](https://github.com/JeffDeCola/my-go-packages)
  * [my-go-tools](https://github.com/JeffDeCola/my-go-tools)
* SYNTAX
  * [golang.org](http://golang.org)
    _- Home base for everything_
  * [golang.org docs](https://golang.org/doc/)
    _- A good collection of docs_
  * [golang.org spec](https://golang.org/ref/spec)
    _- I'll be honest, way to much stuff to make your head spin_
* RUNNING CODE
  * [golang.org go playground](https://play.golang.org/)
    _- Lets you write, compile and share code_
* TUTORIALS
  * [golang.org tour of go](https://tour.golang.org/welcome/1)
    _- A good place to start_
  * [golang.org effective go](https://golang.org/doc/effective_go.html)
    _- A must read to create great things_
  * [tutorialspoint.com](https://www.tutorialspoint.com/go/go_data_types.htm)
    _- A great summary of syntax_
  * [gobyexample.com](https://gobyexample.com/)
    _- The title says it all_
  * [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)
    _- Exactly that_
  * [medium.com golangspec](https://medium.com/golangspec)
    _- A bunch of cool examples_
* GO PACKAGE LISTS
  * [godoc.org](https://godoc.org/)
    _- Both standard and user packages. Also shows popular packages_
  * [golang.org](https://golang.org/pkg/)
    _- Just official standard packages_
* HELP
  * [Go forum](https://forum.golangbridge.org/)
    _- Community forum_

**GO SYNTAX OVERVIEW (BELOW)**
  
* [GO DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-data-types)
* [GO TYPE CONVERSION & TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-type-conversion--type-assertion)
* [VARIABLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#variable)
* [CONSTANT / LITERAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#constant--literal)
* [GROUPING VARIABLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#grouping-variables)
* [ARRAY (Data Structure)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#array-data-structure)
* [SLICE (Data Structure, Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#slice-data-structure-reference-type-make)
* [MAP (Data Structure, Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#map-data-structure-reference-type-make)
* [STRUCT (Data Structure)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#struct-data-structure)
* [POINTER (_new_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#pointer-new)
* [FUNCTION AS A TYPE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#function-as-a-type)
* [FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#function)
* [METHOD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#method)
* [INTERFACE (Reference Type)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#interface-reference-type)
* [GOROUTINES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#goroutines)
* [CHANNELS (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#channels-make)
* [GO OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-operators)
* [CONTROL STRUCTURE / FLOW CONTROL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#control-structure--flow-control)
* [ERROR HANDLING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#error-handling)
* [LOGGING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#logging)
* [FORMAT SPECIFIERS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#format-specifiers)
* [ESCAPE SEQUENCES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#escape-sequences)
* [GO PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-packages)
* [GCO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#gco)

---

## GO DATA TYPES

```go
    // BOOLEAN
        // true, false

    // NUMERIC
        // singed - int8, int16, int32 (rune), int64
        // unsigned - uint8 (byte), uint16, uint32, uint64
        // machine - int, uint, uintptr
        // float - float32, float64
        // complex - complex64, complex128

    // STRING
        // string (immutable array of bytes (or runes))
```

## GO TYPE CONVERSION & TYPE ASSERTION

```go
    // TYPE CONVERSION
    i := 33
    nowAFloat := float32(i) / 2.5                   // Result 13.2 (int to float)

    // TYPE ASSERTION
    var j interface{} = "jeff"
    s, ok := j.(string)
    fmt.Println(s, ok)                              // Prints jeff true
    f, ok := j.(float32)
    fmt.Println(f, ok)                              // Prints 0, false
```

## VARIABLE

```go
    // DECLARE TYPE
    var a string                                    // var name type

    // ASSIGN VALUE
    a = "happy"                                     // name = value

    // DECLARE & ASSIGN (INITIALIZE)
    var b int32 = 22                                // Verbose - var name type = value
    var c = 22                                      // Type Inference
    d := 32                                         // Shorthand Assignment (Preferred)

    // PRINT
    fmt.Println(a, b ,c ,d)                         // happy 22 22 32
```

## CONSTANT / LITERAL

```go
    const a float32 = 3.14                          // Must have Assignment
    const a = 22                                    // Type Inference
```

## GROUPING VARIABLES

```go
    // GROUP DECLARE TYPE
    var a, b string                                 // var name1, name 2 ... type

    // GROUP ASSIGN VALUE
    a = "hi a"                                      // name1 = value
    b = "hi b"                                      // name2 = value

    // GROUP DECLARE & ASSIGN (INITIALIZE)
    var c, d string = "hi c", "hi d"                // Verbose - var name1, name 2 ... type = value1, value2, ...
    var e, f = "hi e", "hi f"                       // Type Inference
    var (                                           // Parenthesis
        g = "hi g"
        h = "hi h"
    )
    i, j := "hi i", "hi j"                          // Group Shorthand Assignment

    // PRINT
    fmt.Println(a, b ,c ,d, e, f ,g, h, i, j)       // hi a hi b hi c hi d hi e hi f hi g hi h hi i hi j
```

## ARRAY (Data Structure)

```go
    // DECLARE TYPE
    var a [2]float32                                // var name [number]type

    // ASSIGN VALUE
    a[0] = 1.1                                      // name[number] = value
    a[1] = 2.0

    // DECLARE & ASSIGN (INITIALIZE)
    var b = [2]float32{1.1, 2.0}                    // Verbose - var name = [number]type{value, value, ...}
    c := [2]float32{1.1, 2.0}                       // Array Shorthand Assignment

    // PRINT
    fmt.Println(a, b, c)                            // [1.1 2] [1.1 2] [1.1 2]
```

## SLICE (Data Structure, Reference Type) (_make_)

```go
    // DECLARE TYPE - NO SIZE
    var a []float64                                 // var name []type

    // ASSIGN VALUE - ADD LENGTH TO SLICE
    a = append(a, 5.7)                              // name = append(name, value, value, ...)
    a = append([]float64{1.1}, a...)                // Trick to prepend to slice

    // DECLARE TYPE - WITH SIZE (make)
    b := make([]string, 1, 25)                      // name := make([]type, length, capacity)

    // ASSIGN VALUE
    b[0] = "hello"                                  // name[index] = value

    // DECLARE & ASSIGN (INITIALIZE)
    var c = []float32{1.1, 2.0}                     // Verbose - var name = []type{value, value, ...}
    d := []float32{3.4, 4.5}                        // Array Shortcut Assignment

    // PRINT
    fmt.Println(a, b, c, d)                         // [1.1 5.7] [hello] [1.1 2] [3.4 4.5]
```

## MAP (Data Structure, Reference Type) (_make_)

```go
    // DECLARE TYPES - THIS IS A NIL MAP - DON'T DO THIS
    var a map[string]int                            // var name map[keytype]valuetype

    // DECLARE TYPES (make)
    var b = make(map[string]int)                    // var name make(map[type]type)
    c := make(map[string]int)                       // name := make(map[type]type)

    // ASSIGN KEY:VALUE
    b["Jill"] = 23                                  // name[key] = value
    b["Bob"] = 34
    b["Mark"] = 28
    c["Jill"], c["Bob"], c["Mark"] = 23, 34, 28

    // DECLARE & ASSIGN KEY:VALUE (INITIALIZE)
    var d = map[string]int{                         // Verbose - var name = map[type]type {key:value, ...}
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }
    e := map[string]int{                            // Array Shortcut Assignment
        "Jill": 23,
        "Bob":  34,
        "Mark": 28,
    }

    // PRINT
    fmt.Println(a, b, c, d, e)                      // map[Jill:23 Bob:34 Mark:28] (For all of the maps)

    // DELETE A KEY
    delete(e,"Jill")                                // Delete key "Jill"
```

## STRUCT (Data Structure)

 ```go
    // CREATE STRUCT TYPE
    type Rect struct {                              // Define a struct
        w, h int
        }

    // DECLARE TYPE
    var r1 Rect                                     // Create a struct
    r2 := new(Rect)                                 // Returns Pointer

    // ASSIGN VALUE TO FIELDS
    r1.w = 2                                        // name.field = value
    r1.h = 4
    r2.w, r2.h = 3, 5

    // DECLARE & ASSIGN (INITIALIZE)
    var r3 Rect = Rect{2, 4}                        // Verbose (Don't use)
    var r4 = Rect{2, 4}                             // Type Inference - var name = structName{value, ....}
    r5 := Rect{w: 2, h: 4}                          // Shorthand Assignment
    r6 := Rect{2, 4}                                // Shorthand Assignment

    // PRINT
    fmt.Println(r1, *r2, r3, r4, r5, r6)            // {2 4} {3 5} {2 4} {2 4} {2 4} {2 4}
```

## POINTER (_new_)

```go
    // DECLARE A POINTER TYPE AND ASSIGN
    a := new(int)                                   // Create int pointer type
    *a = 9                                          // "Contents of a is 9"
    fmt.Println("Content of pointer *a is", *a)     // 9

    // ASSIGN A POINTER TYPE AND ASSIGN (OVERKILL)
    var c *int                                      // YOU DON'T NEED THIS
    b := 33                                         // This is an int
    c = &b                                          // "Address of b" is in pointer a
    fmt.Println("Content of pointer *a is", *a)     // 9

    // ASSIGN A POINTER TO A TYPE INT (INFERRED) - DO THIS
    b := 33                                         // If we have a var int 5
    c := &b                                         // assign c the "address of" b
    fmt.Println("Contents of pointer *c is", *c)    // 33
    fmt.Println("Address of &b is", &b)             // 0x02
    fmt.Println("Contents of pointer c is", c)      // 0x02

    // ASSIGN A POINTER TO A TYPE STRUCT
    var r1 Rect                                     // Create a struct (From above)
    r1.w = 3                                        // I wish it was *r1.w
    r1.h = 5                                        // I wish it was *r1.h
    d := &r1                                        // assign d the "address of" r1
    fmt.Println("Contents of pointer *d is", *d)    // {3 5}
    fmt.Println("Address of &r1 is", &r1)           // &{3 5}
    fmt.Println("Contents of pointer d is", d)      // &{3 5}
```

## FUNCTION AS A TYPE

```go
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
```

## FUNCTION

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

    // ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE
        // See above Function as a Type

    // CLOSURE - RETURN A FUNCTION TO A FUNCTION
        // See above Function as a Type

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

## METHOD

```go
    // BASIC FORMAT (Just a function with receiver)
    func receiver name(parameter list) (return type) {
        stuff
    }

    // Rect is a Rectangle struct type
    type Rect struct {
        w, h float32
    }

    // PASSING STRUCT TO METHOD BY VALUE (COPY) - STRUCT NOT CHANGED
    // Return Area
    func (r Rect) area() float32 {
        return r.w * r.h
    }

    // Scale Struct & Return Area
    func (r Rect) scaleArea(s float32) float32 {
        r.w = r.w * s
        r.h = r.h * s
        return r.area()
    }

    func main() {
        r1 := Rect{2, 4}
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area())       // Rect Area: 2 x 4 = 8
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.scaleArea(3)) // Rect Area: 6 x 12 = 72
    }

    // PASSING STRUCT TO METHOD BY "REFERENCE" (POINTER) - STRUCT CHANGED
    // Return Area
    func (r Rect) area() float32 {
        return r.w * r.h
    }

    // Scale Struct
    func (r *Rect) scaleStruct(s float32) {
        r.w = r.w * s
        r.h = r.h * s
    }

    func main() {
        r1 := Rect{2, 4}
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 2 x 4 = 8
        r1.scaleStruct(3)                                          // Passed by Reference
        fmt.Println("Rect Area:", r1.w, "x", r1.h, "=", r1.area()) // Rect Area: 6 x 12 = 72
    }
```

## INTERFACE (Reference Type)

```go
    // CREATE INTERFACE TYPE
    // STEP 1: Create your data types
    type myStructA struct {
        name string
    }
    type myStructB struct {
        x int
        y int
    }

    // STEP 2: Create methods with same name using your data types as receivers
    func (i myStructA) doThis() {
        fmt.Printf("I'm in doThis() method with receiver myStructA - %v\n", i.name)
    }
    func (i myStructB) doThis() {
        fmt.Printf("I'm in doThis() method with receiver myStructB - %v %v\n", i.x, i.y)
    }

    // STEP 3: Create your interface type with method
    type myInterfacer interface {
        doThis()
    }

    // STEP 4: Create a function that uses this interface as a parameter
    // INTERFACE AS A FUNCTION PARAMETER
    func magic(i myInterfacer) {
        i.doThis()
    }
```

## GOROUTINES

```go
    // GOROUTINES - CONCURRENT THREADS
    func doThis(s string) {
        do stuff
    }
    go doThis("Jeff")                               // Kick off goroutine
```

## CHANNELS (_make_)

```go
    // CHANNELS - GOROUTINE MESSAGE PIPES
    // SEND & RECEIVE SYNTAX
    msgCh := make(chan type, size)                  // name := make(chan type, buffer size)
    msgCh <- type                                   // SEND
    type := <-msgCh                                 // RECEIVE

    // NOT BUFFERED
    func doThis(msgCh <-chan string, t int) {
        rcv := <-msgCh                              // RECEIVE
        do stuff
    }
    msgJeffCh := make(chan string)                  // name := make(chan type)
    go doThis(msgJeffCh)                            // Kick off goroutine
    msgJeffCh <- "Jeff"                             // SEND

    // BUFFERED
    msgCh := make(chan string, 1)                   // name := make(chan type, buffer size)

    // CHANNEL DIRECTION (MORE EXPLICIT)
    func doThis(msgCh <-chan string, t int) {

    // SELECT (CASE) - WAITING FOR BOTH CHANNELS TO BE RECEIVED
    select {
    case msg1 := <-c1:
        fmt.Println("received", msg1)
    case msg2 := <-c1:
        fmt.Println("received", msg2)
```

## GO OPERATORS

```go
    // ARITHMETIC (MATH)
        // +, i, *, /, %,  ++, --

    // RELATIONAL (COMPARE)
        // ==, !=, >, <, >=, <=

    // LOGICAL (BOOLEAN)
        // &&, ||, !

    // BITWISE (BITS)
        // &, |, ^, <<, >>

    // ASSIGNMENT
        // =, +=, -=, *=, /=, %=, <=, >>=, &=, ^=, |=

    // MISCELLANEOUS
        // &, *
```

## CONTROL STRUCTURE / FLOW CONTROL

```go
    // LOOPS
        // FOR LOOP
        for i:=0; i < 8; i++ {
            fmt.Printf("Count is %v\n", i)
        }
        // WHILE LOOP
        i :=0
        for i < 8 {
            fmt.Printf("Count is %v\n", i)
            i++
        }
        // INFINITE LOOP
        for {
            do something forever
        }
        // RANGE LOOP
        myslice := []int{3, 4, 5}
        sum := 0
        for i := range myslice {
            fmt.Printf("Adding %v to sum %v\n", myslice[i], sum)
            sum += myslice[i]
        }
        // BREAK/CONTINUE
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

    // CONDITIONAL
        // IF, IF/ELSE, NESTED IF
        if a == b {
            fmt.Println("equal")
        } else if a > b {
            fmt.Println("higher")
        } else {
            fmt.Println("Lower")
        }
        // SWITCH (CASE)
        switch {
        case (a == b):
            fmt.Println("equal")
        case (a > b):
            fmt.Println("higher")
        default:
            fmt.Println("Lower")
        }
        // DEFER
        func main() {
            defer fmt.Println("world")
            fmt.Println("hello")
        }
        // SELECT (CASE) See Channels above
        // WAITING FOR BOTH CHANNELS TO BE RECEIVED
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c1:
            fmt.Println("received", msg2)
```

## ERROR HANDLING

```go
    // USING THE ERRORS (github.com/pkg/errors) PACKAGE (NOT STANDARD PACKAGE)
    // YOUR FUNCTION THAT RETURNS AN ERROR TYPE
    func a(x int) (int, error) {
        if x == 42 {
            // Make your error
            return -1, errors.New("can't work with 42")
        }
        return x + 3, nil
    }

    // YOUR ERROR CHECKER
    func checkErr(err error) {
        if err != nil {
            fmt.Printf("Error is %+v\n", err)
            log.Fatal("ERROR:", err)
        }
    }

    func main() {
        //MULTIPLE RETURN VALUES (Simpler form)
        r, err = b(42)
        checkErr(err)
        fmt.Println("Returned", r)
    }
```

## LOGGING

```go
    // USING THE LOGRUS (github.com/sirupsen/logrus) PACKAGE (NOT STANDARD PACKAGE)
    log.Panic("I'm bailing. Calls panic() after logging.")
    log.Fatal("Bye. Calls os.Exit(1) after logging.")
    log.Error("Something failed but I'm not quitting.")
    log.Warn("You should probably take a look at this.")
    log.Info("Something noteworthy happened!")
    log.Debug("Useful debugging information.")
    log.Trace("Something very low level.")
```

## FORMAT SPECIFIERS

```go
    // VALUE OF DEFAULT FORMAT
        // %v
            // %t                                   // boolean
            // %d                                   // signed
            // %d                                   // unsigned
            // %g or %f                             // float
            // %g or %f                             // complex
            // %s                                   // string
            // %p                                   // pointer
            // %p                                   // channel

    // BOOLEAN
        // %t

    // NUMERIC
        // %b, %c, %d, %o, %q, %x, %V, %U           // signed, unsigned, machine
        // %b, %e, %E, %f, %F, %g, %G               // float, complex

    // STRING
        // $s, %q, %x, %X

    // SLICE
        // %p

    // POINTER
        // %p
```

## ESCAPE SEQUENCES

```go
    // COMMON ONES USED IN FORMATTING
        // \n                                       // newline
        // \?                                       // ? character
        // \b                                       // backspace
        // \"                                       // " character
```

## GO PACKAGES

```go
    // GO GET A PACKAGE AND USE IT
    // As an example
    go get -u -v github.com/golang/protobuf/protoc-gen-go

    // Lets use it
    import (
        "fmt"
        "github.com/golang/protobuf/proto"
    )
    func main() {
        fmt.Println(proto.WireStartGroup)           // 3 (A constant in the package)
    }

    //  LETS CREATE A CUSTOM PACKAGE
    // Create a go package jeffshapes with methods, types, etc....
    package jeffshapes

    // Like above, get the package for your environment
    go get -u -v github.com/JeffDeCola/my-go-packages/jeffshapes

    // Like above, use in your go code
    import github.com/JeffDeCola/my-go-packages/jeffshapes
```

## GCO

```go
    package main

    /*
    int Add(int a, int b){
        return a+b;
    }
    */
    import "C"
    import "fmt"

    func main() {
        a := C.int(10)
        b := C.int(20)
        c := C.Add(a, b)
        fmt.Println(c)                              // 30
    }
```
