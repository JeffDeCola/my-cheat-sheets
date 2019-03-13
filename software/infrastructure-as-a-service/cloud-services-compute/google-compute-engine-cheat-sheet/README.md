# GOOGLE COMPUTE ENGINE (GCE) CHEAT SHEET

`google compute engine (gce)` _part of google Cloud Platform, like
[amazon web services (aws)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/amazon-web-services-cheat-sheet),
or
[microsoft-azure](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/microsoft-azure-cheat-sheet),
provides high performance scalable VMs (Virtual Machines)._

gcp info,

* [Google Cloud Platform Documentation](https://cloud.google.com/docs)
* [Google Cloud Platform SDK Reference (gcloud)](https://cloud.google.com/sdk/gcloud/reference/)

gce info,

* [Google Compute Engine Documentation](https://cloud.google.com/compute/docs)
* [Google Compute Engine SDK Reference (gcloud compute)](https://cloud.google.com/sdk/gcloud/reference/compute/)


View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## GOOGLE CLOUD PLATFORM (GCP) OVERVIEW

Google Cloud Platform is absolutely massive with a host of
hundreds of powerful products from storage to compute to
data analytics.

Here are some highlights of gcp,

* AI and Machine Learning
* API Management
* Cloud Services Platform
* Compute (Focus of this cheat sheet)
* Data Analytics
* Databases
* Developers Tools
* Internet of Things (IoT)
* Management Tools
* Media
* Migration
* Networking
* Security
* Storage

This cheat sheet will focus on Google Compute Engine (gce),
scalable, high-performance VMs.

## GET ACCOUNT AT GOOGLE CLOUD PLATFORM

First setup a
[google cloud platform](https://cloud.google.com/)
account.

Create a Project which will have a name and ID.
Then submit your billing information.

* [Redeem a Promotion](https://console.cloud.google.com/billing/redeem)
* [Dashboard](https://console.cloud.google.com/home/dashboard)

## GOOGLE COMPUTE ENGINE (GCE) OVERVIEW

Here is an illustration of GCE.

![IMAGE -  google compute engine overview - IMAGE](../../../../../docs/pics/google-computer-engine-overview.jpg)

There are 4 main section of gce:

* Images
* Instance Templates
* Instance Groups
* Instances

## GCE REGIONS AND ZONES

Compute is located in regions across the globe.

* ASIA
  * asia-east1 (Taiwan)
  * asia-east2 (Hong Kong)
  * asia-northeast1 (Tokyo)
  * asia-south1 (Mumbai)
  * asia-southeast1 (Singapore)
* AUSTRALIA
  * australia-southeast1 (Sydney)
* EUROPE
  * europe-north1 (Finland)
  * europe-west1 (Belgium)
  * europe-west2 (London)
  * europe-west3 (Frankfurt)
  * europe-west4 (Netherlands)
  * europe-west6 (Zurich)
* NORTH AMERICA
  * northamerica-northeast1 (Montreal)
* SOUTH AMERICA
  * southamerica-east1 (Sao Paulo)
* UNITED STATES
  * us-central1 (Iowa)
  * us-east1 (South Carolina)
  * us-east4 (Northern Virginia)
  * us-west1 (Oregon)
  * us-west2 (Los Angeles)

Then each region may have a few zones.

## GCE MACHINE TYPES, PRICING & REGIONS

To get an idea on pricing, here are some predefined standard machine 
types, These prices also vary by region.

| MACHINE TYPE    | V. CPUs |   MEMORY |  PRICE/MONTH |  PREEMPTIBLE |
|:----------------|--------:|---------:|-------------:|-------------:|  
| n1-standard-1   |       1 |   3.75GB |         ~$25 |          ~$8 |
| n1-standard-2   |       2 |    7.5GB |         ~$58 |         ~$17 |
| n1-standard-8   |       8 |     30GB |        ~$233 |         ~$70 |
| n1-standard-64  |      64 |    240GB |      ~$1,867 |        ~$561 |
|                 |         |          |              |              |
| n1-highmem-2    |       2 |     13GB |         ~$72 |         ~$21 |
| n1-highmem-4    |       4 |     26GB |        ~$145 |         ~$43 |
| n1-highmem-8    |       8 |     52GB |        ~$290 |         ~$87 |
| n1-highmem-64   |      64 |    416GB |      ~$2,323 |        ~$699 |
|                 |         |          |              |              |
| n1-highcpu-2    |       2 |    1.8GB |         ~$43 |         ~$13 |
| n1-highcpu-4    |       4 |    3.6GB |         ~$86 |         ~$26 |
| n1-highcpu-8    |       8 |    7.2GB |        ~$173 |         ~$52 |
| n1-highcpu-64   |      64 |     57GB |      ~$1,391 |        ~$418 |
|                 |         |          |              |              |
| f1-micro        |       1 |    0.6GB |       ~$4.65 |          ~$3 |
| g1-small        |       1 |    1.7GB |      ~$15.74 |          ~$6 |
|                 |         |          |              |              |

`f1-micro` and `g1-small` machine types offer bursting capabilities that
allow instances to use additional physical CPU for short periods of time.

There are many other models depending on what you need.

* [pricing](https://cloud.google.com/compute/pricing)

## GCE FREE TIER (f1-micro)

There is a free `f1-micro` tier as follows,

* 1 non-preemptible `f1-micro` VM instance per month in one of the following US regions:
  * Oregon: us-west1
  * Iowa: us-central1
  * South Carolina: us-east1
* 30 GB of `Standard persistent disk storage` per month.
* 5 GB of `snapshot storage per month`. Limited to the following regions:
  * Oregon: us-west1
  * Iowa: us-central1
  * South Carolina: us-east1
  * Taiwan: asia-east1
  * Belgium: europe-west1
* Network Traffic Limits:
  * You must use Premium Tier. You cannot use Standard Tier.
  * Traffic must be sent from a GCP region in North America.
  * You can send up to 1GB of egress traffic, in aggregate,
    to regions except for those in Oceania and China.

Preemptible VM instances are not included in the Google Cloud Platform Free Tier.

## INSTALL GCLOUD SOFTWARE DEVELOPMENT KIT (SDK)

`gcloud` SDK is a set of tools for Cloud Platform. It contains `gcloud`,
`gsutil`, and `bq` command-line tools, which you can use to access Google Compute Engine,
Google Cloud Storage, Google BigQuery, and other products
and services from the command-line.

There is great documentation online to install the SDK on your platform
[here](https://cloud.google.com/sdk/)

You will be asked to attached to a project.

When complete, check your active configuration,

```bash
gcloud config list
```

## BASIC GCP COMMANDS

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

* [Compute Engine SDK Reference (gcloud compute)](https://cloud.google.com/sdk/gcloud/reference/compute/)

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

### STEP 1 - Start it

```bash
gcloud beta emulators pubsub start \
    --data-dir="/home/jeff/.config/gcloud/emulators/pubsub"
```

### STEP 2 - Call evn-init

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
You can query this metadata server programmatically for information such as,

* The instance's host name
* Instance ID
* Startup scripts, and
* Custom metadata

ssh onto your instance and perform the following

Relative to `http://metadata.google.internal/computeMetadata/v1/project/`

```bash
curl -s http://metadata.google.internal/computeMetadata/v1/\
project/project-id \
-H "Metadata-Flavor: Google"
```

Relative to `http://metadata.google.internal/computeMetadata/v1/instance/`

```bash
curl -s http://metadata.google.internal/computeMetadata/v1\
/instance \
    -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1\
/instance/hostname \
    -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1\
/instance/machine-type \
    -H "Metadata-Flavor: Google"
curl -s http://metadata.google.internal/computeMetadata/v1\
/instance/scheduling/preemptible \
    -H "Metadata-Flavor: Google"
```

Wait for a change,

```bash
curl http://metadata.google.internal/computeMetadata/v1\
/instance/maintenance-event?wait_for_change=true \
    -H 'Metadata-Flavor: Google'
```

## SSH INTO YOUR INSTANCE

You can ssh onto your instance.

Become root,

```bash
sudo su -
```
