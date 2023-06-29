# DATA TYPES

_Types or data types are a classification of data,
that tells the compiler how to use that data.
It also tells the computer how much space the data type
occupies and how the bit pattern is stored._

tl;dr,

```go
// BOOLEAN

    // VALUES
        // true, false

    // DECLARE TYPE
    var a bool                                      // var name type

    // ASSIGN VALUE
    a = true                                        // name = value

    // DECLARE & ASSIGN (INITIALIZE)
    var b = true                                    // Verbose - var name = value
    c := false                                      // Shorthand Assignment

    // PRINT
    fmt.Println(a, b, c)                            // true false true

// NUMERIC

    // TYPES
        // singed - int8, int16, int32 (rune), int64
        // unsigned - uint8 (byte), uint16, uint32, uint64
        // machine - int, uint, uintptr
        // float - float32, float64
        // complex - complex64, complex128

    // DECLARE TYPE
    var a int32                                     // var name type

    // ASSIGN VALUE
    a = 23                                          // name = value

    // DECLARE & ASSIGN (INITIALIZE)
    var b = 23                                      // Verbose - var name = value
    c := 23                                         // Shorthand Assignment

    // PRINT
    fmt.Println(a, b, c)                            // 23 23 23


// STRING (immutable array of bytes (or runes))

    // TYPE
        // string

    // DECLARE TYPE
    var a string                                    // var name string

    // ASSIGN VALUE
    a = "hello"                                     // name = value

    // DECLARE & ASSIGN (INITIALIZE)
    var b = "good"                                  // Verbose - var name = value
    c := "night"                                    // Shorthand Assignment

    // PRINT
    fmt.Println(a, b, c)                            // hello good night

// ARRAY (Data Structure)

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

// SLICE (Data Structure, Reference Type) (_make_)

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

// MAP (Data Structure, Reference Type) (_make_)

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

// STRUCT (Data Structure)

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

// POINTER (_new_)

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
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#overview)
* [BOOLEAN](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#boolean)
* [NUMERIC](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#numeric)
* [STRING](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#string)
* [ARRAY (Data Structure)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#array-data-structure)
* [SLICE (Data Structure & Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#slice-data-structure--reference-type-make)
* [MAP (Data Structure & Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#map-data-structure--reference-type-make)
* [STRUCT (Data Structure)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#struct-data-structure)
* [POINTER (_new_)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#pointer-new)
* [FUNCTION AS A TYPE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#function-as-a-type)
* [INTERFACE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#interface)
* [CHANNEL (Reference Type) (_make_)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/data-types.md#channel-reference-type-make)

## OVERVIEW

There are 4 data types in go.

* `Boolean`
* `Numeric`
* `String`
* `Derived` (Array, slice, map, struct, pointer, function, interface, channel)

`Data structures` are arrays, slices, maps and structs.  They store data.

`Reference types` are slice, map and channel,
meaning they are passed by reference/address.
Reference types are build upon an underlining data structure
(stores data). We are passing the address of these underlying
data structures. Also can use use _make_ to initialize reference type
before use.

## BOOLEAN

* `true`
* `false`

Boolean uses [logical operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#logical-boolean).

* `&&`, `||`, `!`

[Relational operators](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/operators.md#relational-compare)
(compare) returns a boolean (true or false).

* `==`, `!=`, `<`, `>`, `>=`, `<=`

## NUMERIC

Integers, floating point and complex numbers.

### INTEGER

Based on bit size and sign.

* Signed
  * `int8`
  * `int16`
  * `int32` (`rune`)
  * `int64`

* Unsigned
  * `uint8` (`byte`)
  * `uint16`
  * `uint32`
  * `uint64`

* Machine Dependent Types (Depends on your machine)
  * `int` _(The most popular and inferred)_
  * `uint`
  * `uintptr`

Ranges of some integer types,

```txt
`uint16` is 2^16-1 (0 to 65,535).
`int16` is 2^15-1 (-32,768 to 32,767).
`uint64` is 2^65-1 (0 to 18,446,744,073,709,551,615).
`int64` is 2^64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).
```

`Rune` is nice because it's a 4 bytes which is a character
of any language in the world.  UTF-8 uses 1-4 bytes where each number corresponds
to a character.

Use single quote for runes,

```go
fmt.Println('e') // Prints 105
```

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

### COMPLEX NUMBERS

* Complex number types
  * `complex64`
  * `complex128`

## STRING

`string` types are an immutable array of bytes (or runes).
They are just numbers.

* A `back-tick` string can contain newlines and double quotes.
* A `double quotes` string can contain special characters.

Since they are just arrays, you can index into a string.

As an example `e` is the rune number 101. Hence,

```go
    a := "hello"
    fmt.Println(string(a[1])) // Prints e
    fmt.Println(a[1])         // Prints 101
    fmt.Println('e')          // Prints 101
    fmt.Println("e")          // Prints e
```

## ARRAY (Data Structure)

Arrays are,

* A data structure (holds data)
* Do not change in size (not dynamic)
* A numbers sequence of elements of a single type
* A list/collection identified by an index
* Zero-based indexes

They are really not used that much.

### ARRAY - BASIC FORMAT

The basic verbose format is,

```go
var name = [number]type{value, value, ...}
```

A simple method is as follows,

```go
a := [2]float32{1.1, 2.0}
```

Here is the syntax,

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

Here is printing an array,

```go
testscores := [3]float64{78.3, 98.9, 85.4}
fmt.Printf("a is %v\n", testscores)             // [78.3 98.9 85.4]
fmt.Printf("a is %v\n", testscores[1])          // 98.9
fmt.Printf("a is %v\n", testscores[1:2])        // [98.9]
fmt.Printf("a is %v\n", testscores[:3])         // [78.3 98.9 85.4]
fmt.Printf("a is %v\n", testscores[2:])         // [85.4]
```

Remember a string is an array of bytes. Here is a good example,

```go
func main() {
    testString := "Happy Birthday Jeff"
    fmt.Printf("The first word of testString is %s\n", firstWord(testString))
}

func firstWord(str string) (word []byte) {
    for i := range str {
        if str[i] == ' ' {
            break
        } else {
            word = append(word, str[i])
        }
    }
    return word
}
```

Find length and capacity. They will be the same in an array,

```go
len(name)
cap(name)
```

## SLICE (Data Structure & Reference Type) (_make_)

Slices are for lists,

* A data structure (holds data)
* A reference type (pass by "reference")
* Must use _make_ to initialize reference type
* Dynamic - Changes in size
* Has length and capacity
* Built on an `underlying array`
  If you add things (append), a performance is hit because
  new arrays will be added (will double the size of the capacity)
* A list/collection identified by an index
* Value of uninitialized slice is nil

### SLICE - BASIC FORMAT

The basic verbose format is,

```go
var name = []type{value, value, ....}
```

The preferred method is as follows (where length and capacity are both at 25),

```go
m := make([]int, 25)
```

Here is the syntax,

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

### SLICES ARE REFERENCE TYPES

Like maps and channels, slices are reference types, meaning
slices are always passed by "reference/address".

```go
func change(a []string) {
    a[0] = "goodbye"
}
func main() {
    m := []string{"hello"}
    fmt.Println(m) // ["hello"]
    change(m)
    fmt.Println(m) // ["goodbye"]
}
```

### BUILT ON ARRAYS - make

Make makes a slice of length x and capacity (which is the length of the
underlying array) y,

```go
m := make([]string, length, capacity)
```

Both length and capacity are the same (100),

```go
m := make([]string, 100) // same as above
```

You could actually make a slice form an array,

```go
n := make([]int, 10, 18)
n := new([18]int)[0:10]
```

### VARIADIC PARAMETERS

Slices are used with variadic parameters (a function that will
take an arbitrary number of ints as arguments),

```go
func sum(n ...int) int {
    total := 0
    for _, f := range n {
        total += f
    }
    return total
}

func main() {
    data := []int{43, 44, 55, 66}
    fmt.Println(sum(data...))
}
```

### LENGTH AND CAPACITY

Find length,

```go
len(name)
```

Find capacity of array,

```go
cap(name)
```

Last example, these are all the same. Note capacity of array used for slice
is 16 in second example.  That's so cool.

```go
    var x [10]int
    for i := 0; i < len(x); i++ {
        x[i] = i + 20
    }
    fmt.Println(x, len(x), cap(x)) // [21 22 23 24 25 26 27 28 29 30] 10 10

    var y []int
    for i := 0; i < 10; i++ {
        y = append(y, i+20)
    }
    fmt.Println(y, len(y), cap(y)) // [21 22 23 24 25 26 27 28 29 30] 10 16

    z := make([]int, 10)
    for i := 0; i < 10; i++ {
        z[i] = i + 20
    }
    fmt.Println(z, len(z), cap(z)) // [21 22 23 24 25 26 27 28 29 30] 10 10
}
```

## MAP (Data Structure & Reference Type) (_make_)

Maps are,

* A data structure (holds data)
* A reference type (pass by "reference")
* Must use _make_ to initialize reference type
* key/value storage (like a data base) (dictionary)
* Unordered group of elements of the same type
* Value of uninitialized map is nil
* Map are build on a hash table

### MAP - BASIC FORMAT

The basic verbose format is,

```go
var name = map[type]type {key:value, key:value, ...}
```

The preferred method is as follows (where length and capacity are both at 25),

```go
m1 := map[type]type{key:value, key:value, ...}
```

Here is the syntax,

```go
// DECLARE TYPES - THIS IS A NIL MAP - DON'T DO THIS
var m1 map[string]int                           // var name map[type]type

// DECLARE TYPES (make)
var m1 = make(map[string]int)                   // var name make(map[type]type)
m2 := make(map[string]int)                      // name := make(map[type]type)

// ASSIGN KEY:VALUE
m1["Jill"] = 23                                 // name[key] = value
m1["Bob"] = 34
m1["Mark"] = 28
m2["Jill"], m2["Bob"], m2["Mark"] = 23, 34, 28

// DECLARE & ASSIGN KEY:VALUE (INITIALIZE)
var m1 = map[string]int{                        // Verbose - var name = map[type]type {key:value, key:value, ...}
    "Jill": 23,
    "Bob":  34,
    "Mark": 28,
}
m1 := map[string]int{                           // Array Shortcut Assignment
    "Jill": 23,
    "Bob":  34,
    "Mark": 28,
}
```

### MAPS ARE REFERENCE TYPES

Like slices and channels, maps are reference types, meaning
maps are always passed by reference/address.

```go
func change(a map[int]string) {
    a[1] = "happy"
    a[2] = "birthday"
}
func main() {
    m := map[int]string{1: "hello", 2: "goodbye"}
    fmt.Println(m) // map[1:hello 2:goodbye]
    change(m)
    fmt.Println(m) // map[1:happy 2:birthday]
}
```

### DELETE A KEY/VALUE

Delete a key/value,

```go
delete(a,1)
delete(m1,"Jeff")
```

## STRUCT (Data Structure)

* A data structure (holds data)
* Composite types.
* A struct is a collection of fields.

Elements of different types and start with capital letter.
Anything with a capital letter is exported from package.

The basic verbose format is,

```go
var name = structName{value, value, ....}
```

A struct is a data structure.

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

A pointer is just a variable that holds the address (of memory)
of a value.

* `*` is the `"contents of"`
* `&` is `"address of"`

### POINTER -  BASIC FORMAT

Here is an example,

```go
// DECLARE A POINTER TYPE AND ASSIGN
a := new(int)                                   // Create int pointer type
*a = 9                                          // "Contents of a is 9"
fmt.Println("Content of pointer *a is", *a)     // 9

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

![IMAGE - go pointers - IMAGE](../../../../docs/pics/go-pointers.jpg)

### POINTER USES

Cool uses for pointers are,

* Allocate space for a variable.
* Pass by "reference" to a function to change parameters value outside function.

### POINTER PASS BY REFERENCE

As an example,

```go
func change(b *int) {
    *b = 6
}

func main() {
    a := 5  // If we have a var int 5
    b := &a // b is the "address of" a

    fmt.Println(*b) // 5
    change(b)
    fmt.Println(*b) // 6
}
```

For some pointer examples, refer to my repo my-go-examples
[here](https://github.com/JeffDeCola/my-go-examples#basic-syntax).

## FUNCTION AS A TYPE

Functions can be a type. So like types, I can use
the variables (the scope) the function lives in.

There is a better explanation of this on the cheat sheet
[closures](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/functions.md#closure-func-expression-and-anonymous-func).

But here are two methods that can be used,

### ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE

```go
// ASSIGN ANONYMOUS FUNCTION (func LITERAL) TO A VARIABLE
a, b := 3, 9
add := func() int {                             // Anonymous func as a type (no name)
    return a + b                                // returns int
}

fmt.Println(add())                              // 12
a = 9
fmt.Println(add())                              // 18
```

The anonymous function (function literal) has use of `a` and `b` because
of the scope.  This assignment of an anonymous function (function literal)
to a variable is called `closure`.

### CLOSURE - RETURN A FUNCTION TO A FUNCTION

In this method, the function scope acts just like an assigned variable.
Think of the function like a variable, because that's what is really is,

```go
// RETURN A FUNCTION TO A FUNCTION
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

## INTERFACE

Syntactic way to have multiple structs do the same thing differently,

Basic format,

```go
// CREATE INTERFACE TYPE
type Name interface {
    methodName()
    ...
}
```

See the interface [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md).

## CHANNEL (Reference Type) (_make_)

tbd

Basic Format,

```go
tbd
```

Like slices and maps, channels are reference types, meaning
channels are always passed by reference/address.

See channel [cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md).
