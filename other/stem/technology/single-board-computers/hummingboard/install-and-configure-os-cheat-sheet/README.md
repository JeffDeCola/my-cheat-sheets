# INSTALL AND CONFIGURE OS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Install and configure an OS distribution on your Hummingboard._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#install)
  * [METHOD 1 - FLASH IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#method-1---flash-image)
  * [METHOD 2 - OLD-WAY - "IGNITION"](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#method-2---old-way---ignition)
* [BASIC CONFIGURATION (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#basic-configuration-debian)
* [UPDATE AND UPGRADE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#update-and-upgrade)
  * [CHANGE PASSWORD FOR USER DEBIAN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#change-password-for-user-debian)
  * [RENAME YOUR HUMMINGBOARD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#rename-your-hummingboard)
  * [ADD NEW USER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#add-new-user)
  * [UPDATE PATH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#update-path)
  * [GET GIT & GITHUB WORKING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#get-git--github-working)
* [INSTALL GO & GO TOOLS (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#install-go--go-tools-debian)
* [CONNECT TO WIFI(DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#connect-to-wifi-debian)

Documentation and Reference

* Hummingboard
  [specs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/specifications-cheat-sheet)
  and
  [install](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet)
* Raspberry Pi
  [specs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/raspberry-pi/specifications-cheat-sheet)
  and
  [install](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/technology/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet)
* [Create DNS Server using BIND Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/operating-systems/linux/dns-cheat-sheet/create-dns-server-using-bind.md)

## INSTALL

How to install an OS on your hummingboard.

### METHOD 1 - FLASH IMAGE

Get the images
[here](https://images.solid-run.com/IMX6/).

On windows, I use
[Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
to burn the image on a microSD card.

### METHOD 2 - OLD-WAY - "IGNITION"

The `Ignition` installer from SolidRun lets you easily select
your OS distribution of choice.

To use this, ignition must be burned on a microSD card.

* Download the `Ignition` image
  [here](https://www.solid-run.com/downloads/ignition).
* Get a way to get the image on a microSD
  [Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
* With your computer, use `Win32DiskImager`
  to burn the image on a microSD card.

I would label this microSD card since its easy to forget.

Place the `Ignition` microSD in your hummingboard and boot it up.
You must have your hummingboard connected to the internet
(wired or wireless) to use this software.

`Ignition` will look at
[SolidRun's GitHub Repo](https://github.com/SolidRun/ignition-imx6)
for `i.MX6` OS distribution `ignition installer scripts`.

Choose a distro such as:

* Debian (Manual install
  [here](https://images.solid-build.xyz/IMX6/Debian))
  * Wheezy CLI
  * Wheezy Desktop
  * **Jessie CLI** (I like this one).
  * Jessie Desktop
* GeexBoX
* LibreELEC
* OpenELEC.tv
* ArchLinux
* Fedora-22
* Android-kitkat 4.4.4

Place another microSD card in your hummingboard and install
your image.  Do not write over the `Ignition` microSD card
(like I did).

Note, depending on your hummingboard version, not all OS
distributions will work.

Complete list of `ignition installer scripts` for various OS
[here](https://wiki.solid-run.com/doku.php?id=tag:ignition&do=showtag&tag=ignition).

## BASIC CONFIGURATION (DEBIAN)

This is a short list,

Initial user and password,

```txt
user: debian
pw: debian
```

## UPDATE AND UPGRADE

```bash
sudo apt-get update
sudo apt-get upgrade
```

### CHANGE PASSWORD FOR USER DEBIAN

If you want, change your password for root/debian,

```bash
passwd
```

### RENAME YOUR HUMMINGBOARD

Rename your host,

```bash
sudo nano /etc/hostname
sudo nano /etc/hosts
```

check,

```bash
hostname
```

### ADD NEW USER

For security add a new user,

```bash
sudo adduser jeff
sudo usermod -aG sudo jeff
```

or

```bash
sudo adduser jeff
sudo visudo
pi    ALL=(ALL) ALL
jeff  ALL=(ALL) ALL
```

### UPDATE PATH

If you have an older copy of debian, add sbin
in `.bashrc` file,

```bash
# I also needed to add this
PATH=$PATH:/usr/local/sbin
PATH=$PATH:/usr/sbin
PATH=$PATH:/sbin
```

### GET GIT & GITHUB WORKING

Get git,

```bash
sudo apt install git
git version
```

Create keys,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github (<HOSTNAME/MACHINE NAME>)"
```

I added the comment from where the key came from because it gets confusing.
Place the Raspi public key at github.

I also like to add the hostname/machine name so I know where the commits came from,

Add to auth agent,

```bash
eval "$(ssh-agent)"
ssh-add ~/.ssh/id_rsa
ssh-add -L
```

```bash
git config --global user.name "Jeff DeCola (<HOSTNAME/MACHINE NAME>)"
git config --global user.email <YOUR-EMAIL>
git config --global core.editor nano
git config --global push.default simple
git config --global pull.rebase false
```

Check configuration,

```bash
git config --list
```

Refer to my
[cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/source-version-control/git-cheat-sheet)
for more info.

## INSTALL GO & GO TOOLS (DEBIAN)

I like golang, so here is how you install it,

Let's install `ver 1.12.7`  to `/usr/local`,

```bash
FileName='go1.12.7.linux-armv6l.tar.gz'
wget https://storage.googleapis.com/golang/$FileName
tar -xvf $FileName
sudo mv go /usr/local
rm $FileName
```

Add to your `~/.bashrc` file,

```bash
PATH=$PATH:$HOME/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOROOT/bin
CDPATH=$GOPATH/src/github.com
```

check,

```bash
go env
go version
```

For go tools, when you get the source,
you should already have `gofmt` and `godocs`.

To install other go tools, refer to my
[go-cheat-sheet](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/go-cheat-sheet/install-and-configure.md#install-go-tools).

## CONNECT TO WIFI (DEBIAN)

As a side note your router should always use WPA2-PSK with AES encryption.
So we will connect to that.

CONNMAN is installed on this system. Please dont use ifup/down.
Instead try connman-curses or connmanctl, which are both preinstalled!

You can see it running,

```bash
systemctl list-unit-files | grep connman
```

First install iwlist,

```bash
sudo apt install wireless-tools
```

First scan the wifi in your area,

```bash
sudo iwlist wlan0 scan
sudo iwlist wlan0 scan | grep <YOUR-WIFI-NAME>
```

You should be able to see the network you want.

Edit,

```bash
sudo nano /etc/wpa_supplicant/wpa_supplicant.conf
```

With the following for your `WPA2-PSK (AES)`,

```txt
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=US

network={
    ssid="myssid"
    psk="mypasskey"
    proto=RSN
    key_mgmt=WPA-PSK
    pairwise=CCMP
    auth_alg=OPEN
}
```

Reference this file in,

```bash
sudo nano /etc/network/interfaces
```

With,

```bash
allow-hotplug wlan0
iface wlan0 inet dhcp
wpa-conf /etc/wpa_supplicant/wpa_supplicant.conf
```

Check your connection (you may have to reboot),

```bash
ifconfig wlan0
```
