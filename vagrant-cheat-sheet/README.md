# VAGRANT CHEAT SHEET

`vagrant` _creates and configures portable dev environments._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW

Vagrant mirrors production environments (by providing same OS,
packages, users, configuration, etc...) giving devs the
flexibility to use there own tools (browser, IDE, editor, etc...

Vagrant matches the dev environment to the production environment.
"Well it worked on my computer" is a statement of the past.

## INSTALL & CHECK VERSION

```bash
vagrant version
```

## VAGRANTFILE

Vagrants uses a declarative config file which describes your
software requirements, packages, OS configuration, users, etc..

