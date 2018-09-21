# TYPE CONVERSION / TYPE CASTING

Converting one data type to another data type.

## BASIC FORMAT

```
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