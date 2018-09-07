# WSL - BASH ON UBUNTU ON WINDOWS CHEAT SHEET

`wsl - bash on ubuntu on windows` _allows you to install linux
on a windows machine._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## WSL (WINDOWS SUBSYSTEM FOR LINUX)

The WSL (Windows System for Linux) allows you to run linux on Windows.

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
sudo do-release-upgrade -d
```

I would make a restore point and System Image first before doing any
update like this.  I've learned to never really trust Windows.

## SETUP CODE DEVELOPMENT ENVIRONMENT ON WINDOWS

Since you now have linux running on Windows, why don't you set up
a Code Development Environment.

It actually works.  Yeah, I know right.

Check out my cheat sheet on how to setup
[visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environments/visual-studio-code-cheat-sheet)
on Windows with go.
