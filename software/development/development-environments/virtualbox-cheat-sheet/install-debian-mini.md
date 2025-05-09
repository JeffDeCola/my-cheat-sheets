# INSTALL DEBIAN MINI CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Basic steps to install a Debian distribution using bash without a desktop on VirtualBox._

Table of Contents

* [INSTALL DEBIAN 11 MINI](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#install-debian-11-mini)
  * [DOWNLOAD .iso IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#download-iso-image)
  * [VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#virtualbox---create-new-vm-and-attach-iso-image)
  * [START VM](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#start-vm)
  * [Q & A](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#q--a)
  * [VIRTUALBOX - REMOVE .iso IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#virtualbox---remove-iso-image)
* [FIRST LOGIN AS JEFF & CONFIGURE AS ROOT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#first-login-as-jeff--configure-as-root)
  * [CONFIGURE AS ROOT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#configure-as-root)
  * [CONFIGURE AS JEFF](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#configure-as-jeff)
* [MORE SETTINGS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#more-settings)
  * [SHARED SETTINGS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#shared-settings)
  * [VIRTUALBOX - REMOVE GUEST ADDITIONS.iso IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#virtualbox---remove-guest-additionsiso-image)
  * [HOME NETWORK](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#home-network)
* [OPTIONAL INSTALLS & CONFIGURATIONS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#optional-installs--configurations)
  * [CONNECT TO GITHUB AND GET YOUR REPOS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#connect-to-github-and-get-your-repos)
  * [GIT AWARE PROMPT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#git-aware-prompt)
  * [SSH LOGIN VIA KEYS NOT PASSWORD](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md#ssh-login-via-keys-not-password)

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install ubuntu with GNOME desktop](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)
* [Common debian distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## INSTALL DEBIAN 11 MINI

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

### DOWNLOAD .iso IMAGE

* GET .iso IMAGE
  * Download image from [official website](https://www.debian.org/distrib/)

### VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE

* NEW VM
  * Name "VB-Debian-11-Mini"
  * Chose Debian 64-bit (2048 MB RAM, 20.00 GB Disk, .vdi, dynamically allocated)
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

### START VM

* **START VM**
* INSTALL
  * Use "Install" (non graph)

### Q & A

* ANSWER QUESTIONS  
  * root password
  * Hostname "VB-Debian-11-Mini"
  * Domain name "debian11.com"
  * User Jeff DeCola
  * Timezone
  * User entire disk as partition
  * Make sure you don't pick a desktop (use space bar to uncheck)
  * etc...

### VIRTUALBOX - REMOVE .iso IMAGE

* **CLOSE VM**
* VM SETTINGS  
  * Remove image in Settings -> Storage
* **START VM**

## FIRST LOGIN AS JEFF & CONFIGURE AS ROOT

### CONFIGURE AS ROOT

* **START VM**
* LOGIN
  * Login as jeff
* SU ROOT
  * `su -` (- gives you an environment)
* INSTALL sudo
  * `apt-get -S sudo`
* ADD jeff to sudo
  * `adduser jeff sudo`
  * check with `groups jeff`
* EXIT ROOT
  * `exit`

### CONFIGURE AS JEFF

* CONFIGURE PROMPT
  * nano .bashrc uncomment `force_color_prompt=yes"`  
* UPDATE
  * `sudo apt update`
  * `sudo apt-get upgrade`
* CHECK GUEST ADDITIONS INSTALLED (NO DESKTOP)
  * check `lsmod | grep vboxguest`
  * Not sure how it got installed
* INSTALL GUEST ADDITIONS (NO DESKTOP) (I DID NOT CHECK THIS)
  * `sudo mkdir -p /mnt/guestadditions`
  * `sudo mount /dev/cdrom /mnt/guestadditions`
  * `cd /mnt/guestadditions`
  * `sudo ./VBoxLinuxAdditions.run`
  * `reboot`
  * check `lsmod | grep vboxguest`
* ENABLE SSH SERVICE
  * `sudo apt-get install openssh-server`
* CHECK SERVICES RUNNING
  * `systemctl status sshd.service`
* CHECK ETH CONNECTIONS
  * Note: debian uses ip tools rather than net-tools
  * `ip addr`
* CHECK HOSTNAME
  * `hostname`

## MORE SETTINGS

### SHARED SETTINGS

* **CLOSE VM**
* VM MENU - DRAG AND DROP
  * (NOT AVAILABLE) Devices->Shared ClipBoard->Bidirectional
  * (NOT AVAILABLE) Devices->Drag and Drop->Bidirectional
* CREATE SHARED FOLDER ON WINDOWS
  * Create shared folder on windows
* VM MENU - CONFIGURE WINDOWS SHARED FOLDER
  * Settings->Shared Folders
    * Add folder "VB-Debian-11-Mini"
    * Check Auto-mount
* **START VM**
* MAKE SURE YOUR PART OF vboxsf GROUP
  * `groups`
  * `sudo usermod -a -G vboxsf jeff`
* SHARED FOLDER IS HERE
  * `cd /media/sf_VB-Debian-11-Mini`
* CREATE SYMBOLIC LINK IN YOUR HOME DIRECTORY
  * `sudo ln -sf /media/sf_VB-Debian-11-Mini /home/jeff/shared`

### VIRTUALBOX - REMOVE GUEST ADDITIONS.iso IMAGE

* **CLOSE VM**
* VM SETTINGS
  * Remove guest additions .iso image in Settings -> Storage
* **START VM**

### HOME NETWORK

* BRIDGE MODE
  * Since we are in bridge mode, I like to configure my home router to set the same ip address

## OPTIONAL INSTALLS & CONFIGURATIONS

### CONNECT TO GITHUB AND GET YOUR REPOS

* SSH INTO VM
  * It is easier to ssh into the box to copy paste commands
  * From another computer `ssh <ip>`
* CREATE KEYS
  * `ssh-keygen -t rsa -b 4096 -C "Keys for Github (VB-Debian-11-Mini)"`
* ADD KEYS TO SSH AUTH AGENT  
  * `ssh-add ~/.ssh/id_rsa`
  * If agent not running `eval "$(ssh-agent)"`
  * Check `ssh-add -L`
* ADD PUBLIC KEY TO GITHUB
  * Copy/Paste public key (.ssh/id_rsa.pub) at github
* CONNECT TO GITHUB
  * `ssh -T git@github.com`
* INSTALL GIT
  * `sudo apt-get install git`
* GIT CONFIGURATION SETTINGS
  * `git config --global user.name "Jeff DeCola (VB-Debian-11-Mini)"`
  * `git config --global user.email <YOUR_EMAIL>`
  * `git config --global core.editor nano`
  * `git config --global push.default simple`
  * `git config --global pull.rebase false`
  * Check with `git config --list`
* CLONE REPO
  * `mkdir development`
  * `cd development`
  * `git clone git@github.com:JeffDeCola/<REPO NAME>.git`

### GIT AWARE PROMPT

* INSTALL
  * I like to use [this](https://github.com/jimeh/git-aware-prompt)
  * `mkdir ~/.bash`
  * `cd ~/.bash`
  * `git clone https://github.com/jimeh/git-aware-prompt.git`
* EDIT .bashrc
  * `export GITAWAREPROMPT=~/.bash/git-aware-prompt`
  * `source "${GITAWAREPROMPT}/main.sh"`
  * `PS1="\${debian_chroot:+(\$debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\W\[\033[00m\] \[$txtcyn\]\$git_branch\[$txtred\]\$git_dirty\[$txtrst\]\$ "`

### SSH LOGIN VIA KEYS NOT PASSWORD

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
