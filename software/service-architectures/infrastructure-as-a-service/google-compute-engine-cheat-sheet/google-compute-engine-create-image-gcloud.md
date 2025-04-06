# CREATE A CUSTOM IMAGE USING GCLOUD CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Create a custom image using gcloud is a way
to create a custom image on gce._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-gcloud.md#overview#overview)

## OVERVIEW

I personally do not like this way since it requires a
lot of steps that are not that script friendly.

Use
[create custom image using packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-packer.md)
instead.

But for completeness, as a high level overview,

* Start up an `instance` using an existing `image`.
  As a side note, no instance template or instance group are needed.
* ssh into your new instance and add whatever you want -
  installs, docker, etc...(this is a little ugly).
* When you're happy, stop the `instance` so that it can shut down
  and stop writing the data to the `persistent boot disk`.
* Create `image` from the `persistent boot disk`.

To perform the above steps, you can use the console/ssh or write scripts using `gcloud`.

Again, I understand I didn't get into details and didn't write
any scripts to do this.  Its because I find
[create custom image using packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-packer.md)
much more elegant.
