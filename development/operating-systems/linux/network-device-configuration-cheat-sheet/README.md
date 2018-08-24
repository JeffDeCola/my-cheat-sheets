# NETWORK DEVICE CONFIGURATION CHEAT SHEET

`network device configuration` _will help you manage
your network devices._

There are two methods we will go over,

* ifupdown
* netplan

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

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

## ifupdown METHOD - CONFIGURE YOUR NETWORK DEVICE

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

## ifupdown METHOD - START/STOP networking.service

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
[systemd systemctl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/operating-systems/linux/systemd-systemctl-cheat-sheet).

## netplan METHOD - CONFIGURE YOUR NETWORK DEVICE

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

you can check it,

```bash
netplan --debug generate.
```

## netplan METHOD - netplan apply

```bash
sudo netplan apply
```

Recheck your devices,

```bash
ifconfig -a
```

You should see your new static ip address.
