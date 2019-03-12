# GO TOOLS CHEAT SHEET

```
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`go-tools` _are many tools that can be used with go._

The tools are organized alphabetically as follows:

* [gotests](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-tools-cheat-sheet#gotests)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## GOTESTS

Automatically generates tests.

Install or update,

```bash
go get github.com/cweill/gotests/...
```

This will place binary in `$GOPATH/bin`.

Generate test file `filename_test.go`,

```bash
gotests -w -all filename.go
```

Run,

```bash
go test -v -cover .
```
