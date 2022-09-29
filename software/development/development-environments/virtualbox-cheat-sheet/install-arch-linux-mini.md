# INSTALL ARCH LINUX CHEAT SHEET

_Basis steps to install Arch Linux using zsh without a desktop on VirtualBox._

Documentation and Reference

* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install ubuntu with GNOME desktop](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)

## ARCH LINUX MINI

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://archlinux.org/)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* NEW VM
  * Name "VB-Arch-Linux-Mini"
  * Chose Arch Linux 64-bit (2048 MB RAM, 20 GB Disk, .vdi, dynamically allocated)
* ATTACH IMAGE
  * Settings->Storage with Controller: IDE
  * Attach .iso image
* DISPLAY
  * Settings->Display->Screen
  * Video Memory: 128MB
  * Graphics Controller: VBoxSVGA
  * Enable 3D Acceleration enabled
  * Scale Factor 200%
* SET BRIDGE
  * The VM will receive it's own IP address if DHCP is enabled in the network
  * Settings -> Network -> Adapter 1
    * `Bridged Adapter`
    * `Realtek Gaming GbE (GIGabit Ethernet) Family Controller`

**START VM**

* **START VM**
* SELECT BOOT
  * Select "Arch Linux install medium (x86_64, BIOS)"
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
* INSTALL VI
  * `pacman -S vi`
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
    * /etc/hostname should not contain comments or empty lines.
  * `nano /etc/hosts`
    * #ip-address       hostname.domain.org     hostname"
    * "127.0.0.1        localhost.localdomain   VB-Arch-Linux-Mini localhost"
    * "::1              localhost.localdomain   VB-Arch-Linux-Mini localhost"
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

**SHARED SETTINGS**

* VM MENU
  * Devices->Shared Folders->Shared Folder Settings
    * Pick where you want this folder
  * (NOT AVAILABLE) Devices->Shared ClipBoard->Bidirectional
  * (NOT AVAILABLE) Devices->Drag and Drop->Bidirectional

**VIRTUALBOX - REMOVE .iso IMAGE**

* **CLOSE VM**
* VM SETTINGS  
  * Remove image in Settings -> Storage
* **START VM**

**FIRST BOOT AS ROOT & CONFIGURE**

* LOGIN
  * Login as root
* INSTALL sudo & Zsh
  * `pacman -S sudo zsh`
* ADD jeff USING ZSH
  * `useradd -m -g users -s /usr/bin/zsh jeff`
  * `passwd jeff`
* ADD USER TO SUDOERS
  * `EDITOR=vi visudo`
    * add "jeff ALL=(ALL) ALL"
    * uncomment "%sudo ALL =(ALL:ALL) ALL"
    * uncomment "%wheel ALL=(ALL:ALL) ALL"
    * Check file with `EDITOR=vi visudo -C`
* ADD TO GROUPS
  * `usermod -aG wheel jeff`
  * `usermod -aG vboxsf jeff`
  * check with `groups jeff`
* **CLOSE VM**

**FIRST LOGIN AS JEFF & CONFIGURE**

* **START VM**
* LOGIN
  * Login as jeff
* Q & A FOR ZSH
  * Since it's first time you can go through Q & A  
* CONFIGURE PROMPT
  * nano .zshrc add `prompt="%F{green}%n@%m:%F{cyan}%1~ %F{white}$ "`  
* UPDATE
  * `sudo pacman -Syu`
* CHECK SERVICES RUNNING
  * `systemctl status vboxservice.service`
  * `systemctl status dhcpcd.service`
  * `systemctl status sshd.service`
  * `systemctl status systemd-swap.service`
* CHECK ETH CONNECTIONS
  * Note: arch linux uses ip tools rather than net-tools
  * `ip addr`
* CHECK HOSTNAME
  * `sudo pacman -S inetutils`
  * `hostname`

**YOUR HOME NETWORK**

* BRIDGE MODE
  * Since we are in bridge mode, I like to configure my home router to set the same ip address

## OPTIONAL INSTALLS & CONFIGURATIONS

**CONNECT TO GITHUB AND GET YOUR REPOS**

* SSH INTO VM
  * It is easier to ssh into the box to copy paste commands
  * From another computer `ssh <ip>`
* CREATE KEYS
  * `ssh-keygen -t rsa -b 4096 -C "Keys for Github (VB-Arch-Linux-Mini)"`
* ADD KEYS TO SSH AUTH AGENT  
  * `ssh-add ~/.ssh/id_rsa`
  * If agent not running `eval "$(ssh-agent)"`
  * Check `ssh-add -L`
* ADD PUBLIC KEY TO GITHUB
  * Copy/Paste public key (.ssh/id_rsa.pub) at github
* CONNECT TO GITHUB
  * `ssh -T git@github.com`
* INSTALL GIT
  * `sudo pacman -S git`
* GIT CONFIGURATION SETTINGS
  * `git config --global user.name "Jeff DeCola (VB-Arch-Linux-Mini)"`
  * `git config --global user.email <YOUR_EMAIL>`
  * `git config --global core.editor nano`
  * `git config --global push.default simple`
  * `git config --global pull.rebase false`
  * Check with `git config --list`
* CLONE REPO
  * `mkdir development`
  * `cd development`
  * `git clone git@github.com:JeffDeCola/<REPO NAME>.git`

**GIT AWARE PROMPT**

* INSTALL
  * I like to use [this](https://github.com/joeytwiddle/git-aware-prompt) for zsh
  * `mkdir ~/.bash`
  * `cd ~/.bash`
  * `git clone https://github.com/jimeh/git-aware-prompt.git`
* EDIT .bashrc
  * `export GITAWAREPROMPT=~/.bash/git-aware-prompt`
  * `source "${GITAWAREPROMPT}/main.sh"`
  * `PROMPT='%F{green}%n@%m:%F{cyan}%1~%{$txtcyn%}$git_branch%{$txtred%}$git_unknown_count%{$txtrst%} %F{white}$ '`

**SSH LOGIN VIA KEYS NOT PASSWORD**

* EDIT sshd_config
  * `sudo nano /etc/ssh/sshd_config`
  * Port 22
  * PubkeyAuthentication yes
  * PasswordAuthentication no
  * AuthorizedKeysFile .ssh/authorized_keys
* RESTART SERVICE  
  * `sudo systemctl restart sshd.service`
  * `sudo systemctl status sshd.service`
* AUTHORIZED KEYS
  * If you want to ssh into this machine
  * Add public keys from other hosts in `.ssh/authorized_keys`
  