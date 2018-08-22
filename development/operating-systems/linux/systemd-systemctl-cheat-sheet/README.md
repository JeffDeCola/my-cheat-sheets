# SYSTEMD SYSTEMCTL CHEAT SHEET

`systemd-systemctl` _controls what programs run when Linux boots up._

Note: Ubuntu 14.04 uses upstart as the init system.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SYSTEMCTL

`systemctl` is a systemd utility which is responsible for
controlling the systemd and service manager.

Check your versions,

```bash
systemd --version
systemctl --version
```

Check where binaries are,

```bash
whereis systemd 
whereis systemctl
```

Check if systemd is running,

```bash
ps -eaf | grep [s]ystemd
```

## UNITS (SERVICES, MOUNT POINTS, SOCKETS AND DEVICES)

Systemctl accepts the following as units,

* services (.service)
* mount point (.mount) 
* sockets (.socket) 
* devices (.device)

List all units,

```bash
systemctl list-unit-files
```

List all services,

```bash
systemctl list-unit-files --type=service
```

List all services enabled (starts on linux boot),

```bash
systemctl list-unit-files --type=service | grep enabled
```

## CREATE, START AND STOP A SERVICE - BY WAY OF AN EXAMPLE

Lets do this by way of creating a service that runs at boot.
How about something that prints "Hi $USER, #",
when linux boots

Create a shell script,

```bash
nano say-hi.sh
```

with,

```sh
#!/bin/bash

count=1

while true
do
    echo "Hi" $USER $count
    sleep 1
    count=$((count+1))
done
```

Test it,

```bash
sh say-hi.sh
```

May have to add permissions to run `chmod 664`.

Now create a .service file with your path to `say-hi.sh`,

```bash
sudo nano /etc/systemd/system/say-hi.service
```

with,

```text
[Unit]
Description=say-hi-test

[Service]
ExecStart=your/path/to/say-hi.sh

[Install]
WantedBy=multi-user.target
```

May also have to `chmod 664` .

Check your service,

```bash
sudo systemd-analyze verify say-hi
```

Now see if the service is disabled,

```bash
systemctl list-unit-files --type=service | grep say-hi
```

Or check the status,

```bash
systemctl status say-hi
```

Lets start it,

```bash
sudo ystemctl start say-hi
```

Ways to check if it working,

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

Lets stop it,

```bash
sudo systemctl stop say-hi.service
```

Check its stopped,

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

## ENABLE A SERVICE AT BOOT

Now comes the meat and potatoes; getting the
service to run at boot.

First enable the service to start at boot,

```bash
systemctl enable say-hi
```

Reboot your system.

Now check the status,

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

Disable your service from boot,

```bash
sudo systemctl disable say-hi
```

Now stop it,

```bash
sudo systemctl stop say-hi
```



