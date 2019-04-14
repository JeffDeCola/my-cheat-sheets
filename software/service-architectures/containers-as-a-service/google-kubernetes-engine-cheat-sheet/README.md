# GOOGLE KUBERNETES ENGINE (GKE) CHEAT SHEET

`google kubernetes engine (gke)` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/cloud-services/google-cloud-platform-cheat-sheet)
_alows you to deploy, manage, and scale containerized applications on Kubernetes._

Part of four main compute engines at `gcp`,

* FaaS - Google Cloud Functions
  [(gcf)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/google-cloud-functions-cheat-sheet)
* PaaS - Google App Engine
  [(gae)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/google-app-engine-cheat-sheet)
* CaaS - Google Kubernetes Engine
  [(gke)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/google-kubernetes-engine-cheat-sheet)
* IaaS - Google Compute Engine
  [(gce)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/compute/google-compute-engine-cheat-sheet)

Documentation and reference,

* [My kubernetes cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
* [Google Kubernetes Engine Documentation](https://cloud.google.com/kubernetes-engine/docs/)
* [Quickstart](https://cloud.google.com/kubernetes-engine/docs/quickstart)
* [Google Cloud Container SDK Reference (gcloud container)](https://cloud.google.com/sdk/gcloud/reference/container/)
* [kubectl is Kubernetes cli](https://kubernetes.io/docs/reference/kubectl/overview/)

My repo example using `gke` is
[hello-go-deploy-gke](https://github.com/JeffDeCola/hello-go-deploy-gke).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

A node is simply a Virtual Machine running on the cluster.

A node pool is simply a “pool,” of machines with the same configuration.

All GKE clusters come with a default pool, and the default and
the minimum number of nodes in a pool is 3.

At gke, you are not charged for the master node.

## FREE RESOURCE

No cluster management fee for clusters of all sizes. Each user node is charged at
standard Compute Engine pricing.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## CREATE A 1-NODE KUBERNETES CLUSTER ON GCE

Kubernetes clusters nodes will be put on `gce` VM instance.

### STEP 1 - USE gcloud TO CREATE YOUR CLUSTER

I'm going to spark up an affordable 1 node pool.
I did this via the console.

* g1-small
* 1 node
* us-west1-a
* Boot 10GB
* preemptible

Should be around $5 per month.

In addition to setting up GKE we need to add a couple firewall rules to
allow the outside world to hit HTTP ports on our nodes.

```bash
gcloud container clusters create jeffs-3-node-cluster \
--disk-size=10
--disk-type=pd-standard
--
```

You should see 3 instances in `gce`.

### AUTHENTIFICATE 

To interact with the cluster, you need to authenticate. Run the following command:

```bash
gcloud container clusters get-credentials jeffs-3-node-cluster --zone us-west1-a --project jeffs-project-174816
```

This will make a `~/.kube` configuration folder.





