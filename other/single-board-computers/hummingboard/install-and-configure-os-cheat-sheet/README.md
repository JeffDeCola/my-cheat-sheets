# INSTALL AND CONFIGURE OS CHEAT SHEET

`install and configure OS` _distribution on your Hummingboard._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL "IGNITION" ON A microSD CARD

The `Ignition` installer from SolidRun lets you easily select
your OS distribution of choice.

To use this, it must be burned on a microSD card.

* First, get the `Ignition` image
[here](https://www.solid-run.com/downloads/ignition).
* Second, get a way to get the image on a microSD 
[Win32DiskImager](https://sourceforge.net/projects/win32diskimager).
* Last, with your computer, use `Win32DiskImager`
to burn the image on a microSD card.

I would label this card since its easy to forget.

## USE "IGNITION" TO INSTALL AN OS ON A microSD FOR YOUR HUMMINGBOARD

Place the `Ignition` microSD in your hummingboard and boot it up.
You must have your hummingboard connected to the internet
(wired or wireless) to use this software.

The software will look at SolidRun's GitHub
[repo](https://github.com/SolidRun/ignition-imx6)
for imx.6 ignition scripts for distributions.

Choose a distro such as:
* Debian
* GeexBox
* LibreELEC
* OpenELEC.tv
* ArchLinux
* Fedora-22
* Android-kitkat 4.4.4

Place another microSD card in your hummingboard and install
your image.  Do not write over the `Ignition` microSD card,
like i did.

## DEBIAN IMAGES for i.MX6

Some debian images are listed
[here](https://images.solid-build.xyz/IMX6/Debian)

## FULL LIST OF DISTRIBUTIONS FROM SOLIDRUN

Complete list of ignition installer scripts for various OS
[here](https://wiki.solid-run.com/doku.php?id=tag:ignition&do=showtag&tag=ignition)
