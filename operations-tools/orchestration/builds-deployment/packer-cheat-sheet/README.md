# PACKER CHEAT SHEET

`packer` _is a tool for creating VM or container images (for multiple platforms)
from a single source configuration file._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## A VIRTUAL MACHINE VS CONTAINER

Virtual Machine:

* Must use a Hypervisor emulated Virtual Hardware
* Needs a guest OS
* Takes a lot of system resources

Container:

* Use a shared host OS
* You must use that OS
* Less Recources and lightweight

## INSTALL

Goto this website to install.  Must be 64-bit machine.

[https://www.packer.io/downloads.html](https://www.packer.io/downloads.html)

After install check version,

```bash
packer version
```

## SUPPORTS

Packer supports building images for Amazon, Google, VirtualBox, etc...

## TEMPLATE FILE EXAMPLE

`testcase.json` example,

???????

`Variables` are just that.

`Builders` are ???

`Provisioners` are ???

Validate template file,

```bash
packer validate example.json
```

## RUN EXAMPLE

```bash
packer build testcase.json
```