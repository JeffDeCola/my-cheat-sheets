# TYPE CONVERSION / TYPE CASTING

Converting one data type to another data type.

* [BASIC FORMAT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/type-conversion-type-casting.md#basic-format)

## BASIC FORMAT

```go
type_name(expression)
```

As an example,

```go
var sum int = 17
var count int = 5
var mean float32

mean = float32(sum) / float32(count)
fmt.Println(mean) // Prints 3.4
```

An example of converting a string into utf-8,

```go
fmt.Println([]byte("hello")) // [104 101 108 108 111]
```
