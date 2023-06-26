# INSTALL & CONFIGURE

OK, lets get go...ing.  Yep, that just happened.

tl;dr

```bash
# INSTALL
FileName='go1.19.2.linux-amd64.tar.gz'
FileName='go1.19.2.darwin-amd64.tar.gz'
FileName='go1.19.2.linux-arm64.tar.gz'
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo mv go /usr/local
rm $FileName
# CONFIGURE .bashrc
# JEFF ADDED FOR GO
export PATH=$PATH:/usr/local/go/bin
export GOBIN=/home/jeff/go/bin
export PATH=$PATH:$GOBIN
# CHECK
go version
# INSTALL GO TOOLS
# Use vs code go extension
```

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install)
  * [INSTALL WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-windows)
  * [INSTALL macOS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-macos)
  * [INSTALL LINUX](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-linux)
  * [INSTALL RASPBERRY PI OS (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-raspberry-pi-os-debian)
* [UPGRADE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#upgrade)
  * [UPGRADE WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#upgrade-windows)
  * [UPGRADE macOS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#upgrade-macos)
  * [UPGRADE LINUX & RASPBERRY PI OS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#upgrade-linux--raspberry-pi-os)
* [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure)
  * [CONFIGURE WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-windows)
  * [CONFIGURE LINUX & RASPBERRY PI OS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-linux--raspberry-pi-os)
  * [CONFIGURE BASH ON UBUNTU ON WINDOWS (WSL2)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#configure-bash-on-ubuntu-on-windows-wsl2)
* [INSTALL GO TOOLS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-go-tools)

## INSTALL

[Binary and source installs](https://golang.org/doc/install) are
located here for windows, linux or mac. I would not install from a package.

### INSTALL WINDOWS

Just use the installer found
[Binary and source installs](https://golang.org/doc/install).

### INSTALL macOS

For macOS just use a different FileName,

```bash
FileName='go1.19.2.darwin-amd64.tar.gz'
```

### INSTALL LINUX

Go lives in `/usr/local/go`.

To install there, the tarball format is,

```bash
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Hence,

```bash
FileName='go1.19.2.linux-amd64.tar.gz'
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo mv go /usr/local
rm $FileName
```

### INSTALL RASPBERRY PI OS (DEBIAN)

For Raspberry Pi just use a different FileName,

```bash
FileName='go1.19.2.linux-arm64.tar.gz'
```

## UPGRADE

This is easy.

### UPGRADE WINDOWS

Just use the installer found
[Binary and source installs](https://golang.org/doc/install).

### UPGRADE macOS

Remove your old version,

```bash
which go
sudo rm -rf /usr/local/go
```

Follow the steps above.

### UPGRADE LINUX & RASPBERRY PI OS

Remove your old version from `/usr/local/go` and replace.

```bash
which go
sudo rm -rf /usr/local/go
```

Follow the install steps above.

## CONFIGURE

Let's set up the go paths.

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

Old way I used to do it and depreciated,

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

Old way I used to do it and depreciated,

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
