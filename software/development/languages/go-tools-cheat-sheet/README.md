# GO TOOLS CHEAT SHEET

`go-tools` _are tools that can be used with go._

The tools are organized alphabetically as follows:

* [gotests](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-tools-cheat-sheet#gotests)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## GOTESTS

Automatically generates tests.

Use go get to install and update,

```bash
go get github.com/cweill/gotests/...
```

This will place binary in `$GOPATH/bin`.

Generate test file `filename_test.go`,

```bash
gotests -w -all filename.go
```

* -w writes to test file
* -all generate test for all functions and methods.

Run,

```bash
go test -v -cover .
```
