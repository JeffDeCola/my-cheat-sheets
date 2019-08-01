# TYPE CONVERSION & TYPE ASSERTION

Converting one data type to another data type and
proving its actually that type.

* [TYPE CONVERSION (TYPE CASTING)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/type-conversion-type-casting.md#type-conversion-type-casting)
* [TYPE ASSERTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/type-conversion-type-casting.md#type-assertion)

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
