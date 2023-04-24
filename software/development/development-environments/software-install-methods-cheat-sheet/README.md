# SOFTWARE INSTALL METHODS CHEAT SHEET

_A very high level view on various software/app install methods._

Table of Contents

* [SOURCE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#source)
* [BINARY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#binary)
* [INSTALLER / PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#installer--packages)
* [VIRTUAL MACHINE / CONTAINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/software-install-methods-cheat-sheet#virtual-machine--container)

## SOURCE

The purest form of the install. Download source code
(e.g. _go, python c, etc..._) and compile to a binary.

* PROS
  * The latest software compiled correctly on your machine
* CONS
  * Configuration and provisioning
  * Updating software requires a recompile

## BINARY

Download the binary (e.g. _something.exe_) that has been already
compiled for your machine. For example,
[I install golang](https://github.com/JeffDeCola/my-linux-shell-scripts/blob/master/software/go-install-new-version/go-install-new-version.sh)
this way on my machine.

* PROS
  * Pre-compiled for your OS
* CONS
  * May be an older version

## INSTALLER / PACKAGES

An installer/provisioning type program is
used to install your App.

Linux has many types of software that can be used
as installers (and I use this term very loosely)
since they are much more than just installation (e.g. _BOSH and
[ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet)_).

Linux systems use packages, that contains everything a particular
program needs to run. (e.g. _`apt-get` is used in ubuntu_).

* PROS
  * Can tailor install to your machine
* CONS
  * None

## VIRTUAL MACHINE / CONTAINER

Sometimes its easier just to grab a `virtual machine` or `container`
where someone has already done all, if not most of the work.

For example, [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet)
will allow you to quickly fire up a VM.
Or use [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/builds-deployment-containers/docker-cheat-sheet)
to get a container with the App you want.

* PROS
  * Easy to fire up and shut down
* CONS
  * Need VM/Container software installed on machine
  * Configuration and provisioning can be tricky
  * Not really a good long term solution
  * Find a new image for another version
