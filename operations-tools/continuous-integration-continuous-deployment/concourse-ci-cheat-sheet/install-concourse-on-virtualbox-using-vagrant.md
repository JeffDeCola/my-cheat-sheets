# INSTALL CONCOURSE CI ON VIRTUALBOX USING VAGRANT

This method has been depreciated at concourse, but I'm sure you can
find vagrant files that others have made.

Get vagrant for Windows,

[https://www.vagrantup.com](https://www.vagrantup.com)

Get the Concourse Vagrant files.

To run, enter vagrant directory, from Windows cmd prompt type,

```bash
vagrant up
```

To upgrade vagrant box (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant box prune                           # remove the old Vagrant boxes
vagrant up 
```

Reinstal the new fly.exe for Windows and Linux.