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

`gce` requires a lot, I mean a lot of setup.
But what if you just don't really care about
the guts (infrastructure) and just want to deploy and App.  Well `gae`
is for you.  You bring the code, they handle the rest.

On a side note, can `gae` run a service?  I would say yes.  But that's
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
  * A higher level of abstraction, Focus is on App.
  * It auto scales for you
  * Google worries about infrastructure, you worry about code.
  * Simply deploy your code and platform does the rest
  * App engine will create more instances as needed
  * You don't manage/update/etc... the OS
  * Just upload code and gae does the rest
* [google kubernetes engine (gke)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-kubernetes-engine.md)
  * Containers
  * Immutable OS (Unable to be changed - Do not modify the OS)
  * Autoscaling
  * GCE Resources integrated
* [google compute engine (gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
  * IaaS
  * You have full control/responsibility for server
  * Custom machine types - You create and configure your own virtual machine instances
  * Direct access to OS
  * Manage OS and updates as needed

Good for,

*gae
  * Small services
  * Larger scale high performance service
*gke
  * Micro services
  * Container services 
  * Plan to cross cloud
*gae
 * Web services with large scaling
 * Quick scaling

Here is a high-level illustration,

![IMAGE -  gce-vs-gke-vs-gae - IMAGE](../../../../docs/pics/gce-vs-gke-vs-gae.jpg)

## TWO ENVIRONMENTS

* STANDARD ENVIRONMENT
  * Based on container instances running on Google's infrastructure.
  * Containers are preconfigured with one of several available runtimes.
* FLEXIBLE ENVIRONMENT
  * The App Engine flexible environment automatically scales your app up and down while balancing the load.
  * Allows you to customize the runtime and even the operating system of your virtual machine using Dockerfiles

## INSTALL

Must install the app-engine for gcloud,

```bash
gcloud components update
```

## STEP 1 - TEST (ON  LOCAL DEV SERVER)

Test on the local development server

```bash
dev_appserver.py app.yaml
```

```bash
http://localhost:8080/
```

## STEP 2 - DEPLOY

```bash
gcloud app deploy
```

## STEP 3 - VIEW

```bash
gcloud app browse
```
