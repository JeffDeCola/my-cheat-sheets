# DUAL BOOT UBUNTU AND WINDOWS

_Some tricks to help you if you run into issues._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SIMPLE INSTALL

I got my image
[here]().

Create a bootable ubuntu using windows rufus.

## AN ISSUE I HAD WITH RST

RST is Intel Rapid Storage Techology,  I needed to change it to AHCI.


CREATE RECOVERY DISK ON THUMBDRIVE
You will use this to do fresh reinstall windows

CREATE BOOTABLE UBUNTU THUMB
use rufus and set to GPT partition scheme.

ENTER BIOS - PRESS F2 METHOD
Enter UEFI (BIOS) through pressing ESC or F2
1) In 'Boot' tab: 'Disable Fastboot'
2) In Security tab set secure boot to disabled
3) In Advanced tab disable Intel Rapid Storage Techology (RST) - Change to AHCI
Press F10 to save & exit

WON'T BOOT
Stick in recover, do fresh install windows.  Uses the above settings.
The goal - RST is not used, but AHCI is.

BOOT UBUNTU
Immediately press ESC or F2 again
In 'Boot' tab: your USB drive should be listed - change the order
Press F10 to save & exit
