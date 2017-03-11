# CONCOURSE CHEAT SHEET

`concourse` _is a pipelined CI (Continuous Integration) for
Software Builds._

Concourse's main goal is to run tasks.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

[My Concourse task examples](https://jeffdecola.github.io/my-concourse-ci-tasks/)

## INSTALL AND RUN CONCOURSE ON VIRTUALBOX

Get vargrant for windows

https://www.vagrantup.com

Get the Concourse Vagrant files. 

To run, from Windows cmd prompt type,

```bash
vagrant up
```

To upgrade (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant destroy                             # remove the old Vagrant box
vagrant up 
```

Reinstal the new fly.exe for Windows and Linux.


## INSTALL FLY ON WINDOWS

Open concourse,

http://192.168.100.4:8080/

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

## INSTALL FLY ON LINUX

Open Concourse,

http://192.168.100.4:8080/

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

## BASIC STRUCTURE OF CONCOURSE

See a more detailed example at
[my-concourse-ci-tasks](https://jeffdecola.github.io/my-concourse-ci-tasks/).

The following diagram illustrates compase running a task called seach-and-replace.

* `pipiline.yml` Grabs a repo from github
* `config.yml` Configures task
    * Grabs golang docker image 
    * Sets up inputs/outputs into task container
* `task.sh` does the task

![IMAGE - concourse cheat sheet structure - IMAGE](docs/pics/Concourse-structure.jpg)
