# INSTALL AND CONNECT FLY ON LINUX

There are two ways to get fly, From your gui or using curl.

## FROM GUI

Open Concourse,

[http://192.168.100.4:8080](http://192.168.100.4:8080)

or like I configured it,

[http://192.168.100.6:8080](http://192.168.100.6:8080)

In bottom right download latest versions of fly.exe for Linux,

Place in your bin folder with a version (I
find this easier in the long run),

```bash
mkdir $HOME/bin
mv $HOME/Downloads/fly $HOME/bin/fly-x.x.x
```

## USING CURL

I always have issues with Ubuntu on Bash on Windows so this is easier for me,

```bash
curl -Lo flytemp.tgz https://github.com/concourse/concourse/releases/download/v6.7.1/fly-6.7.1-linux-amd64.tgz
chmod +x flytemp.tgz
tar -xvzf flytemp.tgz
rm flytemp.tgz
cp fly $HOME/bin/fly-6.7.1
```

Old,

```bash
curl -Lo flytemp https://github.com/concourse/concourse/releases/download/v3.2.1/fly_linux_amd64 && chmod +x flytemp && mv flytemp $HOME/bin/fly-3.2.1
```

## FINISH THE CONFIGURATION

Make a symbolic link to your version,

```bash
ln -s fly-x.x.x fly
```

Check version,

```bash
fly -version
```

Login fly to concourse,

```bash
fly -t ci login -c http://192.168.100.4:8080/
fly -t ci login -c http://192.168.100.6:8080/
```

To authenticate, you will handshake with a token.

The name and token are stored in `~/.flyrc`.
