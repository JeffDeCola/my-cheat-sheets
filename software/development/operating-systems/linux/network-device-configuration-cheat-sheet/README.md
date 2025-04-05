# NETWORK DEVICE CONFIGURATION CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Network device configuration will help you manage
your network devices._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#overview)
* [SEE WHAT DEVICES/INTERFACE NAMES YOU HAVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#see-what-devicesinterface-names-you-have)
* [ifupdown METHOD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#ifupdown-method)
  * [CONFIGURE YOUR NETWORK DEVICE FOR IFUPDOWN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#configure-your-network-device-for-ifupdown)
  * [START/STOP networking.service](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#startstop-networkingservice)
* [netplan METHOD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#netplan-method)
  * [CONFIGURE YOUR NETWORK DEVICE FOR NETPLAN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#configure-your-network-device-for-netplan)
  * [netplan apply](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#netplan-apply)

## OVERVIEW

There are two methods we will go over,

* [ifupdown](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#ifupdown-method)
  (older method)
* [netplan](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet#netplan-method)
  (newer method)

## SEE WHAT DEVICES/INTERFACE NAMES YOU HAVE

See your network devices and their configurations,

```bash
ifconfig -a
```

The files are located here,

```bash
/sys/class/net
```

Note, that newer version of ubuntu have changed `eth0` / `eth1`
to interface names like `enp0s3`.

## ifupdown METHOD

An older method to configure network interfaces.

### CONFIGURE YOUR NETWORK DEVICE FOR IFUPDOWN

Edit this file,

```bash
sudo nano /etc/network/interfaces
```

For example, if you want to have a
Static IP of `192.168.100.5` on network device `eth1`,

```text
auto eth1
iface eth1 inet static
    address 192.168.100.5
    netmask 255.255.255.0
```

### START/STOP networking.service

The file,

```bash
/lib/systemd/system/networking.service
```

Restart/Status using `systemctl`,

```bash
sudo systemctl restart networking.service
systemctl status networking.service
```

Recheck your devices,

```bash
ifconfig -a
```

You should see your new static ip address.

For more information on services refer to my cheat sheet
[systemd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/services/systemd-cheat-sheet).

## netplan METHOD

A newer method to configure network interfaces.

### CONFIGURE YOUR NETWORK DEVICE FOR NETPLAN

Edit this .yaml file,

```bash
sudo nano /etc/netplan/NAME.yaml
```

For example, if you want to have a
Static IP of `192.168.100.5` on network device `enp0s8`,
add the following under `ethernets`,

```YAML
network:
 version: 2
 renderer: networkd
 ethernets:
        ...
   enp0s8:
     dhcp4: no
     dhcp6: no
     addresses: [192.168.100.5/24]
        ...
```

You can check it,

```bash
netplan --debug generate.
```

### netplan apply

```bash
sudo netplan apply
```

Recheck your devices,

```bash
ifconfig -a
```

You should see your new static ip address.
