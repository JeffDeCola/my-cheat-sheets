# GOOGLE KUBERNETES ENGINE CHEAT SHEET

`google kubernetes engine (gke)` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet)
_alows you to deploy, manage, and scale containerized applications on Kubernetes._

Part of three compute engines at `gcp`,

* App engine
  [(gae)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/platform-as-a-service/cloud-services-app/google-cloud-platform-cheat-sheet/google-app-engine.md)
  PaaS
* Container/Kubernetes engine (gke)
  IaaS/PaaS or CaaS
* Compute engine
  [(gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
  IaaS

Documentation and reference,

* [Google Kubernetes Engine Documentation](https://cloud.google.com/kubernetes-engine/docs/)
* [Quickstart](https://cloud.google.com/kubernetes-engine/docs/quickstart)
* [Google Cloud Container SDK Reference (gcloud container)](https://cloud.google.com/sdk/gcloud/reference/container/)

My Repo example using `gke` is [hello-go-deploy-gke](https://github.com/JeffDeCola/hello-go-deploy-gke).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

This really sits in the middle of IaaS and PaaS. Some people like to say
Container as a Service (CaaS).  I find this a little unnecessary, I think
we have enough acronyms in this space.

Really great for runing docker containeers in a manage Kubernetes environment.

Kubernetes performs the 
* automation, 
* orchestration, 
* management and 
* deployment 

of your containers. More information here.
## FREE RESOURCE

No cluster management fee for clusters of all sizes. Each user node is charged at
standard Compute Engine pricing.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## GCE, GKE & GAE (THE ENGINES ON GCP)

What are the main differences between `google app engine`, 
`google kubernetes engine` and `google compute engine`?

* [google app engine (gae)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/platform-as-a-service/cloud-services-app/google-cloud-platform-cheat-sheet/google-app-engine.md)
  PaaS
  * A higher level of abstraction. Serverless. Focus is on your code.
  * Auto scales for you. Will create more instances as needed.
  * Google worries about infrastructure, you worry about code.
    Simply deploy your code and platform does the rest.
  * You don't manage or update the OS.
* google kubernetes engine (gke) IaaS/PaaS or CaaS
  * A step up from `gce` that uses Containers to manage your App.
  * Immutable OS (Unable to be changed - Can't modify the OS).
  * Autoscaling.
  * GCE Resources integrated. Kubernetes runs on `gce`.
* [google compute engine (gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/infrastructure-as-a-service/cloud-services-compute/google-cloud-platform-cheat-sheet/google-compute-engine.md)
  IaaS.
  * You have full control/responsibility for server.
  * Create your own VM instance by allocating hardware specific resources
    (e.g. RAM, CPU, Storage).
  * Direct access to OS.
  * Manage OS and updates as needed.

So what is this all good for,

* `gae`
  * Web services with large scaling.
  * Quick scaling.
* `gke`
  * Micro services.
  * Container services.
  * Plan to cross cloud.
* `gce`
  * Small services.
  * Larger scale high performance service.

Here is a high-level illustration,

![IMAGE -  gce-vs-gke-vs-gae - IMAGE](../../../../docs/pics/gce-vs-gke-vs-gae.jpg)


