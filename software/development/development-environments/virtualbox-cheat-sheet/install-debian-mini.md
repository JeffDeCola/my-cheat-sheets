# INSTALL DEBIAN MINI CHEAT SHEET

_Basis steps to install Debian distribution using zsh without a desktop on VirtualBox._

Documentation and Reference

* [Common debian distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## DEBIAN 11 MINI

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://www.debian.org/distrib/)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* CREATE VM
  * Name "VB-Debian-11-Mini"
  * Chose Debian 64-bit (2048 MB RAM, 21.07 GB Disk, .vdi, dynamically allocated)
* VM SETTINGS
  * Attach .iso image in Settings -> Storage

**START VM**

* START VM
  * Use "Install" (non graph)

**Q & A**
  
* ANSWER QUESTIONS  
  * root password
  * Use "debian11.com"
  * User Jeff DeCola
  * Host "VB-Debian-11-Mini"
  * timezones'
  * User entire disk as partition
  * Make sure you don;t pick a desktop
  * etc...
* **CLOSE VM**

**VIRTUALBOX - REMOVE .iso IMAGE**

* VM SETTINGS  
  * Remove image in Settings -> Storage

**VIRTUALBOX - NETWORK**

* SETTING MENU
  * Devices->Shared Clipboard->Bidirectional
  NETWORKING - CHANGE NAT TO BRIDGE
  * Settings -> Network - Adapter -> Bridged Adapter

**FIRST LOGIN AS JEFF**

* LOGIN
  * Login as jeff

**FIRST BOOT & CONFIGURE**

* LOGIN
  * Login as jeff
* CHECK NETWORK  
  * `ping.google.com`
* INSTALL sudo & Zsh
  * `pacman -S sudo zsh`??
* ADD GROUPS
  * `sudo nano /etc/sudoers`??
  * add "jeff ALL=(ALL) ALL"??
  * uncomment "%sudo ALL =(ALL:ALL) ALL"??
  * uncomment "%wheel ALL=(ALL:ALL) ALL"??
  * `usermod -aG wheel jeff`??
  * `usermod -aG vboxsf jeff`?
  * check with `groups jeff`??
* PROMPT
  * nano .zshrc add `PS1="%F{green}%n@%m:%F{cyan}%1~ %F{white}$"`  ??
* UPDATE
  * `sudo pacman -Syu`??