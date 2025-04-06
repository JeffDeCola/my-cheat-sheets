# GOOGLE CLOUD SPANNER CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Google cloud spanner, which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet),
is a horizontally scalable relational database
service from google._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#overview)
* [FREE RESOURCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#free-resource)
* [SCALES HORIZONTALLY (Adding machines)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#scales-horizontally-adding-machines)
* [GOOGLE CLOUD PLATFORM](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#google-cloud-platform)
* [CREATE AN INSTANCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#create-an-instance)
* [CREATE A DATABASE IN THAT INSTANCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#create-a-database-in-that-instance)
* [WRITE TO DATABASE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#write-to-database)
* [READ FROM DATABASE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-cloud-spanner.md#read-from-database)

Documentation and Reference

* [cloud spanner documentation](https://cloud.google.com/spanner/docs/)

## OVERVIEW

Cloud Spanner is a fully managed, mission-critical, relational database service
that offers transactional consistency at global scale, schemas, SQL
(ANSI 2011 with extensions), and automatic, synchronous replication
for high availability.

## FREE RESOURCE

Not sure of any.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## SCALES HORIZONTALLY (Adding machines)

Normally relational databases scale vertically
(i.e. adding more power like CPU and RAM)
But Cloud Spanner scales horizontal from one to
thousands of nodes (i.e. adding more machines).

## GOOGLE CLOUD PLATFORM

You must have an account for google cloud platform.  This is not the scope
of this cheat sheet.

Note, google cloud spanner is obviously billable.

## CREATE AN INSTANCE

First create an instance, which is an allocation of resources (i.e. A machine).

You could use the ui or gcloud cli,

```bash
gcloud spanner instances create test-instance --config=regional-us-central1 \
--description="Test Instance" --nodes=1
```

Nodes allow you to horizontally scale the instance.

## CREATE A DATABASE IN THAT INSTANCE

As above, you could use the ui or gcloud cli,

```bash
go run snippet.go createdatabase projects/$GCLOUD_PROJECT/instances/test-instance/databases/example-db
```

Where snippet.go is [here](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/spanner/spanner_snippets/snippet.go)

## WRITE TO DATABASE

tbd

## READ FROM DATABASE

tbd
