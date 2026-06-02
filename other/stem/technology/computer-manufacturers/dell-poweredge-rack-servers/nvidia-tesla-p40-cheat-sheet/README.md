# NVIDIA TESLA P40 CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_NVIDIA Tesla P40 — install, configure, and use as a CUDA device
on a Dell PowerEdge R730._

tl;dr

```text
ON VM WITH P40
    nvidia-smi                            # Full P40 overview
    watch -n1 nvidia-smi                  # Live refresh every second
    nvidia-smi dmon                       # Live stats (utilization, temp, power)
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#overview)
* [CONFIGURE NVIDIA TESLA P40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#configure-nvidia-tesla-p40)
* [ADD NVIDIA DRIVERS TO VM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#add-nvidia-drivers-to-vm)

Documentation and Reference

* [poweredge r730](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#poweredge-r730-cheat-sheet)
* [proxmox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#proxmox-cheat-sheet)

## OVERVIEW

The Tesla P40 is a Pascal-architecture compute GPU — 24GB
GDDR5, 3840 CUDA cores, 250W TDP, **passively cooled with no
fans of its own** (it relies on chassis airflow), and **no
display outputs** (compute only, not usable for video). It
slots into a full-height, full-length PCIe x16 slot.

In my homelab the P40 is installed in a Dell PowerEdge R730
and exposed to a Proxmox VM via PCIe passthrough for CUDA
workloads (e.g. Ollama).

* For the host chassis (iDRAC, RAID, fan noise), see the
  [poweredge r730](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#poweredge-r730-cheat-sheet)
  cheat sheet.
* For VM passthrough setup (IOMMU, vfio-pci, VM config), see
  the [proxmox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#proxmox-cheat-sheet)
  cheat sheet.

![IMAGE - dell-poweredge-r730-2016 - IMAGE](../../../../../../docs/pics/other/dell-poweredge-r730-2016.svg)

## CONFIGURE NVIDIA TESLA P40

Physical install notes for the R730:

* The P40 needs a **CPU-style power cable** (EPS 8-pin, not a
  standard 6+2 PCIe connector). I bought one separately - it was not
  included with the bare card.
* On my R730 the P40 sits in **PCIe slot 3**.
* The card runs **stock** with no aftermarket cooling shroud,
  relying on the R730's chassis fans for airflow. This is why
  fan noise needs to be managed.

After installing the P40 and bringing up Proxmox, confirm the
host sees the card:

```bash
lspci | grep -i nvidia
```

Should return something like (the PCI address `82:00.0` may
differ depending on which slot the card is in):

```text
82:00.0 3D controller: NVIDIA Corporation GP102GL [Tesla P40] (rev a1)
```

For full details on the card:

```bash
lspci -v | grep -A20 "Tesla P40"
```

To make the P40 usable as a CUDA device inside a VM, the host
still needs IOMMU and vfio-pci passthrough configuration — see
the [proxmox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#proxmox-cheat-sheet)
cheat sheet for both the host-side passthrough setup and VM
PCI assignment.

## ADD NVIDIA DRIVERS

This step runs **inside the VM**, not on the Proxmox host. The
host deliberately doesn't have NVIDIA drivers - it blacklisted
nouveau and bound the P40 to `vfio-pci` so the VM could claim
the card cleanly. Now that the VM owns the P40 via passthrough,
the VM needs the actual NVIDIA driver to use it.

On Ubuntu (this VM runs Ubuntu 24.04):

```bash
sudo apt update
sudo apt install -y nvidia-driver-570 nvidia-utils-570
sudo reboot
```

Driver `570` is the current LTS line as of 2026. To check
what's currently available:

```bash
apt-cache search nvidia-driver | grep -E '^nvidia-driver-[0-9]+'
```

Or let Ubuntu pick the recommended driver automatically:

```bash
sudo ubuntu-drivers autoinstall
sudo reboot
```

If Secure Boot is enabled on the VM, the installer will prompt
you to set a **MOK (Machine Owner Key) password**. On next
boot, you'll be dropped into the MOK manager — choose
**Enroll MOK**, enter that password, and continue. Without
this, the NVIDIA kernel modules are unsigned and won't load.
The alternative is to disable Secure Boot in the VM's OVMF
firmware settings (Proxmox UI → VM → Options → BIOS).

After reboot, confirm the VM sees the P40:

```bash
nvidia-smi
```

The key things to verify in the output:

* **`Tesla P40`** appears under `Name`
* **`24576 MiB`** shown for memory (the full 24GB)
* **CUDA Version** populated (any version is fine; the driver
  reports the highest CUDA version it supports)

Note: this driver install gives you the NVIDIA driver and
basic utilities (`nvidia-smi`, etc.). For running Ollama or
other prebuilt CUDA workloads, that's all you need — they
bundle the CUDA runtime. For native CUDA development on the
VM, also install `nvidia-cuda-toolkit`.
