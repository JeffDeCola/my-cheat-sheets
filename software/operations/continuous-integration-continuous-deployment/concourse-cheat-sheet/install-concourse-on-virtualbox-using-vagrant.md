# INSTALL CONCOURSE ON VIRTUALBOX USING VAGRANT

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
**** OLD - ARCHIVED INFORMATION ****
```

_This method has been depreciated at concourse, but I'm sure you can
find vagrant files that others have made._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-on-virtualbox-using-vagrant.md#install)
* [BRIDGE CONCOURSE NETWORK TO HOME LAN (198.168.1.x)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-on-virtualbox-using-vagrant.md#bridge-concourse-network-to-home-lan-1981681x)

Documentation and Reference

* [concourse](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/README.md)

## INSTALL

Get vagrant for Windows [here](https://www.vagrantup.com).

Search for a concourse vagrant file
[here](https://app.vagrantup.com/boxes/search).
This [one](https://app.vagrantup.com/concourse/boxes/lite) works,
but is in older version of concourse.

To run your vagrant box, enter vagrant directory, and from the
Windows cmd prompt type,

```bash
vagrant up
```

To upgrade vagrant box (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant box prune                           # remove the old Vagrant boxes
vagrant up
```

Reinstall the new `fly.exe`.

## BRIDGE CONCOURSE NETWORK TO HOME LAN (198.168.1.x)

Add in Vagrantfile,

```txt
config.vm.network "public_network", bridge: 'Intel(R) I210 Gigabit
Network Connection', use_dhcp_assigned_default_route: true
```

Or you could make this change in Virtualbox.

Instead of NATS, you use a bridged adapter,

```txt
Settings -> Network - Adapter -> Bridged Adapter
```

Either way, your router (on the host machine's network) will choose an
IP for you.  You can go into that router to save the
IP based on the mac address if you would like.

Also, do not set your machine with a static IP.
Let your router assign one.
