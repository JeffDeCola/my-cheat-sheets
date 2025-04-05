# INSTALL & CONFIGURE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_OK, lets get go...ing.  Yep, that just happened._

tl;dr

```bash
# INSTAL/UPGRADE LINUX
FileName='go1.20.5.linux-amd64.tar.gz'
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo rm -rf /usr/local/go
sudo mv go /usr/local
rm $FileName
# CONFIGURE .bashrc
export PATH=$PATH:/usr/local/go/bin
export GOBIN=/home/jeff/go/bin
export PATH=$PATH:$GOBIN
# CHECK
go version
```

Table of Contents

* [INSTALL/UPGRADE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#installupgrade)
  * [LINUX](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#linux)
  * [RASPBERRY PI OS (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#raspberry-pi-os-debian)
  * [macOS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#macos)
  * [WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#windows)
* [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure)
  * [CONFIGURE LINUX & RASPBERRY PI OS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-linux--raspberry-pi-os)
  * [CONFIGURE BASH ON UBUNTU ON WINDOWS (WSL2)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-bash-on-ubuntu-on-windows-wsl2)
  * [CONFIGURE WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-windows)
* [INSTALL GO TOOLS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-go-tools)

Documentation and Reference

* [go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/go-cheat-sheet#go-cheat-sheet)
  main page
* [go-install-new-version](https://github.com/JeffDeCola/my-linux-shell-scripts/tree/master/go/go-install-new-version)
  is an install script I wrote

## INSTALL/UPGRADE

[Binary and source installs](https://golang.org/doc/install) are
located here for windows, linux or mac. I would not install from a package.

I wrote a script
[go-install-new-version](https://github.com/JeffDeCola/my-linux-shell-scripts/tree/master/go/go-install-new-version)
to do this automatically.

### LINUX

Go lives in `/usr/local/go`.

To install there, the tarball format is,

```bash
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Hence,

```bash
FileName='go1.20.5.linux-amd64.tar.gz'
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo rm -rf /usr/local/go
sudo mv go /usr/local
rm $FileName
```

### RASPBERRY PI OS (DEBIAN)

For Raspberry Pi just use a different FileName,

```bash
FileName='go1.20.5.linux-arm64.tar.gz'
```

### macOS

```bash
FileName='go1.20.5.darwin-amd64.pkg'
wget https://storage.googleapis.com/golang/$FileName
sudo installer -pkg $FileName -target /usr/local
rm $FileName
```

### WINDOWS

Just use the installer found
[Binary and source installs](https://golang.org/doc/install).

## CONFIGURE

Let's set up the go paths.

### CONFIGURE LINUX & RASPBERRY PI OS

I place the following in `.bashrc`,

New way without GOROOT and GOPATH (depreciated).
I still explicitly define GOBIN even though it's the default,

```bash
# JEFF ADDED FOR GO
export PATH=$PATH:/usr/local/go/bin
export GOBIN=/home/jeff/go/bin
export PATH=$PATH:$GOBIN
```

Old way I used to do it (depreciated),

```bash
# JEFF ADDED FOR GO
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
```

### CONFIGURE BASH ON UBUNTU ON WINDOWS (WSL2)

I place the following in `.bashrc`,

New way without GOROOT and GOPATH (depreciated).
I still explicitly define GOBIN even though it's the default,

```bash
# JEFF ADDED FOR GO
export PATH=$PATH:/usr/local/go/bin
export GOBIN=/home/jeff/go/bin
export PATH=$PATH:$GOBIN
```

Old way I used to do it (depreciated),

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

### CONFIGURE WINDOWS

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

## INSTALL GO TOOLS

I use the go extension in vs code to install go tools.

When you get the source, you should get the go tools
`gofmt` and `godocs`.

They usually live in the following locations,

* WINDOWS
  * C:\Users\Jeff\go\bin
* macOS and Ubuntu
  * /usr/local/go/bin
  * /home/jeff/go/bin
