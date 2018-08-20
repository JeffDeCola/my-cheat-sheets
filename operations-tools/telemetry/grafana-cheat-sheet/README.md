# GRAFANA CHEAT SHEET

`grafana` _is an open platform for analytics and monitoring (metrics dashboard)_

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

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

## SLACK NOTIFICAITON

Send slack a grafana notifcation (from an alert).

At Slack setup an incomming webhook.

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

