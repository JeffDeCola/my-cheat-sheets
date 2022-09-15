# INSTALL ARCH LINUX CHEAT SHEET

_Basis steps to install Arch Linux using zsh without a desktop on VirtualBox._

## ARCH LINUX MINI

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://archlinux.org/)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* CREATE VM
  * Name "VB-Arch-Linux-Mini"
  * Chose ARch Linux 64-bit (2048 MB RAM, 21.07 GB Disk, .vdi, dynamically allocated)
* VM SETTINGS
  * Attach .iso image in Settings -> Storage

**START VM**

* START VM
  * Use "Boot Arch Linux (x86_64)"
* CHECK NETWORK  
  * `ping.google.com`

**SETUP DISK PARTITION & MOUNT**

* PARTITION DISK
  * `fdisk /dev/sda`
    * n -- new partition
    * p -- primary disk
    * 1 -- partition number
    * set first sector
    * set last sector (use whole disk)
    * w -- write partition table and quit
  * Check with `fdisk -l`
* FORMAT NEW DISK sda1
  * `mkfs.ext4 /dev/sda1`
* MOUNT NEW DISK
  * `mount /dev/sda1 /mnt`

**INSTALL ARCH LINUX BASE SYSTEM**

* INSTALL BASE PACKAGES ON /mnt
  * `pacstrap /mnt base base-devel linux linux-firmware linux-headers`
* INSTALL OTHER PACKAGES ON /mnt  
  * `pacstrap /mnt netctl dhcpcd dhclient`
* GENERATE FILESYSTEM TABLE ON /mnt
  * `genfstab -U /mnt >>/mnt/etc/fstab`

**CONFIGURE SETTINGS ON NEW DRIVE**

* CHANGE ROOT TO NEW SYSTEM
  * `arch-chroot /mnt`
* INSTALL NANO
  * `pacman -S nano`  
* LANGUAGE
  * `nano /etc/locale.gen`, uncomment en_US.UTF-8 UTF-8
  * `locale-gen`
  * `nano /etc/locale.conf`, add LANG=en_US.UTF-8
* TIMEZONE
  * `ls /usr/share/zoneinfo`
  * `ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime`
  * `hwclock --systohc --utc`
* HOSTNAME
  * `nano /etc/hostname`
  * "VB-Arch-Linux-Mini"
* CHANGE ROOT PASSWORD
  * `passwd`

**ENABLE SERVICES ON NEW DRIVE**

* ENABLE CLIENT DHCP SERVICE
  * `systemctl enable dhcpcd`
* ENABLE SSH SERVICE
  * `pacman -S openssh`
  * `systemctl enable sshd.service`
* ENABLE GUEST ADDITIONS SERVICE (NO DESKTOP)
  * `pacman -S virtualbox-guest-utils-nox`
  * `systemctl enable vboxservice`
* ENABLE SWAP SERVICE
  * `pacman -S systemd-swap`
  * `systemctl enable systemd-swap`

**BOOT LOADER & EXIT**

* INSTALL GRUB (A BOOT LOADER)
  * `pacman -S grub os-prober`
  * `grub-install --target=i386-pc /dev/sda` (correct target even for 64-bit)
  * `grub-mkconfig -o /boot/grub/grub.cfg`
* EXIT & UMOUNT
  * `exit`
  * `umount -R /mnt`
* **CLOSE VM**

**VIRTUALBOX - REMOVE .iso IMAGE**

* VM SETTINGS  
  * Remove image in Settings -> Storage

**FIRST BOOT & CONFIGURE**

* LOGIN
  * Login as root
* INSTALL sudo & Zsh
  * `pacman -S sudo zsh`
* ADD USER USING ZSH
  * `useradd -m -g users -s /usr/bin/zsh jeff`
  * `passwd jeff`
  * `sudo nano /etc/sudoers`
    * add "jeff ALL=(ALL) ALL"
    * uncomment "%sudo ALL =(ALL:ALL) ALL"
    * uncomment "%wheel ALL=(ALL:ALL) ALL"
  * `usermod -aG wheel jeff`
  * `usermod -aG vboxsf jeff`
  * check with `groups jeff`
* **CLOSE VM**

**VIRTUALBOX - NETWORK**

* SETTING MENU
  * Devices->Shared Clipboard->Bidirectional
  NETWORKING - CHANGE NAT TO BRIDGE
  * Settings -> Network - Adapter -> Bridged Adapter

**FIRST LOGIN AS JEFF**

* LOGIN
  * Login as jeff
* PROMPT
  * nano .zshrc add `PS1="%F{green}%n@%m:%F{cyan}%1~ %F{white}$"`  
* UPDATE
  * `sudo pacman -Syu`
* CHECK SERVICES RUNNING
  * `systemctl status vboxservice.service`
  * `systemctl status dhcpcd.service`
  * `systemctl status sshd.service`
  * `systemctl status systemd-swap.service`
