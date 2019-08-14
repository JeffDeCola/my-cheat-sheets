# ERROR HANDLING & LOGGING

Using the non-standard `github.com/pkg/errors` and
the `github.com/sirupsen/logrus` package.

## ERROR HANDLING

I use the go package
[errors](http://github.com/pkg/errors).

Go doesn’t have exceptions, so it doesn’t have try, catch or anything similar.
So how do we handle errors?

* Functions return errors just like any other value. `Multiple return values`
  distinguish errors from normal return values.
* Errors are the last return value and have type error, a built-in interface.
* Errors are handled by checking the value(s) returned from a function.
* A nil in the error type means no error

As and example of a function that has multiple return values, including type error.

```go
// YOUR FUNCTION THAT RETURNS AN ERROR TYPE
func a(x int) (int, error) {
    if x == 42 {
        // Make your error
        return -1, errors.New("can't work with 42")
    }
    return x + 3, nil
}

// YOUR ERROR CHECKER
func checkErr(err error) {
    if err != nil {
        fmt.Printf("Error is %+v\n", err)
        log.Fatal("ERROR:", err)
    }
}

func main() {
    //MULTIPLE RETURN VALUES (Simpler form)
    r, err = b(42)
    checkErr(err)
    fmt.Println("Returned", r)
}
```

Refer to
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/errors)
for an example using `github.com/pkg/errors` package.

## LOGGING

I use the go package
[logrus](https://github.com/sirupsen/logrus).

A logger or logging give you information from a running program.
Typically it uses stderr or a file.

I like the following levels,

* Panic
* Fatal
* Error
* Warning
* Info
* Debug
* Trace

Refer to
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus)
for an example using `logrus` package.
