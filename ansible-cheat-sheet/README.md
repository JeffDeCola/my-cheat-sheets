# ANSIBLE CHEAT SHEET

`ansible` _is a configuration management and provisioning tool for deployin a server._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL MAC

```bash
brew install ansible
```

Version,

```bash
ansible --version
```

## SETUP ON MAC - INVENTORY FILE

Make an inventory file of your hosts,

`nano /etc/ansible/hosts`,

```bash
p-stack-to-graph-d9rf
p-stack-to-graph-d9rf:2222
```

`nano ~/.ansible.cfg`,

```bash
[defaults]
inventory = /etc/ansible/hosts
```

## COMMANDS

All commands have the same format,

```bash
ansible server_or_group -m module_name -a arguments
```

If you didn't have an inventory file you would have to do,

```bash
ansible all -i myserver.com, -m ping
```

But since we have an inventory file, this al works

```bash
ansible all -m ping
```

## PLAYBOOK

The template playbook looks like,

```yml
---
- hosts: [target hosts]
  remote_user: [yourname]
  tasks:
    - [task 1]
    - [task 2]
```