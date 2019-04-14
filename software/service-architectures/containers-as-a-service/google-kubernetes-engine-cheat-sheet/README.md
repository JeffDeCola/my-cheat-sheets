# GOOGLE KUBERNETES ENGINE (GKE) CHEAT SHEET

`google kubernetes engine (gke)` _which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/cloud-services/google-cloud-platform-cheat-sheet)
allows you to deploy, manage and scale containerized applications on Kubernetes._

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

Each node is a VM instance on `gce`.

A node pool is a “pool,” of machines with the same configuration.

At `gke`, you are not charged for the master node.

Refer to my
[My kubernetes cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
for more information about Kubernetes.

## FREE RESOURCE

No cluster management fee for clusters of all sizes. 

Each user node is charged at standard `gce` pricing.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## STEP 1 - CREATE/DELETE KUBERNETES CLUSTER ON GKE

Since each node Each node is a VM instance on `gce`, to keep costs down we
would create one node.

I would use preemptible disk to keeps the costs even lower.

`gke` will not charge you for the master node.

You can use the console or gcloud to create a cluster.

I like to script, so I'm using gcloud,

```bash
gcloud container --project "$GCP_JEFFS_PROJECT_ID" \
    clusters create jeffs-gke-cluster-hello-go-deploy-gke \
    --addons HorizontalPodAutoscaling,HttpLoadBalancing \
    --disk-size "10" \
    --disk-type "pd-standard" \
    --enable-autorepair \
    --enable-autoupgrade \
    --enable-cloud-logging \
    --enable-cloud-monitoring \
    --image-type "COS" \
    --machine-type "f1-micro" \
    --no-enable-ip-alias \
    --num-nodes "3" \
    --preemptible \
    --username "admin" \
    --zone "us-west1-a" 
```

You should see 3 instance in `gce`.

To destroy,

```bash
gcloud container clusters delete jeffs-gke-cluster
```

## STEP 2 - CONNECT TO CLUSTER 

To interact with the cluster, you need to authenticate. Run the following command:

This following command configures `kubectl` to use the
cluster you created.

```bash
gcloud container clusters get-credentials jeffs-gke-cluster-hello-go-deploy-gke \
    --zone us-west1-a \
    --project $GCP_JEFFS_PROJECT_ID
```

This will make a `~/.kube` configuration folder.

## STEP 3 - DEPLOY A DOCKER IMAGE TO GKE CLUSTER

Lets deploy from a docker image (that has port 8080) at Dockerhub,

```bash
kubectl run jeffs-web-counter \
    --image "jeffdecola/hello-go-deploy-gke:latest" \
    --port "8080"
```

The deployment will run in a pod in one of the nodes.

## STEP 4 - EXPOSE TO THE WORLD

Using a `gke` load balancer,

```bash
kubectl expose deployment hello-server
    --type LoadBalancer \
    --port 80 \
    --target-port 8080
```

## STEP 5 - INSPECT AND DELETE

Inspect your service,

```bash
kubectl get service jeffs-web-counter
```

Delete your service,

```bash
kubectl delete service jeffs-web-counter
```
