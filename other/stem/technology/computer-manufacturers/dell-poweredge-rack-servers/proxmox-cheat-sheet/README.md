# PROXMOX CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How to install Proxmox on a Dell PowerEdge, create VMs and LXCs,
and configure GPU passthrough._

tl;dr

```text
PROXMOX (run on host)
    VM management (qm)
        qm list                                  # All VMs
        qm status 102                            # VM 102 status
        qm start 102                             # Start VM
        qm stop 102                              # Stop VM (hard)
        qm shutdown 102                          # Shutdown VM (graceful)
        qm reboot 102                            # Reboot VM
        qm terminal 102                          # Serial console
    Storage
        pvesm status                             # All storage pools
        pvesm list SSD-Fast                      # Contents of SSD-Fast
        pvesm list SAS-Data                      # Contents of SAS-Data
        df -h                                    # Disk usage
    Hardware sensors (via R730 BMC)
        ipmitool sdr                             # All sensors
        ipmitool sdr type Fan                    # Fan speeds
        ipmitool sdr type Temperature            # Temperatures
    Power chassis (via R730 BMC)
        ipmitool chassis power status            # Power state
        ipmitool chassis power on                # Power on
        ipmitool chassis power off               # Power off
        ipmitool chassis power reset             # Hard reboot
    GPU passthrough verify
        lspci | grep -i nvidia                   # Confirm GPU visible
        lspci -nnk | grep -A3 "82:00.0"          # Verify GPU bound to vfio-pci
        dmesg | grep -e DMAR -e IOMMU            # Verify IOMMU enabled
    Fan control
        systemctl status fan-control.service     # Verify fan override service
    Power
        shutdown -h now                          # Shutdown Proxmox host
```

Table of Contents

* [INSTALL AND CONFIGURE PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#install-and-configure-proxmox)
  * [MAKE PROXMOX USB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#make-proxmox-usb)
  * [BOOT FROM USB](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#boot-from-usb)
  * [INSTALL PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#install-proxmox)
  * [UPDATE PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#update-proxmox)
* [SSD-FAST STORAGE (MAIN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#ssd-fast-storage-main)
* [SAS-DATA STORAGE (BULK)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#sas-data-storage-bulk)
  * [PARTITION HDD VIRTUAL DISK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#partition-hdd-virtual-disk)
  * [ADD A FILESYSTEM AND MOUNT IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#add-a-filesystem-and-mount-it)
  * [ADD SAS-DATA TO PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#add-sas-data-to-proxmox)
* [CREATE A VM - UBUNTU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#create-a-vm---ubuntu)
* [CREATE A LXC - DEBIAN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#create-a-lxc---debian)
* [BACKUP VMs/LXCs USING PROXMOX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#backup-vmslxcs-using-proxmox)
* [CONFIGURE PROXMOX FOR PACKER BUILDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#configure-proxmox-for-packer-builds)
* [NVIDIA P40 GPU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#nvidia-p40-gpu)
  * [IOMMU SETUP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#iommu-setup)
  * [BIND GPU FOR PASSTHROUGH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#bind-gpu-for-passthrough)
  * [CREATE VM WITH GPU PASSTHROUGH - ollama](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#create-vm-with-gpu-passthrough-ollama)

Documentation and Reference

* [poweredge r730](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#poweredge-r730-cheat-sheet)
* [nvidia tesla p40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#nvidia-tesla-p40-cheat-sheet)

## INSTALL AND CONFIGURE PROXMOX

Let's install proxmox on the R730 and set up some VMs and containers.

### MAKE PROXMOX USB

Make a proxmox USB stick,

* Use an 8GB+ USB stick
* Download Proxmox VE ISO from
  [proxmox.com/en/downloads](http://proxmox.com/en/downloads)
* Flash it to a USB drive using Rufus (Windows)

### BOOT FROM USB

Done via iDRAC virtual console (no monitor or keyboard
attached to the R730).

* Plug the USB stick into the R730 (physical)
* From your workstation, open iDRAC at `https://<idrac-ip>`
* Launch Virtual Console from the iDRAC dashboard
* Cold-reboot the server (iDRAC Power → Reset System, or the
  physical button)
* Press **F11** via the virtual keyboard at the Dell POST
  screen to get the One-Shot BIOS Boot Menu
* Select your USB drive to boot

### INSTALL PROXMOX

Still in the iDRAC virtual console, the Proxmox installer
will boot from the USB and present an installer menu,

* Select **Install Proxmox VE (Graphical)**
* Target disk: the SSD virtual disk (`/dev/sda`, ~744GB —
  the SSD-Fast RAID 10 from the R730 cheat sheet). Don't pick
  `/dev/sdb` — that's SAS-Data.
* Network config:
  * IP/CIDR: `192.168.20.135/24`
  * Gateway: `192.168.20.1`
  * DNS: `192.168.20.1` (or whatever you use)
  * Hostname FQDN: `r730.proxmox`
* Install — takes about 5–10 minutes

When the installer finishes, it prompts to reboot. Eject the
USB stick (or unmap the virtual media in iDRAC) and let the
server boot into Proxmox.

If the hostname didn't stick (sometimes the installer's
hostname doesn't fully propagate), set it manually after first
login,

```bash
hostnamectl set-hostname r730.proxmox
```

Verify,

```bash
hostnamectl status
cat /etc/hosts
```

### UPDATE PROXMOX

Back on your workstation, copy your SSH public key to Proxmox
so you don't have to type the password every time,

```bash
ssh-copy-id root@192.168.20.135
```

Then SSH in,

```bash
ssh root@192.168.20.135
```

Before updating, switch from the enterprise repo (paid
subscription required) to the no-subscription repo,

```bash
# Disable enterprise repos
sed -i 's/^deb/#deb/' /etc/apt/sources.list.d/pve-enterprise.list
sed -i 's/^deb/#deb/' /etc/apt/sources.list.d/ceph.list 2>/dev/null || true

# Add no-subscription repo (replace 'bookworm' with your Debian
# codename — 'trixie' for Proxmox 9.x; check with cat /etc/os-release)
echo "deb http://download.proxmox.com/debian/pve bookworm pve-no-subscription" \
  > /etc/apt/sources.list.d/pve-no-subscription.list
```

Now update,

```bash
apt update && apt dist-upgrade -y
```

## SSD-FAST STORAGE (MAIN)

SSD-Fast (`/dev/sda`, 744 GiB RAID 10) is the **main disk** —
think of it as the C: drive on a Windows box. Proxmox itself
boots from here, and the LVM thin pool that holds VM disks
lives here too.

The Proxmox installer set this up automatically during
INSTALL PROXMOX — no manual setup needed. The installer
carved `/dev/sda` into,

* `/dev/sda1` — 1MB BIOS boot
* `/dev/sda2` — 1GB EFI system
* `/dev/sda3` — 743GB Linux LVM (the `pve` volume group)
  * `root` — 96GB Proxmox OS
  * `swap` — 8GB swap
  * `data` — ~610GB LVM thin pool (VM disks live here)

Verify the layout from the Proxmox shell,

```bash
lsblk
vgs
lvs
```

## SAS-DATA STORAGE (BULK)

SAS-Data (`/dev/sdb`, 3.3 TiB RAID 5) is the **bulk disk** —
the equivalent of a D: drive holding everything that doesn't
need to live on the fast SSD: backups, ISOs, LXC templates,
overflow VM disks.

Unlike SSD-Fast, the installer didn't touch this array. We
need to partition it, format it, mount it, and register it
with Proxmox manually.

### PARTITION HDD VIRTUAL DISK

Run `fdisk -l` to see your disks,

```bash
fdisk -l
```

You'll see two disks:

* `/dev/sda` = SSD RAID 10 (744GB) — Proxmox OS lives here
  * `/dev/sda1` — 1007K BIOS boot
  * `/dev/sda2` — 1G EFI System
  * `/dev/sda3` — 743G Linux LVM
* `/dev/sdb` = SAS RAID 5 (3.27TB) — no partitions yet

Create a single partition using the full drive,

```bash
apt install parted -y
parted /dev/sdb mklabel gpt
parted /dev/sdb mkpart primary 0% 100%
```

Run `fdisk -l` again to confirm `/dev/sdb1` now appears.

### ADD A FILESYSTEM AND MOUNT IT

We put XFS directly on `/dev/sdb1` — one big directory that holds everything
(VM disks, backups, ISOs). No LVM needed.

Format the partition with XFS,

```bash
mkfs.xfs /dev/sdb1
```

Mount it — tell Linux this partition lives at `/mnt/sas-data`,

```bash
mkdir -p /mnt/sas-data
mount /dev/sdb1 /mnt/sas-data
```

Make it auto-mount on every reboot,

```bash
echo '/dev/sdb1 /mnt/sas-data xfs defaults 0 2' >> /etc/fstab
```

Verify it mounted correctly and check the size,

```bash
df -h
```

You should see `/dev/sdb1` mounted at `/mnt/sas-data` with ~3.3TB available.

### ADD SAS-DATA TO PROXMOX

Tell Proxmox to use this directory for all bulk storage,

```text
Datacenter → Storage → Add → Directory
ID:        SAS-Data
Directory: /mnt/sas-data
Content:   Disk Image, Backup, ISO Image, Snippets
```

## CREATE A VM - UBUNTU

Get the latest version, I got,

```bash
wget -P /mnt/sas-data/template/iso/ \
  https://releases.ubuntu.com/24.04/ubuntu-24.04.4-live-server-amd64.iso
```

Head to your Proxmox UI and click Create VM in the top right.

General tab,

```text
Node: r730
VM ID: 101
Name: ubuntu-general
click start at boot
```

OS tab,

```text
Storage: SAS-Data
ISO Image: ubuntu-24.04.4-live-server-amd64.iso
Type: Linux
Version: 6.x - 2.6 Kernel
```

System tab,

```text
Machine: q35
BIOS: OVMF (UEFI)
EFI Storage: SSD-Fast
SCSI Controller: VirtIO SCSI single
Qemu Agent: check this
```

Disks tab,

```text
Bus/Device: SCSI
Storage: SSD-Fast
Disk size: 40
Cache: Write back
Discard: checked
SSD emulation: checked
```

CPU tab,

```text
Sockets: 1
Cores: 4
Type: host
```

Memory tab,

```text
Memory: 8192 (that's 8GB)
Ballooning: checked
Minimum memory: 512
```

Network tab,

```text
Bridge: vmbr0
Model: VirtIO (paravirtualized)
Firewall: checked
```

Review and create.

Now start the VM,

* Select VM 101 in the left panel
* Click Start (top right)
* Then click Console to open the display

Go through the Ubuntu setup process.

> NOTE: Uncheck "Set up this disk as an LVM group"
> Just "Use an entire disk".

When you first log in, update and upgrade your distro,

```bash
sudo apt update && sudo apt upgrade -y
```

## CREATE A LXC - DEBIAN

First we need a Debian 12 LXC template. This is a special Debian
made for Proxmox. In the Proxmox UI,

```text
Click on your node (the server name) in the left panel
Click local storage
Click CT Templates
Click Templates button at the top
Search for debian-12 and download it
```

Once this template is downloaded, click Create CT at the top
right of the Proxmox UI.

General tab,

```text
CT ID: 201
Hostname: lxc-debian-tailscale
Password: (set something strong, you'll need it)
SSH key: Can create this later
Unprivileged container: NO (uncheck it — make it privileged)
```

Template tab,

```text
Storage: local
Template: debian-12-standard
```

Disks tab,

```text
Storage: SAS-Data
Size: 4GB
```

CPU tab,

```text
Cores: 1
```

Memory tab,

```text
RAM: 256
Swap: 256
```

Network tab,

```text
Bridge: vmbr0
IPv4: DHCP (we assign static IP at router)
```

DNS tab,

```text
Leave as default
```

Review and create.

Inside the LXC, you may have to allow root logins,

```bash
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
systemctl restart sshd
```

When you first log in, update and upgrade your distro,

```bash
apt update && apt upgrade -y
```

May also want to create your ssh keys,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github (lxc-debian-tailscale)"
```

## BACKUP VMs/LXCs USING PROXMOX

Once you get some VMs set up, you can add a backup,

```text
Datacenter → Backup → Add
```

Then set,

```text
Storage: SAS-Data
Schedule: 0 2 * * * (2am daily)
Mode: Snapshot
Compression: ZSTD
Pick the VMs, LXCs you want for this backup
```

Under retention for this backup pick what you want. I picked,

```text
Ollama VM — 2 days since it's 117GB
Everything else — 7 day retention
```

Hit create.

## CONFIGURE PROXMOX FOR PACKER BUILDS

If you use Packer to build VM templates, on Proxmox create a Packer user,

```bash
pveum user add packer@pam --comment "Packer build user"
```

Create a role with required permissions,

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

Assign the role to the Packer user at the root level,

```bash
pveum aclmod / -user packer@pam -role Packer
```

Assign SDN permissions,

```bash
pveum aclmod /sdn/zones/localnetwork -user packer@pam -role Packer
```

Create the API token,

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

## NVIDIA P40 GPU

I installed an NVIDIA Tesla P40 in the R730 for local LLM
inference (Ollama). Three steps to make the card usable from a
VM,

1. **IOMMU SETUP** — turn on passthrough capability at the kernel level.
2. **BIND GPU FOR PASSTHROUGH** — claim the GPU for `vfio-pci` so the host won't.
3. **CREATE VM WITH GPU PASSTHROUGH - OLLAMA** — hand the GPU to VM 102.

For the card itself (specs, install, NVIDIA driver setup), see
the
[nvidia tesla p40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#nvidia-tesla-p40-cheat-sheet)
cheat sheet.

### IOMMU SETUP

IOMMU lets Proxmox hand a PCI device directly to a VM. Without
it, only emulated access works.

Edit grub,

```bash
nano /etc/default/grub
```

Set the cmdline,

```text
GRUB_CMDLINE_LINUX_DEFAULT="quiet intel_iommu=on iommu=pt"
```

Apply, add the VFIO modules, reboot,

```bash
update-grub

cat >> /etc/modules << 'EOF'
vfio
vfio_iommu_type1
vfio_pci
vfio_virqfd
EOF

reboot
```

Verify after reboot,

```bash
dmesg | grep -e DMAR -e IOMMU                       # look for: DMAR: IOMMU enabled
find /sys/kernel/iommu_groups/ -type l | sort -V    # GPU should be in its own group
```

### BIND GPU FOR PASSTHROUGH

IOMMU is on, but `nouveau` still grabs the GPU at boot. Bind
the GPU to `vfio-pci` so the host ignores it.

Find the vendor:device ID,

```bash
lspci -nn | grep -i nvidia
# e.g. 82:00.0 ... [Tesla P40] [10de:1b38]
```

Blacklist nouveau and bind to vfio-pci (replace `10de:1b38`
with your ID),

```bash
echo "blacklist nouveau" >> /etc/modprobe.d/blacklist.conf
echo "options vfio-pci ids=10de:1b38" >> /etc/modprobe.d/vfio.conf
update-initramfs -u
reboot
```

Verify (use your PCI address),

```bash
lspci -nnk | grep -A3 "82:00.0"
# look for: Kernel driver in use: vfio-pci
```

### CREATE VM WITH GPU PASSTHROUGH - OLLAMA

Create VM 102 following CREATE A VM - UBUNTU above (do not
power it on yet). Then,

VM 102 → Hardware → Add → PCI Device,

```text
Device: 0000:82:00.0    ← your GPU
All Functions: checked
ROM-Bar: checked
PCI-Express: checked
Primary GPU: leave unchecked
```

VM 102 → Hardware → Add → Serial Port,

```text
Serial Port: 0
```

Edit the VM config,

```bash
nano /etc/pve/qemu-server/102.conf
```

Add at the top,

```text
args: -cpu host,kvm=off
```

Make sure the hostpci line reads,

```text
hostpci0: 0000:82:00.0,pcie=1,rombar=1
```

Add anywhere,

```text
vga: std
```

Start the VM and install the NVIDIA drivers — see the
[nvidia tesla p40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#add-nvidia-drivers-to-vm)
cheat sheet, ADD NVIDIA DRIVERS section.
