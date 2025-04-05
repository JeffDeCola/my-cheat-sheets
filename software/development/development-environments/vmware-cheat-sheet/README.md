# VMWare CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_VMWare is a "virtualization" product, which allows you to
run an operating systems._

Table of Contents

* [VIRTUAL MACHINE (VM) vs CONTAINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vmware-cheat-sheet#virtual-machine-vm-vs-container)
* [INSTALL FREE VERSION OF ESXi](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vmware-cheat-sheet#install-free-version-of-esxi)

Documentation and Reference

* [Virtualbox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet)
* [Docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/builds-deployment-containers/docker-cheat-sheet)
* [github webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## VIRTUAL MACHINE (VM) vs CONTAINER

The following diagram shows the difference between Virtual Machines
and Containers,

![IMAGE - virtual-machine-vs-docker-container - IMAGE](../../../../docs/pics/software/development/virtual-machine-vs-docker-container.svg)

Virtual Machines

* Must use a Hypervisor emulated Virtual Hardware
* May or may not use a guest OS
* Takes a lot of system resources
* Takes up a lot of memory

Containers

* Uses a shared host OS
* You must use that OS
* Less resources and lightweight

## INSTALL FREE VERSION OF ESXi

Download and install from [virtualbox](https://customerconnect.vmware.com/en/evalcenter?p=free-esxi8).

Make sure you get the free license key.
