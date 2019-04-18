# GOOGLE CLOUD STORAGE CHEAT SHEET

`google cloud storage` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
_is a scalable, fully managed, highly reliable, and cost-efficient
object/blob store. Used for pictures, videos, objects, blobs,
and other unstructured data._

Documentation and reference,

* [Your Buckets](https://console.cloud.google.com/storage/browser)
* [Google Cloud Storage Documentation](https://cloud.google.com/storage/docs/)
* [Quickstart using console](https://cloud.google.com/storage/docs/quickstart-console)
* [Quickstart using gsutil](https://cloud.google.com/storage/docs/quickstart-gsutil)
* [Google Cloud Storage gsutil tool documentation](https://cloud.google.com/storage/docs/gsutil)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Cloud Storage allows world-wide storage and retrieval
of any amount of data at any time.

Really just a place to keep your junk, I mean stuff.

Your stuff is placed in buckets.

`gsutil` is a python application that lets you access
Cloud Storage from the command line.

## FREE RESOURCE

As of my last update, the free resources are,

* 5 GB-months of Regional Storage (US regions only)
* 5000 Class A Operations per month
* 50000 Class B Operations per month
* 1 GB network egress from North America to all region destinations per month
  (excluding China and Australia)
* Only Available on following US regions:
  * Oregon: us-west1
  * Iowa: us-central1
  * South Carolina: us-east1

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## CREATE A BUCKET USING gsutil

Your stuff is placed in a bucket. You must choose your own,
globally-unique, bucket name.

```bash
gsutil mb -l us-west1 gs://my-awesome-bucket/
```

* `mb` is make bucket.
* `-l` is location.
* `gs` I assume means google storage.

list all your buckets,

```bash
gsutil ls
```

## UPLOAD/DOWNLOAD FILES/FOLDERS TO/FROM BUCKET

Upload file,

```bash
gsutil cp filename gs://my-awesome-bucket
```

Upload file to a folder,

```bash
gsutil cp filename gs://my-awesome-bucket/folder
```

Download a file,

```
gsutil cp gs://my-awesome-bucket/filename filename
```

List contents of bucket,

```bash
gsutil ls gs://my-awesome-bucket
gsutil ls -l gs://my-awesome-bucket
```

## OTHER COMMANDS

You can see these are just linux commands,

`mb`, `cp`, `ls`, `rm`, `mv`.

Full list [here](https://cloud.google.com/storage/docs/gsutil).
