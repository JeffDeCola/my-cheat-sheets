# GO TOOLS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Some tools that can be used with go._

Table of Contents

* [GOTESTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-tools-cheat-sheet#gotests)

## GOTESTS

`gotests` automatically generates unit tests.

Use `go get` to install and update,

```bash
go get -v -u github.com/cweill/gotests/...
```

This will place binary in `$GOPATH/bin`.

Generate test file `filename_test.go`,

```bash
gotests -w -all filename.go
```

* `-w` writes to test file
* `-all` generate test for all functions and methods

Run,

```bash
go test -v -cover .
```
