# WSL2 - BASH ON UBUNTU ON WINDOWS CHEAT SHEET

_`wsl2 allows you to install linux on a windows machine._

Table of Contents,

* [WSL2 (WINDOWS SUBSYSTEM FOR LINUX)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#wsl2-windows-subsystem-for-linux)
* [INSTALL UBUNTU ON WINDOWS FROM APP STORE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#install-ubuntu-on-windows-from-app-store)
* [FILE LOCATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#file-locations)
* [ADD SUDO FOR USER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#add-sudo-for-user)
* [FIX DIRECTORY COLOR IF HARD TO READ](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#fix-directory-color-if-hard-to-read)
* [CONFIGURE SSH (NORMAL PORT 22) ON WSL2 (EASY WAY)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#configure-ssh-normal-port-22-on-wsl2-easy-way)
  * [START SSH SERVER ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#start-ssh-server-on-windows)
  * [MAKE WINDOWS OPENSSH TO OPEN WSL2 BASH SHELL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#make-windows-openssh-to-open-wsl2-bash-shell)
  * [CONFIGURE SSHD ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#configure-sshd-on-windows)
  * [WE ARE USING WINDOWS .ssh DIRECTORY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#we-are-using-windows-ssh-directory)
  * [TEST IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#test-it)
* [CONFIGURE SSH (PORT 2222) ON WSL2 (TOO MUCH WORK)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#configure-ssh-port-2222-on-wsl2-too-much-work)
  * [MAKE SSH KEYS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#make-ssh-keys)
  * [INSTALL OPENSSH-SERVER IN WSL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#install-openssh-server-in-wsl)
  * [CONFIGURE SSH PORT NUMBER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#configure-ssh-port-number)
  * [START SSH SERVICE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#start-ssh-service)
  * [FORWARD PORTS INTO WSL2](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#forward-ports-into-wsl2)
  * [TEST](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/windows/wsl2-bash-on-ubuntu-on-windows-cheat-sheet#test)
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

## SSH SERVER

How to ssh into/from wsl2.

### METHOD 1 - USE WINDOWS SSH SERVER

Let's make this simple and use windows ssh server.
So we will be using windows ssh keys and authentification.

Then we'll just change windows shell for ssh to wsl2 bash shell.

FILES LOCATIONS OVERVIEW

* C:\Programdata\ssh
  * sshd_config
* C:\Users\jeffry\.ssh
  * authorized_keys
  
* /home/jeff/.ssh
  * config
  * keys
  * known_hosts

#### START SSH SERVER ON WINDOWS

Open an admin PowerShell prompt to see if you have the server running,

```bash
Get-WindowsCapability -Online | ? Name -like 'OpenSSH*'
```

If OpenSSH.Server NOTPresent, add it,

```bash
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
```

Start service,

```bash
Start-Service sshd
```

Start service automatically,

```bash
Set-Service -Name sshd -StartupType 'Automatic'
```

Check if it's running,

```bash
 get-service sshd
```

#### MAKE WINDOWS OPENSSH USE WSL2 BASH SHELL

This is the magic.
Update the registry via powershell or you could use regedit,

```bash
New-ItemProperty -Path "HKLM:\SOFTWARE\OpenSSH" -Name DefaultShell -Value "C:\WINDOWS\System32\bash.exe" -PropertyType String -Force
```

#### CONFIGURE SSHD ON WINDOWS

```bash
start-process notepad C:\Programdata\ssh\sshd_config
```

You may do things like, I would leave authorized_keys commented so it's in programdata.

```bash
Port 22
PubkeyAuthentication yes
PasswordAuthentication no
# AuthorizedKeysFile .ssh/authorized_keys
```

Restart,

```bash
restart-service sshd
get-service sshd
```

#### ADD THE KEYS YOU WANT TO USE

Place keys in C:\Programdata\ssh

```bash
Get-Service ssh-agent
set-service ssh-agent StartupType ‘Automatic’
Start-Service ssh-agent
ssh-add C:\Users\jeffry\.ssh\id_rsa
ssh-add C:\Users\jeffry\.ssh\id_ecdsa
```

If you have issue, try StartType as manual,

```bash
Get-Service -Name ssh-agent | Set-Service -StartupType Manual
Get-Service ssh-agent | Select StartType
```

Check,

```bash
ssh-add -L
```

#### WE ARE USING WINDOWS .ssh DIRECTORY

We are not using WSL2 .ssh directory, we are using windows located in
`Users\jeffry\.ssh`.

#### ADD AUTHORIZED KEYS FROM OTHER MACHINES

```bash
start-process notepad C:\Users\jeffry\.ssh\authorized_keys
```

#### TEST IT

Now test your ssh from another machine on the network and use your windows
username and password. NOT wsl2.

```bash
ssh -v -p 22 user@<HOSTNAME>
ssh -v -p 22 jeffry@192.168.20.122
```

### METHOD 2 - USE WSL2 SSH SERVER (TOO MUCH WORK)

**NOTE: After thinking about it, this is too much work, especially since the**
**IP of the WSL will change on reboot.**

Windows only forwarding ports, and uses WSL2's Linux OpenSSH and
authenticates against Linux.

The WSL version of Ubuntu uses the old init.d style scripts
for most services. So we'll use service ssh start,
restart, status, stop, etc....

#### MAKE SSH KEYS

If you don't have ssh keys, make them,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github (TK2-PC Bash on Ubuntu on Windows)"
ssh-add ~/.ssh/id_rsa
```

If having issue with ssh-add, you may need to,

```bash
eval `ssh-agent -s`
ssh-add
```

#### INSTALL OPENSSH-SERVER IN WSL

Check if you have ssh installed,

```bash
sudo service ssh status
```

If it is not recognized, install it,

```bash
sudo apt install openssh-server
```

#### CONFIGURE SSH PORT NUMBER

Edit `/etc/ssh/sshd_config` to add port and listen to any adapter,

```bash
sudo nano /etc/ssh/sshd_config
Port 2222
ListenAddress 0.0.0.0
PubkeyAuthentication no
PasswordAuthentication yes
```

#### START SSH SERVICE

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

#### FORWARD PORTS INTO WSL2

Now let's make other machines on the network to ssh into wsl2.
This is because the network interface we see in wsl2 is a virtual
interface that does not match the physical interface that Windows manages.

**Step 1** - Create firewall rule to allow incoming traffic on port 2222.
I used a powershell but you can you `Windows Defender Firewall` GUI,

```bash
New-NetFirewallRule -Name sshd -DisplayName 'Jeff OpenSSH Server (sshd) for WSL' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 2222
```

You can remove it using,

```bash
Remove-NetFireWallRule -DisplayName 'Jeff OpenSSH Server (sshd) for WSL';
```

You can check via `Windows Defender Firewall` rules on Windows GUI.

**Step 2** - Now route incoming traffic on the physical interface to the WSL interface
via a power forwarding rule. I use a static IP to make this easier.
Again, you can use powershell or `Windows Defender Firewall` GUI,

First get ip of your WSl

```bash
ip addr
```

Then use that IP for connectaddress,

```bash
netsh interface portproxy delete v4tov4 listenaddress=0.0.0.0 listenport=2222 protocol=tcp
netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=2222 connectaddress=172.31.206.81 connectport=2222
netsh interface portproxy show v4tov4;
```

NOTE: **This ip changes on reboot, so you need script or do this every reboot**

You check via `Windows Defender Firewall` rules on Windows GUI.

#### TEST

Now test your ssh from another machine on the network,

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
