# VAGRANT CHEAT SHEET

`vagrant` _creates and configures portable dev environments._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW

Vagrant mirrors production environments (by providing same OS,
packages, users, configuration, etc...) giving devs the
flexibility to use there own tools (browser, IDE, editor, etc...

Vagrant matches the dev environment to the production environment.
"Well it worked on my computer" is a statement of the past.

## INSTALL

goto,

[vagrant downloads](https://www.vagrantup.com/downloads.html)

## CHECK VERSION

```bash
vagrant version
```

## FIND BOXES

Search for boxes here,

[vagrant box search](https://app.vagrantup.com/boxes/search)


## VAGRANTFILE

Vagrants uses a declarative config file which describes your
software requirements, packages, OS configuration, users, etc..

A vagrant file could look life the following.

```code
Vagrant.configure("2") do |config|
  config.vm.box = "concourse/lite"
end
```

To chose a particular box version use,


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


