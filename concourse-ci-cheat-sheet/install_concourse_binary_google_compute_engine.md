# INSTALL CONCOURSE BINARY GOOGLE COMPUTE ENGINE

First get your ubuntu instance running on Google Copmpute Engine
using at least 1 CPU.

ssh into you instance.

## CONCOURSE BINARY AND KEYS

Get the concourse binary,

```bash
wget https://github.com/concourse/concourse/releases/download/v2.7.3/concourse_linux_amd64
```

mv to /usr/local/bin and chmod 775,

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

## INSTALL AND CONFIGURE POSTGRESSQL

Install postgresSQL Server,

```bash
sudo apt-get update
sudo apt-get -y install postgresql postgresql-client postgresql-contrib
sudo -u postgres psql postgres
> \password postgres
> enter password
> admin
> CREATE EXTENSION adminpack;
> \q
```

Edit pg_hba.conf,

```bash
sudo nano /etc/postgresql/9.5/main/pg_hba.conf
```

Make sure you have,

`host all all 127.0.0.1/32 md5`

Optionally I think I could of done,

`host concourse concourse 127.0.0.1/32 md5`
`host atc concourse 127.0.0.1/32 md5`

Make sure you have these two lines (I believe this is done),

`host all all 127.0.0.1/32 md5`
`host all all ::1/128 md5`

Edit postgresql.conf,

```bash
sudo nano /etc/postgresql/9.5/main/postgresql.conf
```

Replace `listen_addresses = 'localhost'` with `listen_addresses = '*'`

Restart postgres,

```bash
sudo service postgresql restart
```

Change user postgres

```bash
sudo su - postgres
```

Open psql as user postgres,

```bash
psql
```

```bash
create database atc;
create database concourse;
create user concourse password 'admin';
```

Check that role was created correctly,

```bash
SELECT rolname FROM pg_roles;
```

Exit psql `ctrl+D` and Exit postgres user `exit`.

## START CONCOURSE WEB

Now lets start a Single node, local Postgres. Spin up the ATC, 
listening on port 8080, with some basic auth configured, and a TSA listening on port 2222.
This assumes you have a local Postgres server running on the default port (5432) 
with an atc database, accessible by the current user.

```bash
concourse web \
  --basic-auth-username USERNAME \
  --basic-auth-password PASSWORD \
  --session-signing-key session_signing_key \
  --tsa-host-key tsa_host_key \
  --tsa-authorized-keys authorized_worker_keys \
  --external-url http://10.240.0.79 \
  --postgres-data-source postgres://concourse:admin@localhost/concourse &
```

## START CONCOURSE WORKER

```bash
  sudo concourse worker \
  --work-dir /opt/concourse/worker \
  --tsa-host 127.0.0.1 \
  --tsa-public-key tsa_host_key.pub \
  --tsa-worker-private-key worker_key &
```

Check it out

`http://<ip>:8080/`

## FLY

Login to fly using "main" team,

```bash
fly -t ci login -c http://<ip>:8080
```

Enter the above username and password.
