# GOOGLE CLOUD PUB/SUB CHEAT SHEET

`google cloud pub/sub` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/cloud-services/google-cloud-platform-cheat-sheet)
is a fully-managed real-time messaging service that allows you to
send and receive messages between independent applications._

Documentation and reference,

* [Google Cloud Pub/Sub Documentation](https://cloud.google.com/pubsub/docs/)
* [Quickstart using console](https://cloud.google.com/pubsub/docs/quickstart-console)
* [Quickstart using gcloud pubsub](https://cloud.google.com/pubsub/docs/quickstart-cli)
* [Google Cloud Pub/Sub SDK Reference (gcloud pubsub)](https://cloud.google.com/sdk/gcloud/reference/pubsub/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Cloud Pub/Sub is a fully-managed real-time messaging service that allows you
to send and receive messages between independent applications.

## FREE RESOURCE

As of my last update, the free resources are,

* 10GB of messages per month.

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

$(gcloud beta emulators pubsub env-init)

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

