# WSL2 - BASH ON UBUNTU ON WINDOWS CHEAT SHEET

_`wsl2 allows you to install linux on a windows machine._

Table of Contents,

* [WSL2 (WINDOWS SUBSYSTEM FOR LINUX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#wsl2-windows-subsystem-for-linux)
* [INSTALL UBUNTU ON WINDOWS FROM APP STORE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#install-ubuntu-on-windows-from-app-store)
* [FILE LOCATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#file-locations)
* [ADD SUDO FOR USER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#add-sudo-for-user)
* [FIX DIRECTORY COLOR IF HARD TO READ](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#fix-directory-color-if-hard-to-read)
* [CONFIGURE SSH (PORT 2222) ON WSL2](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#configure-ssh-port-2222-on-wsl2)
* [SETUP CODE DEVELOPMENT ENVIRONMENT ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#setup-code-development-environment-on-windows)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## WSL2 (WINDOWS SUBSYSTEM FOR LINUX)

The WSL2 (Windows System for Linux) allows you to run linux on Windows.

Ubuntu runs on top of windows with its own /home folder.

## INSTALL UBUNTU ON WINDOWS FROM APP STORE

A popular distro of linux is ubuntu.  People refer to this
as `Bash on Ubuntu on Windows`.

Turn on Developer Mode,

* `Click Start->Settings->Privacy & Security->For developers`

Add Windows Subsystem for Linux,

* Press Windows key + R then type: optionalfeatures.exe then hit Enter
* Check "Windows Subsystem for Linux"
* Check "Virtual Machine Platform"
* Restart

Bios settings,

* Must have virtualization  enabled in bios

Update WSL2 Linux Kernal,

* May have to update WSL2 Linux Kernal. Download
  [here](https://learn.microsoft.com/en-us/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package)

From Powershell, set wsl2 as your default version,

```txt
`wsl --set-default-version 2`
```

Install from APP Store

* Go to the windows App store
  [here](https://www.microsoft.com/en-us/p/ubuntu/9nblggh4msv6?activetab=pivot%3aoverviewtab)
  and install.

Check your version,

```bash
lsb_release -a
```

If you already have it installed and want to update to latest version,

```bash
sudo do-release-upgrade
```

I would make a restore point and System Image first before doing any
update like this.  I've learned to never really trust Windows.

## FILE LOCATIONS

Where does Windows keep the files?

For Ubuntu 14.04/16.04,

```txt
C:\Users\<WindowsNAME>\AppData\Local\lxss\home\<bashusername>
```

For Ubuntu 18.04+ (From Windows Store),

```txt
C:\Users\<WindowsNAME>\AppData\Local\Packages\<SOMETHING>\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\<bashusername>
```

For Ubuntu 20.04+ it's under "linux folder".

Linux can work with windows files, but windows CAN NOT work with linux files.

## ADD SUDO FOR USER

Check if your in sudo group,

```bash
groups
```

If not add,

```bash
sudo adduser jeff sudo
```

If not run visudo,

```bash
sudo visudo
```

Add,

```txt
jeff    ALL=(ALL:ALL) ALL
jeff    ALL = NOPASSWD : ALL
```

## FIX DIRECTORY COLOR IF HARD TO READ

Place in .bashrc,

```bash
LS_COLORS='ow=01;36;40'
export LS_COLORS
```

For more information, refer to my
[LS_COLORS cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ls_colors-cheat-sheet)

## CONFIGURE SSH (PORT 2222) ON WSL2

The WSL version of Ubuntu uses the old init.d style scripts
for most services. So we'll use service ssh start,
restart, status, stop, etc....

If you don't have ssh keys, make them,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github ADDIE-PC"
ssh-keygen - A # Generates RSA DSA ECDSA ED25519
ssh-add
```

If having issue with ssh-add, you may need to,

```bash
eval `ssh-agent -s`
ssh-add
```

Check if you have ssh installed,

```bash
sudo service ssh status
```

If it is not recognized, install it,

```bash
sudo apt install openssh-server
```

Edit `/etc/ssh/sshd_config` to add port and listen to any adapter,

```bash
sudo nano /etc/ssh/sshd_config
Port 2222
ListenAddress 0.0.0.0
PubkeyAuthentication no
PasswordAuthentication yes
```

Start service,

```bash
sudo service ssh restart
sudo service ssh status
```

How to start service without asking for a password.

```bash
sudo sh -c "echo '${USER} ALL=(root) NOPASSWD: /usr/sbin/service ssh start' >/etc/sudoers.d/service-ssh-start"
```

Test it starts without asking for password,

```bash
sudo service ssh start
```

Quick ssh test from local machine using powershell,

```bash
ssh -p 2222 localhost
```

Now let's make other machines on the network to ssh into wsl2.
This is because the network interface we see in wsl2 is a virtual
interface that does not match the physical interface that Windows manages.

Step one - Create firewall rule to allow incoming traffic on port 2222.
I used a powershell but you can you `Windows Defender Firewall` GUI,

```bash
New-NetFirewallRule -Name sshd -DisplayName 'Jeff OpenSSH Server (sshd) for WSL' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 2222
```

You check via `Windows Defender Firewall` rules on Windows GUI.

Step 2 - Now route incoming traffic on the physical interface to the WSL interface
via a power forwarding rule. I use a static IP to make this easier.
Again, you can use powershell or `Windows Defender Firewall` GUI,

```bash
netsh interface portproxy delete v4tov4 listenaddress=0.0.0.0 listenport=2222 protocol=tcp
netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=2222 connectaddress=192.168.10.122 connectport=2222
```

You check via `Windows Defender Firewall` rules on Windows GUI.

Done, now test your ssh from another machine on the network,

```bash
ssh -v -p 2222 user@<HOSTNAME>
ssh -v -p 2222 jeff@192.168.20.122
```

## SETUP CODE DEVELOPMENT ENVIRONMENT ON WINDOWS

Since you now have linux running on Windows, why don't you set up
a Code Development Environment.

Check out my cheat sheet on how to setup
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
on Windows with go.
