# DERIVED DATA TYPES

The derived data types are built on the basic three data types
(Boolean, Numeric and String).

The Data Types in go,

* Boolean (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Numeric (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* String (See previous [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/data-types.md))
* Derived
  * Array
  * Slice
  * Map
  * Struct
  * Pointer
  * Function (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/functions.md))
  * Interface (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/interfaces.md))
  * Channel (See own [Section](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/languages/go-cheat-sheet/concurrency-channels.md))

## ARRAY

Really just an array of a variable; a data structure.

The basic format is,

```
name = [number]type{assignment, assignment....}
```  

Declare,

students := []string{}

Assign

```
testscores := [3]float64{78.3, 98.9, 85.4}
```

All arrays are zero based so `testscores[0]` is 78.3.

For example,

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

## SLICE (ARRAY SUB TYPE) (make)

Making an initial array of n length long out of a
total capacity,

```go
// initial slice is 10 long out of 100
a := make([]int, 10, 100)
fmt.Printf("a is %v\n", a)

// adding to a slice
a = append(a, 20)
fmt.Printf("a is %v\n", a)
```

Assignment, rather then using make,

```go
c := []int{5, 6}
fmt.Println(c)
c = append(c, 7)
fmt.Println(c)
```

Output is,

```
[5 6]
[5 6 7]
```

## MAP (key:value)

Really key/value pairs, like a database.

```go
ages := map[string]int{
    "Jill": 23,
    "Bob":  34,
    "Mark": 28,
}

fmt.Println(ages["Jill"])
```

## STRUCT

Elements of different types and start
with capital letter.  Because
anything with a capital letter is exported.

type Rect struct {
    w,h,float64
}

## POINTER

A pointer is just a variable that holds the address (of memory)
of a value.

`*` is the contents of operator
`&` is address of operator.


