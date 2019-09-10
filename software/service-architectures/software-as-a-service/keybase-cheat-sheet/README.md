# KEYBASE CHEAT SHEET

`keybase` is secure groups, files and chat for everyone. It also has
encrypted private and group repos.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL

The easiest way to install is being going to the website
[here](https://keybase.io/)

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
