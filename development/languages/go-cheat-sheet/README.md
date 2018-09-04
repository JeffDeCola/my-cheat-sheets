# GO CHEAT SHEET

`go` _._

This is a very abbreviated cheat-sheet highlighting the main
syntax of go.

Go is an open source language developed by google. Its concurrency
mechanisms allows Apps to get the most out of multi core and
networked systems.

The cheat sheet is broken up into the following sections,

* REFERENCES / DOCUMENTATION
* INSTALL / CONFIGURE
  * For Linux, Bash on Ubuntu on Windows, Windows._
* BASIC CONCEPT
  * _Packages._
* TYPES
  * _Basic Types(variables, constants), Type Conversion, arrays, slices, structs, pointers, maps._
* CONTROL STRUCTURES / FLOW CONTROL
  * _Loops (for loop), Conditional Statements (if / else, switch, defer?)._
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
PATH=$PATH:$GOROOT/bin
```

### CONFIGURE - BASH ON UBUNTU ON WINDOWS

I place this is `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=/mnt/c/Users/<WINDOWSNAME>/home/<USERNAME>
PATH=$PATH:$GOROOT/bin
```

The trick is to have a different project directory as
shown with my GOPATH. DO NOT work in the home directory.

### CONFIGURE - WINDOWS

If you install it will automatically set the Windows environment
variables as,

```bash
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```
On a side note, here is my cheat sheet
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)
on how to setup VS Code on Windows with Go.

## **************** BASIC CONCEPT ***************************

### PACKAGES

## ***************** TYPES ************************************

### BASIC TYPES

variable, constant

## TYPE CONVERSION

### ARRAYS - COLLECTION TYPES?

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
