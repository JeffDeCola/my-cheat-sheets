# GOOGLE APP ENGINE (GAE) CHEAT SHEET

`google app engine (gae)` _provides serverless application platform
for apps and backends._

Documentation and reference,

* [Google App Engine Documentation](https://cloud.google.com/appengine/docs)
* [Quickstart standard environment using go](https://cloud.google.com/appengine/docs/standard/go/)
* [Quickstart flexible environment using go](https://cloud.google.com/appengine/docs/flexible/go/)
* [Google App Engine SDK Reference (gcloud compute)](https://cloud.google.com/sdk/gcloud/reference/app/)

My Repo example is [hello-go-deploy-gae](https://github.com/JeffDeCola/hello-go-deploy-gae).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

## FREE RESOURCE (standard environment only)

As of my last update, the free resources are,

* 28 frontend instance hours per day
* 9 backend instance hours per day
* 5 GB Cloud Storage
* 1 GB of egress per day
* Shared memcache
* 1000 search operations per day, 10 MB search indexing
* 100 emails per day

The free tier is available only for the `Standard Environment`.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## DIFFERENCE BETWEEN GCE

Some main differences between `google compute engine`,

* `google app engine (gae)`
  * PaaS
  * A higher level of abstraction
  * Simply deploy your code and platform does the rest
  * App engine will create more instances as needed
  * You don't manage/update/etc... the OS
  * Just upload code and gae does the rest
* [google compute engine (gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
  * IaaS
  * You have full control/responsibility for server
  * You create and configure your own virtual machine instances
  * Direct access to OS
  * Manage OS and updates as needed



