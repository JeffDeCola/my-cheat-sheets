# postgreSQL CHEAT SHEET

`postgreSQL` _is a open source object-relational database system._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL AND CONFIGURE postgreSQL on UBUNTU/DEBIAN

Install postgresSQL Server,

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

Exit psql `ctrl+D` 

Exit postgres user `exit`.

## USE

Become user postgres,

```bash
sudo su - postgres
```

Open psql as user postgres,

## GET THIS CODE
 
 
Select * From Case ac
join templates as t
    on +.ID = c.teamplte_ID
    when +.ID = {}

