# INSTALL UBUNTU WITH GNOME DESKTOP CHEAT SHEET

_Basis steps to install Ubuntu distribution using bash with a GNOME desktop on VirtualBox._

Table of Contents,

* [UBUNTU 22.04 LTS WITH GNOME DESKTOP](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md#ubuntu-2204-lts-with-gnome-desktop)
* [OPTIONAL INSTALLS & CONFIGURATIONS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-ubuntu-with-gnome-desktop.md#optional-installs--configurations)

Documentation and Reference

* [Install arch linux mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-arch-linux-mini.md)
* [Install debian mini](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-debian-mini.md)
* [Install windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/virtualbox-cheat-sheet/install-windows.md)
* [Common ubuntu distros](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/common-distributions-cheat-sheet)

## UBUNTU 22.04 LTS WITH GNOME DESKTOP

**DOWNLOAD .iso IMAGE**

* GET .iso IMAGE
  * Download image from [official website](https://ubuntu.com/download/desktop)

**VIRTUALBOX - CREATE NEW VM AND ATTACH .iso IMAGE**  

* NEW VM
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

**SHARED SETTINGS**

* **CLOSE VM**
* VM MENU - DRAG AND DROP
  * Devices->Shared ClipBoard->Bidirectional
  * Devices->Drag and Drop->Bidirectional
* CREATE SHARED FOLDER ON WINDOWS
  * Create shared folder on windows
* VM MENU - CONFIGURE WINDOWS SHARED FOLDER
  * Settings->Shared Folders
    * Add folder "VB-Ubuntu-2204-GNOME"
    * Check Auto-mount
* **START VM**
* MAKE SURE YOUR PART OF vboxsf GROUP
  * `groups`
  * `sudo usermod -a -G vboxsf jeff`
* SHARED FOLDER IS HERE
  * `cd /media/sf_VB-Ubuntu-2204-GNOME`
* CREATE SYMBOLIC LINK IN YOUR HOME DIRECTORY
  * `sudo ln -sf /media/sf_VB-Ubuntu-2204-GNOME /home/jeff/shared`

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
  * `ssh-keygen -t rsa -b 4096 -C "Keys for Github (VB-Ubuntu-2204-GNOME)"`
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
  * `git config --global user.name "Jeff DeCola (VB-Ubuntu-2204-GNOME)"`
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
  * I like to use [this](https://github.com/jimeh/git-aware-prompt)
  * `mkdir ~/.bash`
  * `cd ~/.bash`
  * `git clone https://github.com/jimeh/git-aware-prompt.git`
* EDIT .bashrc
  * `export GITAWAREPROMPT=~/.bash/git-aware-prompt`
  * `source "${GITAWAREPROMPT}/main.sh"`
  * `PS1="\${debian_chroot:+(\$debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\W\[\033[00m\] \[$txtcyn\]\$git_branch\[$txtred\]\$git_dirty\[$txtrst\]\$ "`

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
  
**ADD A PRINTER**

Pain in the butt and may not work anymore,

* SHARED NETWORK
  * Setup a shared network printer on Windows.
  * Must also use a private network.
* INSTALL
  * `sudo apt-get install lsb`
  * `sudo apt-get update`
  * `sudo apt-get install python3-smbc`
  * `sudo apt-get install smbclient`
* PRINTER DRIVER
  * Get printer driver from their website, usually a .deb file
* FIND PRINTER  
  * SEarch on your network (e.g. 192.168.1.x)
  * e.g. smb://192.168.1.115/SHARE-EPSON-XP-630-Series
* CONFIG  
  * Also need Username, Workgroup and password
