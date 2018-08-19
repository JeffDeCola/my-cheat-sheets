# GCE (GOOGLE COMPUTE ENGINE) CLI CHEAT SHEET

`(gce) google compute engine` _part of google Cloud Platform, 
like [aws](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/servers/compute/amazon-web-services-cheat-sheet),
provides high performance scalable VMs (Virtual Machines)._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

[Google Cloud Platform Reference](https://cloud.google.com/sdk/gcloud/reference/)

[Compute Engine Reference](https://cloud.google.com/sdk/gcloud/reference/compute/)

## BASIC GCP (GOOGLE CLOUD PLATFORM) COMMANDS

GCP Help,

```bash
gcloud help
```

Version information for GCP SDK Components,

```bash
gcloud config list
```

See your active configuration,

```bash
gcloud config list
```

## BASIC GCE COMMANDS RELATING TO YOUR VMs

List your images,

```bash
gcloud compute images list
```

List your instance templates,

```bash
gcloud compute instance-templates list
```

List your instance groups,

```bash
gcloud compute instance-groups list
gcloud compute instance-groups managed list
gcloud compute instance-groups unmanged list
```

List your instances/VMs,

```bash
gcloud compute instances list
```

## GCE HEALTH CHECK COMMANDS

List your health checks at GCE,

```bash
gcloud compute health-checks list
gcloud compute http-health-checks list
```

Describe the health check settings,

```bash
gcloud compute health-checks describe general-http-healthcheck-8080
gcloud compute http-health-checks describe p-jeff-test
```

Add your instance group to an existing GCE Health Check,

You can go into the browser and do it, or via CLI using,

```bash
gcloud beta compute instance-groups managed set-autohealing [IMAGE_NAME] \
    --http-health-check [HEALTHCHECK_NAME] \
    --initial-delay 120 \
   --zone [ZONE]
```

## SHOW EXTERNAL IPs FOR PROJECT

Listing static external IP addresses,


```bash
gcloud compute addresses list
```

gcloud compute addresses list

## OTHER BASIC GCE COMMANDS

GCE Help,

```bash
gcloud help compute
```

List of Machines Types per zone,

```bash
gcloud compute machine-types list
```

List zones,

```bash
gcloud compute zones list
```

## TO CREATE AN IMAGE, INSTANCE-TEMPLATE, INSTANCE GROUP OR INSTANCE/VM

Outside the scope of this cheat-sheet, but this will point you
in the right direction,

```bash
gcloud help compute images create
gcloud help compute instance-templates create
gcloud help compute instance-groups managed create
gcloud help compute instance-groups unmanaged create
gcloud help compute instances create
```

## GCLOUD BETA - PUBSUB

Beta versions of gcloud commands such as ones for pubsub.

List whats available

```bash
gcloud help beta
```

List pubsub topics,

```bash
gcloud beta pubsub topics list
```

List pubsub subscriptions,
This lists everything, so it can be a long list.

```bash
gcloud beta pubsub subscriptions list
```

## PUB/SUB LOCAL EMULATOR

[Local pub/Sub Install](https://cloud.google.com/pubsub/docs/emulator)

Must install Java JRE

```bash
sudo apt-get update
sudo apt-get install default-jre
```

STEP 1 - Start it,

```bash
gcloud beta emulators pubsub start \
    --data-dir="/home/jeff/.config/gcloud/emulators/pubsub"
```

STEP 2 - Call evn-init

Make your code call the API running in the local
instance instead of the production API, hence
run the env-init command in another terminal,

$(gcloud beta emulators pubsub env-init)

To see command line arguments,

```bash
gcloud beta emulators pubsub --help
```

## METADATA SERVER QUERY

Every instance stores its metadata on the metadata server.
You can query this metadata server programmatically for information such as 

* The instance's host name
* Instance ID
* Startup scripts, and
* Custom metadata

ssh onto your instance and perform the following

Relative to `http://metadata.google.internal/computeMetadata/v1/project/`

```bash
curl -s http://metadata.google.internal/computeMetadata/v1/project/project-id  -H "Metadata-Flavor: Google"
```

Relative to `http://metadata.google.internal/computeMetadata/v1/instance/`

```bash
curl -s http://metadata.google.internal/computeMetadata/v1/instance -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1/instance/hostname -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1/instance/machine-type -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1/instance/scheduling/preemptible  -H "Metadata-Flavor: Google"
```

wait for a change,

```bash
curl http://metadata.google.internal/computeMetadata/v1/instance/maintenance-event?wait_for_change=true -H 'Metadata-Flavor: Google'
```

## SSH INTO YOUR INSTANCE

You can ssh onto your instance,

Become root,

```bash
sudo su -
```
