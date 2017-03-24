# GCE (GOOGLE COMPUTE ENGINE) CHEAT SHEET

`google compute engine` _part of google Cloud Platform, provides
high performance scalable VMs (Virtual Machines)._

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
