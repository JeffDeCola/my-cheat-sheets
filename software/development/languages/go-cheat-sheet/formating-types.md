# FORMATTING TYPES

How to make the output look pretty.

* [FORMAT SPECIFIERS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/formating-types.md#format-specifiers)
* [ESCAPE SEQUENCES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/formating-types.md#escape-sequences)

## FORMAT SPECIFIERS

There is a large list [here](https://golang.org/pkg/fmt/).

Default types using %v,

* `%v` Default format
* `%+v` Adds field names
  * `%t` boolean
  * `%d` signed
  * `%d` unsigned
  * `%f` or `%g` float
  * `%f` or `%g` complex
  * `%s` string
  * `%p` pointer
  * `%p` channel

Common ones,

* `%t` boolean
* `%b` integer base 2
* `%d` integer base 10
* `%x` integer base 16
* `%f` float (no exponent)
* `%g` float
* `%g` complex
* `%s` string
* `%p` pointer
* `%p` channel
* `%p` slice - address of 0th element in base 16 notation, with leading 0

As an example,

```go
var myboolean bool = true
var myint int = -33
var myuint32 uint32 = 33
var myfloat32 float32 = 3.1425
var mycomplex64 complex64 = 33.23423432
var mystring string = "Jeff"
var myslice = []string{"jeff", "larry"}
var mymap = map[int]string{1: "cat", 2: "dog"}
mypointer := new(string)
*mypointer = "jeffry"

fmt.Printf("myboolean    %v or %t\n", myboolean, myboolean)     // true or true
fmt.Printf("myint        %v or %d\n", myint, myint)             // -33 or -33
fmt.Printf("myuint32     %v or %d\n", myuint32, myuint32)       // 33 or 33
fmt.Printf("myfloat32    %v or %g\n", myfloat32, myfloat32)     // 3.1425 or 3.1425
fmt.Printf("myfloat32    %f or %.2f\n", myfloat32, myfloat32)   //  3.142500 or 3.14
fmt.Printf("mycomplex64  %v or %g\n", mycomplex64, mycomplex64) // (33.234234+0i) or (33.234234+0i)
fmt.Printf("mystring     %v or %s\n", mystring, mystring)       // Jeff or Jeff
fmt.Printf("myslice     %v or %s\n", myslice, myslice)          // [jeff larry] or [jeff larry]
fmt.Printf("myslice     %v or %s\n", myslice[0], myslice[1])    // jeff or larry
fmt.Printf("mymap       %v or %s\n", mymap, mymap[1])           // map[1:cat 2:dog] or cat
fmt.Printf("mypointer   %v or %p\n", mypointer, mypointer)      // 0xc42000e1e0 or 0xc42000e1e0
fmt.Printf("mypointer   %v or %s\n", *mypointer, *mypointer)    // jeffry or jeffry
fmt.Printf("mypointer   %v or %p\n", &mypointer, &mypointer)    // 0xc42000c028 or 0xc42000c028
```

## ESCAPE SEQUENCES

These are really constants, but I put them here because they are mainly used in formatting.

Popular ones,

* `\n` newline
* `\?` ? character
* `\b` backspace
* `\"` " character
