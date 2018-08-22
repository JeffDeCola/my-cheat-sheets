# SYSTEMD SYSTEMCTL CHEAT SHEET

`systemd-systemctl` _controls what programs run when Linux boots up._

Note, Ubuntu 14.04 uses Upstart as the init system,
systemd is for Ubuntu 14.10+. 

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SYSTEMCTL

`systemctl` is a systemd utility which is responsible for
Controlling the systemd system and service manager.

Check your version,

```bash
systemd --version
systemctl --version
```

Checck where binaries are,

```bash
whereis systemd 
whereis systemctl
`
check if systemd is running,

```bash
ps -eaf | grep [s]ystemd
`

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

## SERVICE - CREATE, START AND STOP A SERVICES - BY WAY OF AN EXAMPLE

Lets do this by way of creating a service that runs at boot.
How about something that prints "Hi $USER, have a great day",
when linux boots

Create a shell script such as the following,

```bash
nano say-hi.sh
```

```base
#!/bin/bash

count=1

while true
do
    echo "Hi" $USER $count
    sleep 1
    count=$((count+1))
done
```

```bash
sh say-hi.sh
```

May have to add permissions to run `chmod 664`.

Now create a .service file with your path to say-hi.sh,

```bash
sudo nano /etc/systemd/system/say-hi.service
```

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

Now see all the services and you should see it disabled,

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

Check it working bu looking at the tail (switch -f) of the service logs,

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

Lets stop it

```bash
sudo systemctl stop say-hi.service
```

Cehck its stoped counting.

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

## ENABLE A SERVICE AT BOOT

Now comes the meat and potatoes, getting the
service to run at boot.

```bash
systemctl enable say-hi
```

It will create a symlink,

Test it be rebooting and checking the status,

```bash
journalctl -f
systemctl status say-hi
systemctl is-active say-hi
```

To disable your service jefffrom boot, just disable it,

```bash
sudo systemctl disable say-hi
```

And you still need to stop it

```bash
sudo systemctl stop say-hi
```



