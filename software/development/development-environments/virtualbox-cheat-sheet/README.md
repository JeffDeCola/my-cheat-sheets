# VIRTUALBOX CHEAT SHEET

`virtualbox` _a "virtualization" product, allows you to
run an operating system (guest) on top of your existing
operating system (host).  You can think of it as a sandbox of an OS
running on your computer_.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## VIRTUAL MACHINE (VM) vs DOCKER CONTAINER

The following diagram shows the difference between a Virtual Machine
and a Docker Container.

![IMAGE - virtual-machine-vs-docker-container - IMAGE](../../../docs/pics/virtual-machine-vs-docker-container.jpg)

Virtual Machine:

* Must use a Hypervisor emulated Virtual Hardware.
* Needs a guest OS.
* Takes a lot of system resources.
* Takes up a lot of memory.

Container:

* Uses a shared host OS.
* You must use that OS.
* Less Resources and lightweight.

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
sudo /media/cdrom/VBoxLinuxAdditions.run
reboot
```

You could also do it via the shared folder,

When you update Virtualbox it automatically puts the updated disk image here:
'C:\Program Files\Oracle\VirtualBox\VBoxGuestAdditions.iso'

Copy 'C:\Program Files\Oracle\VirtualBox\VBoxGuestAdditions.iso' to shared
VM folder

```bash
sudo mkdir /media/cdrom
sudo mkdir /media/temp
sudo cp /media/<sharedfolder>/VBoxGuestAdditions.iso /media/temp
sudo chmod 775 /media/temp/VBoxGuestAdditions.iso

sudo mount -o loop /media/temp/VBoxGuestAdditions.iso /media/cdrom
sudo /media/cdrom/VBoxLinuxAdditions.run
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

Add the group,

```bash
sudo usermod --append --groups vboxsf USERNAME
reboot
```

## SHARED CLIPBOARD & DRAG AND DROP

The following is required to use these features:

* Shared folder enabled.
* Guest Additions installed.
* Desktop like GNOME.

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

Restart/Status `networking.service` using `systemclt`
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

## ALLOW HOST ACCESS TO VM

Create a VirtualBox Host-Only Ethernet Adapter and configure
the VM,

```txt
Settings -> Network - Adapter -> Host-Only Adapter.
```

When you add a static IP, your Host will be able to access the VM.

## ALLOW LOCAL NETWORK ACCESS TO VM

You use a bridged adapter,

```txt
Settings -> Network - Adapter -> Bridged Adapter.
```

I did not explore this feature yet.

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

Step 1 - Make a backup (clone).

Step 2 - Create new virtual Machine that is bigger.

Step 3 - Find UUID of your old and new VM.

```bash
VBoxManage list hdds
"C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" list hdds
```

Step 4 - run VBoxManage clonemedium.

```bash
VBoxManage clonemedium <source-guid> <destinatin-guid> --existing
"C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" clonemedium b1e21e90-xxx 391354e8-xxx --existing
```

Step 5 - Download gparted-live*.iso

Insert .iso into the new virtual machine and reboot it.
Grow your partition and save changes.
You may have to delete partitions (swap and /dev/sd2) to grow `/dev/sda1`.
