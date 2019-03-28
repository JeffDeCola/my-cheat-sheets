# GOOGLE APP ENGINE (GAE) CHEAT SHEET

`google app engine (gae)` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet)
provides a serverless application platform that you can use to run an App._

Part of three compute engines at `gcp`,

* Compute engine
  [(gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
* Container engine
  [(gke)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-kubernetes-engine.md)
* App engine (gae)

Documentation and reference,

* [Google App Engine Documentation](https://cloud.google.com/appengine/docs)
* [Quickstart standard environment using go](https://cloud.google.com/appengine/docs/standard/go/)
* [Quickstart flexible environment using go](https://cloud.google.com/appengine/docs/flexible/go/)
* [Google App Engine SDK Reference (gcloud compute)](https://cloud.google.com/sdk/gcloud/reference/app/)

My Repo example is [hello-go-deploy-gae](https://github.com/JeffDeCola/hello-go-deploy-gae).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

GCE requires a lot of setup.  But what if you just don't really care about
the guts (infrastructure) and just want to deploy and App.  Well `gae`
is for you.

On a side note, can `gae` run a service.  I would say yes.  But that's
not really the point of `gae`.

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

## GCE, GKE & GAE

Main differences between `google compute engine`, `google kubernetes engine`
and `google app engine` are,

* [google app engine (gae)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-app-engine.md)
  * PaaS
  * A higher level of abstraction
  * Simply deploy your code and platform does the rest
  * App engine will create more instances as needed
  * You don't manage/update/etc... the OS
  * Just upload code and gae does the rest
* [google compute engine (gke)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-kubernetes-engine.md)
  * Containers
  * Immutable OS (Unable to be changed - Do not modify the OS)
* [google compute engine (gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
  * IaaS
  * You have full control/responsibility for server
  * You create and configure your own virtual machine instances
  * Direct access to OS
  * Manage OS and updates as needed

Here is a high-level illustration,

![IMAGE -  gce-vs-gke-vs-gae - IMAGE](../../../../docs/pics/gce-vs-gke-vs-gae.jpg)
