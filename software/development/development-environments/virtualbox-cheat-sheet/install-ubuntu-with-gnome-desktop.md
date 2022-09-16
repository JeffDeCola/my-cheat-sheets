# INSTALL UBUNTU WITH GNOME DESKTOP CHEAT SHEET

_Basis steps to install Ubuntu distribution using zsh with a GNOME desktop on VirtualBox._

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)
* [Common ubuntu distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## UBUNTU 22.04 WITH GNOME DESKTOP

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://ubuntu.com/download/desktop)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* CREATE VM
  * Name "VB-Ubuntu-22.04-GNOME"
  * Chose Arch Linux 64-bit (?? MB RAM, 100 GB Disk, .vdi, dynamically allocated)
* ATTACH IMAGE
  * Settings->Storage with Controller: IDE
  * Attach .iso image
* DISPLAY
  * Settings->Display->Screen
  * Video Memory: 128MB
  * Graphics Controller: VBoxSVGA
  * Enable 3D Acceleration enabled
  * Scale Factor 200%

**START VM**

* **START VM**
* SELECT BOOT
  * Select "Arch Linux install medium (x86_64, BIOS)"
* CHECK NETWORK  
  * `ping.google.com`
