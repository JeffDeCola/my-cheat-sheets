# INIT SysV SCRIPT CHEAT SHEET

`init SysV script` are depreciated and systemd should be used.
But systemd can still run these old init sysV scrips, so
it is here for reference.

tl;dl,

```bash
# START/STOP
sudo /etc/init.d/say-hi start
sudo /etc/init.d/say-hi stop
# CHECK
ps -efa|grep init
```

Current way to start a process at boot,

* [systemd systemctl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/systemd-systemctl-cheat-sheet)

Note: Ubuntu 14.04 uses upstart.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INIT IS DEAD

Even though you have the script in `/etc/init.d`,
these are controlled by `systemd`.

## CREATE, START & STOP A SERVICE (BY WAY OF AN EXAMPLE)

Lets do this by way of creating a shell script that
runs at boot. How about something that simply prints
`Hi $USER, #`, when linux boots.

### CREATE A SHELL SCRIPT

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
    echo "Hi, I'm using init:" $USER $count
    sleep 1
    count=$((count+1))
done
```

Test it,

```bash
sh say-hi.sh
```

May have to add permissions to run `chmod 775`.

### CREATE A INIT SysV SCRIPT

The scripts are kept in `/etc/init.d`.

```bash
sudo nano /etc/init.d/say-hi
```

with,

```text
#!/bin/bash
# say-hi daemon
# chkconfig: 345 20 80
# description: say-hi jeffs daemon
# processname: say-hi

DAEMON_PATH="/home/jeff"

DAEMON=say-hi.sh
DAEMONOPTS=""

NAME=say-hi
DESC="My daemon description"
PIDFILE=/var/run/$NAME.pid
SCRIPTNAME=/etc/init.d/$NAME

case "$1" in
start)
	printf "%-50s" "Starting $NAME..."
	cd $./DAEMON_PATH
	PID=`$DAEMON $DAEMONOPTS > /dev/null 2>&1 & echo $!`
	#echo "Saving PID" $PID " to " $PIDFILE
        if [ -z $PID ]; then
            printf "%s\n" "Fail"
        else
            echo $PID > $PIDFILE
            printf "%s\n" "Ok"
        fi
;;
status)
        printf "%-50s" "Checking $NAME..."
        if [ -f $PIDFILE ]; then
            PID=`cat $PIDFILE`
            if [ -z "`ps axf | grep ${PID} | grep -v grep`" ]; then
                printf "%s\n" "Process dead but pidfile exists"
            else
                echo "Running"
            fi
        else
            printf "%s\n" "Service not running"
        fi
;;
stop)
        printf "%-50s" "Stopping $NAME"
            PID=`cat $PIDFILE`
            cd $DAEMON_PATH
        if [ -f $PIDFILE ]; then
            kill -HUP $PID
            printf "%s\n" "Ok"
            rm -f $PIDFILE
        else
            printf "%s\n" "pidfile not found"
        fi
;;

restart)
  	$0 stop
  	$0 start
;;

*)
        echo "Usage: $0 {status|start|stop|restart}"
        exit 1
esac
```

Add to init,

```bash
sudo update-rc.d say-hi defaults
```

### START/STOP YOUR SERVICE

```bash
sudo /etc/init.d/say-hi start
sudo /etc/init.d/say-hi stop
```
check,

```bash
ps -efa|grep init
```
