# REDIS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Redis is an open source non-relational (NoSQL) database system
using key:value storage._

tl;dr

```bash
# VERSIONS
redis-server --version
redis-cli --version
# MANUAL START
redis-server /etc/redis/6379.conf
# CHECK IF ALIVE
redis-cli ping
ps aux | grep -i redis
# SHUTDOWN saves the data file /var/redis/6379/jeff-redis.rdb
redis-cli shutdown
# CHECK LOG
tail /var/log/redis_6379.log
# LINUX BOOT - USING INIT
sudo /etc/init.d/redis_6379 start
sudo /etc/init.d/redis_6379 stop
ps aux | grep -i redis
# MACOS BOOT - USING LAUNCHD
launchctl load ~/Library/LaunchAgents/local.redis.redis-server.plist
launchctl unload ~/Library/LaunchAgents/local.redis.redis-server.plist
launchctl list | grep redis
# SET/GET/DEL
redis-cli
set jeff "whats up"
get jeff
del jeff
exit
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#overview)
* [INSTALL REDIS (SERVER) & REDIS-CLI (CLIENT)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#install-redis-server--redis-cli-client)
* [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#configure)
* [TEST YOU CAN RUN IT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#test-you-can-run-it)
* [START ON BOOT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#start-on-boot)
* [REDIS-CLI COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#redis-cli-commands)
* [TO USE WITH GO](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/redis-cheat-sheet#to-use-with-go)

Documentation and Reference

* [redis website](https://redis.io/)
* Full list of
  [redis-cli commands](https://redis.io/commands#hash)
* To use with go, refer to my repo
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples?tab=readme-ov-file#databases)
* A great list of using
  [go with databases](https://github.com/gostor/awesome-go-storage)
* Check out my relational database cheat sheet on
  [postgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/database/postgresql-cheat-sheet)

## OVERVIEW

Redis is an open source key-value store,
often referred to as a NoSQL database.
The essence of a key-value store is the ability
to store some data, called a value, inside a key.
This data can later be retrieved only if we know the
exact key used to store it.

There are a few ways to interact with your database (server),

* Command Line
  * `redis-server` - The server itself
  * `redis-sentinel`
  * `redis-benchmark`
  * `redis-check-aof`
  * `redis-check-dump`
* redis-cli (client)
* [github.com/go-redis/redis go library](https://github.com/go-redis/redis)

Here is an illustration,

![IMAGE - redis-server - IMAGE](../../../../../docs/pics/software/development/redis-server.svg)

## INSTALL REDIS (SERVER) & REDIS-CLI (CLIENT)

Install from
[here](https://redis.io/download).

### LINUX FROM SOURCE

```bash
wget http://download.redis.io/releases/redis-5.0.5.tar.gz
tar xzf redis-5.0.5.tar.gz
rm redis-5.0.5.tar.gz
cd redis-5.0.5
make
cd ..
sudo mv redis-5.0.5 /usr/lib
```

Add command line tools path ~/.bashrc

```bash
export PATH=/usr/lib/redis-5.0.5/src:$PATH
```

### MACOS FROM SOURCE

```bash
curl -O http://download.redis.io/releases/redis-5.0.5.tar.gz
tar xzf redis-5.0.5.tar.gz
rm redis-5.0.5.tar.gz
cd redis-5.0.5
make
cd ..
sudo mv redis-5.0.5 /usr/local/lib
```

Add command line tools path ~/.bashrc

```bash
export PATH=/usr/local/lib/redis-5.0.5/src:$PATH
```

### CHECK VERSIONS

Server,

```bash
redis-server --version
```

Client,

```bash
redis-cli --version
```

## CONFIGURE

There are a few simples steps to properly configure redis
on your machine.  We will be using port 6379.

### CONFIGURATION FILE

The configuration template file is `redis.conf`, located in
`/usr/lib/redis-5.0.5`.

Place in `/etc/redis`,

```bash
sudo mkdir /etc/redis
sudo cp /usr/lib/redis-5.0.5/redis.conf /etc/redis/6379.conf
```

Or for macOS,

```bash
sudo mkdir /etc/redis
sudo cp /usr/local/lib/redis-5.0.5/redis.conf /etc/redis/6379.conf
```

Edit,

```bash
sudo nano /etc/redis/6379.conf
```

* Set daemonize to yes (default set to no).
* Set the pidfile to `/var/run/redis_6379.pid`.
* Change the port accordingly.
* Set your preferred loglevel.
* Set the logfile to `/var/log/redis_6379.log`.
* Set the dir to `/var/redis/6379` and dbfilename to `jeff-redis.rdb`.

### DATABASE STORAGE AREA

Create the database storage area,

```bash
sudo mkdir /var/redis
sudo mkdir /var/redis/6379
```

But most data is kept in memory for performance.

You could,

* Issue the save command
* Edit your config file

the default is,

```txt

save 900 1
save 300 10
save 60 10000
```

After 900 sec (15 min) if at least 1 key changed
After 300 sec (5 min) if at least 10 keys changed
After 60 sec if at least 10000 keys changed

## TEST YOU CAN RUN IT

```bash
redis-server /etc/redis/6379.conf
```

Check if the redis-server server is running,

```bash
ps aux | grep -i redis
redis-cli ping
redis-cli shutdown
```

Shutdown saves the data file.

Check log if you have issues,

```bash
tail /var/log/redis_6379.log
```

## START ON BOOT

Lets start the redis database server on boot.

### LINUX - USE THE OLD SysV INIT SCRIPT

SysV init script are old and really should be updated with systemd
.service scripts but since redis still use them,
we will do it.  Note the `systemd` actually runs these scripts,
not the old init system.

For more information on `systemd` and `init` refer to my cheat sheet
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/systemd-cheat-sheet)

Create an init.d script,

```bash
sudo cp /usr/lib/redis-5.0.5/utils/redis_init_script /etc/init.d/redis_6379
```

Edit,

```bash
sudo nano /etc/init.d/redis_6379
```

Change the following to look like,

```bash
EXEC=/usr/lib/redis-5.0.5/src/redis-server
CLIEXEC=/usr/lib/redis-5.0.5/src/redis-cli
```

Add to init,

```bash
sudo update-rc.d redis_6379 defaults
```

Start/stop (As mentioned above systemd is actually running the service manager),

```bash
sudo /etc/init.d/redis_6379 start
sudo /etc/init.d/redis_6379 stop
```

### MACOS - LAUNCHD

Start redis-server using `launchd`.

For more information on launchd refer to my cheat sheet
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/macos/launchd-cheat-sheet)

Create a configuration property list file for your daemon (redis).
You can call it what you like.

```bash
nano ~/Library/LaunchAgents/local.redis.redis-server.plist
```

with,

```json
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>local.redis.redis-server</string>
    <key>Program</key>
    <string>/usr/local/lib/redis-5.0.5/src/redis-server</string>
    <key>ProgramArguments</key>
    <array>
        <string>/etc/redis/6379.conf</string>
    </array>
    <key>KeepAlive</key>
    <true/>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardErrorPath</key>
    <string>/var/log/redis_6379.log</string>
    <key>StandardOutPath</key>
    <string>/var/log/redis_6379.log</string>
</dict>
</plist>
```

Lets load it,

```bash
launchctl load ~/Library/LaunchAgents/local.redis.redis-server.plist
```

Check if launchd launched it (does not mean redis is working),

```bash
launchctl list | grep redis
```

### CHECK IF RUNNING

Check if the redis-server server is running,

```bash
ps aux | grep -i redis
redis-cli ping
```

Check log if issues,

```bash
tail /var/log/redis_6379.log
```

## REDIS-CLI COMMANDS

Start,

```bash
redis-cli
```

Set/get/delete,

```bash
set jeff "whats up"
get jeff
del jeff
exit
```

But it probably makes more sense to be more organized,

Set/get more organized,

```bash
set people:jeff "whats up"
set people:larry "yo yo"
get people:jeff
exit
```

Expire data,

```bash
set bye "this will expire in 20 seconds"
expire bye 20
get bye
```

Lists,

* RPUSH puts the new value at the end of the list
* LPUSH puts the new value at the start of the list
* LRANGE gets a subset of the list
* LLEN returns the current length of the list.
* LPOP removes the first element from the list and returns it.
* RPOP removes the last element from the list and returns it.

```bash
rpush friends "Alice"
rpush friends "Bob"
lpush friends "Sam"
lrange friends 0 -1
llen friends
```

Set is like a list, but no order,

```bash
sadd superpowers "flight"
sadd superpowers "x-ray vision"
sadd superpowers "reflexes"
smembers superpowers
```

Full list of redis-cli commands
[here](https://redis.io/commands#hash).

## TO USE WITH GO

Refer to my repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/database/redis).
