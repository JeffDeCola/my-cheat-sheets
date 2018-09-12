# FORMATTING TYPES

How to make the output look pretty.

## READ INPUT (bufio.NewReader PACKAGE)

Get an enter string as input,

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter a String")
str, err := reader.ReadString('\n')
if err != nil {
    panic("error: newline not found")
}

fmt.Println(str)
```

## SCAN INPUT (fmt.Scan PACKAGE)

```go
word1, word2 := "", ""
count, err := fmt.Scan(&word1, &word2)
if err != nil {
    panic(err)
}
fmt.Printf("You entered %v.  There are %v words. \n", word1+" "+word2, count)
```

## FORMAT SPECIFIERS

There is a large list [here](https://golang.org/pkg/fmt/).

Some of the common ones,

* `%v` Default format.
  * boolean - %t
  * signed - %d
  * unsigned - %d
  * float -%g
  * complex - %g
  * string - %s
  * pointer - %p
  * channel - %p
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

Popular ones,

* `\n` newline.
* `\?` The ? character.
* `\b` backspace.
* `\"` The " character.