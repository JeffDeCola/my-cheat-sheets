# POWEREDGE R730 CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Configure a Dell PowerEdge R730 — iDRAC, RAID virtual disks, and fan noise mitigation._

tl;dr

```text
iDRAC (racadm)
    racadm serveraction powerdown                                # Power off
    racadm serveraction powerup                                  # Power on
    racadm serveraction powercycle                               # Hard reboot
    racadm get System.ThermalSettings.ThirdPartyPCIFanResponse   # Verify fan response setting
    racadm set iDRAC.Users.2.Password <new-password>             # Change password
    racadm getsysinfo                                            # Server info (model, service tag, etc.)
    racadm getniccfg                                             # iDRAC network config
    racadm getsel                                                # System event log
    racadm racreset                                              # Reset iDRAC if it gets wedged
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#overview)
* [CONFIGURE IDRAC](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#configure-idrac)
* [CONFIGURE RAID VIRTUAL DRIVES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#configure-raid-virtual-drives)
* [FIX FAN NOISE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#fix-fan-noise)
  * [Attempted: racadm thermal settings](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#attempted-racadm-thermal-settings)
  * [Solution: IPMI override via systemd](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/poweredge-r730-cheat-sheet#solution-ipmi-override-via-systemd)

Documentation and Reference

* [nvidia tesla p40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/nvidia-tesla-p40-cheat-sheet#nvidia-tesla-p40-cheat-sheet)
* [proxmox](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/computer-manufacturers/dell-poweredge-rack-servers/proxmox-cheat-sheet#proxmox-cheat-sheet)

## OVERVIEW

![IMAGE - dell-poweredge-r730-2016 - IMAGE](../../../../../../docs/pics/other/dell-poweredge-r730-2016.svg)

## CONFIGURE IDRAC

iDRAC (Integrated Dell Remote Access Controller) is the R730's
management processor. It runs independently of the
main system and lets you control the server (power, BIOS, RAID,
virtual console) over the network even when the OS is off or
the server is powered down.

iDRAC has its own dedicated ethernet port on the back of the
R730, its own IP address, and its own login.

There are two ways to access iDRAC, ssh or browser.
Default factory credentials are `root` / `calvin`.
SSH'ing into iDRAC drops you into the `racadm` shell,
**not** a Linux shell. Linux commands like `pvesm status` or
`ls` will return `COMMAND NOT RECOGNIZED`. Only `racadm`
commands work.

Useful racadm commands.

```bash
racadm getsysinfo               # Server info (model, service tag, IPs, BIOS)
racadm getniccfg                # iDRAC network config
racadm getsel                   # System event log
racadm racreset                 # Reset iDRAC if it gets wedged
```

## CONFIGURE RAID VIRTUAL DRIVES

The R730 uses the embedded **PERC H730 Mini** controller in
RAID mode (not HBA mode). I set mine as follows,

* **SSD-Fast** (800GB raw / ~610 GiB thin pool for VMs)
  * RAID 10, 4x400GB SSD
  * 1x400GB SSD hot spare
  * Hosts Proxmox itself plus the LVM thin pool used for VM disks
* **SAS-Data** (3.6TB raw / 3.3 TiB usable)
  * RAID 5, 4x1.2TB HDD
  * 1x1.2TB HDD hot spare
  * Mounted at `/mnt/sas-data` (XFS) for ISOs, backups, and bulk data

The SSD array is sized for fast VM disk I/O; the SAS array is
the bulk cold-storage tier.

Confirm the controller status first,

```text
Storage → Controllers → PERC H730 Mini
  * Battery Backup Unit (BBU) — should show Present / Ready
  * Cache Memory — should show 1GB (standard on H730 Mini)
```

The BBU is what makes **Write Back** caching safe — if the
server loses power mid-write, the battery preserves the cache
contents until power returns. If the BBU isn't Ready, the
controller falls back to **Write Through** (slower, but no
data risk). The steps below assume BBU is Ready and use Write
Back.

Create the first virtual disk,

```text
Storage → Virtual Disks → Create Virtual Disk
  * Virtual Disk Name:                    SSD-Fast
  * Controller:                           PERC H730 Mini (embedded)
  * Layout:                               RAID 10
  * Media Type:                           SSD
  * Stripe Element Size:                  64KB (default)
  * Read Policy:                          Adaptive Read Ahead
  * Write Policy:                         Write Back (BBU required)
  * Disk Cache Policy:                    Enabled
  * T10 Protection Information Capability: Disabled
```

Assign the hot spare,

```text
Storage → Virtual Disks → Manage
  * Right-click the 5th SSD (your spare drive)
  * Select Assign Dedicated Hot Spare
  * Assign it to the SSD-Fast virtual disk
```

Verify in **Physical Disks** that the drive is now marked as a
hot spare.

Repeat the process for the second virtual disk with these
changes,

```text
  * Virtual Disk Name:    SAS-Data
  * Layout:               RAID 5
  * Media Type:           HDD
  (all other settings the same)
```

Assign one of the 1.2TB HDDs as a dedicated hot spare for
SAS-Data using the same process.

If you're rebuilding and the drives have a previous RAID
configuration on them, you'll need to clear it first:

```text
Storage → Controllers → PERC H730 Mini → Setup
  * Clear Foreign Configuration (if drives came from another array)
  * Clear Configuration (wipes all RAID config — destructive)
```

After the OS is installed and the host is up, you can verify
the layout from the Proxmox shell,

```bash
lsblk                # See the RAID virtual disks as sda, sdb
pvesm status         # See how Proxmox is using them
vgs                  # Volume group capacity (SSD-Fast → pve VG)
lvs                  # Logical volumes (thin pool, VM disks)
df -h                # Filesystem usage (SAS-Data mount)
```

On this R730, `lsblk` shows `sda` (744 GiB, the SSD-Fast RAID
10) carved into Proxmox root, swap, and a ~610 GiB LVM thin
pool for VM disks. `sdb` (3.3 TiB, the SAS-Data RAID 5) is a
single XFS partition mounted at `/mnt/sas-data`.

## FIX FAN NOISE

The R730's thermal management ramps the fans up when it detects
an unknown PCIe card (the P40), because iDRAC has no profile for
it. The P40 manages its own thermals — it doesn't need the chassis
fans at full speed — so the noise is unnecessary.

Make sure **both power supplies** are used. The R730 runs the fans
harder when only one PSU is connected.

### Attempted: racadm thermal settings

The "official" Dell knob is to tell iDRAC not to use third-party
PCIe response curves:

```bash
# SSH into iDRAC
racadm set System.ThermalSettings.ThirdPartyPCIFanResponse 0
racadm get System.ThermalSettings.ThirdPartyPCIFanResponse   # verify
```

Optionally also cap the minimum fan speed:

```bash
racadm set System.ThermalSettings.MinimumFanSpeed 25
```

**On this R730 these settings didn't fully resolve the fan
ramp-up.** If they work for you, great. If not, see the IPMI
override below.

### Solution: IPMI override via systemd

Bypass iDRAC's thermal logic entirely and pin the fans to 25%
using raw IPMI commands. The override doesn't persist across
reboots, so it gets wrapped in a systemd service that re-applies
it on every boot.

The fan control script,

```bash
cat << 'EOF' > /usr/local/bin/fan-control.sh
#!/bin/bash
# Disable automatic fan control
ipmitool raw 0x30 0x30 0x01 0x00
# Set fans to 25%
ipmitool raw 0x30 0x30 0x02 0xff 0x19
EOF

chmod +x /usr/local/bin/fan-control.sh
```

The systemd service that runs it on every boot,

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

Enable and verify,

```bash
systemctl daemon-reload
systemctl enable fan-control.service
systemctl start fan-control.service
systemctl status fan-control.service
```
