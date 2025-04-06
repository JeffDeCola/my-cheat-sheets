# AMAZON ELASTIC CONTAINER SERVICE FOR KUBERNETES (EKS) CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Amazon elastic container service for kubernetes is an
[aws](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/amazon-web-services-cheat-sheet)
service that provides Containers as a Service (CaaS)._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#overview)
* [FREE RESOURCE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#free-resource)
* [STEP 1 - CREATE/DELETE KUBERNETES CLUSTER ON EKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#step-1---createdelete-kubernetes-cluster-on-eks)
* [STEP 2 - CONNECT TO CLUSTER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#step-2---connect-to-cluster)
* [STEP 3 - DEPLOYMENT (DEPLOY DOCKERHUB IMAGE TO CLUSTER)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#step-3---deployment-deploy-dockerhub-image-to-cluster)
* [STEP 4 - SERVICE (EXPOSE CONTAINER TO THE WORLD)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#step-4---service-expose-container-to-the-world)
* [KUBERNETES DASHBOARD](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet#kubernetes-dashboard)

Amazon Service Architectures

* FaaS - AWS Lambda
  [(aws lambda)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/aws-lambda-cheat-sheet)
* PaaS - AWS Elastic Beanstalk
  [(aws elastic beanstalk)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/aws-elastic-beanstalk-cheat-sheet)
* CaaS - Amazon Elastic Container Service
  [(amazon ecs)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-cheat-sheet)
  [(amazon eks)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet)
* IaaS - Amazon Elastic Compute Cloud
  [(amazon ec2)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/amazon-elastic-compute-cloud-cheat-sheet)

Documentation and Reference

* [my kubernetes cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
* [Amazon EKS Documentation](https://aws.amazon.com/eks/)
* my repo example using amazon eks is
[hello-go-deploy-amazon-eks](https://github.com/JeffDeCola/hello-go-deploy-amazon-eks)

## OVERVIEW

Amazon Elastic Container Service for Kubernetes makes it easy for
you to use Kubernetes on AWS without needing to install and operate
your own.

Refer to my
[my kubernetes cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/kubernetes-cheat-sheet)
for more information about Kubernetes.

## FREE RESOURCE

tbd

Full list of [free aws services](https://aws.amazon.com/free/).

## STEP 1 - CREATE/DELETE KUBERNETES CLUSTER ON EKS

Since each node is a VM instance on `ec2`, to keep costs down we'll
create a few nodes as needed.

Lets also use a preemptible disk????? to keeps the costs even lower.

As stated above `ec2` will not charge you for the master node.

You can use the console or ??? to create a Kubernetes Cluster.
I like to script, so I'm using ????,

```bash
?????????????????????
jeffs-eks-cluster-hello-go-deploy-eks
$MACHINE_TYPE
$NUM_NODES
$GCP_JEFFS_PROJECT_ID
```

You should see $NUM_NODES VM instances in `ec2`.

To destroy your Kubernetes Cluster,

```bash
???
```

Reference for this ?? SDK command ??

## STEP 2 - CONNECT TO CLUSTER

To interact with the cluster, you need to hook it up to kubectl.

This following command configures `kubectl` to use the
cluster you created.

```bash
???
```

This will make a `???` configuration folder.

## STEP 3 - DEPLOYMENT (DEPLOY DOCKERHUB IMAGE TO CLUSTER)

To deploy you can use a kubectl command (see below) or a yaml file.
I like using a yaml file.

Lets use a docker image (that has port 8080 exposed) from Dockerhub.
I'm using docker image `hello-go-deploy-eks` from my repo
[hello-go-deploy-amazon-eks](https://github.com/JeffDeCola/hello-go-deploy-amazon-eks),

```bash
kubectl create -f deploy.yaml
```

This is your yaml file,

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: jeffs-web-counter-deployment
spec:
  selector:
    matchLabels:
      app: jeffs-web-counter
  replicas: 2 # How many pods I want to create (Default 1 container per pod)
  template:
    metadata:
      labels:
        app: jeffs-web-counter
    spec:
      containers:
      - name: jeffs-web-counter
        image: jeffdecola/hello-go-deploy-eks:latest
        ports:
        - containerPort: 8080
```

The deployment will run in a pod in one of the nodes.

`eks` likes to call this a `workload`.

Inspect your deployment,

```bash
kubectl get deployments
kubectl get deployment jeffs-web-counter-deployment
```

Delete your deployment,

```bash
kubectl delete deployment jeffs-web-counter-deployment
```

## STEP 4 - SERVICE (EXPOSE CONTAINER TO THE WORLD)

Services are endpoints that export ports to the outside world.

Using a `gke` load balancer you can expose your container to the word,
Google calls this a `service`.

```bash
kubectl create -f service.yaml
```

This is your yaml file,

```yaml
apiVersion: v1
kind: Service
metadata:
  name: jeffs-web-counter-service
spec:
  selector:
    app: jeffs-web-counter
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
      nodePort: 31000
  type: LoadBalancer
```

Inspect your service,

```bash
kubectl get services
kubectl get service jeffs-web-counter
```

Delete your service,

```bash
kubectl delete service jeffs-web-counter
```

## KUBERNETES DASHBOARD

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
