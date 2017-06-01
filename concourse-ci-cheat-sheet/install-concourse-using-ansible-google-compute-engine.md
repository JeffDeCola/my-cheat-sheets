# INSTALL CONCOURSE USING AN ANSIBLE ROLE ON GCE

The following will install and configure concourse as 2 web and 4 workers.

The following ansible role repos are used,

* [ansible-concourse](https://github.com/ahelal/ansible-concourse)
* [postgresql](https://github.com/ANXS/postgresql)

Another way to install concourse is to use the binary method shown [here](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/concourse-ci-cheat-sheet/install_concourse_binary_google_compute_engine.md).

## CLONE ROLES REPOS TO /ROLE

In /roles,

```bash
git clone git@github.com:ahelal/ansible-concourse.git
git clone git@github.com:ANXS/postgresql.git
```

In the future, may want to consider using git submodules.

## CREATE FOLLOWING INSTANCES (2X WEB/ 4X WORKER)

First get the followiong ubuntu instances running on Google Compute Engine:

* p-concourse-web-01
* p-concourse-web-02
* p-concourse-worker-01
* p-concourse-worker-02
* p-concourse-worker-03
* p-concourse-worker-04

GCE Labels can be used to identify the hosts in the next step.  Or just make an inventory file.

## LOCAL INVENTORY (HOSTS) FILE

Your inventory file should look like,

```bash
[concourse]
p-concourse-web-01.gce.px
p-concourse-web-02.gce.px
p-concourse-worker-01.gce.px
p-concourse-worker-02.gce.px
p-concourse-worker-03.gce.px
p-concourse-worker-04.gce.px

[concourse-web]
p-concourse-web-01.gce.px
p-concourse-web-02.gce.px

[concourse-worker]
p-concourse-worker-01.gce.px
p-concourse-worker-02.gce.px
p-concourse-worker-03.gce.px
p-concourse-worker-04.gce.px
```

Tell ansible where you host file is,

`nano ~/.ansible.cfg`,

```bash
[defaults]
inventory = /host/file/path
```

## SSH INTO TO GET KEYS IN ~/.ssh/known_hosts

ssh into the following to put keys in `~/.ssh/known_hosts`.

```bash
ssh p-concourse-web-01.gce.px
ssh p-concourse-web-02.gce.px
ssh p-concourse-worker-01.gce.px
ssh p-concourse-worker-02.gce.px
ssh p-concourse-worker-03.gce.px
ssh p-concourse-worker-04.gce.px
```

## GENERATE KEYS FOR 2 WEB AND 4 WORKERS

Generate keys for web and 4 workers using the keygen shell
script located in `/roles/ansible-concourse/keys`.

Run `/roles/ansible-concourse/keys/key.sh`.

Place keys in `concourse.yml` file.

For a future step, may want to use hidden ansible files.

## CREATE THE CONCOURSE PLAYBOOK

Create `concourse.yml`.  This playbook creates 2X Web / 4X Workers.

```yml
---
- name: Create postgreSQL database
  hosts: concourse-web
  become: yes
  vars:
    postgresql_version: 9.6
    # List of databases to be created (optional)
    # Note: for more flexibility with extensions use the postgresql_database_extensions setting.
    postgresql_databases:
      - name: concourse
        owner: postgres             # optional; specify the owner of the database
        hstore: yes                 # flag to install the hstore extension on this database (yes/no)
        uuid_ossp: yes              # flag to install the uuid-ossp extension on this database (yes/no)
        citext: yes                 # flag to install the citext extension on this database (yes/no)
        encoding: 'UTF-8'           # override global {{ postgresql_encoding }} variable per database
        lc_collate: 'en_US.UTF-8'   # override global {{ postgresql_locale }} variable per database
        lc_ctype: 'en_US.UTF-8'     # override global {{ postgresql_ctype }} variable per database
    # List of users to be created (optional)
    postgresql_users:
      - name: concourseci
        pass: xxxxxxxx
        encrypted: no               # denotes if the password is already encrypted.
  roles:
    - postgresql

- name: Create Web Nodes
  hosts: concourse-web
  become: True
  vars:
    concourseci_postgresql_user                 : "concourseci"
    concourseci_postgresql_pass                 : "xxxxxxxx"
    concourseci_postgresql_db                   : "concourse"
    CONCOURSE_WEB_BIND_IP                       : "0.0.0.0"
    CONCOURSE_WEB_TSA_BIND_IP                   : "0.0.0.0"
    concourseci_version                         : "v3.0.1"
    CONCOURSE_WEB_EXTERNAL_URL                  : "http://{{ inventory_hostname }}:8080"
    CONCOURSE_WEB_BASIC_AUTH_USERNAME           : "XXXXXXXXXX"
    CONCOURSE_WEB_BASIC_AUTH_PASSWORD           : "XXXXXXXXXX"
  roles:
    - ansible-concourse

- name: Create Worker Nodes
  hosts: concourse-worker
  become: True
  roles:
    - ansible-concourse
```

Note that ‘inventory_hostname’ is a magic variable that indicates
the current host you are looping over in the host loop.

Latest concourse version is [here](https://github.com/concourse/concourse/releases).

## RUN PLAYBOOK

```bash
ansible-playbook concourse.yml
```

## TESTING AND TROUBLESHOOTING - LOCALLY

Test the WEBs,

[http://p-concourse-web-01:8080/](http://p-concourse-web-01:8080/)

[http://p-concourse-web-02:8080/](http://p-concourse-web-02:8080/)


## TESTING AND TROUBLESHOOTING - POSTGRESQL - ON THE INSTANCE

```bash
ssh p-concourse-web-01
```

The database created by ansible-concourse role is called: `concourse`
username: `concourseci`
password as: `xxxxxxxx`

You can check,

```bash
sudo su - postgres
psql
\l
```

## TESTING AND TROUBLESHOOTING - CONCOURSE - ON THE INSTANCE

```bash
ssh p-concourse-web-01
```

Check if process running,

```bash
ps -aux | grep concourse
```

Check service,

```bash
systemctl | grep concourse
systemctl status concourse-web.service
```

## HOW ANSIBLE KICKS OFF CONCOURSE

There are two scripts (web or worker) that start and stop the service,

```bash
sudo /etc/init.d/concourse-web start
sudo /etc/init.d/concourse-worker start
```

which calls,

```bash
start-stop-daemon --start --background --name con-web --chdir /opt/concourseci/bin --chuid concourseci --group concourseci --exec /opt/concourseci/bin/concourse-web
```

that in turn calls,

```bash
exec /opt/concourseci/bin/concourse web >> /var/log/concourse/concourseci_web.log 2>&1
```

Check systemd log,

```bash
journalctl -xe
journalctl -f
```

## CONNECT LOCAL MACHINE TO CONCOURSE via FLY

To connect your local machine to concourse goto
[http://p-concourse-web-01:8080/](http://p-concourse-web-01:8080/)
and donwload the `fly` executable.

`fly` is your direct link to concourse and is used to
upload your ci pipelines.

Place `fly` in a /bin area of your choice and check the version,

```bash
fly -v
```

Now connect your local machine to concourse using team `ci`,

```bash
fly -t ci login -c http://p-concourse-web-01:8080/
```

username: `XXXXXXXXX`

password: `XXXXXXXXX`

You're all set.

## LOGIN TO CONCOURSE (HTTP)

[http://p-concourse-web-01:8080/](http://p-concourse-web-01:8080/)

username: `XXXXXXXXXX`

password: `XXXXXXXXXX`
