# VIRTUALBOX CHEAT SHEET

`virtualbox` _creates and configures portable dev environments._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL ON WINDOWS

Download and install[virtualbox](https://www.virtualbox.org/).

## INSTALL GUEST ADDITIONS FROM COMMAND LINE

Get the latest `VBoxGuestAdditions.iso` image file from
[virtualbox](http://download.virtualbox.org/virtualbox).

Start your VM.  Go to the menu item
`Devices -> Insert Guest Additions CD image`
to mount the ISO image.

From the terminal, run the following commands,

```bash
sudo mkdir -p /media/cdrom
sudo mount /dev/cdrom /media/cdrom
sudo apt install -y dkms build-essential linux-headers-generic linux-headers-$(uname -r)
sudo sh /media/cdrom/VBoxLinuxAdditions.run  
reboot
```

After reboot, check that it installed,

```bash
ls /opt
```

You may need to add yourself to the vboxsf group,

First check,

```bash
groups
```

To add,

```bash
sudo usermod --append --groups vboxsf USERNAME
reboot
```

## DRAG AND DROP

Must have a shared folder.

