# PACKER CHEAT SHEET

`packer` _is a tool from HashiCorp that automates the building
of custom machine `images` (for multiple platforms) from a single
configuration file._

My Repo example is [hello-go-deploy-gce](https://github.com/JeffDeCola/hello-go-deploy-gce).

Documentation and reference,

* My cheat sheet,
  [create a custom image using packer on gce](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine-create-image-packer.md)
* [Packer Documentation](https://www.packer.io/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## AUTOMATES MACHINE IMAGE BUILDS

A machine `image` is a single static unit that contains a 
pre-configured operating system and installed software
which is used to quickly create new running machines.

Once you set up a configuration file, you run `packer` and it
does the rest. It fits perfectly in a `continuous deployment` pipeline.

The beauty is, it automates the building of your custom machine `image`
(you can also say provisioned image, but I like custom image).
And once set up, you can generate different machine `images` for
different platforms.

Packer supports building `images` for
Amazon EC2, CloudStack, DigitalOcean, Docker,
Google Compute Engine, Microsoft Azure, QEMU,
VirtualBox, VMware, and more.

Here is a high level view of packer,

![IMAGE -  packer high level view - IMAGE](../../../../../docs/pics/packer-high-level-view.jpg)

## INSTALL

Go [here](https://www.packer.io/downloads.html)
to install.  Must use a 64-bit machine.

For Linux, its actually really simple to install
the binary, just download, unzip and place in `/usr/bin`.

For macOS, place the binary in `/usr/local/bin`. This is due to
System Integrity Protection (SIP).  SIP makes `/usr/bin` read-only.

After install check the version,

```bash
packer version
```

## TEMPLATE FILE (CONFIGURATION)

There is only one configuration file that packer uses.
It is called a template in packer terminology.

It has three main sections,

* `variables` - Just variables to make your life easier.
* `builders` - Define and configure your builders (e.g. gce, aws, azure etc...)
* `provisioners` - This is the magic, where you add/install stuff.

### BASIC FORMAT OF A TEMPLATE FILE

By way of example, here is a very simple template file used
to build a simple image at `gce`,

```json
{
    "variables": {
        "zone": "us-central1-a",
    },

    "builders": [
        {   
            "type": "googlecompute",
            "account_file": "account.json",
            "project_id": "my project",
            "source_image": "debian-7-wheezy-v20150127",
            "ssh_username": "packer",
            "zone": "{{user `zone`}}"
        }
    ],

    "provisioners": [
        {
            "type": "file",
            "source": "./welcome.txt",
            "destination": "/home/ubuntu/"
        },
        {
            "type": "shell",
            "script": "./example.sh"
        }
    ]
}
```

### VARIABLES

Variables are just that, variables. It makes life easier.
they are used in the builder section.

You can also pass in variables to packer via
the packer command line interface.

### BUILDERS

* Defines the builders that will be used to create machine `images`
  For example `gce`, `aws`, etc...
* Configures each of those builders

### PROVISIONERS (ADD YOUR STUFF)

Install and configure software.  This stage is also known
as the provision step.

## VALIDATE THE TEMPLATE FILE

Before you kick off a build, validate the template file,

```bash
packer validate packer-template-file.json
```

## RUN THE TEMPLATE FILE (LETS BUILD OUR IMAGE)

Lets build our custom machine image,

```bash
packer build -force packer-template-file.json
```

`-force` removes artifacts from previous build.

You can pass in variables to `packer`
if you don't want to keep them in your
template file.

For example,

```bash
packer build -force \
    -var 'region=us-central1' \
    -var 'zone=us-central1-b' \
    packer-template-file.json
```

There are also lots of command line switches,
but I like to keep everything in my
template file.

Again, to see a working example, go to my repo
[hello-go-deploy-gce](https://github.com/JeffDeCola/hello-go-deploy-gce).
