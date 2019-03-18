# CREATE IMAGE USING GCLOUD CHEAT SHEET

`create custom image using gcloud` _is one of two ways
to create a custom image on `gce`._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

I personally do not like this way since it requires a
lot of steps that are not script friendly.

Use 
[create custom image using packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine-create-image-packer.md)
instead.

But as a high level overview,

* Start up an `instance` using an existing `image`.
* ssh into that instance and add whatever you want (ugly).
* When you are happy, stop the `instance` so that it can shut down
  and stop writing any data to the `persistent disk`.
* Create `image` from instance's `persistent disk` or `snapshot`.

You can script this, but I don't like the process.