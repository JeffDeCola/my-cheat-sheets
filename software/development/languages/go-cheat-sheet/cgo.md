# CGO (CALLING C WITH GO)

_`cgo` lets go call C code._

Table of Contents

* [EXAMPLE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/cgo.md#example)

## EXAMPLE

`import "C"` is immediately preceded by a comment, that comment,
called the preamble, is used as a header when compiling the C parts
of the package.

As an example,

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
    fmt.Println(c) // 30
}
```

You can run this example from my repo `my-go-examples`
[here](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/simple-c-code).

To use libraries like `stdio.h`, check out this example
[here](https://github.com/JeffDeCola/my-go-examples/tree/master/cgo/simple-c-code-using-stdio).
