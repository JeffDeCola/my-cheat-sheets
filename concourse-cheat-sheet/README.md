# CONCOURSE CHEAT SHEET

`concourse` _is a pipelined CI (Continuous Integration) for
Software Builds._

Concourse's main goal is to run tasks.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL AND RUN CONCOURSE ON VIRTUALBOX

Get vargrant for windows

https://www.vagrantup.com

Get the concourse vagrant files. 

To run, from windows cmd prompt type,

```bash
vagrant up
```

To upgrade (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant destroy                             # remove the old Vagrant box
vagrant up 
```

Reinstal the new fly.exe for windows and linux.


## INSTALL FLY ON WINDOWS

Open concourse,

http://192.168.100.4:8080/

In bottom right download latest verions of fly.exe for linux,

Place in,

```bash
C:\Program Files (x86)\Concourse\fly.exe
```

Login fly to concourse,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -t ci login -c http://192.168.100.4:8080/
```

Check version,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -version
```

## INSTALL FLY ON LINUX

Open concourse,

http://192.168.100.4:8080/

In bottom right download latest verions of fly.exe for linux,

Place in,

```bash
mkdir $HOME/bin
install $HOME/Downloads/fly $HOME/bin
```

Login fly to concourse,

```bash
fly -t ci login -c http:// 192.168.100.4:8080/
```

Check version,

```bash
fly -version
```
