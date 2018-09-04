# GO CHEAT SHEET

*** THIS CHEAT SHEET IS CURRENTLY BEING DEVELOPED - CHECK BACK SOON ***

`go` _???_

This is a very abbreviated cheat-sheet highlighting the main
syntax of go.

Go is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems.

The cheat sheet is broken up into the following sections,

* REFERENCES / DOCUMENTATION

* INSTALL / CONFIGURE
  * _For Linux, Bash on Ubuntu on Windows, Windows._

* BASIC CONCEPTS
  * _packages, run, build, install._

* TYPES
  * _integer, floating, string, boolean._

* DECLARATION
  * _Variables, constant._ 

Type Conversion, arrays, slices, structs, pointers, maps._

* CONTROL STRUCTURES / FLOW CONTROL
  * _Loops (for loop), Conditional Statements (if/else, switch, defer?)._

* FUNCTIONS
  * _Stand on own. Passing Parameters by reference or by value._

* METHODS
  * _Attached to data._

* INTERFACES
  * _ ??._

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

## ************** INSTALL AND CONFIGURE ************

### INSTALL

[Installs](https://golang.org/doc/install) for windows, linux or mac.

Instructions for binary and source.  I would not install from a package.

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

The trick is to have a different project directory as
shown with my GOPATH. Do not work in the home directory.

### CONFIGURE - WINDOWS

If you install it will automatically set the Windows environment
variables as,

```text
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```
On a side note, my cheat sheet
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)
explains how to setup VS Code on Windows with Go extensions
and a bash terminal.

## **************** BASIC CONCEPT ***************************

### PACKAGES

Everything in go is a package.

[List of go packages](http://golang.org/pkg).

Or use go tool,

```bash
go doc fmt
```

### RUN

```bash
go run hello.go
```

### BUILD

Compiles packages and dependencies (does not install results),

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
  * `int` (MOST POPULAR)
  * `uint` 
  * `uintptr`

Ranges of some integer types,

```txt
`uint16` is 2^16-1 (0 to 65,535).
`int16` is 2^15-1 (-32,768 to 32,767).
`uint64` is 2^65-1 (0 to 18,446,744,073,709,551,615)
`int64` is 2^64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).
```

### FLOATING POINT

Floating point are stored as `significand × base^exponent`
where base is usually 10.

It means the decimal point floats.

* Floating point types,
  * `float32` (significand is 24 bits)
  * `float64` (significand is 53 bits)

* Complex number types,
  * `complex64`
  * `complex128`

### STRING

`string` types are arrays of bytes or runes.

A back-tick string can contain newlines.

A double quotes can contain special characters.

Since they are just arrays, you can index into a string.

### BOOLEAN

* true
* false

Boolean operators
* && || !

Relational Operators returns a boolean (true or false)
*  == != < < >= <=

## ***************** DECLARATION ************************************

Variables always has a single type and can be assigned.

```txt
var name type = assignment
```

E.g.

```go
var greeting string = "hello, world"
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

### LOOPS

for loop

### CONDITIONAL STATEMENTS

if /else, switch ,defer??

## ****************** FUNCTIONS ******************************

Stand on own, black box.

### Passing Parameters by reference or by value_

## ****************** METHODS ******************************

Attached to data.

Using fuction as methods of structures?

## ****************** INTERFACES  ******************************

## ****************** CONCURRENCY  ******************************
