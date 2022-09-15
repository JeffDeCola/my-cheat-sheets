# VIRTUALBOX CHEAT SHEET

_VirtualBox is a "virtualization" product, which allows you to
run an operating system (guest) on top of your existing
operating system (host)._

Table of Contents

* [VIRTUAL MACHINE (VM) vs DOCKER CONTAINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#virtual-machine-vm-vs-docker-container)
* [INSTALL ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#install-on-windows)
* [INSTALL GUEST ADDITIONS FROM COMMAND LINE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#install-guest-additions-from-command-line)
* [SHARED CLIPBOARD & DRAG AND DROP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#shared-clipboard--drag-and-drop)
* [CONFIGURE STATIC IP IN UBUNTU RUNNING ON VIRTUALBOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#configure-static-ip-in-ubuntu-running-on-virtualbox)
* [ALLOW HOST ACCESS TO VM (HOST-ONLY)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#allow-host-access-to-vm-host-only)
* [ALLOW LOCAL NETWORK ACCESS TO VM (BRIDGE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#allow-local-network-access-to-vm-bridge)
* [INCREASE YOUR VIDEO MEMORY TO 256MB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#increase-your-video-memory-to-256mb)
* [INCREASE THE SIZE OF YOUR VIRTUAL DISK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#increase-the-size-of-your-virtual-disk)
* [SSH ONTO UBUNTU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#ssh-onto-ubuntu)
* [FTP ONTO UBUNTU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#ftp-onto-ubuntu)
* [ADD A PRINTER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/virtualbox-cheat-sheet#add-a-printer)

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install ubuntu with GNOME desktop](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md)
* [Install Windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows-11.md)

* View a list of all my-cheat-sheets on my
  [github webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## VIRTUAL MACHINE (VM) vs DOCKER CONTAINER

The following diagram shows the difference between a Virtual Machine
and a Docker Container,

![IMAGE - virtual-machine-vs-docker-container - IMAGE](../../../../docs/pics/virtual-machine-vs-docker-container.jpg)

Virtual Machine

* Must use a Hypervisor emulated Virtual Hardware
* Needs a guest OS
* Takes a lot of system resources
* Takes up a lot of memory

Container

* Uses a shared host OS
* You must use that OS
* Less resources and lightweight

## INSTALL ON WINDOWS

Download and install from [virtualbox](https://www.virtualbox.org/).

## INSTALL GUEST ADDITIONS FROM COMMAND LINE

Get the latest `VBoxGuestAdditions.iso` image file from
[virtualbox](http://download.virtualbox.org/virtualbox).

You can download from,

```bash
http://download.virtualbox.org/virtualbox/6.0.10/VBoxGuestAdditions_6.0.10.iso
```

Sometimes when you update Virtualbox it automatically puts the updated disk
image here:
`'C:\Program Files\Oracle\VirtualBox\VBoxGuestAdditions.iso'`.
But not all the time.

### USING SHARED FOLDER

I like installing it via the shared folder.
Place you VBoxGuestAdditions.iso' in your shared folder.

```bash
sudo mkdir /media/cdrom
sudo mkdir /media/temp
sudo cp /media/<SHAREDFOLDER>/VBoxGuestAdditions.iso /media/temp
sudo chmod 775 /media/temp/VBoxGuestAdditions.iso

sudo mount -o loop /media/temp/VBoxGuestAdditions.iso /media/cdrom
sudo /media/cdrom/VBoxLinuxAdditions.run
reboot
```

After reboot, check that it installed,

```bash
ls /opt
```

You may need to add yourself to the vboxsf group.

First check,

```bash
groups
```

Add the group,

```bash
sudo usermod --append --groups vboxsf USERNAME
reboot
```

### USING THE MENU

You could do this. Start your VM.  Go to the menu item
`Devices -> Insert Guest Additions CD image`
to mount the ISO image.  But this does not always work.

Now from the terminal, run the following commands,

```bash
sudo mkdir -p /media/cdrom
sudo mount /dev/cdrom /media/cdrom
sudo apt install -y dkms build-essential linux-headers-generic \
  linux-headers-$(uname -r)
sudo /media/cdrom/VBoxLinuxAdditions.run
reboot
```

## SHARED CLIPBOARD & DRAG AND DROP

The following is required to use these features,

* Shared folder enabled
* Guest Additions installed
* Desktop like GNOME

## CONFIGURE STATIC IP IN UBUNTU RUNNING ON VIRTUALBOX

See your network devices and their configurations,

```bash
ifconfig -a
```

Note, that newer versions of ubuntu have changed `eth0` / `eth1`
to interface names like `enp0s3`.

Configure your network configuration file
 `/etc/network/interfaces` for a static IP address `192.168.100.5`,

```text
auto enp0s3
iface enp0s3 inet static
    address 192.168.100.5
    netmask 255.255.255.0
```

Restart/Status `networking.service` using `systemctl`
or reboot your machine,

```bash
sudo systemctl restart networking.service
systemctl status networking.service
```

More info about `systemctl` at my cheat sheet
[systemd systemctl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/systemd-systemctl-cheat-sheet).

Recheck your devices,

```bash
ifconfig -a
```

You should see your new static ip address.

For more information about configuring network devices, goto my cheat-sheet
[network-device-configuration](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet).

## ALLOW HOST ACCESS TO VM (HOST-ONLY)

Create a VirtualBox Host-Only Ethernet Adapter and configure
the VM,

```txt
Settings -> Network - Adapter -> Host-Only Adapter
```

Also set your machine with a static IP.

Refer to my cheat sheet on
[network device configuration](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet)
on how to add a static IP.

## ALLOW LOCAL NETWORK ACCESS TO VM (BRIDGE)

This is actually quite easy.

Instead of NATS, you use a bridged adapter,

```txt
Settings -> Network - Adapter -> Bridged Adapter
```

Your router (on the host machine's network) will choose an
IP for you.  You can go into that router to save the
IP based on the mac address if you would like.

Also, do not set your machine with a static IP.
Let your router assign one.

## INCREASE YOUR VIDEO MEMORY TO 256MB

```bash
VBoxManage modifyvm "Name of VM" --vram 256
"C:\Program Files\Oracle\VirtualBox\VBoxManage.exe"  modifyvm "Name of VM" --vram 256
```

## INCREASE THE SIZE OF YOUR VIRTUAL DISK

Let's say you run `df -h` and you're running out of room.

The increase is in MB, so if you want to increase 10GB or ~10,000MB
you pick a number that is a power of 2. 2 to the power of 13 is 8192.
2 to the power of 14 is 16384.  Lets use the latter.

Step 1 - Make a backup (clone)

Step 2 - Create new virtual Machine that is bigger

Step 3 - Find UUID of your old and new VM

```bash
VBoxManage list hdds
"C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" list hdds
```

Step 4 - run VBoxManage clonemedium

```bash
VBoxManage clonemedium <source-guid> <destinatin-guid> --existing
"C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" clonemedium b1e21e90-xxx 391354e8-xxx --existing
```

Step 5 - Download gparted-live*.iso

Insert .iso into the new virtual machine and reboot it.
Grow your partition and save changes.
You may have to delete partitions (swap and /dev/sd2) to grow `/dev/sda1`.

## SSH ONTO UBUNTU

On Virtual box add port forwarding for the ubuntu Virtual Machine.
Network settings and click the Port Forwarding button. Add a new Rule.
From virtual box configure in settings,

* Host port 3022, guest port 22, name ssh, other left blank.

From ubuntu start a ssh server,

```bash
sudo apt-get install openssh-server
```

To use just ssh to [127.0.0.1:3022](127.0.0.1:3022).

## FTP ONTO UBUNTU

```bash
sudo apt-get install vsftpd
nano /etc/vsftpd.conf
```

Just use sftp on port 3022.

## ADD A PRINTER

Pain in the butt,

Step 1 - Windows

First setup a shared network printer on Windows.
Must also use a private network.

Step 2 - Get some software loaded on ubuntu

```bash
sudo apt-get install lsb
sudo apt-get update
sudo apt-get install python3-smbc
sudo apt-get install smbclient
```

Step 3 - Get printer driver from printer website. Most likely a .deb file

Step 4 - Then you need to find the printer

You could search for network printer on `192.168.1.x` or just type it in.

smb://192.168.1.115/SHARE-EPSON-XP-630-Series

Also need Username, Workgroup and password.
