# INSTALL CONCOURSE USING DOCKER-COMPOSE

_This is a cheat sheet on how to install concourse using docker-compose._

Table on Contents

* [INSTALL USING DOCKER-COMPOSE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#install-using-docker-compose)
* [USER A STATIC IP TO ACCESS OUTSIDE YOUR MACHINE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#user-a-static-ip-to-access-outside-your-machine)
* [ALLOW LOCAL NETWORK ACCESS TO VM (BRIDGE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#allow-local-network-access-to-vm-bridge)
* [IF YOU GET AN IMAGE ERROR](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#if-you-get-an-image-error)
* [START YOUR DOCKER-COMPOSE CONTAINERS AUTOMATICALLY AT BOOT/CRASH (USING SYSTEMD)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#start-your-docker-compose-containers-automatically-at-bootcrash-using-systemd)
* [USER YOUR OWN POSTGRES DATABASE RATHER THEN DOCKER](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose.md#user-your-own-postgres-database-rather-then-docker)

## INSTALL USING DOCKER-COMPOSE

Get the `docker-compose.yml` file from concourse,

```bash
curl -O https://concourse-ci.org/docker-compose.yml
```

Edit the `docker-compose.yml` file to your liking.

```yaml
CONCOURSE_EXTERNAL_URL: http://192.168.20.112:8080
comment out # cgroup: host
```

Run docker-compose up. This will run two docker containers:
concourse (web and worker) and concourse-db.

```bash
docker-compose up -d
```

Note that this will pull two docker images,

* postgres
* concourse/concourse

```bash
docker images
```

Default team is main and the User/Password is test/test.

To stop all docker containers is simple,

```bash
docker-compose down
```

## USER A STATIC IP TO ACCESS OUTSIDE YOUR MACHINE

If you want to access this from outside you box,
first setup a static IP to a network device.
I have a cheat sheet on that called
[network device configuration](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/network-device-configuration-cheat-sheet)

Remember to match the static IP to what's in your
`docker-compose.yml` file,

```text
    - CONCOURSE_EXTERNAL_URL=http://192.168.100.6:8080
```

## ALLOW LOCAL NETWORK ACCESS TO VM (BRIDGE)

This is actually quite easy.

Instead of NATS, you use a bridged adapter,

```txt
Settings -> Network - Adapter -> Bridged Adapter
```

Your router (on the host machine's network) will choose an
IP for you.  You can go into that router to save the
IP based on the mac address if you would like.

Remember to match the static IP to what's in your
`docker-compose.yml` file,

Also, you do not need set your machine with a static IP.
Let your router assign one.

## IF YOU GET AN IMAGE ERROR

Start fresh, so prune the containers and volumes,

```bash
sh clean-up.sh
```

or

```bash
docker container prune
docker volume prune
```

And make sure to stop all running containers,

```bash
docker-compose down
```

## START YOUR DOCKER-COMPOSE CONTAINERS AUTOMATICALLY AT BOOT/CRASH (USING SYSTEMD)

I wrote a cheat sheet on this which has this info;
[systemd systemctl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/systemd-systemctl-cheat-sheet)

The goal is to start `start-concourse.sh` on boot/crash.

Note, may have to `chmod 755` this file.

Create your `concourse.service` file and
place in `/etc/systemd/system`,

```bash
sudo nano /etc/systemd/system/concourse.service
```

```yaml
[Unit]
Description=Concourse service with docker compose
Requires=docker.service
After=docker.service

[Service]
Restart=always

WorkingDirectory=/your/working/path

ExecStartPre=/your/working/path/clean-up.sh
ExecStart=/your/working/path/start-concourse.sh
ExecStop=/your/working/path/stop-concourse.sh

[Install]
WantedBy=multi-user.target
```

May also have to `chmod 664` this file.

```bash
sudo chmod 644 /etc/systemd/system/concourse.service
```

Tell the system about this new unit service,

```bash
sudo systemctl daemon-reload
```

Start and Enable on boot,

```bash
sudo systemctl start concourse
sudo systemctl enable concourse
```

If you want, check how systemctl is doing,

```bash
journalctl -f
```

Changes will come to effect on a reboot,

```bash
sudo reboot
```

## USER YOUR OWN POSTGRES DATABASE RATHER THEN DOCKER

Instead of using docker postgres, lets use our own on the server itself.

Install (unless you already have it),

```bash
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```

Create user concourse and database atc via the user postgres,

```bash
sudo -u postgres createuser -P concourse
```

Prompt to enter the password `concourse-password`.

```bash
sudo -u postgres createdb --owner=concourse atc
```

Or in psql,

```text
CREATE DATABASE yourdbname;
CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpass';
GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;
```

Check you have a database. Switch to user postgres and
enter `psql`,

```bash
sudo su - postgres
psql
\l
```

And list the users,

```bash
\du
```

Edit,

```bash
sudo nano /etc/postgresql/<VERSION>/main/pg_hba.conf
```

Add,

```text
 host all all 192.168.100.6/24 md5
```

Where `type database user address method`.

Hence we must add a user concourse,

```bash
sudo adduser --system --group concourse
```

To restart postgres,

```bash
sudo service postgresql restart
```

You can get more info on postgreSQL at my cheat-sheet
[postgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/database/postgreSQL-cheat-sheet).
