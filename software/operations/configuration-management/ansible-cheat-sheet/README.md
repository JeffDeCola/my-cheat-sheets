# ANSIBLE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Ansible is a configuration management and provisioning
tool for deploying a server/service._

Table of Contents

* [WHAT CAN YOU DO WITH ANSIBLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#what-can-you-do-with-ansible)
* [ANSIBLE USES SSH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#ansible-uses-ssh)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#install)
  * [macOS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#macos)
  * [UBUNTU](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#ubuntu)
* [INVENTORY FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#inventory-file)
* [COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#commands)
* [MODULES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#modules)
* [AD HOC COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#ad-hoc-commands)
* [PLAYBOOK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#playbook)
* [ROLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/configuration-management/ansible-cheat-sheet#roles)

## WHAT CAN YOU DO WITH ANSIBLE

* Configuration Management.
* Application deployment.
* Provisioning (make available).
* Orchestration (automation).
* Security.
* Continuous Delivery.

## ANSIBLE USES SSH

Ansible depends on SSH access to the servers you are managing.
There is no client software needed on host your want to manage.

## INSTALL

### macOS

The control node.

```bash
brew install ansible
```

Version,

```bash
ansible --version
```

### UBUNTU

```bash
sudo apt-get install software-properties-common
sudo apt-add-repository ppa:ansible/ansible
sudo apt-get update
sudo apt-get install ansible
```

## INVENTORY FILE

Ansible uses an inventory file to determine what hosts to work against.

Make an inventory file of your hosts (use hostname or ip),

`nano /etc/ansible/hosts`,

```bash
p-stack-to-graph-ksf5
p-stack-to-graph-ksf5:2222
```

Tell ansible where you host file is,

`nano ~/.ansible.cfg`,

```bash
[defaults]
inventory = /etc/ansible/hosts
```

Your public SSH key should be located in authorized_keys on those systems.

There are many more things you can do with a inventory file.

## COMMANDS

All commands have the same format,

```bash
ansible server_or_group -m module_name -a arguments
```

If you didn't have an inventory file you would have to do,

```bash
ansible all -i p-stack-to-graph-ksf5, -m ping
```

* all look at all the host
* -m --module-name like ping
* -i what is the inventory path

But since we have an inventory file (the hosts), the following works fine,

```bash
ansible all --module-name ping
ansible all -m ping
```

## MODULES

Ansible’s way of abstracting certain system management or configuration tasks.

Control things you automate.

There are over 450 modules [here](http://docs.ansible.com/ansible/modules_by_category.html).

## AD HOC COMMANDS

Use the module commands to send a command to the host,

```bash
ansible all --module-name command --args "uptime"
ansible all -m command -a "uptime"
ansible all -m command -a "/bin/date"
```

## PLAYBOOK

Allow you to organize your configuration and management
tasks in simple, human-readable files.

Playbooks can be combined with other playbooks and organized into
Roles which allow you to define sophisticated infrastructures and
then easily provision and manage them.

* Playbooks contain plays.
* Plays contain tasks.
* Tasks call modules.
* Tasks run sequentially.
* Handlers are triggered by tasks and are ran once, at the end of the plays.

A playbook template looks like,

```yml
---
- hosts: [target hosts]
  remote_user: [yourname]
  tasks:
    - [task 1]
    - [task 2]
```

Lets create a task,

`nano test.yml`

```yml
---
- hosts: all
  tasks:
    - name: Ensure git is installed
      apt: name=git state=installed
      sudo: yes

    - name: Install cheat sheets repo
      git: repo=https://github.com/JeffDeCola/my-cheat-sheets.git
           dest=~/my-cheats
      remote_user: jeffdecola
      sudo: false

    - name: Copy file
      command: cp ~/my-cheats/ansible-cheat-sheet/README.md ~/copied-this-README.MD
      remote_user: jeffdecola
      sudo: false
```

Run,

```bash
ansible-playbook test.yml
```

## ROLES

Roles are a special kind of playbook.

Folders are as follows,

* `defaults`    _The default variables (lower priority variables)._
  * main.yml
* `files`       _If need to add files._
* `handlers`    _Targets for notify directives, and are almost always
  associated with services._
  * main.yml
* `meta`        _Meta data._
  * main.yml
* `tasks`       _Series of Ansible plays to install, configure, and run software._
  * main.yml
* `templates`   _Similar to files except that templates support modification._
* `vars`        _The default variables (higher priority variables)._
  * main.yml
