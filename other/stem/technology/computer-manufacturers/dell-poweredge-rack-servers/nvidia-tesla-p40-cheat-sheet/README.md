# NVIDIA TESLA P40 CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Configure a Dell PowerEdge R730 with an NVIDIA Tesla P40 for GPU passthrough._

tl;dr

```text
PROXMOX
    GPU verify
        lspci | grep -i nvidia                                   # Confirm P40 visible
        lspci -v | grep -A20 "Tesla P40"                         # Full P40 info
        lspci -nnk | grep -A3 "82:00.0"                          # Verify P40 bound to vfio-pci
        dmesg | grep -e DMAR -e IOMMU                            # Verify IOMMU enabled
    Fan control
        systemctl status fan-control.service                     # Verify fan override service
ON VM WITH P40
    nvidia-smi                                                   # Full P40 overview
    watch -n1 nvidia-smi                                         # Live refresh every second
    nvidia-smi dmon                                              # Live stats (utilization, temp, power)
```

Table of Contents

TBD

Documentation and Reference

* [poweredge r730](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#poweredge-r730-cheat-sheet)
* [proxmox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#proxmox-cheat-sheet)

## OVERVIEW

![IMAGE - dell-poweredge-r730-2016 - IMAGE](../../../../../../docs/pics/other/dell-poweredge-r730-2016.svg)

## CONFIGURE NVIDIA TESLA P40

The Tesla P40 is a Pascal-architecture compute GPU — 24GB
GDDR5, 3840 CUDA cores, 250W TDP, **passively cooled with no
fans of its own** (it relies on chassis airflow), and **no
display outputs** (compute only, not usable for video). It
slots into a full-height, full-length PCIe x16 slot.

Physical install notes for the R730:

* The P40 needs a **CPU-style power cable** (EPS 8-pin, not a
  standard 6+2 PCIe connector). I bought one separately — it's
  not included with the bare card.
* On this R730 the P40 sits in **PCIe slot 3**.
* The card runs **stock** with no aftermarket cooling shroud,
  relying on the R730's chassis fans for airflow. This is why
  fan noise needs to be managed (see FIX FAN NOISE below).

After installing the P40 and installing Proxmox, confirm the
host sees the card:

```bash
lspci | grep -i nvidia
```

Should return something like:

```text
82:00.0 3D controller: NVIDIA Corporation GP102GL [Tesla P40] (rev a1)
```

For full details on the card:

```bash
lspci -v | grep -A20 "Tesla P40"
```

To make the P40 usable as a CUDA device inside a VM, the host
needs two configuration steps before the VM is created:

1. **IOMMU setup** — enable the Intel IOMMU and load the VFIO
   kernel modules so passthrough is possible.
2. **Bind P40 for CUDA passthrough** — blacklist the nouveau
   driver on the host and bind the P40 directly to vfio-pci so
   the host won't claim it.

## FIX FAN NOISE

The iDRAC has no clue what this new hardware is so it ramps up the fans.
But this is unnecessary since the P40 manages it's own temperatures.

Disable third party PCIe fan response

ssh into iDRAC (ssh root@192.168.20.134)

```text
racadm set system.thermalsettings.ThirdPartyPCIFanResponse 0
```

Verify it took effect

```text
racadm get System.ThermalSettings.ThirdPartyPCIFanResponse
```

Also note, you want both power supplies connected to
reduce fan noise.

Also may want to try capping the minimum fan speed at 25%
instead of letting iDRAC decide.

```text
racadm set System.ThermalSettings.ThirdPartyPCIFanResponse 0
racadm set System.ThermalSettings.MinimumFanSpeed 25
```

OK, it didn't really work, so I did a IPMI override, but now
I had to create a script that re-applies the IPMI override on every boot.

```bash
cat << 'EOF' > /usr/local/bin/fan-control.sh
#!/bin/bash
# Disable automatic fan control
ipmitool raw 0x30 0x30 0x01 0x00
# Set fans to 25%
ipmitool raw 0x30 0x30 0x02 0xff 0x19
EOF
```

```bash
chmod +x /usr/local/bin/fan-control.sh
```

```bash
cat << 'EOF' > /etc/systemd/system/fan-control.service
[Unit]
Description=Dell Fan Control Override
After=multi-user.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/fan-control.sh

[Install]
WantedBy=multi-user.target
EOF
```

```bash
systemctl daemon-reload
systemctl enable fan-control.service
systemctl status fan-control.service
```

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

