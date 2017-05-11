# DNS CHEAT SHEET

`dns` _(Domain Name Server) resolves hosts/domains into IP addresses
and visa versa.  Basically a large database.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## NAMESERVERS

Computers that run DNS are called name servers.

## BIND - NAMERSERVER FOR LINUX

BIND (Berkley Internet Naming Daemon) is the most common
program used for maintaining a name server on Linux.

Installing BIND is out of the scope of this cheat-sheet.

All DNS configurations for BIND are located under `/etc/bind`.

* /etc/bind/named.conf - Primary configuration.

* /etc/bind/db.root -  describes the root name servers in the world.

## DNSmasq - LIGHTWEIGHT NAMESERVER AND DHCP FOR LINUX

Another name server that is simple and lightweight is
`DNSmasq` that also does DHCP.

## CACHING NAMERSERVER

Sometimes a name server will just cahce off another name server.

## CLIENT DNS

The client DNS is called a DNS resolver.  And the config file is
`resolv.conf`.

## CLIENT DNS CONFIGURATION - /ETC/RESOLV.CONF

The configuration file for DNS resolver (the client DNS).

There are 3 directives in this file:

nameserver # Points to the IP address of the Name Server(s) you want to use.
domain     # Domain Name of local host
search     # Search list for host name lookup.  you need this to resolve short host names.

A note about searc.  Search resolves short hostnames.

For example if you have,

`search whatever.com`

and if you ssh into something like `ssh p-test`.  It will actually try,

`ssh p-test.whatever.com`

## CHECK THAT DNS IS WORKING

host ubuntu.com
dig ubuntu.com
ping ubuntu.com
nslookup ubuntu.com

## RESOLVCONF PROGRAM 

The service file is `/lib/systemd/system/resolvconf.service`.

`resolvconf` manages resolv.conf files from multiple sources,
such as DHCP and VPN clients.

The problem, multiple programs need to dynamically modify
the resolv.conf file they can step on each other and the
file can become out-of-sync. The resolvconf program addresses
this problem. It acts as an intermediary between programs that supply
nameserver information and programs that use this information.

So bottom line, `we can't edit /etc/resolv.conf`.

And its actually linked to,

`/etc/resolv.conf -> ../run/resolvconf/resolv.conf`

## METHOD 1 - ADDING TO RESOLV.CONF - USE TAIL FILE

`sudo nano /etc/resolvconf/resolv.conf.d/tail`

nameserver 1.2.3.4

`sudo resolveconf -u`

## METHOD 2 - ADDING TO RESOLV.CONF - USE INTERFACES FILE

To add something like nameserver to the file edit

sudo nano /etc/network/interfaces

iface lo inet loopback
    dns-nameservers 12.34.56.78 12.34.56.79

## TO PREVENT UPDATES

sudo nano /etc/dhcp/dhclient-enter-hooks.d/nodnsupdate

```bash
#!/bin/sh
make_resolv_conf() {
    :
}
```

chmod +x /etc/dhcp/dhclient-enter-hooks.d/nodnsupdate

When dhclient runs -- either on reboot or when you
manually run it loads this script nodnsupdate.
This script overrides an internal function called
make_resolv_conf() that would normally overwrite
resolv.conf and instead does nothing.
