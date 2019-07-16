# INSTALL AND CONFIGURE OS CHEAT SHEET

`install and configure OS` _distribution on your Raspberry Pi._

* [DOWNLOAD RASPBIAN IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#download-raspbian-image)
* [WRITE IMAGE TO microSD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#write-image-to-microsd)
* [BASIC CONFIGURATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#basic-configuration)
* [INSTALL GO & GO TOOLS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#install-go--go-tools)
* [CONNECT TO WIFI](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#connect-to-wifi)

Some fun things you can do with you raspberry pi,

* [Create DNS Server Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/create-dns-server-cheat-sheet)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## DOWNLOAD RASPBIAN IMAGE

Raspbian is the official OS of the raspberry pi.
Raspbian comes pre-installed with plenty of software for education,
programming and general use. It has Python, Scratch, Sonic Pi,
Java and more.
Raspbian is a free operating system based on Debian, optimized for
the Raspberry Pi hardware. Raspbian comes with over 35,000 packages:
pre-compiled software bundled in a nice format for easy installation.

Get an image of Raspbian OS
[here](https://www.raspberrypi.org/downloads/raspbian/).

Some caveats,

* Use the lite version if you don't want a desktop
* Don't use the full version if you don't have a Pi 4

## WRITE IMAGE TO microSD

I use  [Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
Make sure to use a microSD card of at least 4GB.

## BASIC CONFIGURATION

This is a short list,

Initial user and password,

```txt
user: pi
pw: raspberry
```

### UPDATE AND UPGRADE

```bash
sudo apt-get update
sudo apt-get upgrade
```

If you have issues try,

```bash
sudo apt update --allow-releaseinfo-change
```

### WHAT OS VERSION ARE YOU USING

To find the OS release,

```bash
cat /etc/os-release
```

### CHANGE PASSWORD FOR USER PI

Change your password for user pi,

```bash
passwd
```

### RENAME YOUR RASPI

Rename your host,

```bash
sudo nano /etc/hostname
sudo nano /etc/hosts
```

### TURN ON SSH

I would do this first since you probably want to ssh into it,

```bash
sudo systemctl enable ssh
sudo systemctl start ssh
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

### USE SSH KEYS TO SSH INTO RASPI

If you don't want to type the password all the time when you ssh
into your raspi, you can use your ssh keys.

You first need to get client's ssh public key into
the known hosts file on your rasp pi.  From your client,\
the following command does everything for you,

```bash
ssh-copy-id <USERNAME>@<IP-ADDRESS>
```

Again, you do not do this command on the Raspi.
You should now be able to ssh into the Raspi
with having to enter your password.

## INSTALL GO & GO TOOLS

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

## CONNECT TO WIFI

As a side note your router should always use WPA2-PSK with AES encryption.
So we will connect to that.

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

Check your connection (you may have to reboot),

```bash
ifconfig wlan0
```
