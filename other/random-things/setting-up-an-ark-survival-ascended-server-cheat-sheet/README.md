# SETTING UP AN ARK: SURVIVAL ASCENDED (ASA) SERVER

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Setting up an ARK: Survival Ascended (ASA) server on a linux box (with proton)
using the cdp1337/ARKSurvivalAscended-Linux
community install script._

```text
Manage server:          sudo /home/steam/ArkSurvivalAscended/manage.py
Check status:           htop or sudo ss -ulnp | grep 7708
View logs:              sudo tail -f /home/steam/ArkSurvivalAscended/ShooterGame.log
Gameplay Config:        /home/steam/ArkSurvivalAscended/Game.ini
Player Facing Config:   /home/steam/ArkSurvivalAscended/GameUserSettings.ini
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#overview)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#install)
  * [STEP 1 - RUN THE CDP1337 INSTALL SCRIPT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#step-1---run-the-cdp1337-install-script)
  * [STEP 2 - START YOUR SERVER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#step-2---start-your-server)
  * [KEY FILE LOCATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#key-file-locations)
* [FIREWALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#firewall)
* [PORT FORWARDING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#port-forwarding)
  * [MAP PORTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#map-ports)
  * [MULTI-ROUTER TOPOLOGY EXAMPLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#multi-router-topology-example)
* [CONFIGURATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#configuration)
  * [SERVER PASSWORD (FRIENDS ONLY)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#server-password-friends-only)
  * [GAME PLAY CONFIG](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#game-play-config)
  * [PLAYER FACING CONFIG](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#player-facing-config)
* [CONNECTING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#connecting)
  * [FROM LOCAL NETWORK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#from-local-network)
  * [FROM THE INTERNET](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#from-the-internet)
* [BACKUPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#backups)
  * [MANUAL BACKUP](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#manual-backup)
  * [AUTOMATED DAILY BACKUPS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#automated-daily-backups)
  * [BACKUP LOCATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#backup-locations)
* [ADMIN CHEAT COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#admin-cheat-commands)
  * [ENABLE/DISABLE CHEATS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#enabledisable-cheats)
  * [COMMON COMMANDS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#common-commands)
  * [SPAWNING ITEMS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#spawning-items)
  * [TAMING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#taming)
  * [SERVER MANAGEMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/random-things/setting-up-an-ark-survival-ascended-server-cheat-sheet/#server-management)
  
Documentation and Reference

* [cdp1337 GitHub](https://github.com/cdp1337/ARKSurvivalAscended-Linux)
* [ARK Official Wiki - Dedicated Server Setup](https://ark.wiki.gg/wiki/Dedicated_server_setup)
* [SteamDB - App 2430930](https://steamdb.info/app/2430930/info/)

## OVERVIEW

ARK: Survival Ascended does not have a native Linux dedicated server.
The community script by cdp1337 uses Proton (a Windows compatibility layer)
to run the Windows server binary on Linux. It handles SteamCMD, Proton,
systemd, and provides a management console for starting/stopping servers,
switching maps, managing mods, backups, and more.

## INSTALL

### STEP 1 - RUN THE CDP1337 INSTALL SCRIPT

This one command installs everything: SteamCMD, Proton, ARK server files,
systemd services, and the management console.

```bash
sudo su -c "bash <(wget -qO- https://raw.githubusercontent.com/cdp1337/ARKSurvivalAscended-Linux/main/dist/server-install-debian12.sh)" root
```

It will ask you:

* **Community name**: Server Name (e.g. JeffsArkServer)
* **Enable whitelist**: No for a private server with password
* **Enable multi-server cluster support**: No unless you want maps on diff servers
* **Enable system firewall (UFW)**: Yes

### STEP 2 - START YOUR SERVER

Open the management console

```bash
sudo /home/steam/ArkSurvivalAscended/manage.py
```

You'll see a menu with all available maps. Press the number for the map you
want (e.g. `5` for Ragnarok), then press `S` to start it.

Press `E` to enable auto-start so the server starts on boot.

### KEY FILE LOCATIONS

* Game files: `/home/steam/ArkSurvivalAscended/AppFiles/`
* Config: `/home/steam/ArkSurvivalAscended/GameUserSettings.ini`
* Game config: `/home/steam/ArkSurvivalAscended/Game.ini`
* Log: `/home/steam/ArkSurvivalAscended/ShooterGame.log`
* Admin list: `/home/steam/ArkSurvivalAscended/admins.txt`
* Services: `/home/steam/ArkSurvivalAscended/services/`

## FIREWALL

The install script handles UFW rules automatically (ports 7701-7710/udp
and 27001-27010/tcp). You can verify with:

```bash
sudo ufw status
```

Make sure SSH is allowed:

```bash
sudo ufw allow 22/tcp
```

## PORT FORWARDING

For friends to connect from the internet, you need to forward the ports
on your router(s) to your server's local IP.

### MAP PORTS

Each map gets its own port:

| Map               | Game Port (UDP) | RCON Port (TCP) |
|-------------------|-----------------|-----------------|
| TheIsland_WP      | 7701            | 27001           |
| Aberration_WP     | 7702            | 27002           |
| BobsMissions_WP   | 7703            | 27003           |
| ScorchedEarth_WP  | 7704            | 27004           |
| TheCenter_WP      | 7705            | 27005           |
| Extinction_WP     | 7706            | 27006           |
| Astraeos_WP       | 7707            | 27007           |
| Ragnarok_WP       | 7708            | 27008           |
| Valguero_WP       | 7709            | 27009           |
| LostColony_WP     | 7710            | 27010           |

Forward the game port (UDP) for whichever map you're running to your
server's local IP on each router in the chain.

### MULTI-ROUTER TOPOLOGY EXAMPLE

If you have multiple routers like me

```text
Internet
Pepsi (192.168.1.x)
Mookie (192.168.10.x)
King (192.168.20.x)
Server (192.168.20.115)
```

Each router forwards the same port(s) to the next hop:

* **Pepsi**: forward port 7708 UDP → Mookie's IP on the 1.x network
* **Mookie**: forward port 7708 UDP → King's IP on the 10.x network
* **King**: forward port 7708 UDP → 192.168.20.115

**NOTE**: Only forward the game port (UDP).
Do NOT forward the RCON/query port (TCP).
The query port is what Steam uses to list your server publicly.
If you don't forward it,
your server stays invisible in the server browser.
Friends connect via direct IP instead.

I would also add fail2ban to stop anyone from ssh attacks.
It's stock anyone from trying password more than 5 times.
Locks them out for 10 minutes.

```bash
sudo apt install fail2ban -y
```

Check it's working

```bash
sudo systemctl status fail2ban
```

## CONFIGURATION

### SERVER PASSWORD (FRIENDS ONLY)

> THIS DOES NOT WORK YET - THERE IS A BUG - SKIP FOR NOW

Edit the config file:

```bash
sudo nano /home/steam/ArkSurvivalAscended/GameUserSettings.ini
```

Add under `[ServerSettings]`:

```ini
ServerPassword=yourpassword
```

### GAME PLAY CONFIG

Game.ini (/home/steam/ArkSurvivalAscended/Game.ini)
This is deeper "game rule" overrides. Things like:

* Engram overrides (unlock/lock specific engrams)
* Loot crate quality and contents
* Per-level stat multipliers for dinos and players
* NPC replacements (swap or remove specific dinos)
* Stack size overrides
* Custom recipes

### PLAYER FACING CONFIG

GameUserSettings.ini (/home/steam/ArkSurvivalAscended/GameUserSettings.ini)
are basically anything you'd tweak to change the gameplay experience.

* Server name, passwords
* Taming speed, harvest amount, XP multipliers
* Day/night cycle speed
* Player stat multipliers (health, stamina, weight)
* Difficulty level
* Max players

## CONNECTING

### FROM LOCAL NETWORK

Open ARK: Survival Ascended, open the console (Tab key), and type

```text
open 192.168.20.115:7708
```

### FROM THE INTERNET

Your friends use your public IP (found in the manage.py service details):

```text
open <your ip>:7708
```

> NOTE: The password be here if it worked

## BACKUPS

### MANUAL BACKUP

In the management console (`manageark`), press `B` to create a backup.

### AUTOMATED DAILY BACKUPS

Created a backup script

```bash
nano /home/arkserver/backup_ark.sh
```

```bash
#!/bin/bash

BACKUP_DIR=/home/arkserver/ark-backups
SAVE_DIR=/home/steam/ArkSurvivalAscended/AppFiles/ShooterGame/Saved/SavedArks/Ragnarok_WP
DATE=$(date +%Y-%m-%d_%H-%M)

mkdir -p $BACKUP_DIR

# Create compressed backup
sudo tar -czf $BACKUP_DIR/ragnarok_backup_$DATE.tar.gz -C $SAVE_DIR .

# Delete backups older than 7 days
find $BACKUP_DIR -name "*.tar.gz" -mtime +7 -delete

echo "Backup complete: ragnarok_backup_$DATE.tar.gz"
```

Set up cron job to run daily at 4am:

```bash
sudo crontab -e
```

Add this line to backup every 2 hours:

```text
0 */2 * * * /home/arkserver/backup_ark.sh
```

Test it now to make sure it works:

```bash
sudo sh /home/arkserver/backup_ark.sh
```

Then check:

```bash
ls -la /home/arkserver/ark-backups/
```

### BACKUP LOCATIONS

* Automated backups: `/home/arkserver/ark-backups/`
* Nitrado import files: `/home/arkserver/files-from-nitrado/`
* Original fresh server backup: `/home/steam/ArkSurvivalAscended/AppFiles/ShooterGame/Saved/SavedArks/Ragnarok_WP_ORIGINAL_BACKUP/`

## ADMIN CHEAT COMMANDS

### ENABLE/DISABLE CHEATS

Enable/disable cheats

```text
enablecheats YOUR_ADMIN_PASSWORD
admincheat disablecheats
```

### COMMON COMMANDS

| Command | What it does |
|---|---|
| `cheat fly` | Fly mode |
| `cheat walk` | Back to walking |
| `cheat ghost` | No-clip through walls |
| `cheat god` | Invincible |
| `cheat infinitestats` | Unlimited health/stamina/food/water/weight |
| `cheat giveresources` | Gives 50 of every resource |
| `cheat addexperience 1000 0 0` | Give yourself XP |
| `cheat settimeofday 12:00` | Set time to noon |
| `cheat destroywilddinos` | Wipe all wild dinos (forces respawn) |
| `cheat saveworld` | Force save the world |
| `cheat teleport` | Teleport where you're looking |
| `cheat setplayerpos X Y Z` | Teleport to coordinates |

### SPAWNING ITEMS

| Command | What it does |
|---|---|
| `cheat gfi MetalIngot 100 0 0` | 100 metal ingots |
| `cheat gfi StoneWall 100 0 0` | 100 stone walls |
| `cheat gfi MetalWall 100 0 0` | 100 metal walls |
| `cheat gfi CementPaste 100 0 0` | 100 cementing paste |
| `cheat gfi Polymer 100 0 0` | 100 polymer |
| `cheat gfi Electronics 100 0 0` | 100 electronics |

### TAMING

| Command | What it does |
|---|---|
| `cheat forcetame` | Instantly tame what you're looking at |
| `cheat dotame` | Tame without riding ability |
| `cheat settargetdinocolor 0 1` | Change dino color (region 0, color 1) |

### SERVER MANAGEMENT

| Command | What it does |
|---|---|
| `cheat broadcast message` | Send message to all players |
| `cheat kickplayer EOSID` | Kick a player |
| `cheat saveworld` | Force save |
| `cheat listplayers` | Show connected players |
| `cheat whoami` | Show your EOS ID | Opus 4.6Extended
