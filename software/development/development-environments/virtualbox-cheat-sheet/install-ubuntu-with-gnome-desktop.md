# INSTALL UBUNTU WITH GNOME DESKTOP CHEAT SHEET

_Basis steps to install Ubuntu distribution using zsh with a GNOME desktop on VirtualBox._

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)
* [Common ubuntu distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## UBUNTU 22.04 LTS WITH GNOME DESKTOP

Sadly, you must type the following in the command line. You can't copy/paste into terminal.
Yup, it stinks.

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://ubuntu.com/download/desktop)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* CREATE VM
  * Name "VB-Ubuntu-2204-GNOME"
  * Chose Ubuntu 64-bit (4096 MB RAM, 100 GB Disk, .vdi, dynamically allocated)
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
  * Select "Install"

**Q & A**
  
* ANSWER QUESTIONS
  * Timezone
  * Hostname "VB-Ubuntu-2204-GNOME"
  * User Jeff DeCola
  * etc...
* **CLOSE VM (WHEN IT ASKED TO RESTART)**

**VIRTUALBOX - REMOVE .iso IMAGE**

* VM SETTINGS  
  * Remove image in Settings -> Storage

**VIRTUALBOX - ATTACH GUEST ADDITIONS.iso IMAGE**

* VM SETTINGS
  * Attach guest additions .iso image in Settings -> Storage

**FIRST LOGIN AS JEFF & CONFIGURE**

* **START VM**
* LOGIN
  * Login as Jeff DeCola
* MORE Q & A  
* BRING UP TERMINAL
* CHECK jeff part of sudo
  * `groups jeff`
* UPDATE
  * `sudo apt update`
  * `sudo apt-get upgrade`
* INSTALL GUEST ADDITIONS
  * `sudo mkdir -p /mnt/guestadditions`
  * `sudo mount /dev/cdrom /mnt/guestadditions`
  * `cd /mnt/guestadditions`
  * `sudo ./VBoxLinuxAdditions.run`
  * `sudo reboot`
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

**DRAG AND DROP SETTINGS**

* VM MENU
  * Devices->Drag and Drop->Bidirectional

**DISPLAY (AUTO RESIZE)**  

* VM MENU
  * This should now be available
  * View->Auto Resize Guest Display
  * Make sure your Host Settings->Display is 200% (Help with native resolution like 4K)

**VIRTUALBOX - REMOVE GUEST ADDITIONS.iso IMAGE**

* **CLOSE VM**
* VM SETTINGS
  * Remove guest additions .iso image in Settings -> Storage
* **START VM**

**YOUR HOME NETWORK**

* BRIDGE MODE
  * Since we are in bridge mode, I like to configure my home router to set the same ip address

## OPTIONAL INSTALLS & CONFIGURATIONS

**INSTALL VS CODE AND SYNC**

* DOWNLOAD AND INSTALL
  * Download vs Code from browser and install
* SYNC TO YOUR GITHUB ACCOUNT  
* COPY/ PAST workspace file from another machine

**CONNECT TO GITHUB AND GET YOUR REPOS**

* SSH INTO VM
  * It is easier to ssh into the box to copy paste commands
  * From another computer `ssh <ip>`
* CREATE KEYS
  * `ssh-keygen -t rsa -b 4096 -C "Keys for Github VB-Ubuntu-2204-GNOME"`
  * `ssh-add ~/.ssh/id_rsa`
* ADD PUBLIC KEY TO GITHUB
  * Copy/Paste public key (.ssh/id_rsa.pub) at github
* CONNECT TO GITHUB
  * `ssh -T git@github.com`
* INSTALL GIT
  * `sudo apt-get install git`
* GIT CONFIGURATION SETTINGS
  * `git config --global user.name "Jeff DeCola (VB-Ubuntu-2204-GNOME)"`
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
