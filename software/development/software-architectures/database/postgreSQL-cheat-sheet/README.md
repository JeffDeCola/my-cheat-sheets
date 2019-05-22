# postgreSQL CHEAT SHEET

`postgreSQL` _is a open source object-relational database system._

Documentation and reference,

* [postgreSQL website](https://www.postgresql.org/)
* You use this command a lot to run psql `sudo -u postgres psql`
* To use with go, refer to my repo
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/database/postgreSQL)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## HISTORY

PostgreSQL is an object-relational database management system
(ORDBMS) based on POSTGRES, Version 4.2, developed at the
University of California at Berkeley Computer Science Department.

The origins of PostgreSQL date back to 1986 as part of the POSTGRES
project at the University of California at Berkeley and has more than
30 years of active development on the core platform.

### OVERVIEW

There are a few ways to interact with your database (server),

* Command Line
* psql (client)
* go library
* pgAdmin 4

Here is an illustration,

![IMAGE - postgreSQL-server - IMAGE](../../../../../docs/pics/postgreSQL-server.jpg)

## INSTALL POSTGRES (SERVER), PSQL (CLIENT) & PGADMIN4

Install from
[here](https://www.postgresql.org/download).

### UBUNTU/DEBIAN

Install postgreSQL Server,

```bash
sudo apt-get update
sudo apt-get -y install postgresql postgresql-client postgresql-contrib
sudo apt-get -y install pgadmin3
```

Add command line tools path ~/.bashrc

```bash
export PATH=/usr/lib/postgresql/9.5/bin:$PATH
export PATH=/usr/lib/postgresql/10/bin:$PATH
```

### MAC OS

I downloaded the interactive 
[installer](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads)
by EnterpriseDB which includes,

* PostgreSQL Server
* pgAdmin - graphical tool for managing database
* Stackbuilder - package manager
* Command Line Tools (e.g. createdb)

Add command line tools path ~/.bashrc

```bash
export PATH=/Library/PostgreSQL/11/bin:$PATH
```

You could also use brew,

```bash
brew update
brew install postgresql
```

### CHECK VERSION

Server,

```bash
postgres -V
```

Client,

```bash
psql --version
```

pgadmin,

```bash
pgadmin3 -v
```

### pgAdmin

`pgAdmin` is a graphical tool for managing and
developing your databases.

Install from [here](https://www.pgadmin.org/download/)

It will run in your browser
[http://127.0.0.1:50088/browser/]http://127.0.0.1:50088/browser/)

## CONFIGURE

You may need to configure depending how you installed.

Enter postgres using psql as user postgres (yeah its a little confusing),

```bash
sudo -u postgres psql postgres
> \password postgres
> "admin"
> CREATE EXTENSION adminpack;
> \q
```

Check `pg_hba.conf`,

```bash
sudo nano /etc/postgresql/<VERSION>/main/pg_hba.conf
```

Edit `postgresql.conf`,

```bash
sudo nano /etc/postgresql/<VERSION>/main/postgresql.conf
sudo nano /etc/postgresql/10/main/postgresql.conf
```

Replace `listen_addresses = 'localhost'` with `listen_addresses = '*'`

To restart postgres server,

```bash
sudo service postgresql restart
```

## START/STOP POSTGRESQL SERVER

Check the postgreSQL server is running,

```bash
ps aux | grep -i postgres
```

Start,

```bash
pg_ctl -D /usr/local/var/postgres start
```

Stop,

```bash
pg_ctl -D /usr/local/var/postgres stop
``` 

## USING PSQL (CLIENT)

Become user `postgres`,

```bash
sudo su - postgres
./psql
```

But I like to use,

```bash
sudo -u postgres psql
```

Quit,

```bash
\q
```

## CREATE A USER

Create a user `jeffd`,

Method 1 - Create user using command line as user postgres,

```bash
sudo su - postgres
./createuser jeffd
```

Method 2 - Create user using command line without becoming user postgres,

```bash
sudo -u postgres createuser jeffd
```

Method 3 - Create user using psql (REMEMBER TO END WITH `;`),

```bash
sudo -u postgres psql
CREATE USER jeffd;
CREATE USER jeffd WITH ENCRYPTED PASSWORD 'yourpass';
```

Check user was created,

```bash
sudo -u postgres psql
\du
```

## CREATE A DATABASE

Create database `jeff_db_example`,

Method 1 - Create database using command line as user postgres,

```bash
sudo su - postgres
./creatdb --owner=jeffd jeff_db_example
```

Method 2 - Create database using command line without becoming user postgres,

```bash
sudo -u postgres createdb --owner=jeffd jeff_db_example
```

Method 3 - Create database using psql (REMEMBER TO END WITH `;`),

```bash
sudo -u postgres psql
CREATE DATABASE jeff_db_example3 OWNER jeffd;
```

List all databases from psql,

```bash
sudo -u postgres psql
\l
```

## CONNECT TO A DATABASE

Connect to a `database` using psql,

```bash
sudo -u postgres psql
\c jeff_db_example
```

Then you can do things like create a table,

## CREATE A TABLE (YOUR SCHEMA)

Some types you can do,

```bash
CREATE TABLE (id int primary key not null, first_name text, last_name text);
```

## LIST TABLES

List all tables in a database your connected to,

```bash
\d
```

List a table in database,

```bash
\d people
\d+ people
```

## LIST ROWS OF TABLE

```bash
select * from people;
select last_name from people;
```
