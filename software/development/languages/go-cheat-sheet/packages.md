# PACKAGES

Every thing in go is packages. A package is nothing but a
directory with some code files
It helps us organize our code into folders.

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

For some examples of using packages, checkout my repo
[my-go-example](https://github.com/JeffDeCola/my-go-examples#packages).

## LETS CREATE A CUSTOM PACKAGE

Lets create a simple `jeffshapes` package that will,

* Get the area and circumference for a circle

The code is located in
[my-go-packages](https://github.com/JeffDeCola/my-go-packages#jeffshapes).

First get the package,

```bash
go get -u -v github.com/JeffDeCola/my-go-packages/jeffshapes
```

Then import in your code,

```bash
import github.com/JeffDeCola/my-go-packages/jeffshapes
```
