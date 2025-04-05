# TYPE CONVERSION & TYPE ASSERTION

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Converting one data type to another data type and
proving its actually that type._

Table of Contents

* [TYPE CONVERSION (TYPE CASTING)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/type-conversion-and-type-assertion.md#type-conversion-type-casting)
* [TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/type-conversion-and-type-assertion.md#type-assertion)

Documentation and Reference

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-cheat-sheet)
  main page

## TYPE CONVERSION (TYPE CASTING)

```go
type_name(expression)
```

int to float,

```go
i := 33
nowAFloat := float32(i) / 2.5  // int to float
fmt.Println(nowAFloat)         // Prints 13.2
```

## TYPE ASSERTION

A type assertion provides access to an interface value's underlying
concrete value,

```go
t := interface.(T)
```

```go
var j interface{} = "jeff"
s, ok := j.(string)
fmt.Println(s, ok) // Prints jeff true
f, ok := j.(float32)
fmt.Println(f, ok) // Prints 0, false
```
