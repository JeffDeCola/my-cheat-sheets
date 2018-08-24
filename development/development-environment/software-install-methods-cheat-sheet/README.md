# SOFTWARE INSTALL METHODS CHEAT SHEET

`software install methods` _is a very high level view on
software/app installs._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL METHODS

    * SOURCE
    * BINARY
    * INSTALLER / PACKAGES
    * VIRTUAL MACHINE / CONTAINER

Each have their own strengths and drawbacks.

## SOURCE

The purest form of the install. Download source code
(go, python c, etc...) and compile.

There is most likely a lot of provisioning and configuraiton involved.
To update software, you need to grab the new source.

## BINARY

Download the binary (e.g. something.exe) that has been already
compiled for your machine.

## INSTALLER / PACKAGES

An installer/provisioning type program is
used to install your app.

Linux has many types of software that can be used
as installers (and I use this term very loosly)
since they are much more then just installation.

For example, BOSH and
[ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/configuration-management/ansible-cheat-sheet).

Linux systems use packages, that contains everything a particular
program needs to run.

For example, `apt-get` is used in ubuntu.

## VIRTUAL MACHINES / CONTAINERS

Sometimes its easier just to grab a Virtual Machine or container
which someone who has already done all the work.

For example, [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environment/vagrant-cheat-sheet)
will allow you to quickly fire up a VM.
Or use [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/builds-deployment-containers/docker-cheat-sheet)
to get a container with the app you want.
