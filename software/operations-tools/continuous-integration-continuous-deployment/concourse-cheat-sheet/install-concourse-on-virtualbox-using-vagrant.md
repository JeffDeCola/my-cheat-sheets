# INSTALL CONCOURSE ON VIRTUALBOX USING VAGRANT

This method has been depreciated at concourse, but I'm sure you can
find vagrant files that others have made.

Get vagrant for Windows [here](https://www.vagrantup.com).

Search for a concourse vagrant file
[here](https://app.vagrantup.com/boxes/search).
This [one](https://app.vagrantup.com/concourse/boxes/lite) works,
but is in older version of concourse.

To run your vagrant box, enter vagrant directory, and from the
Windows cmd prompt type,

```bash
vagrant up
```

To upgrade vagrant box (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant box prune                           # remove the old Vagrant boxes
vagrant up
```

Reinstall the new `fly.exe`.
