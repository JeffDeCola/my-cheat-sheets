# GOOGLE CLOUD SPANNER CHEAT SHEET

```
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`google cloud spanner` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet)
is a horizontally scalable relational database
service from google._

Documentation and reference,

* [Cloud Spanner Documentation](https://cloud.google.com/spanner/docs/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

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

TBD

## READ FROM DATABASE

TBD