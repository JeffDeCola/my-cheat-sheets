# postgreSQL CHEAT SHEET

`postgreSQL` _is a open source object-relational database system._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL AND CONFIGURE postgreSQL on UBUNTU/DEBIAN

Install postgresSQL Server,

```bash
sudo apt-get update
sudo apt-get -y install postgresql postgresql-client postgresql-contrib
```

check version,

```bash
psql --version
```

Enter postgres using psql,

```bash
sudo -u postgres psql postgres

> \password postgres
> "admin"
> CREATE EXTENSION adminpack;
> \q
```

Check pg_hba.conf,

```bash
sudo nano /etc/postgresql/<VERSION>/main/pg_hba.conf
```

Edit postgresql.conf,

```bash
sudo nano /etc/postgresql/<VERSION>/main/postgresql.conf
```

Replace `listen_addresses = 'localhost'` with `listen_addresses = '*'`

To restart postgres,

```bash
sudo service postgresql restart
```

Open `psql` as user `postgres`,

```bash
sudo su - postgres
```

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

## USING PSQL

Become user `postgres`,

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

## CREATE A USER

For example, create user concourse,

```bash
sudo su - postgres
createuser concourse
```

or

```bash
sudo -u postgres createuser concourse
```

## CREATE A DATABASE / CONNECT TO A DATABASE

Create a database,

```bash
sudo -u postgres createdb --owner=concourse atc
```

```bash
sudo su - postgres
creatdb --owner=concourse <DBNAME>
```

or,

```bash
createdb --owner=concourse <DBNAME>
```

List all databases,

```bash
psql
\l
```

Connect to a `database` as user `postgres`,

```bash
\c <DBNAME>
```

## CREATE TABLE

Some types you can do,

```bash
create table people (id int primary key not null, first_name text, last_name text);
```

## LIST TABLE

List all tables in a database,

```bash
\d
```

List a table in database,

```bash
\d people
\d+ people
```

## LIST ROWS OF TALBE

```bash
select * from people;
select last_name from people;
```