# GOOGLE STACKDRIVER MONITORING CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Google stackdriver monitoring, which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet),
provides visibility into the performance, uptime, and overall
health of cloud-powered applications._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-stackdriver-monitoring.md#overview)
* [FREE RESOURCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-stackdriver-monitoring.md#free-resource)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-stackdriver-monitoring.md#install)

Documentation and Reference

* [google stackdriver monitoring documentation](https://cloud.google.com/stackdriver/docs/)
* [quickstart](https://cloud.google.com/monitoring/docs/quickstart)

## OVERVIEW

Stackdriver Monitoring provides visibility into the performance,
uptime, and overall health of cloud-powered applications.

## FREE RESOURCE

I believe this is always free.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## INSTALL

To install the Stackdriver monitoring agent,

```bash
curl -sSO https://dl.google.com/cloudagents/install-monitoring-agent.sh
sudo bash install-monitoring-agent.sh
```

To install the Stackdriver logging agent:

```bash
curl -sSO https://dl.google.com/cloudagents/install-logging-agent.sh
sudo bash install-logging-agent.sh
```
