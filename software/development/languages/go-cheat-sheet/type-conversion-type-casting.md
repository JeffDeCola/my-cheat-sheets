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

mean = float32(sum)/float32(count)
```

Good example converting into utf-8,

```go
fmt.Println([]byte("hello"))
```