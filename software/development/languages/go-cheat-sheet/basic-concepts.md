# BASIC CONCEPTS

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Go actually forces you to use a particular coding style.  This is a little
painful at first, but actually great in the long run since
it makes everyone's code look and feel the same.
It also will not allow pollution in your code (things you
don't use)._

Table of Contents

* [BASIC STRUCTURE OF GO](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#basic-structure-of-go)
* [BASIC SYNTAX](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#basic-syntax)
* [PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#packages)
* [GO RUN](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#go-run)
* [GO BUILD](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#go-build)
* [GO INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#go-install)
* [THE OBJECT SIDE OF GO](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/basic-concepts.md#the-object-side-of-go)

Documentation and Reference

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-cheat-sheet)
  main page

## BASIC STRUCTURE OF GO

A go program usually has the following things,

* Package Declaration
* Import Packages
* Functions
* Variables
* Statements and Expressions
* Comments

A `statement` is an action (_print, if, calling a func_) and
mostly vertical in your code.

An `expression` produces a value (_x == 2_) and is mostly horizontal
in your code.

## BASIC SYNTAX

* `Tokens` are basically the building blocks of go.  For example,
  `fmt.Println("Hello")` has six tokens.
* `Line Separator` is a statement terminator.  It's like placing a `;`.
* `Comments` start with `/*` and end with `*/` or use `//` for lines.
* `Identifiers` is a name used to identify a variable, function or user defined term.
* `Keywords` are reserved like `func`, `for`, `if`, `return`, etc..
* `Whitespace` is ignored.

## PACKAGES

Everything in go is a package.

Here is a great [list of go packages](http://golang.org/pkg).

Or use a go tool to look up packages,

```bash
go doc fmt
```

`import "fmt"` just imports the package fmt.

Anything with a capital letter in a package is exported.

Refer to my cheat sheet on packages
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

## THE OBJECT SIDE OF GO

Object oriented programing is,

* Properties - What it knows.
* Methods - What it does.

Go is similar,

* Structs - The data structure of the properties.
  What we know.
* Methods -  A special function (with a receiver type)
  that does something with the struct (properties).
