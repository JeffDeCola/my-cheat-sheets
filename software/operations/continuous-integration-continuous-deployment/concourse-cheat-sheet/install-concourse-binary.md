# INSTALL CONCOURSE BINARY

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
**** OLD - ARCHIVED INFORMATION ****
```

_Install concourse using binaries files._

Table of Contents

* [CONCOURSE BINARY AND KEYS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md#concourse-binary-and-keys)
* [INSTALL AND CONFIGURE POSTGRESQL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md#install-and-configure-postgresql)
* [START CONCOURSE WEB (SINGLE NODE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md#start-concourse-web-single-node)
* [START CONCOURSE WORKER (SINGLE MODE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md#start-concourse-worker-single-mode)
* [FLY](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md#fly)

Documentation and Reference

* [concourse](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/operations/continuous-integration-continuous-deployment/concourse-cheat-sheet/README.md)

## CONCOURSE BINARY AND KEYS

Get the latest concourse binary,

```bash
wget https://github.com/concourse/concourse/releases/download/v2.7.3/concourse_linux_amd64
```

mv to `/usr/local/bin` and `chmod 775`,

```bash
sudo mv concourse_linux_amd64 /usr/local/bin/concourse
chmod 775 /usr/local/bin/concourse
```

Generate 3 keys,

```bash
ssh-keygen -t rsa -f tsa_host_key -N ''
ssh-keygen -t rsa -f worker_key -N ''
ssh-keygen -t rsa -f session_signing_key -N ''
```

```bash
cp worker_key.pub authorized_worker_keys
```

## INSTALL AND CONFIGURE POSTGRESQL

Install postgreSQL Server,

```bash
sudo apt-get update
sudo apt-get -y install postgresql postgresql-client postgresql-contrib
sudo -u postgres psql postgres

> \password postgres
> "admin"
> CREATE EXTENSION adminpack;
> \q
```

Check pg_hba.conf,

```bash
sudo nano /etc/postgresql/9.5/main/pg_hba.conf
```

Make sure you have these two lines,

`host all all 127.0.0.1/32 md5`

`host all all ::1/128 md5`

Optionally I think you could do,

`host concourse concourse 127.0.0.1/32 md5`

`host atc concourse 127.0.0.1/32 md5`

Edit postgresql.conf,

```bash
sudo nano /etc/postgresql/9.5/main/postgresql.conf
```

Replace `listen_addresses = 'localhost'` with `listen_addresses = '*'`

Restart postgres,

```bash
sudo service postgresql restart
```

Change user postgres,

```bash
sudo su - postgres
```

Open psql as user postgres,

```bash
psql
> create database atc;
> create database concourse;
> create user concourse password 'admin';
```

Check that role was created correctly,

```bash
SELECT rolname FROM pg_roles;
```

Exit psql `ctrl+D`.

Exit postgres user `exit`.

## START CONCOURSE WEB (SINGLE NODE)

Now lets start a Single node, local Postgres.

Spin up the ATC, listening on port 8080, with some basic auth configured,
and a TSA listening on port 2222.

This assumes you have a local Postgres server running on the default port (5432)
with an atc database, accessible by the current user.

```bash
sudo concourse web \
  --basic-auth-username USERNAME \
  --basic-auth-password PASSWORD \
  --session-signing-key session_signing_key \
  --tsa-host-key tsa_host_key \
  --tsa-authorized-keys authorized_worker_keys \
  --external-url http://<CONCOURSEIP> \
  --postgres-data-source postgres://concourse:admin@localhost/concourse &
```

## START CONCOURSE WORKER (SINGLE MODE)

```bash
  sudo concourse worker \
  --work-dir /opt/concourse/worker \
  --tsa-host 127.0.0.1 \
  --tsa-public-key tsa_host_key.pub \
  --tsa-worker-private-key worker_key &
```

Check it out in a browser,

```bash
http://<CONCOURSEIP>:8080/
```

Check ps,

```bash
ps -aux | grep concourse
```

## FLY

Download fly executable to your machine.

Login to fly using "main" team (default),

```bash
fly -t ci login -c http://<CONCOURSEIP>:8080
```

Enter the above username and password.
