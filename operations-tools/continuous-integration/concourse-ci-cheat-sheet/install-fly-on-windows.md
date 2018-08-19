# INSTALL FLY ON WINDOWS

Open concourse,

[http://192.168.100.4:8080](http://192.168.100.4:8080)

In bottom right download latest verions of fly.exe for Linux,

Place in,

```bash
C:\Program Files (x86)\Concourse\fly.exe
```

Login fly to Concourse,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -t ci login -c http://192.168.100.4:8080/
```

Check version,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -version
```
