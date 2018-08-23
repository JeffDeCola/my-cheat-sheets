# NETWORK DEVICE CONFIGURATION CHEAT SHEET

`network device configuration` _will help you manage
your network devices._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SEE YOUR DEVICES

See your network devices and their configurations,

```bash
ifconfig -a
```

The fiels are located here,

```bash
/sys/class/net
```

## CONFIGURE YOUR NETWORK DEVICE

Edit this file,

```bash
/etc/network/interfaces
```

For example, if you want to have a 
Static IP of `192.168.100.5` on network device `eth1`,


```text
auto eth1
iface eth1 inet static
    address 192.168.100.5
    netmask 255.255.255.0
```

## START/STOP networking.service

The file,

```bash
/lib/systemd/system/networking.service
```

Restart/Status using systemclt,

```bash
sudo systemctl restart networking
systemctl status networking.service
```

Recheck your devices,

```bash
ifconfig -a
```

For more information on services refer to my cheat sheet
[systemd systemctl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/operating-systems/linux/systemd-systemctl-cheat-sheet)


## OTHER COMMANDS

tbd



