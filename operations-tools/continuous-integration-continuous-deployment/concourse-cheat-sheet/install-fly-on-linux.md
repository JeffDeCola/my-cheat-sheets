# INSTALL AND CONNECT FLY ON LINUX

Open Concourse,

[http://192.168.100.6:8080](http://192.168.100.6:8080)

In bottom right download latest versions of fly.exe for Linux,

Place in your bin folder with a version (I
find this easier in the long run),

```bash
mkdir $HOME/bin
mv $HOME/Downloads/fly $HOME/bin/fly-x.x.x
```

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
fly -t ci login -c http://192.168.100.6:8080/
```

To authenticate, you will handshake with a token.

The name and token are stored in `~/.flyrc`.
