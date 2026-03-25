# PROXMOX - INSTALL, CONFIGURE AND CREATE VM

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How to install proxmox on a dell poweredge rack server and create a VM._

tl;dr

```text
ON PROXMOX
    shutdown -h now                       # shutdown enter proxmox
    qm list                               # qm is the proxmox cli for managing vms
    qm status 102
    qm start 102
    qm stop 102
    qm reboot 102
    qm shutdown 102
    qm terminal 102
    pvesm status                          # All storage pools
    pvesm list SSD-Fast                   # Contents of SSD-Fast
    pvesm list SAS-Data                   # Contents of SAS-Data
    df -h                                 # Disk usage
ON VM
    nvidia-smi                            # Full P40 overview
    watch -n1 nvidia-smi                  # live refresh every second
    nvidia-smi dmon                       # live stats (utilization, temp, power)
    lspci -v | grep -A20 "Tesla P40"      # full P40 detail
    lscpu                                 # CPU info
```

Table of Contents

* [MAKE PROXMOX USB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#make-proxmox-usb)
* [BOOT FROM USB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#boot-from-usb)
* [INSTALL PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#install-proxmox)
* [UPDATE PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#update-proxmox)
* [PARTITION HDD VIRTUAL DISK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#partition-hdd-virtual-disk)
* [ADD A FILESYSTEM AND MOUNT IT (SAS-DATA)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#add-a-filesystem-and-mount-it-sas-data)
* [ADD SAS-DATA TO PROXMOX (KEEP BULK DATA AND BACKUPS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#add-sas-data-to-proxmox-keep-bulk-data-and-backups)
* [CREATE A VM - UBUNTU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#create-a-vm---ubuntu)
* [CREATE A LXC - DEBIAN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-install-configure-and-create-vm-cheat-sheet#create-a-lxc---debian)
* [BACKUP VMs/LXCs USING PROXMOX]()

## MAKE PROXMOX USB

* Download Proxmox VE ISO from
  [proxmox.com/en/downloads](http://proxmox.com/en/downloads)
* Flash it to a USB drive using Rufus (Windows)
* Use an 8GB+ USB stick

## BOOT FROM USB

* Plug USB into the R730
* Get the virtual console and keyboard ready
* Reboot the server (cold)
* Press F11 at the Dell POST screen to get the boot menu (use virtual keyboard)
* Bios Boot Menu
  * One-Shot BIOS Boot Menu
  * Select your USB drive to boot

## INSTALL PROXMOX

* Install Proxmox VE (Graphical)
* Target disk: SSD virtual disk
* Set your static IP: 192.168.20.135
* Set hostname "r730.proxmox"
* Install - takes about 5–10 minutes

As a side note, may have to run this to set hostname properly

```bash
hostnamectl set-hostname r730.proxmox
```

To check

```bash
nano /etc/hosts
```

## UPDATE PROXMOX

To be able to ssh into proxmox I copied my public keys to it

```bash
ssh-copy-id root@192.168.20.135
```

```bash
apt update && apt dist-upgrade -y
```

If you hit subscription errors, disable the enterprise repo
and add the no-subscription repo.

## PARTITION HDD VIRTUAL DISK

Run `fdisk -l` to see your disks

```bash
fdisk -l
```

You'll see two disks:

* `/dev/sda` = SSD RAID 10 (744GB) — Proxmox OS lives here
  * `/dev/sda1` — 1007K BIOS boot
  * `/dev/sda2` — 1G EFI System
  * `/dev/sda3` — 743G Linux LVM
* `/dev/sdb` = SAS RAID 5 (3.27TB) — no partitions yet

Create a single partition using the full drive

```bash
apt install parted -y
parted /dev/sdb mklabel gpt
parted /dev/sdb mkpart primary 0% 100%
```

Run `fdisk -l` again to confirm `/dev/sdb1` now appears.

## ADD A FILESYSTEM AND MOUNT IT (SAS-DATA)

We put XFS directly on `/dev/sdb1` — one big directory that holds everything
(VM disks, backups, ISOs). No LVM needed.

Format the partition with XFS

```bash
mkfs.xfs /dev/sdb1
```

Mount it — tell Linux this partition lives at `/mnt/sas-data`

```bash
mkdir -p /mnt/sas-data
mount /dev/sdb1 /mnt/sas-data
```

Make it auto-mount on every reboot

```bash
echo '/dev/sdb1 /mnt/sas-data xfs defaults 0 2' >> /etc/fstab
```

Verify it mounted correctly and check the size

```bash
df -h
```

You should see `/dev/sdb1` mounted at `/mnt/sas-data` with ~3.3TB available.

## ADD SAS-DATA TO PROXMOX (KEEP BULK DATA AND BACKUPS)

Tell Proxmox to use this directory for all bulk storage

```text
Datacenter → Storage → Add → Directory
ID:        SAS-Data
Directory: /mnt/sas-data
Content:   Disk Image, Backup, ISO Image, Snippets
```

## CREATE A VM - UBUNTU

Get the latest version, I got

```bash
wget -P /mnt/sas-data/template/iso/ \
  https://releases.ubuntu.com/24.04/ubuntu-24.04.4-live-server-amd64.iso
```

Head to your proxmox UI and click Create VM in the top right.

General Tab

```text
Node: r730
VM ID: 101
Name: ubuntu-general
click start at boot
```

OS tab:

```text
Storage: SAS-Data
ISO Image: ubuntu-24.04.4-live-server-amd64.iso
Type: Linux
Version: 6.x - 2.6 Kernel
```

System tab:

```text
Machine: q35
BIOS: OVMF (UEFI)
EFI Storage: SSD-Fast
SCSI Controller: VirtIO SCSI single
Qemu Agent: check this
```

Disks tab:

```text
Bus/Device: SCSI
Storage: SSD-Fast
Disk size: 40
Cache: Write back
Discard: checked
SSD emulation: checked
```

CPU tab:

```text
Sockets: 1
Cores: 4
Type: host
```

Memory tab:

```text
Memory: 8192 (that's 8GB)
Ballooning: checked
Minimum memory: 512
```

Network tab:

```text
Bridge: vmbr0
Model: VirtIO (paravirtualized)
Firewall: checked
```

Review and create.

Now start the VM

* Select VM 101 in the left panel
* Click Start (top right)
* Then click Console to open the display

Go threw the ubuntu setup process.

> NOTE: Uncheck "Set up this disk as an LVM group"
> Just "Use an entire disk".

When you first login, you should update and upgrade your distro.

```bash
sudo apt update && sudo apt upgrade -y
```

## CREATE A LXC - DEBIAN

First we need a Debian 12 LXC template.
This is a special debian made for proxmox.
In the Proxmox UI

```text
Click on your node (the server name) in the left panel
Click local storage
Click CT Templates
Click Templates button at the top
Search for debian-12 and download it
```

Once this special template is downloaded
click Create CT button at the top right of the Proxmox UI.

General Tab

```text
CT ID: 201
Hostname: lxc-debian-tailscale
Password: (set something strong, you'll need it)
SSH key: CAn create this later
Unprivileged container: NO (uncheck it — make it privileged)
```

Template Tab

```text
Storage: local
Template: debian-12-standard
```

Disks Tab

```text
Storage: SAS-Data
Size: 4GB
```

CPU tab

```text
Cores: 1
```

Memory Tab

```text
RAM: 256
Swap: 256
```

Network Tab

```text
Bridge: vmbr0
IPv4: DHCP (We assign static IP at router)
```

DNS TAb

```text
Leave as default
```

Review and create.

You may have to allow root logins, from proxmox

```bash
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
systemctl restart sshd
```

When you first login, you should update and upgrade your distro.

```bash
apt update && sudo apt upgrade -y
```

May also want to create your ssh keys

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github (lxc-debian-tailscale)"
```

## BACKUP VMs/LXCs USING PROXMOX

Once you get some VMs setup you can now add a backup

```text
Datacenter → Backup → Add
```

Then set

```text
Storage: SAS-Data
Schedule: 0 2 * * * (2am daily)
Mode: Snapshot
Compression: ZSTD
Pick the VMs, LXC you want for this backup
```

Under retention for this backup pick what you want, I picked

```bash
Ollama VM — 2 days since it's 117GB
Everything else — 7 day retention
```

Hit create
