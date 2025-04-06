# GRAFANA CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Grafana is an open platform for analytics and monitoring (metrics dashboard)._

* [INSTALL GRAFANA SERVER ON DEBIAN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#install-grafana-server-on-debian)
* [ADMIN LOGIN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#admin-login)
* [ENV VARIABLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#env-variables)
* [LOG FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#log-file)
* [DATABASE FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#database-file)
* [CONFIG](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#config)
* [SYSTEMD FOR GRAFANA](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#systemd-for-grafana)
* [GRAFANA-CLI](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#grafana-cli)
  * [GRAFANA HTTP API](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#grafana-http-api)
* [SLACK NOTIFICATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/telemetry/grafana-cheat-sheet#slack-notification)

## INSTALL GRAFANA SERVER ON DEBIAN

Out of the scope for this cheat-sheet.

Refer to [Debian Install](http://docs.grafana.org/installation/debian/).

## ADMIN LOGIN

Out of the box, `admin:admin`

## ENV VARIABLES

Environment vars are located in `/etc/default/grafana-server`.

## LOG FILE

The logs files are located,

```bash
/var/log/grafana
```

## DATABASE FILE

The data base is located,

`/var/lib/grafana/grafana.db`

## CONFIG

The configuration file is located at,

`/etc/grafana/grafana.ini`

## SYSTEMD FOR GRAFANA

The config file is,

/usr/lib/systemd/system/grafana-server.service

systemctl daemon-reload
systemctl start grafana-server
systemctl status grafana-server
sudo service grafana-server restart

Start at boot,

sudo systemctl enable grafana-server.service

## GRAFANA-CLI

`grafana-cli` command is bundled with grafana and allows you to do basic things.

```bash
grafana-cli help
```

list plugins,

```bash
grafana-cli plugins ls
```

### GRAFANA HTTP API

```bash
GET http://x.x.x.x.:3000/dashboard/db/jeff-hello
```

```bash
curl http://admin:admin@x.x.x.x.:3000/api/org
```

## SLACK NOTIFICATION

Send slack a grafana notification (from an alert).

At Slack setup an incoming webhook.

The end result will give you a URL like:

```bash
https://hooks.slack.com/services/xxx/xxx/xxx
```

Test your webhook,

```bash
curl -X POST --data-urlencode \
    'payload={"channel": "#grafana-alert", "username": "jeff", "text": \
    "This is posted to #grafana-alert and comes from a bot named Jeff.", \
    "icon_emoji": ":ghost:"}' \
    https://hooks.slack.com/services/xxx/xxx/xxx
```

To get the slack notification link to works,

Normally its,

```bash
http://localhost:3000/dashboard/db/xxx
```

But the ip is not filled in.

Edit the configuration .ini file

```bash
sudo nano /etc/grafana/grafana.ini
```

```bash
The http port  to use
;http_port = 3000

The public facing domain name used to access grafana from a browser
;domain = localhost
```

Change the port and domain appropriately.

Restart service,

sudo service grafana-server restart
