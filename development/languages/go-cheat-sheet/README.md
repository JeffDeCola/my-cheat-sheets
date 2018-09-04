# GO CHEAT SHEET

*** THIS CHEAT SHEET IS CURRENTLY BEING DEVELOPED - CHECK BACK SOON ***

`go` _is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems._

This is a very abbreviated cheat-sheet highlighting the main
syntax of go.

The cheat sheet is broken up into the following sections,

* REFERENCES / DOCUMENTATION

* INSTALL & CONFIGURE
  * _Linux, Bash on Ubuntu on Windows, Windows._

* BASIC CONCEPTS
  * _packages, run, build, install._

* TYPES
  * _integer, floating point, complex, string, boolean._

* DECLARATION / ASSIGNMENT
  * _variables, type inference, shortcut assignment, group your variables._



Type Conversion, arrays, slices, structs, pointers, maps._



* CONTROL STRUCTURES / FLOW CONTROL
  * _Loops (for loop), Conditional Statements (if/else, switch, defer)._

* FUNCTIONS (Black Box)
  * _Basic Format, passing parameters by reference, by value._

* METHODS (Attached to Data)
  * _Basic Format._

* INTERFACES (????)
  * _Basic Format._

* CONCURRENCY
  * _goroutines._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

### REFERENCES / DOCUMENTATION

* [golang.org](http://golang.org) - _Home base for everything._
* [golang.org docs](https://golang.org/doc/) _A good collection of docs._
* [A tour of go](https://tour.golang.org/welcome/1) _A good place to start._
* [Effective go](https://golang.org/doc/effective_go.html) _A must read to create great things._
* [Plugin for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) _I use
  [visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)._

## ************** INSTALL & CONFIGURE ************

OK, lets get go...ing.  Yep, that just happened.

### INSTALL

Instructions for binary and source installs.

[Installs](https://golang.org/doc/install) for windows, linux or mac.

I would not install from a package.

### CONFIGURE - LINUX

I place this is `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

### CONFIGURE - BASH ON UBUNTU ON WINDOWS

I place this is `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=/mnt/c/Users/<WINDOWSNAME>/home/<USERNAME>
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

The trick is to have a different project/working directory as
shown with my GOPATH. Do not work in the home directory.

### CONFIGURE - WINDOWS

If you install it will automatically set the Windows environment
variables as follows,

```text
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```

On a side note, my cheat sheet
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)
explains how to setup VS Code on Windows with the Go Extensions
and a bash terminal.  Say that ten times fast.

## **************** BASIC CONCEPTS ***************************

Go likes to force you to use a coding style.  This is a little
painful at first, but actually good news in the long run since
it makes everyone's code look and feel the same.

### PACKAGES

Everything in go is a package.

[List of go packages](http://golang.org/pkg).

Or use go tool,

```bash
go doc fmt
```

`import "fmt"` just imports another package.

### RUN

```bash
go run hello.go
```

Just puts the executable in a temp location, then kills it.

### BUILD

Compiles packages and dependencies (does not install results)
and saves your binary,

```bash
go build hello.go
./hello
```

#### INSTALL

Compiles packages and dependencies and places in /bin and /pkg,

```bash
go install hello.go
$GOPATH/bin/hello
```

## ***************** TYPES ************************************

Types or data types are a classification of data,
that tells the compiler how to use that data.

### INTEGER 

Based on bit size and sign.

* Signed,
  * `int8`
  * `int16`
  * `int32`
  * `int64`

* Unsigned,
  * `uint8` (`byte`)
  * `uint16`
  * `uint32` (`rune`)
  * `uint64`

* Machine dependent types (Depends on system),
  * `int` Most popular and inferred.
  * `uint` 
  * `uintptr`

Ranges of some integer types,

```txt
`uint16` is 2^16-1 (0 to 65,535).
`int16` is 2^15-1 (-32,768 to 32,767).
`uint64` is 2^65-1 (0 to 18,446,744,073,709,551,615).
`int64` is 2^64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).
```

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types,
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

### COMPLEX

* Complex number types,
  * `complex64`
  * `complex128`

### STRING

`string` types are arrays of bytes or runes.

A back-tick string can contain newlines.

A double quotes can contain special characters.

Since they are just arrays, you can index into a string.

### BOOLEAN

* `true`
* `false`

Boolean operators
* `&&`, `||`, `!`_

Relational Operators returns a boolean (true or false)
*  `==`, `!=`, `<`, `>`, `>=`, `<=`_

## ***************** DECLARATION / ASSIGNMENT ************************************

Declaration is picking a name for a data type and optionally assigning
that data type a value.

### VARIABLES

Variables always has a single type and may be assigned.

The basic format is,

```txt
var name type
var name type = assignment
```

For example,

```go
var greeting string = "hello, world"
var pi float32 = 3.14
```

### TYPE INFERENCE

The compile will infer a type if none is giving
based on your assignment.

```go
var age = 42         // will infer an int
var pi = 3.14        // Will infer a float64
var cNum = 3 + 5i    // will infer a complex128
```

### SHORTCUT ASSIGNMENT

You may use the shortcut `:=`,

```go
x := 42
```

this is the same as,

```go
var x int = 42;
var x = 42;
```

Shortcuts can not be used outside function.

### GROUP YOUR VARIABLES

```go
var x int = 5
var y int = 8
```

Group also can infer a type,
```
var (
    x=5
    y=8
)
```

But even better, put it on one line,

```go
var x, y = 5, 8
```

But even even better lets use shorthand
(again, only inside a function),

```go
x, y := 5, 8
```






## TYPE CONVERSION

### ARRAYS - COLLECTION TYPES?

All arrays are zero based.

### SLICES (ARRAY SUB TYPE)

### STRUCTS

Elements of differnet types.

### POINTERS

Pointers and structs go hand in hand.

### MAPS







## ***************** CONTROL STRUCTURES / FLOW CONTROL *********************

Control structures are????

### LOOPS

for loop

### CONDITIONAL STATEMENTS

if /else, switch ,defer??

## ****************** FUNCTIONS ******************************

Functions stand on their own, a black box.

### BASIC FORMATS 

The basic format is,

```txt
func name(parameter list - optional)(return type - optional) {
    stuff
}
```

As an example,

```go
func add(a, b int) int {
    return a + b
}
```

Variadic Parameter List - Don't know how many parameters are being passed,

```go
func add(name ...int) int {
    ....
    return sum
}
```

Multiple returns,

```go    
func swap(a, b int) (int, int) {
	x, y := b, a
	return x, y
}
```

Name your returned parameters,

```go
func swap(a, b int) (x int, y int) {
    x, y = b, a
    return
}
```

### PASSING PARAMETERS BY REFERENCE

### PASSING PARAMETERS BY VALUE

## ****************** METHODS ******************************

Methods, unlike functions which stand on their own,
are attached to data.

### BASIC FORMATS

Using function as methods of structures?

## ****************** INTERFACES  ******************************

Interfaces...

### BASIC FORMATS 

## ****************** CONCURRENCY  ******************************
