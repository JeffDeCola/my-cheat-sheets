# PUB SUB LOCAL EMULATOR

[Local pub/Sub Install](https://cloud.google.com/pubsub/docs/emulator)

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
