# GOOGLE CLOUD SPANNER CHEAT SHEET

`google cloud spanner` _is a horizontally scalable relational database
service from google._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## SCALES HORIZONATALLY (Adding machines)

Nornally relational databases scale vertically
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

TBD - Not finished

## READ FROM DATABASE

TBD - Not finished