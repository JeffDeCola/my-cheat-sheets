# PACKAGES

Every thing in go is packages. A package is nothing but a
directory with some code files
It helps us organize our code into folders,

* [godoc.org](https://godoc.org/)
  _- Both standard and user packages. Also shows popular packages._
* [golang.org](https://golang.org/pkg/)
  _-Just official standard packages._

Table of Contents,

* [GO GET A PACKAGE AND USE IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md#go-get-a-package-and-use-it)
* [LETS CREATE A CUSTOM PACKAGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md#lets-create-a-custom-package)

## GO GET A PACKAGE AND USE IT

You already have the regular go packages.
Lets bring another package to your computer and use it,

```go
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

`-u -v` looks for update and is verbose.

The package in now located in your `/src` folder.

Now use the package. I know there is a constant
WireStartGroup that is 3. Hence,

```go
import (
    "fmt"
    "github.com/golang/protobuf/proto"
)

func main() {
    fmt.Println(proto.WireStartGroup)   // 3
}
```

## LETS CREATE A CUSTOM PACKAGE

Lets create a simple shapes package that will,

* Area and circumference for circles

The code is located in
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/shapes).




We will call the package shapes.  Hence the folder name must be called shapes.

    circle.shapes.area()
    rectangle.shapes.area()

In shapes I will put methods.
