# GOOGLE CLOUD PUB/SUB CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Google cloud pub/sub, which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet),
is a fully-managed real-time messaging service that allows you to
send and receive messages between independent applications._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#overview)
* [FREE RESOURCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#free-resource)
* [LOCAL PUB/SUB INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#local-pubsub-install)
* [STEP 1 - START](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#step-1---start)
* [STEP 2 - CALL evn-init](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#step-2---call-evn-init)
* [GCLOUD BETA - PUB/SUB](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-pub-sub.md#gcloud-beta---pubsub)

Documentation and Reference

* [google cloud pub/sub documentation](https://cloud.google.com/pubsub/docs/)
* [quickstart using console](https://cloud.google.com/pubsub/docs/quickstart-console)
* [quickstart using gcloud pubsub](https://cloud.google.com/pubsub/docs/quickstart-cli)
* [google cloud pub/sub sdk reference (gcloud pubsub)](https://cloud.google.com/sdk/gcloud/reference/pubsub/)

## OVERVIEW

Cloud Pub/Sub is a fully-managed real-time messaging service that allows you
to send and receive messages between independent applications.

## FREE RESOURCE

As of my last update, the free resources are,

* 10GB of messages per month

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## LOCAL PUB/SUB INSTALL

[Local Pub/Sub Install](https://cloud.google.com/pubsub/docs/emulator)

Must install Java JRE

```bash
sudo apt-get update
sudo apt-get install default-jre
```

## STEP 1 - START

```bash
gcloud beta emulators pubsub start \
    --data-dir="/home/jeff/.config/gcloud/emulators/pubsub"
```

## STEP 2 - CALL evn-init

Make your code call the API running in the local
instance instead of the production API, hence
run the env-init command in another terminal,

```bash
gcloud beta emulators pubsub env-init
```

To see command line arguments,

```bash
gcloud beta emulators pubsub --help
```

## GCLOUD BETA - PUB/SUB

Beta versions of gcloud commands such as ones for pubsub.

List whats available

```bash
gcloud help beta
```

List pub/sub topics,

```bash
gcloud beta pubsub topics list
```

List pub/sub subscriptions,
This lists everything, so it can be a long list.

```bash
gcloud beta pubsub subscriptions list
```
