# INSTALL AND CONFIGURE OS CHEAT SHEET

`install and configure OS` _distribution on your Hummingboard._

* [INSTALL "IGNITION" ON A microSD CARD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#install-ignition-on-a-microsd-card)
* [USE "IGNITION" TO INSTALL AN OS ON A microSD CARD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#use-ignition-to-install-an-os-on-a-microsd-card)
* [FULL LIST OF DISTRIBUTIONS FROM SOLIDRUN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#full-list-of-distributions-from-solidrun)
* [BASIC CONFIGURATION (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#basic-configuration-debian)
* [INSTALL GO & GO TOOLS (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#install-go--go-tools-debian)
* [CONNECT-TO_WIFI (DEBIAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/hummingboard/install-and-configure-os-cheat-sheet#connect-to-wifi-debian)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL "IGNITION" ON A microSD CARD

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

## USE "IGNITION" TO INSTALL AN OS ON A microSD CARD

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

## FULL LIST OF DISTRIBUTIONS FROM SOLIDRUN

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

Change your password for user pi,

```bash
passwd
```

### RENAME YOUR HUMMINGBOARD

Rename your host,

```bash
sudo nano /etc/hostname
sudo nano /etc/hosts
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

Since this an older copy of debian, I had to add sbin
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

```bash
git config --global user.name "Jeff DeCola (<HOSTNAME/MACHINE NAME>)"
git config --global user.email <YOUR-EMAIL>
git config --global core.editor nano
git config --global push.default simple
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

## CONNECT-TO_WIFI (DEBIAN)

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
