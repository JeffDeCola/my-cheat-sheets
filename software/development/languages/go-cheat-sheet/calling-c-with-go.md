# CGO (CALLING C WITH GO)

Cgo lets go call C code.

## EXAMPLE

`import "C"` is immediately preceded by a comment, that comment,
called the preamble, is used as a header when compiling the C parts
of the package.

As an example,

```go
package main

//int Add(int a, int b){
//    return a+b;
//}
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
