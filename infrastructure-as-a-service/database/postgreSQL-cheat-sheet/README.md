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

## INSTALL ON MAC OS

```bash
brew install postgres
```

## STATE SERVER

```bash
pg_ctl -D /usr/local/var/postgres start
```

``` check
ps aux | grep -i postgres
```

## USE

Become user postgres,

```bash
sudo su - postgres
psql
```

Open psql as user postgres,

```bash
psql -d postgres
```

Quit,

```bash
\q
```

## CREATE AND CONNECT TO A DATABASE

Create a database,

```bash
create database rm
```

List all databases,

```bash
\l
```

Connect to a database,

```bash
\c rm
```

## CREATE TABLE

Some types you can do,

```bash
create table people (id int primary key not null, first_name text, last_name text);
```

## LIST TABLE

list all tables in a database,

```bash
\d
```

list a table in database,

```bash
\d people
\d+ people
```

## LIST ROWS OF TALBE

```bash
select * from people;
select last_name from people;
```