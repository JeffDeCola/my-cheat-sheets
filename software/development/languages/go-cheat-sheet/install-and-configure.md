# INSTALL & CONFIGURE

OK, lets get go...ing.  Yep, that just happened.

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install)
* [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-paths)
* [CHECK](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#check-your-install)
* [INSTALL GO TOOLS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-go-tools)

## INSTALL

[Binary and source installs](https://golang.org/doc/install) are
located there for windows, linux or mac.

I would not install from a package.

### LINUX INSTALL

Let's install `ver 1.12.7`  to `/usr/local`,

```bash
FileName=`go1.12.7.linux-amd64.tar.gz`
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo mv go /usr/local
rm $FileName
```

## CONFIGURE

Let's set up the paths.

### LINUX

I place the following in `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

You probably don't need the GOBIN path, but I
use VS Code and it just makes it clear.

### BASH ON UBUNTU ON WINDOWS (Windows System for Linux)

I place the following in `.bashrc`,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=/mnt/c/Users/<WINDOWSNAME>/home/<USERNAME>
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

For Ubuntu 14.04 and 16.04 (if upgraded),

```txt
C:\Users\<WindowsNAME>\AppData\Local\lxss\home\<bashusername>\.bashrc
```

For Ubuntu 18.04 from Windows Store,

```txt
C:\Users\<WindowsNAME>\AppData\Local\Packages\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\<bashusername>\.bashrc
```

The trick is to have your project/working directory
not in your home directory as shown with my `$GOPATH`.
This is because Windows and linux do not play well together.

### WINDOWS

Do not get confused, this is not being installed on WSL (Windows System for Linux),
this is being installed on Windows OS. The installation should automatically
set the Windows environment variables as follows,

```text
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```

On a side note, my cheat sheet
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
explains `how to setup VS Code on Windows with the Go Extensions
and a bash terminal`.  Say that ten times fast.

## CHECK

Check your go paths and version,

```bash
go env
go version
```

## INSTALL GO TOOLS

When you get the source, you should get the go tools
`gofmt` and `godocs`.

Go tools should be in GOPATH or GOROOT.

They usually live in the following locations,

* WINDOWS
  * C:\Users\Jeff\go\bin & C:\Go\bin
* macOS and Ubuntu
  * ~/bin & /usr/bin & usr/local/go/bin

A few other tools may be useful,

```bash
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u golang.org/x/lint/golint
```
