# BASIC CONCEPTS

Go actually forces you to use a particular coding style.  This is a little
painful at first, but actually very good news in the long run since
it makes everyone's code look and feel the same.

It also will not allow pollution in your code, meaning things you
don't use.  Makes lean readable code.

Its also very fun to code.

## BASIC STRUCTURE OF GO

A go program usually has the following things,

* Package Declaration.
* Import Packages.
* Functions.
* Variables.
* Statements and Expressions.
* Comments.

A statement is an action (print, if, calling a func) and
mostly vertical in your code.

An expression produces a value (x == 2) and is mostly horizontal
in your code.

## BASIC SYNTAX

`Tokens` are basically the building blocks of go.  For example,
`fmt.Println("Hello")` has six tokens.

A `Line Separator` is a statement terminator.  It's like placing a `;`.

`Comments` start with `/*` and end with `*/` or use '//'.

`Identifiers` is a name used to identify a variable, function or user defined term.

`Keywords` are reserved like `func`, `for`, `if`, `return`, etc..

`Whitespace` is ignored.

## PACKAGES

Everything in go is a package.

Here is a huge [list of go packages](http://golang.org/pkg).

Or use a go tool to look up packages,

```bash
go doc fmt
```

`import "fmt"` just imports the package fmt.

Anything with a capital letter in a package is exported.

Refer to the section on packages
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet/packages.md).

## GO RUN

```bash
go run hello.go
```

Just puts the executable in a temp location, then deletes it.

## GO BUILD

Compiles packages and dependencies (does not install results)
and saves your binary (executable) in current directory,

```bash
go build hello.go
./hello
```

## GO INSTALL

Compiles packages and dependencies and places in `/bin` and `/pkg`,

```bash
go install hello.go
$GOPATH/bin/hello
```
