# INSTALL AND CONFIGURE OS CHEAT SHEET

`install and configure OS` _distribution on your Raspberry Pi._

* [DOWNLOAD RASPBIAN IMAGE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#download-raspbian-image)
* [WRITE IMAGE TO microSD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#write-image-to-microsd)
* [BASIC CONFIGURATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/single-board-computers/raspberry-pi/install-and-configure-os-cheat-sheet#basic-configuration)

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

* Use the lite version if you don't want a desktop.
* Don't use the full version if you don't have a Pi 4.

## WRITE IMAGE TO microSD

I use  [Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
Make sure to use a microSD card of at least 4GB.

## BASIC CONFIGURATION

This is a short list,

### CHANGE PASSWORD FOR USER PI

Change your password for user pi,

```bash
passwd
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

### NAME YOUR RASPI

Name your host,

```bash
sudo nano /etc/hostname
sudo nano /etc/hosts
```

### GET GIT & GITHUB WORKING

Get git,

```bash
sudo apt install git
git version
```

Create keys,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github"
```

Place the Raspi public key at github.

Refer to my
[cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/source-version-control/git-cheat-sheet)
for more info.

### USE SSH KEYS TO SSH INTO RASPI

If you don't want to type the password all the time when you ssh
into your raspi, you can use your ssh keys.

You first need to get client's ssh public key into
the known hosts file on your rasp pi.  The following command does
everything for you,

```bash
ssh-copy-id <USERNAME>@<IP-ADDRESS>
```

You should now be able to ssh into the Raspi
with having to enter your password.
