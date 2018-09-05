# INSTALL & CONFIGURE

OK, lets get go...ing.  Yep, that just happened.

## INSTALL

Instructions for binary and source installs.

[Installs](https://golang.org/doc/install) for windows, linux or mac.

I would not install from a package.

## CONFIGURE

### LINUX

I place this is `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

### BASH ON UBUNTU ON WINDOWS

I place this is `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=/mnt/c/Users/<WINDOWSNAME>/home/<USERNAME>
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

The trick is to have a different project/working directory as
shown with my `$GOPATH`. Do not work in the home directory.

### WINDOWS

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
