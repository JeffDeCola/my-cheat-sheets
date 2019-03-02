# PACKAGES

Every thing in go is packages.

It helps us organize our code into folders.

* [godoc.org](https://godoc.org/)
  _- Both standard and user packages._
* [golang.org](https://golang.org/pkg/)
  _-Just official standard packages._

## GO GET A PACKAGE AND USE IT

Bring code to computer and use it.

```go
go get ??????
```

Now use it,

```go
import "path/to/the/package"
```

## LETS CREATE A PACKAGE

When something is Capitalized, its exported outside package.

Lets create a package based on the [interface example](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/interfaces.md).

* Area and circumference for circles
* Volume and surface area for cylinders.

We will call the package shapes.  Hence the folder name must be called shapes.

    circle.shapes.area()
    rectangle.shapes.area()

In shapes I will put methods.
