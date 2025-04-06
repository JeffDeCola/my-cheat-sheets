# KEYBASE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** DEPRECIATED ***
```

_Keybase is secure groups, files and chat for everyone. It also has
encrypted private and group repos._

Table of Contents

* [WHY DO YOU NEED THIS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/communications/keybase-cheat-sheet#why-do-you-need-this)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/communications/keybase-cheat-sheet#install)
  * [LINUX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/communications/keybase-cheat-sheet#linux)
  * [GO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/software-as-a-service/communications/keybase-cheat-sheet#go)

## WHY DO YOU NEED THIS

I use it for the `encrypted repos` hosted at keybase.io.
It is not like github where you can goto the website and do PR, etc.
Its just a place to keep a repo that is encrypted end-to-end.

Hence this program is connected to your keybase account
and will encrypt your data before it is sent to the repo.

Refer to your keybase app to create a repo and then just use normal git.

## INSTALL

The easiest way to install is being going to the website
[here](https://keybase.io/).

### LINUX

```bash
curl --remote-name https://prerelease.keybase.io/keybase_amd64.deb
sudo apt install ./keybase_amd64.deb
run_keybase
```

Check,

```bash
keybase version
```

### GO

```bash
go get -u -v github.com/keybase/client/go/keybase
go install -tags production github.com/keybase/client/go/keybase
```

Run service,

```bash
keybase service &
```

Run client,

```bash
keybase login
```

Check,

```bash
keybase version
```

Install kbfsfuse,

```bash
cd $GOPATH/src/github.com/keybase/client/go/kbfs/kbfsfuse
go install
```

Install git-remote-keybase,

```bash
cd $GOPATH/src/github.com/keybase/client/go/kbfs/kbfsgit/git-remote-keybase
go install
```
