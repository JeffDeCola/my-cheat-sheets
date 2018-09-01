# VIRTUALBOX CHEAT SHEET

`virtualbox` _a "virtualization" product, allows you to
run an operating system (guest) on top of your existing
operating system (host)_.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL ON WINDOWS

Download and install from [virtualbox](https://www.virtualbox.org/).

## INSTALL GUEST ADDITIONS FROM COMMAND LINE

Get the latest `VBoxGuestAdditions.iso` image file from
[virtualbox](http://download.virtualbox.org/virtualbox).

Start your VM.  Go to the menu item
`Devices -> Insert Guest Additions CD image`
to mount the ISO image.

From the terminal, run the following commands,

```bash
sudo mkdir -p /media/cdrom
sudo mount /dev/cdrom /media/cdrom
sudo apt install -y dkms build-essential linux-headers-generic \
  linux-headers-$(uname -r)
sudo sh /media/cdrom/VBoxLinuxAdditions.run
reboot
```

After reboot, check that it installed,

```bash
ls /opt
```

You may need to add yourself to the vboxsf group,

First check,

```bash
groups
```

To add,

```bash
sudo usermod --append --groups vboxsf USERNAME
reboot
```

## DRAG AND DROP

Must have a shared folder.

## CONFIGURE STATIC IP IN UBUNTU RUNNING ON VIRTUALBOX

See your network devices and their configurations,

```bash
ifconfig -a
```

Note, that newer version of ubuntu have changed `eth0` / `eth1`
to interface names like `enp0s3`.

Configure your network configuration file
 `/etc/network/interfaces` for a static IP address `192.168.100.5`,

```text
auto enp0s3
iface enp0s3 inet static
    address 192.168.100.5
    netmask 255.255.255.0
```

Restart/Status `networking.service` using `systemclt`
or reboot your machine,

```bash
sudo systemctl restart networking.service
systemctl status networking.service
```

Recheck your devices,

```bash
ifconfig -a
```

You should see your new static ip address.

For more information goto [network device configuration](),
