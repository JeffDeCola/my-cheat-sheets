# INSTALL DEBIAN MINI CHEAT SHEET

_Basis steps to install Debian distribution using bash without a desktop on VirtualBox._

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install ubuntu with GNOME desktop](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)
* [Common debian distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## DEBIAN 11 MINI

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://www.debian.org/distrib/)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

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

**START VM**

* **START VM**
* INSTALL
  * Use "Install" (non graph)

**Q & A**
  
* ANSWER QUESTIONS  
  * root password
  * Hostname "VB-Debian-11-Mini"
  * Domain name "debian11.com"
  * User Jeff DeCola
  * Timezone
  * User entire disk as partition
  * Make sure you don't pick a desktop (use space bar to uncheck)
  * etc...

**VIRTUALBOX - REMOVE .iso IMAGE**

* **CLOSE VM**
* VM SETTINGS  
  * Remove image in Settings -> Storage
* **START VM**

**FIRST LOGIN AS JEFF & CONFIGURE AS ROOT**

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

**CONFIGURE AS JEFF**

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

**SHARED SETTINGS**

* VM MENU
  * Devices->Shared Folders->Shared Folder Settings
    * Pick where you want this folder
  * (NOT AVAILABLE) Devices->Shared ClipBoard->Bidirectional
  * (NOT AVAILABLE) Devices->Drag and Drop->Bidirectional

**VIRTUALBOX - REMOVE GUEST ADDITIONS.iso IMAGE**

* **CLOSE VM**
* VM SETTINGS
  * Remove guest additions .iso image in Settings -> Storage
* **START VM**

**YOUR HOME NETWORK**

* BRIDGE MODE
  * Since we are in bridge mode, I like to configure my home router to set the same ip address

## OPTIONAL INSTALLS & CONFIGURATIONS

**CONNECT TO GITHUB AND GET YOUR REPOS**

* SSH INTO VM
  * It is easier to ssh into the box to copy paste commands
  * From another computer `ssh <ip>`
* CREATE KEYS
  * `ssh-keygen -t rsa -b 4096 -C "Keys for Github VB-Debian-11-Mini"`
  * `ssh-add ~/.ssh/id_rsa`
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
  * Check with `git config --list`
* CLONE REPO
  * `mkdir development`
  * `cd development`
  * `git clone git@github.com:JeffDeCola/<REPO NAME>.git`

**GIT AWARE PROMPT**

* INSTALL
  * I like to use [this](https://github.com/jimeh/git-aware-prompt)
  * `mkdir ~/.bash`
  * `cd ~/.bash`
  * `git clone https://github.com/jimeh/git-aware-prompt.git`
* EDIT .bashrc
  * `export GITAWAREPROMPT=~/.bash/git-aware-prompt`
  * `source "${GITAWAREPROMPT}/main.sh"`
  * `PS1="\${debian_chroot:+(\$debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\W\[\033[00m\] \[$txtcyn\]\$git_branch\[$txtred\]\$git_dirty\[$txtrst\]\$ "`
