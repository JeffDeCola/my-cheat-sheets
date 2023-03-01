# VAGRANT CHEAT SHEET

`vagrant` _is a tool from HashiCorp that creates and configures (manages)
portable development environments._

tl;dr,

```bash
# VERSION
vagrant version
# RUN
vagrant up
vagrant ssh
ssh -i ~/.vagrant.d/insecure_private_key -p 2222 vagrant@127.0.0.1
# LIST/ADD/REMOVE VAGRANT BOXES
vagrant box list
vagrant box add --name "ubuntu/jeffs-ubuntu" --force ubuntu-box.box
vagrant box remove ubuntu/jeffs-ubuntu
```

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#overview)
* [HOW DOES IT WORK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#how-does-it-work)
* [WHAT IS A VAGRANT BOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#what-is-a-vagrant-box)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#install)
* [CHECK VERSION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#check-version)
* [FIND VAGRANT BOXES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#find-vagrant-boxes)
* [VAGRANTFILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#vagrantfile)
* [START IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#start-it)
* [SSH INTO IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#ssh-into-it)
* [CREATING A VAGRANT BOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#creating-a-vagrant-box)
* [SOME BASIC COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#some-basic-commands)
* [VAGRANT, DOCKER AND PACKER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/vagrant-cheat-sheet#vagrant-docker-and-packer)

Documentation and reference,

* [my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes)
* View a list of all my-cheat-sheets on my
  [github webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW

Put simple, the goal is to **create a common development environment**.

Vagrant provides the same production environment (_OS,
packages, users, configuration, etc..._) giving developers the
flexibility to use there own tools (_browsers, IDEs, editors, etc..._).

Vagrant can match the development environment to the
production environment. _"Well it worked on my computer"_
is a statement of the past.

Vagrant can run on providers such as,

* VirtualBox
* KVM
* Hyper-V
* Docker
* VMware
* AWS

![IMAGE - vagrant-overview - IMAGE](../../../../docs/pics/vagrant-overview.jpg)

## HOW DOES IT WORK

The process is pretty simple,

* STEP 1 - GET BASE BOX/IMAGE
  * Starts with a box/image (for a particular provider) that contains the
    base operating system.
* STEP 2 - CONFIGURE ON PROVIDER
  * Based on your configuration for your provider, fires up a VM/Container.
* STEP 3 PROVISION
  * Provisions it.

All of this information is neatly contained within the Vagrantfile (see below).

## WHAT IS A VAGRANT BOX

The vagrant box is basically all the info needed to
fire up a VM/container on a provider.  Vagrant packages it as such.

## INSTALL

Visit [vagrant downloads](https://www.vagrantup.com/downloads.html).

## CHECK VERSION

```bash
vagrant version
```

## FIND VAGRANT BOXES

Search for vagrant boxes at
[vagrant box search](https://app.vagrantup.com/boxes/search).

## VAGRANTFILE

Vagrants uses a declarative config file which describes your
software requirements, packages, OS configuration, users, etc..

A simple Vagrantfile for firing up a VirtualBox VM on Windows may look like,

```bash
Vagrant.configure("2") do |config|

  # STEP 1 - GET BASE BOX
  config.vm.box = "ubuntu/trusty64"
  config.vm.box_version = "latest"

  # STEP 2 - CONFIGURE ON PROVIDER (WINDOWS VIRTUALBOX)
  config.vm.provider "virtualbox" do |vb|
        vb.name = 'ubuntu-1604-virtualbox-vm'
        vb.gui = true
        vb.memory = 8192
        vb.cpus = 2
  end

  # STEP 3 - PROVISION
  config.vm.provision "shell", inline: "sudo apt-get install htop"

end
```

To see a few example goto my repo
[my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes).

## START IT

Simple type,

```bash
vagrant up
```

## SSH INTO IT

ssh into it,

```bash
vagrant ssh
```

Or you use vagrant insecure keys,

* Your machine - You use the vagrant insecure private key
  located on your machine  `~/.vagrant.d/insecure_private_key`.
* The vagrant box - Uses the default public key in `~/.ssh/authorized_keys`.

This will also add the box fingerprint on your machine in
`~/.ssh/known_host` for IP `127.0.0.1`.

```bash
ssh -i ~/.vagrant.d/insecure_private_key -p 2222 vagrant@127.0.0.1
```

## CREATING A VAGRANT BOX

Refer to my repo [my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes)
that has an example.

## SOME BASIC COMMANDS

List boxes on your machine,

```bash
vagrant box list
```

Add box on your machine,

```bash
vagrant box add --name "ubuntu/jeffs-ubuntu" --force ubuntu-box.box
```

Destroy box on your machine,

```bash
vagrant box remove ubuntu/jeffs-ubuntu
```

## VAGRANT, DOCKER AND PACKER

An illustration of vagrant, docker and packer tools,

![IMAGE -  vagrant docker packer - IMAGE](../../../../docs/pics/vagrant-docker-packer.jpg)
