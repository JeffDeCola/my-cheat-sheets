# BASIC CONCEPTS

Go likes to force you to use a coding style.  This is a little
painful at first, but actually good news in the long run since
it makes everyone's code look and feel the same.

## STRUCTURE OF GO

* Package Declaration
* Import Packages
* Functions
* Variables
* Statements and Expressions
* Comments

### BASIC SYNTAX

`Tokens` are basically the building blocks of go.  For example, `fmt.Println("Hello")` has six tokens.

A `Line Separator` is a statment terminator.  It's like placing a `;`.

`Comments` start with `/*` and end with `*/`.

`Identifiers` is a name used to identify a variable, function or user defined term,

`Keywords` are reserved like `func`, `for`, `if`, `return`, etc..

`Whitespace` is ignored.

## PACKAGES

Everything in go is a package.

[List of go packages](http://golang.org/pkg).

Or use go tool,

```bash
go doc fmt
```

`import "fmt"` just imports another package.

## RUN

```bash
go run hello.go
```

Just puts the executable in a temp location, then kills it.

## BUILD

Compiles packages and dependencies (does not install results)
and saves your binary,

```bash
go build hello.go
./hello
```

### INSTALL

Compiles packages and dependencies and places in /bin and /pkg,

```bash
go install hello.go
$GOPATH/bin/hello
```
