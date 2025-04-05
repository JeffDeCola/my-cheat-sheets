# SSH & KEYS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_ssh and keys (secure socket shell) is a way to securely
communicate (via ssh keys) with a remote computer._

tl;dr

```bash
# KEYS
ssh-keygen -t rsa -b 4096 -C "Keys for Github <MACHINE/HOSTNAME>"
ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
ssh-copy-id user@<IP>
## SSH CONFIGURATION FILE (EASY LOGIN)
nano ~/.ssh/config
# Host stimpy
#    HostName stimpy
#    User jeff
#    IdentityFile ~/.ssh/id_rsa
ssh stimpy
# SSH EXAMPLES
ssh user@<IP>
ssh -v -p 22 user@<HOSTNAME>
ssh -i ~/.vagrant.d/insecure_private_key -p 2222 vagrant@127.0.0.1
ssh stimpy "ls -lat"
# COPY TO REMOTE
scp test1.txt jeff@stimpy:~/
scp -r testdirectory1 jeff@stimpy:~/
# COPY FROM REMOTE
scp jeff@stimpy:~/test2.txt ~/
scp -r jeff@stimpy:testdirectory2 ~/
```

Table of Contents

* [KEYS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#keys)
  * [SECURE CONNECTIONS - HOW KEYS ARE USED](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#secure-connections---how-keys-are-used)
  * [GENERATE KEYS (ssh-keygen)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#generate-keys-ssh-keygen)
  * [KEY FINGERPRINT (ssh-keygen)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#key-fingerprint-ssh-keygen)
  * [ADD PRIVATE KEY TO YOUR SSH AGENT (ssh-add)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#add-private-key-to-your-ssh-agent-ssh-add)
* [SSH SERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#ssh-server)
  * [INSTALL SSH SERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#install-ssh-server)
  * [CONFIGURE SSH SERVER (sshd_config)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#configure-ssh-server-sshd_config)
  * [PLACE PUBLIC KEY ON SERVER MACHINE (authorized_keys)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#place-public-key-on-server-machine-authorized_keys)
  * [TEST](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#test)
* [SSH CLIENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#ssh-client)
  * [INSTALL SSH CLIENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#install-ssh-client)
  * [SSH CONFIGURATION FILE (config)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#ssh-configuration-file-config)
* [SSH EXAMPLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#ssh-examples)
  * [LOGIN TO REMOTE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#login-to-remote)
  * [RUN COMMANDS REMOTELY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#run-commands-remotely)
  * [COPY FILE LOCAL-TO-REMOTE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#copy-file-local-to-remote)
  * [COPY FILE REMOTE-TO-LOCAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#copy-file-remote-to-local)
* [CONNECT TO MACHINE ON PRIVATE NETWORK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#connect-to-machine-on-private-network)
  * [METHOD 1 - PORT FORWARDING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#method-1---port-forwarding)
  * [METHOD 2 - REVERSE SSH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ssh-and-keys-cheat-sheet#method-2---reverse-ssh)

## KEYS

SSH stands for “secure shell” and is what will allow us to
establish a secure connection between two computers.

### SECURE CONNECTIONS - HOW KEYS ARE USED

Sequence,

* You create both a public and private key
* You give the remote computer your public key
* You keep your private key and NEVER share
* The server encrypts the data and sends it to you
* You use your private key to decrypt it

If the data went anywhere else, it would need your private
key to decrypt.

Here is an illustration using ssh for github,

![IMAGE -  ssh keys github overview - IMAGE](../../../../../docs/pics/software/development/ssh-keys-github-overview.svg)

### GENERATE KEYS (ssh-keygen)

Generate rsa key pair,

```bash
ssh-keygen -t rsa -b 4096 -C "Keys for Github <MACHINE/HOSTNAME>"
```

* -t switch is the `rsa` type key (very popular).
* -b switch is the length in bits.
* -C is a comment for this key.

This will create two keys and place them in
`~/.ssh` with default names as `id_rsa` and `id_rsa.pub`.

I'm not a fan of using a paraphrase (password), because I'll
forget. Its up to you if you want an extra level of protection if
by chance someone was on your computer being mischievous or if
someone stole your keys. Hey, maybe you pushed your private key
to github by accident.

### KEY FINGERPRINT (ssh-keygen)

Sometimes it gets confusing with lots of keys, so you
can check your public key fingerprint to make sure your
server is using the right public key.

If you public key is located in `~/.ssh/id_rsa.pub`

```bash
ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
```

### ADD PRIVATE KEY TO YOUR SSH AGENT (ssh-add)

Now add your key,

```bash
ssh-add ~/.ssh/id_rsa
```

If agent not running try,

```bash
eval "$(ssh-agent)"
Check ssh-add -L
```

## SSH SERVER

Another way to communicate to a remote computer without
traditional username/password `login`.

It stands for secure shell and uses port 22.

Remote will mean the ssh server and local will be the ssh client machine.

### INSTALL SSH SERVER

If you want to ssh into a machine, that machine must have a ssh server running.

To install on ubuntu, on the remover machine,

```bash
sudo apt-get install openssh-server
sudo systemctl start ssh
sudo systemctl enable ssh
sudo systemctl status sshd.service
```

### CONFIGURE SSH SERVER (sshd_config)

```bash
sudo nano /etc/ssh/sshd_config
```

And change the following to only allow ssh,

```bash
Port 22
Protocol 2
PermitRootLogin no
MaxAuthTries 6
MaxStartups 3
#PermitRootLogin prohibit-password
PubkeyAuthentication yes
IgnoreRhosts yes
PasswordAuthentification no
PermitEmptyPasswords no
```

Restart ssh service,

```bash
sudo systemctl reload sshd.service
```

If you ever want to add another ssh client, you need to
turn `PasswordAuthentification yes` and then upload

### PLACE PUBLIC KEY ON SERVER MACHINE (authorized_keys)

Push the public key from your client machine to the ssh server,

```bash
ssh-copy-id user@<IP>
```

Or you can login to the remote ssh server machine and manually
add the public key to `.ssh/authorized_keys`.

### TEST

From command line,

```go
ssh user@<IP>
ssh -v -p 22 user@<HOSTNAME>
```

## SSH CLIENT

Add a host file to make things easier.

### INSTALL SSH CLIENT

You already have it, but just in case,

```bash
sudo apt-get install openssh-client
```

### SSH CONFIGURATION FILE (config)

Make it easier to ssh into servers without typing everything,

edit,

```bash
nano ~/.ssh/config
```

with,

```txt
Host stim
    # HostName 192.168.20.102
    HostName stimpy
    User jeff
    IdentityFile ~/.ssh/id_rsa
```

Then just do a short command of,

```bash
ssh stim
```

## SSH EXAMPLES

Some cool examples.

### LOGIN TO REMOTE

From command line,

```go
ssh user@<IP>
ssh -v -p 22 user@<HOSTNAME>
```

Again, no password is needed.

In this example I am using a particular private key and port.
This is useful to ssh on a vagrant made virtualbox VM.

```bash
ssh -i ~/.vagrant.d/insecure_private_key -p 2222 vagrant@127.0.0.1
```

Again, put this info in `~/.ssh/config` as shown above
so you don't have to type all the time.

### RUN COMMANDS REMOTELY

```bash
ssh stimpy "ls -lat"
```

### COPY FILE LOCAL-TO-REMOTE

```bash
scp test1.txt jeff@stimpy:~/
scp -r testdirectory1 jeff@stimpy:~/
```

### COPY FILE REMOTE-TO-LOCAL

```bash
scp jeff@stimpy:~/test2.txt ~/
scp -r jeff@stimpy:testdirectory2 ~/
```

## CONNECT TO MACHINE ON PRIVATE NETWORK

There are two methods to connect to a machine on a private network.
Forward a port on the router or use reverse port forwarding.

### METHOD 1 - PORT FORWARDING

A home router has it's ports closed (for security).
But if you want to remote logic, you can open a ssh port.

You have two IP addresses to find,

* Your WAN IP(i.e. your ISP)
* Your LAN IP (i.e. your local private IP of your device)

To find your WAN ip (your isp), in your browser type,

```bash
www.whatsmyip.org
```

To find your LAN IP (your network), type,

```baby
ip addr
```

Or the old depreciated command,

```bash
ifconfig 
```

Configure your router to forward your ssh port (Usually port 22),

* Log in to your router
* Add a service (SSH is usually one of the default options)
* Select port number (22 by default for SSH)
* Input your local IP address of your device
* Save the updated settings

Test with an outside machine,

```go
ssh user@<IP>
```

Please remember to secure your sshd server via the sshd_config file
as shown above.

### METHOD 2 - REVERSE SSH

This method requires you to know the machine that is going to
login, beforehand.

tbd
