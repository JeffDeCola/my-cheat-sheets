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
* [kubectl (Kubernetes cli)](https://kubernetes.io/docs/reference/kubectl/overview/)

My repo example using `gke` is
[hello-go-deploy-gke](https://github.com/JeffDeCola/hello-go-deploy-gke).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Each node is a VM instance on `gce`.

A node pool is a “pool,” of machines with the same configuration.

At `gke` you are not charged for the master node.

Refer to my
[My kubernetes cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
for more information about Kubernetes.

## FREE RESOURCE

No cluster management fee for clusters of all sizes. 

Each user node is charged the standard `gce` pricing.

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## STEP 1 - CREATE/DELETE KUBERNETES CLUSTER ON GKE

Since each node is a VM instance on `gce`, to keep costs down we'll
create one node.

Lets also use a preemptible disk to keeps the costs even lower.

As stated above `gke` will not charge you for the master node.

You can use the console or gcloud to create a Kubernetes Cluster.
I like to script, so I'm using gcloud,

```bash
gcloud container --project "$GCP_JEFFS_PROJECT_ID" \
    clusters create jeffs-gke-cluster-hello-go-deploy-gke \
    --addons HorizontalPodAutoscaling,HttpLoadBalancing,KubernetesDashboard \
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

To destroy your Kubernetes Cluster,

```bash
gcloud container clusters delete jeffs-gke-cluster
```

Reference for this gcloud SDK command [here](https://cloud.google.com/sdk/gcloud/reference/container/clusters/create)

## STEP 2 - CONNECT TO CLUSTER 

To interact with the cluster, you need to hook it up to kubectl.

This following command configures `kubectl` to use the
cluster you created.

```bash
gcloud container clusters get-credentials jeffs-gke-cluster-hello-go-deploy-gke \
    --zone us-west1-a \
    --project $GCP_JEFFS_PROJECT_ID
```

This will make a `~/.kube` configuration folder.

## STEP 3 - DEPLOY A DOCKER IMAGE TO GKE CLUSTER (CREATE WORKLOAD)

Lets use a docker image (that has port 8080 exposed) from Dockerhub.

I made a docker image `hello-go-deploy-gke` from my repo
[hello-go-deploy-gke](https://github.com/JeffDeCola/hello-go-deploy-gke),

```bash
kubectl run jeffs-web-counter \
    --image "jeffdecola/hello-go-deploy-gke:latest" \
    --port "8080"
```

The deployment will run in a pod in one of the nodes.

`gke` likes to call this a `workload`.  You can deploy as many `workloads` as you want.

## STEP 4 - EXPOSE CONTAINER TO THE WORLD (CREATE SERVICE)

Using a `gke` load balancer you can expose your container to the word,
Google calls this a `service`.

```bash
kubectl expose deployment jeffs-web-counter \
    --type LoadBalancer \
    --port 80 \
    --target-port 8080
```

Inspect your service,

```bash
kubectl get service jeffs-web-counter
```

Delete your service,

```bash
kubectl delete service jeffs-web-counter
```

This will not delete the `workload`, just the `service`.

## KUBERNETES DASHBOARD (THIS IS NICE)

If you noticed we used the addon KubernetesDashboard when we created our cluster.

To use, first get a secret token,

```bash
gcloud config config-helper --format=json | jq -r '.credential.access_token'
```

Then run a proxy,

```bash
kubectl proxy
```

And open in a browser,

[localhost:8001](http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy)
