# INSTALL FLY ON LINUX

Open Concourse,

[http://192.168.100.4:8080](http://192.168.100.4:8080)

In bottom right download latest verions of fly.exe for Linux,

Place in,

```bash
mkdir $HOME/bin
install $HOME/Downloads/fly $HOME/bin
```

Login fly to Concourse,

```bash
fly -t ci login -c http:// 192.168.100.4:8080/
```

Check version,

```bash
fly -version
```