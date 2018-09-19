# DERIVED DATA TYPES

I like to think of derived data types as
special variables built on the basic three data types
(Boolean, Numeric and String).

So you still need to declare type and assign value.

The data types in go,

* Boolean (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/data-types.md))
* Numeric (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/data-types.md))
* String (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/data-types.md))
* Derived (This Section)
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function (as a type)
  * Interface
  * Channel

## ARRAY

Really just an array of a variable; a data structure.

The basic verbose format is,

```
var name = [number]type{value, value....}
```

Here is the syntax,

```go

// DECLARE TYPE
var a [2]float32{}

// ASSIGN VALUE
a[1] = 1.1
a[2] = 2.0

// DECLARE & ASSIGN
var a = [2]float32{1.1, 2.0}                    // Verbose
a := [2]float32{1.1, 2.0}                       // Array Shorthand Assignment
```

Here is printing an array,

```go
testscores := [3]float64{78.3, 98.9, 85.4}
fmt.Printf("a is %v\n", testscores)
fmt.Printf("a is %v\n", testscores[1])
fmt.Printf("a is %v\n", testscores[1:2])
fmt.Printf("a is %v\n", testscores[:3])
fmt.Printf("a is %v\n", testscores[2:])
```

The output is,

```go
a is [78.3 98.9 85.4]
a is 98.9
a is [98.9]
a is [78.3 98.9 85.4]
a is [85.4]
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

## SLICE (ARRAY SUB TYPE)

Making an initial array of n length long out of a
total capacity.

The basic verbose format is,

```
var name = []type{value, value....}
```

Here is the syntax,

```go
// DECLARE TYPE
var a []float64

// ASSIGN VALUE / ADD TO SLICE
a = append(a, 5.7)

// DECLARE & ASSIGN (INITIALIZE)
var a = []float32{1.1, 2.0}                    // Verbose
a := []float32{3.4, 4.5}                       // Array Shortcut Assignment

// ADD TO SLICE
a := append(a, 5.7)                            // Append to different slice
```

## MAP (key:value)

Really key/value pairs, like a database.

```go
// DECLARE TYPES
var m1 = make(map[string]int)
m2 := make(map[string]int)

// ASSIGN KEY:VALUE
m1["Jill"] = 23
m1["Bob"] = 34
m1["Mark"] = 28
m2["Jill"], m2["Bob"], m2["Mark"] = 23, 34, 28

// DECLARE & ASSIGN
var m3 = map[string]int{                        // Verbose
    "Jill": 23,
	"Bob":  34,
	"Mark": 28,
}
m4 := map[string]int{                           // Array Shortcut Assignment
	"Jill": 23,
	"Bob":  34,
	"Mark": 28,
}

fmt.Println(m1, m2, m3, m4)
```

## STRUCT

Elements of different types and start with capital letter.
Anything with a capital letter is exported.

```go
// CREATE STRUCT TYPE
type Rect struct {
    w, h float32
}

// DECLARE TYPE
var r1 Rect
r2 := new(Rect)                                 // Returns Pointer

// ASSIGN VALUE TO FIELDS
r1.w = 6.1
r1.h = 5.0
r2.w, r2.h = 6.1, 6.0

// DECLARE & ASSIGN (INITIALIZE)
var r3 Rect = Rect{6.1, 5.0}                    // Verbose
var r4 = Rect{6.1, 5.0}                         // Type Inference
r5 := Rect{w: 6.1, h: 5.0}                      // Shortcut Assignment
r6 := Rect{6.1, 5.0}                            // Shortcut Assignment

fmt.Println(r1, *r2, r3, r4, r5, r6)
```

## POINTER

A pointer is just a variable that holds the address (of memory)
of a value.

`*` is the `"contents of"`
`&` is `"address of"`

You can,

```go
// CREATE A POINTER TYPE AND ASSIGN
a := new(int)                                   // Create int pointer type
*a = 9                                          // "contents of a is 9"

// ASSIGN A POINTER TO A TYPE
a := 5                                          // If we have a var int 5
b := &a                                         // b is the "address of" a
// var b *int = &a                              // Long form way to write this
// a == *b (both are 5)                         // "contents of" b is a

// ASSIGN A POINTER TO A STRUCT
b := &r1                                        // From struct Rect above
r1.w = 6.1                                      // I feel it should be *r1.w
r1.h = 5.0                                      // I feel it should be *r1.h
```

![IMAGE - go pointers - IMAGE](../../../../docs/pics/go-pointers.jpg)

Cool uses for pointers are,

* Allocate space for a variable.
* Pass by reference to a function to change parameters value outside function.

Another example, 

```go
a := 43
fmt.Println(a)  // 43
fmt.Println(&a) // Ox333333

// b is a pointer
// var b *int = &a
b := &a
fmt.Println(b)  // Ox333333
fmt.Println(*b) // 43

// Make the "contents of" what I'm Pointer to 33
*b = 33
fmt.Println(a) //33
```

## FUNCTION TYPE (FUNC EXPRESSION AND ANONYMOUS FUNC)

When used in a function, acts just like a type.
So I can use the variables int the function it lives.

```go

a, b := 3, 9

add := func() int {
    return a + b
}
```

This is also called func expression.

The func does not have a name and that is called
anonymous func.

## INTERFACE TYPE

Syntactic way to have multiple strucs do the same thing differently,

```go
// CREATE INTERFACE TYPE
type Name interface {
    methodName()
    ...
}
```

See own [section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md).

## CHANNEL TYPE

Basic Format,

```go
tbd
```

See own [section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/concurrency-channels.md).
