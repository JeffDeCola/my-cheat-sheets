# PROXMOX CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How to install proxmox on a dell poweredge rack server and create a VM._

tl;dr

```text
PROXMOX
    Storage / VM
        pvesm status                                             # Storage status
    Hardware sensors
        ipmitool sdr                                             # All sensors
        ipmitool sdr type Fan                                    # Fan speed
        ipmitool sdr type Temperature                            # Temperatures
    Power chassis
        ipmitool chassis power status                            # Power state
        ipmitool chassis power on                                # Power on
        ipmitool chassis power off                               # Power off
        ipmitool chassis power reset                             # Hard reboot
    GPU verify
        lspci | grep -i nvidia                                   # Confirm P40 visible
        lspci -v | grep -A20 "Tesla P40"                         # Full P40 info
        lspci -nnk | grep -A3 "82:00.0"                          # Verify P40 bound to vfio-pci
        dmesg | grep -e DMAR -e IOMMU                            # Verify IOMMU enabled
    Fan control
        systemctl status fan-control.service                     # Verify fan override service


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

tbd

Documentation and Reference

* [poweredge r730](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#poweredge-r730-cheat-sheet)
* [nvidia tesla p40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#nvidia-tesla-p40-cheat-sheet)

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

## CONFIGURE PROXMOX FOR PACKER BUILDS

On Proxmox create a packer user:

```bash
pveum user add packer@pam --comment "Packer build user"
```

Create a role with required permissions:

```bash
pveum role add Packer -privs \
  "VM.Allocate \
  VM.Clone \
  VM.Config.CDROM \
  VM.Config.CPU \
  VM.Config.Cloudinit \
  VM.Config.Disk \
  VM.Config.HWType \
  VM.Config.Memory \
  VM.Config.Network \
  VM.Config.Options \
  VM.Audit \
  VM.PowerMgmt \
  Datastore.AllocateSpace \
  Datastore.AllocateTemplate \
  Datastore.Audit \
  Sys.Modify \
  SDN.Use \
  VM.Console \
  VM.GuestAgent.Audit \
  VM.GuestAgent.Unrestricted"
```

Assign the role to the packer user at the root level:

```bash
pveum aclmod / -user packer@pam -role Packer
```

Assign SDN permissions:

```bash
pveum aclmod /sdn/zones/localnetwork -user packer@pam -role Packer
```

Create the API token:

```bash
pveum user token add packer@pam mytoken --privsep=0
```

| Permission                     | Why Needed                    |
|--------------------------------|-------------------------------|
| `VM.Allocate`                  | Create/delete VMs             |
| `VM.Clone`                     | Clone VMs                     |
| `VM.Config.*`                  | Configure VM hardware         |
| `VM.Audit`                     | Read VM config                |
| `VM.PowerMgmt`                 | Start/stop VMs                |
| `VM.Console`                   | Send boot keystrokes via VNC  |
| `VM.GuestAgent.Audit`          | Query VM IP via guest agent   |
| `VM.GuestAgent.Unrestricted`   | Full guest agent access       |
| `Datastore.*`                  | Upload ISO, allocate disk     |
| `SDN.Use`                      | Attach VM to network bridge   |
| `Sys.Modify`                   | System-level modifications    |




### IOMMU SETUP

IOMMU (Input/Output Memory Management Unit) is what allows
Proxmox to hand a physical PCI device (like the P40) directly
to a VM as if the device were plugged straight into that VM's
motherboard. Without IOMMU, the host kernel always owns the
device and the VM can only access it through emulation.

This section turns on the capability — the next section
binds the specific P40 to it.

First, edit grub (the bootloader config the server reads at
power-on):

```bash
nano /etc/default/grub
```

Change the `GRUB_CMDLINE_LINUX_DEFAULT` line to:

```text
GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on iommu=pt"
```

* `intel_iommu=on` — activates the Intel CPU's IOMMU hardware.
* `iommu=pt` — passthrough mode (only enables IOMMU for
  devices that need it, leaving everything else alone for
  performance).

Apply the change:

```bash
update-grub
```

Next, add the VFIO kernel modules so they load on every boot:

```bash
nano /etc/modules
```

Append:

```text
vfio
vfio_iommu_type1
vfio_pci
vfio_virqfd
```

* `vfio` — the core framework that makes passthrough possible.
* `vfio_iommu_type1` — creates the memory isolation boundary
  around the VM.
* `vfio_pci` — the module that actually claims a PCI device
  away from the host.
* `vfio_virqfd` — handles interrupt signaling from the device
  (legacy; merged into vfio core on kernel 6.x+, harmless to
  list).

Reboot:

```bash
reboot
```

Verify IOMMU is active after reboot:

```bash
dmesg | grep -e DMAR -e IOMMU
```

Look for:

```text
DMAR: IOMMU enabled
```

Confirm the P40 is in its own IOMMU group (so it can be passed
through cleanly without dragging other devices along):

```bash
find /sys/kernel/iommu_groups/ -type l | grep -i 82:00
```

The P40 (and only the P40) should appear under a single group
number. If other devices share its group, passthrough is still
possible but requires passing the whole group together — or
applying the ACS override patch.

### BIND P40 FOR CUDA PASSTHROUGH

IOMMU is now on, but the host's default NVIDIA driver
(`nouveau`) will still try to grab the P40 at boot. This
section tells the host to ignore the P40 entirely and hand it
to `vfio-pci`, which holds it ready for a VM.

First, find the P40's vendor:device ID:

```bash
lspci -nn | grep -i nvidia
```

You should see something like:

```text
82:00.0 3D controller [0302]: NVIDIA Corporation GP102GL [Tesla P40] [10de:1b38] (rev a1)
```

The `[10de:1b38]` is the ID we need — `10de` is NVIDIA's
vendor ID, `1b38` is the P40 specifically. The `82:00.0` part
is the PCI address (it may differ on your R730 depending on
which slot the card is in).

Blacklist nouveau and bind the P40 to vfio-pci:

```bash
echo "blacklist nouveau" >> /etc/modprobe.d/blacklist.conf
echo "options vfio-pci ids=10de:1b38" >> /etc/modprobe.d/vfio.conf
update-initramfs -u
reboot
```

After reboot, verify the P40 is now bound to vfio-pci (use
your actual PCI address):

```bash
lspci -nnk | grep -A3 "82:00.0"
```

You should see:

```text
82:00.0 3D controller [0302]: NVIDIA Corporation GP102GL [Tesla P40] [10de:1b38]
        Subsystem: NVIDIA Corporation Device [10de:11d9]
        Kernel driver in use: vfio-pci
        Kernel modules: nvidiafb, nouveau
```

The line that matters is `Kernel driver in use: vfio-pci`.
That confirms the P40 is bound correctly and ready for VM
passthrough.

Note: you'll reboot twice across these two sections (once
after IOMMU setup, once after vfio-pci binding). IOMMU has to
be active before vfio-pci can claim the device meaningfully,
so the two-step reboot is intentional.




## CREATE VM WITH GPU PASSTHROUGH

Create a vm but do not power it on yet. In Proxmox UI, go to.

VM 102 → Hardware → Add → PCI Device

```bash
Device: 0000:82:00.0    ← your P40
All Functions: checked
ROM-Bar: checked
PCI-Express: checked
Primary GPU: leave unchecked
```

Now we need to add a serial port

VM 102 → Hardware → Add → Serial Port

```text
Serial Port: 0
```

Edit the VM Config on Proxmox Host

```bash
nano /etc/pve/qemu-server/102.conf
```

Add this line at the very top:

```text
args: -cpu host,kvm=off
```

Then find the `hostpci` line (the P40 entry) and make sure it looks like this:

```text
hostpci0: 0000:82:00.0,pcie=1,rombar=1
```

Then add this line anywhere in the file:

```bash
vga: std
```

Start VM and when finished that come back and install the drivers.


### CONFIGURE PROXMOX TO USE CUDA CORES

 We need to configure proxmox so a VM will be able to run
 Nvidia drivers and use the P40 for CUDA.

```bash
echo "blacklist nouveau" >> /etc/modprobe.d/blacklist.conf
echo "options vfio-pci ids=10de:1b38" >> /etc/modprobe.d/vfio.conf
update-initramfs -u
reboot
```

Check that it worked

```bash
lspci -nnk | grep -A3 "82:00.0"
```

You're looking for Kernel driver in use: vfio-pci.
That confirms the P40 is bound correctly and ready for
passthrough.


### ADD PCI DEVICE TO VM

tbd

### EDIT VM CONFIG FILE

tbd

## ADD NVIDIA DRIVERS TO VM

```bash
sudo apt install -y nvidia-driver-550 nvidia-utils-550
sudo reboot
```

If there are issues, you may have to get rid of secure boot.

Check if your p40 is alive on you vm

```bash
nvidia-smi
```

