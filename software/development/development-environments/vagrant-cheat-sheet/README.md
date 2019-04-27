# VAGRANT CHEAT SHEET

`vagrant` _is a tool from HashiCorp that creates and configures (manages) portable development environments._

Check out my repo [my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Put simple, the goal is to create a common development environment.

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

## HOW DOES WORK?

The process is pretty simple,

* STEP 1 - GET BASE BOX/IMAGE
  * Starts with a box/image (for a particular provider) that contains the base operating system.
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

## CREATING A VAGRANT BOX

Refer to my repo [my-vagrant-boxes](https://github.com/JeffDeCola/my-vagrant-boxes)
that has an example.

##  SOME BASIC COMMANDS

List boxes on your machine,

```bash
vagrant box list
```
