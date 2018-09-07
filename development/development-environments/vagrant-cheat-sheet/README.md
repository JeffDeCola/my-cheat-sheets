# VAGRANT CHEAT SHEET

`vagrant` _creates and configures portable development environments._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Vagrant provides the same production environment (_OS,
packages, users, configuration, etc..._) giving developers the
flexibility to use there own tools (_browsers, IDEs, editors, etc..._).

Vagrant can match the development environment to the
production environment. _"Well it worked on my computer"_
is a statement of the past.

## INSTALL

Visit [vagrant downloads](https://www.vagrantup.com/downloads.html).

## CHECK VERSION

```bash
vagrant version
```

## FIND BOXES

Search for vagrant boxes at
[vagrant box search](https://app.vagrantup.com/boxes/search)

## VAGRANTFILE

Vagrants uses a declarative config file which describes your
software requirements, packages, OS configuration, users, etc..

A vagrant file could look life the following,

```code
Vagrant.configure("2") do |config|
  config.vm.box = "concourse/lite"
end
```

## YOUR LOCAL VAGRANT BOXES

To list boxes on your machine,

```bash
vagrant box list
```

To update all boxes,

```bash
vagrant box update
```

To update a particular box,

```bash
vagrant box update --box NAME
```

To destroy all old boxes,

```bash
vagrant box destroy
```
