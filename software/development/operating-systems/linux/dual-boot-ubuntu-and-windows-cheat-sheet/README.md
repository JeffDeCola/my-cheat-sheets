# DUAL BOOT UBUNTU & WINDOWS

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Installing Ubuntu on PC that already has Windows installed._

Table of Contents

* [INSTALL USING A BOOTABLE THUMB DRIVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dual-boot-ubuntu-and-windows-cheat-sheet#install-using-a-bootable-thumb-drive)
* [AN ISSUE I HAD WITH INTEL RST](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dual-boot-ubuntu-and-windows-cheat-sheet#an-issue-i-had-with-intel-rst)

## INSTALL USING A BOOTABLE THUMB DRIVE

I got my Ubuntu from
[here](https://ubuntu.com/download/desktop).

I created a bootable ubuntu using windows
[rufus](https://rufus.ie/)
(A free application for Windows that can be used to format
and create bootable USB thumb drives)
The MBR partition scheme should work.

You may have to enter your bios to change the boot order.

## AN ISSUE I HAD WITH INTEL RST

If your computer uses RST (Intel Rapid Storage Technology), you
may need to change to AHCI (Advanced Host Controller Interface).
The **trick is to reinstall Windows with AHCI**.  Don't worry, it's not that hard.

First you need two USB thumb drives,

* Create a `Windows Recovery Disk` on a USB thumb drive
on the computer you want to dual boot on. Search for `recovery drive`
in Windows and follow the steps
* Create your `bootable Ubuntu USB thumb drive`
using rufus (see above) with the GPT partition scheme

Second, depending on your BIOS, you may need to,

* Disable `Fastboot`
* Set `Secure Boot` to disabled
* Change Intel Rapid Storage Technology (RST) to AHCI

Third, Windows probably won't boot, so put in your Windows Recovery
Drive and do a fresh install.
**Now AHCI is being used, not RST.**  This was the trick.

Last, after you have Windows reinstalled (using AHCI),

* Put in `bootable Ubuntu USB thumb drive`
* Reboot
* Enter BIOS
* Choose Ubuntu to boot
* Save exit

Now you should be able to dual boot install Ubuntu on your rig.  Congrats.
