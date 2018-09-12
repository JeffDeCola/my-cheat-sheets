# SOFTWARE INSTALL METHODS CHEAT SHEET

`software install methods` _is a very high level view on
how to install software/Apps._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL METHODS

There are four basic methods to install software, each having
their own strengths and weaknesses,

* [SOURCE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#source)
* [BINARY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#binary)
* [INSTALLER / PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#installer--packages)
* [VIRTUAL MACHINE / CONTAINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#virtual-machines--containers)

## SOURCE

The purest form of the install. Download source code
(e.g. _go, python c, etc..._) and compile.

There is most likely a lot of provisioning and configuration involved.
To update software, you need to grab the new source.

## BINARY

Download the binary (e.g. _something.exe_) that has been already
compiled for your machine.

## INSTALLER / PACKAGES

An installer/provisioning type program is
used to install your App.

Linux has many types of software that can be used
as installers (and I use this term very loosely)
since they are much more then just installation (e.g. _BOSH and
[ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/configuration-management/ansible-cheat-sheet)_).

Linux systems use packages, that contains everything a particular
program needs to run. (e.g. _`apt-get` is used in ubuntu_).

## VIRTUAL MACHINE / CONTAINER

Sometimes its easier just to grab a virtual machine or container
where someone has already done all, if not most of the work.

This solution is good for a quick startup to get yourself going, but
not really good for a long term solution since software updates,
provisioning can be really tricky.

For example, [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet)
will allow you to quickly fire up a VM.
Or use [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/builds-deployment-containers/docker-cheat-sheet)
to get a container with the App you want.
