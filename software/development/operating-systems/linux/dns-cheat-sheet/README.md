# DNS CHEAT SHEET

`dns` _(Domain Name Server) resolves hosts/domains into IP addresses
and visa versa.  Basically a large database._

* [DNS SERVER (name server)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dns-cheat-sheet#dns-server-name-server)
* [DNS CLIENT (resolver)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/dns-cheat-sheet#dns-client-resolver)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## DNS SERVER (NAME SERVER)

Keeps the list of names/IP addresses.  There are many many types
of name servers that do many different things living on many
different systems. I will only briefly talk about the ones
that may live on your linux box.

Some popular ones are,

* [BIND](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/operating-systems/linux/dns-cheat-sheet/create-dns-server-using-bind.md)
* NSD (Name Server Daemon)
* djbdns (tinydns)
* DNSmasq

### BIND (A POPULAR NAME SERVER FOR LINUX)

BIND (Berkley Internet Naming Daemon) is the most common
program used for maintaining a name server on Linux.
I wrote a cheat sheet on
[how to install BIND](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/operating-systems/linux/dns-cheat-sheet/create-dns-server-using-bind.md).

## DNS CLIENT (RESOLVER)

The client DNS is called a DNS resolver and is your like to
a name server.  The config file is `resolv.conf`.

### CONFIGURATION (/etc/resolv.conf)

The configuration file for DNS resolver (the client DNS).

There are 3 directives in this file:

* nameserver _Points to the IP address of the Name Server(s)._
* domain     _Domain Name of local host._
* search     _Search list for host name lookup.  This resolves short host names._

Search resolves short hostnames. For example if you have,

`search whatever.com`

And you ssh into something like `ssh p-test`.  It will actually do,

`ssh p-test.whatever.com`

### CHECK THAT YOUR DNS CLIENT IS WORKING PROPERLY

```bash
host ubuntu.com
dig ubuntu.com
ping ubuntu.com
nslookup ubuntu.com
```

### RESOLVCONF (An App used by Ubuntu and Debian)

Used by ubuntu and debian to "edit your resolve.conf".

The service file is `/lib/systemd/system/resolvconf.service`.

`resolvconf` manages resolv.conf files from multiple sources,
such as DHCP and VPN clients.

The problem, multiple programs need to dynamically modify
the resolv.conf file they can step on each other and the
file can become out-of-sync. The resolvconf program addresses
this problem. It acts as an intermediary between programs that supply
nameserver information and programs that use this information.

So bottom line, we can't edit `/etc/resolv.conf`.

The file is actually linked to,

`/etc/resolv.conf -> ../run/resolvconf/resolv.conf`

But there are two methods to solve this.

#### METHOD 1 - ADDING TO RESOLV.CONF - USE TAIL FILE

Again, when using resolvconf program, you can't edit
`/etc/resolv.conf`, so this is method 1 to get around that.

`sudo nano /etc/resolvconf/resolv.conf.d/tail`

nameserver 1.2.3.4

You can restart if you want.

`sudo resolveconf -u`

#### METHOD 2 - ADDING TO RESOLV.CONF - USE INTERFACES FILE

To add something like nameserver to the file edit

sudo nano /etc/network/interfaces

iface lo inet loopback
    dns-nameservers 12.34.56.78 12.34.56.79

### NETWORK MANAGER (An App used by Centos and Red Hat)

Used by centos and red hat.

Status,

```bash
systemctl status NetworkManager
```

Install,

```bash
sudo apt-get -y install network-manager
sudo yum -y install NetworkManager
```

Start,

```bash
sudo systemctl start NetworkManager.service
```

Enable restarting the network manager when the system reboots,

```bash
sudo systemctl enable NetworkManager.service
```

Configure,

```bash
/etc/NetworkManager/NetworkManager.conf
```

Now use with /etc/resolve.conf

### PREVENT UPDATES FROM DHCP

When dhclient runs; either on reboot or when you
manually run it loads this script nodnsupdate.
This script overrides an internal function called
make_resolv_conf() that would normally overwrite
resolv.conf and instead does nothing.

sudo nano /etc/dhcp/dhclient-enter-hooks.d/nodnsupdate

```bash
#!/bin/sh
make_resolv_conf() {
    :
}
```

chmod +x /etc/dhcp/dhclient-enter-hooks.d/nodnsupdate
