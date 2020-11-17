# INSTALL & CONFIGURE OS CHEAT SHEET

`install & configure OS` _distribution on your Raspberry Pi._

* [DOWNLOAD IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#download-image)
* [WRITE IMAGE TO microSD CARD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#write-image-to-microsd-card)
* [CONFIGURE YOUR OS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#configure-your-os)
* [INSTALL GO & GO TOOLS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#install-go--go-tools)
* [CONNECT TO WIFI](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#connect-to-wifi)
* [INCREASE SWAP SPACE FROM 100M TO 2G](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#increase-swap-space-from-100m-to-2g)
* [INPUT & OUTPUT GPIO PINS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#input--output-using-gpio-pins)

Some fun things you can do with your raspberry pi,

* [Create DNS Server using BIND Cheat Sheet](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/operating-systems/linux/dns-cheat-sheet/create-dns-server-using-bind.md)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## DOWNLOAD IMAGE

Get an image of the `Raspberry Pi OS` (previously called Raspbian),

* [32-bit](https://www.raspberrypi.org/downloads/raspbian/)
* [64-bit (beta test version)](https://www.raspberrypi.org/forums/viewtopic.php?f=117&t=275370)

Raspberry Pi OS comes pre-installed with plenty of software for education,
programming and general use. It has Python, Scratch, Sonic Pi, Java and more.

Some caveats,

* Use the lite version if you don't want a desktop
* Don't use the full version if you don't have a Pi 4

## WRITE IMAGE TO microSD CARD

I use  [Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
Make sure to use a microSD card of at least 4GB.

They recommend to use the [Raspberry Pi Imager](https://www.raspberrypi.org/downloads/)
for an easy way to install Raspberry Pi OS and other operating systems to a
microSD card.

## CONFIGURE YOUR OS

What I would do.

### Initial Username and Password

```txt
user: pi
pw: raspberry
```

### Update and Upgrade

```bash
sudo apt-get update
sudo apt-get upgrade
```

If you have issues try,

```bash
sudo apt update --allow-releaseinfo-change
```

### What OS version are you using

To find the OS release,

```bash
cat /etc/os-release
```

### Change Password for PI

Change your password for user pi,

```bash
passwd
```

### Rename your RasPi

Rename your host,

```bash
sudo nano /etc/hostname
sudo nano /etc/hosts
```

### Turn on ssh

I would do this first since you probably want to ssh into it,

```bash
sudo systemctl enable ssh
sudo systemctl start ssh
```

### Add a New User

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

### Get Git and GitHub Working

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

### Use ssh Keys to ssh into your RasPI

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

I had an issue with the Raspi changing the mac address for wlan0.
Yeah, it makes no sense.  My fix was,

```bash
sudo nano /etc/rc.local
```

with,

```bash
sudo ifconfig wlan0 hw ether AE:EF:49:11:81:41 # Spoofed MAC
sudo ifconfig wlan0 up # Start
sudo dhclient wlan0 # Start DHCP
```

## INCREASE SWAP SPACE FROM 100M TO 2G

Increase you swap from 100M to 2GB.

Edit,

```bash
sudo nano /etc/dphys-swapfile
```

Change to,

```txt
CONF_SWAPSIZE=2000
```

Stop and start,

```bash
sudo /etc/init.d/dphys-swapfile stop
sudo /etc/init.d/dphys-swapfile start
```

## INPUT & OUTPUT USING GPIO PINS

Some specs,

* GPIO is General Purpose Input / Output
* The newer RasPi's have 40 pin GPIOs
* The GPIO pins can be used as either digital inputs or outputs
* All the pins have 3.3V logic levels and are not 5V-safe
* There are 5V and 3.3V power pins

For fun I wrote a golang program to control these pins
[raspi-gpio](https://github.com/JeffDeCola/my-go-examples/tree/master/single-board-computers/raspi-gpio).

Here is the GPIO pin map I got from
[raspberrypi.org](raspberrypi.org),

![IMAGE - raspberry-pi-gpio-pins.png - IMAGE](../../../../docs/pics/raspberry-pi-gpio-pins.png)

This may help,

![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](../../../../docs/pics/raspberry-pi-input-and-output-using-gpio-pins.jpg)
