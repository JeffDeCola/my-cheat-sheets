# FORMATTING TYPES

How to make the output look pretty.

## FORMAT SPECIFIERS

There is a large list [here](https://golang.org/pkg/fmt/).

Some of the common ones,

* `%v` all types - the value in a default format.
* `%s` string - the uninterpreted bytes of the string or slice.
* `%f` floating point - decimal point but no exponent, e.g. 123.456.
* `%p` slice - address of 0th element in base 16 notation, with leading 0.

```go
a := 3.14234234
b := "happy birthday"
c := []"make a slice" ????????????????????????????????
a := make([]int, 10, 100)

fmt.Printf("%v\n", a)
fmt.Printf("%.2f\n", a)
fmt.Printf("%s\n", b)
fmt.Printf("%s\n", c) ??????????????????????????
```

## ESCAPE SEQUENCES

These are really constants, but I put them here because they are mainly used in formatting.

Popular ones

* `\n` newline.
* `\?` The ? character.
* `\b` backspace.
* `\"` The " character.