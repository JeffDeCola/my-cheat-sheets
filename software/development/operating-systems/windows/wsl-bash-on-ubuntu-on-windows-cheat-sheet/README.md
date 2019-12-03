# WSL - BASH ON UBUNTU ON WINDOWS CHEAT SHEET

`wsl (bash on ubuntu on windows)` _allows you to install linux
on a windows machine._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## WSL (WINDOWS SUBSYSTEM FOR LINUX)

The WSL (Windows System for Linux) allows you to run linux on Windows.

Ubuntu runs on top of windows with its own /home folder.

## UBUNTU ON WINDOWS

A popular distro of linux is ubuntu.  People refer to this
as `Bash on Ubuntu on Windows`.

It's actually really simple now, go to the windows store
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

## HOME

Where does Windows keep the files?

For Ubuntu 14.04/16.04,

```
C:\Users\<WindowsNAME>\AppData\Local\lxss\home\<bashusername>
```

For Ubuntu 18.04 (From Windows Store),

```
C:\Users\<WindowsNAME>\AppData\Local\Packages\<SOMETHING>\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\<bashusername>
```

## ALLOW SUDO PASSWORD FOR JEFF

Make yourself sudo,

```bash
sudo adduser jeff sudo
```

Then run visudo,

```bash
sudo visudo
```

Add,

```txt
jeff    ALL=(ALL:ALL) ALL
jeff  ALL = NOPASSWD : ALL
```

## FIX DIRECTORY COLOR HARD TO READ

Place in .bashrc,

```bash
LS_COLORS='ow=01;36;40'
export LS_COLORS
```

For more information, refer to my
[LS_COLORS cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ls_colors-cheat-sheet)

## HOME FOLDER

Linux can work with windows files, but windows CAN NOT work with linux files.

So I keep everything in a fake "home\jeff" folder and don't touch
these homes,

```txt
C:\Users\<WindowsNAME>\AppData\Local\lxss\home\<bashusername>
C:\Users\<WindowsNAME>\AppData\Local\Packages\<SOMETHING>\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\<bashusername>
```

Mine is,

```txt
/mnt/c/Users/Jeff/home/jeff
```

I keep my code here. I suggest you do something similar.

## SETUP CODE DEVELOPMENT ENVIRONMENT ON WINDOWS

Since you now have linux running on Windows, why don't you set up
a Code Development Environment.

It actually works.  Yeah, I know right.

Check out my cheat sheet on how to setup
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
on Windows with go.
