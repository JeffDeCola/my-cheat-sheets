# AMAZON ELASTIC COMPUTE CLOUD (EC2) CHEAT SHEET

`amazon elastic compute cloud (ec2)` _which is part of
[aws](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/cloud-services/amazon-web-services-cheat-sheet)
provides IaaS._

Part of five main compute engines at `aws`,

* FaaS - AWS Lambda
  [(aws lambda)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/function-as-a-service/aws-lambda-cheat-sheet)
* PaaS - AWS Elastic Beanstalk
  [(aws elastic beanstalk)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/platform-as-a-service/aws-elastic-beanstalk-cheat-sheet)
* CaaS - Amazon Elastic Container Service
  [(amazon ecs)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-cheat-sheet)
  [(amazon eks)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/containers-as-a-service/amazon-elastic-container-service-for-kubernetes-cheat-sheet)
* IaaS - Amazon Elastic Compute Cloud
  [(amazon ec2)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/compute/amazon-elastic-compute-cloud-cheat-sheet)

Documentation and reference,
* Your
  [Images](??),
  [Instances Templates](??),
  [Instances Groups](??) and
  [Instances](??)
  ([VM boot disks](??)
  via the console.
* [Amazon EC2 Documentation](https://aws.amazon.com/ec2/)

My repo example using `amazon ec2` is
[hello-go-deploy-amazon-ec2](https://github.com/JeffDeCola/hello-go-deploy-amazon-ec2).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

In a nutshell, `ec2` allows you to deploy a VM instance from an `image`.
And your VM instance can contain Apps, services, containers, etc...

`ec2` offers scale, performance and value that allows
you to easily launch large compute clusters on Amazon's infrastructure.
There are no upfront investments.  Pay what you use.

## WHAT YOU NEED IN ORDER TO USE EC2

In order to use `ec2`, you will need,

* ACCOUNT - AWS Account (with billing).
* CREDENTIALS - User Access Keys ($HOME/.aws/credentials).
* API - `aws` cli or client libraries.
* YOUR WORLD AT AMAZON - Virtual Private Cloud (VPC) at amazon.
* CONNECT TO VM - ssh keys ($HOME/.ssh/jeffs-aws-ssh-keys.pem).

Refer to
[amazon web services (aws)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-architectures/infrastructure-as-a-service/cloud-services/amazon-web-services-cheat-sheet)
to get all this set up.

## FREE RESOURCE (t2.micro)

`t2.micro` has 1 CPU and 1 GB RAM using High
frequency Intel Xeon processors.  Normally cost about $9
per month.

Full list of [free aws services](https://aws.amazon.com/free/).

## EC2 REGIONS AND ZONES

Compute is located in regions across the globe.

* ASIA PACIFIC
  * ap-northeast-1 (Tokyo)
  * ap-northeast-2 (Seoul)
  * ap-northeast-3 (Osaka-Local)
  * ap-southeast-1 (Singapore)
  * ap-southeast-2 (Sydney)
  * ap-south-1 (Mumbai)
* EUROPE
  * eu-central-1 (Frankfurt)
  * eu-west-1 (Ireland)
  * eu-west-2 (London)
  * eu-west-3 (Paris)
  * eu-north-1 (Stockholm)
* NORTH AMERICA / CANADA
  * ca-central-1 (Central)
* SOUTH AMERICA
  * sa-east-1 (Sao Paulo)
* UNITED STATES
  * us-east-1 (Northern Virginia)
  * us-east-2 (Ohio)
  * us-west-1 (Northern California)
  * us-west-2 (Oregon)

Each region may have a few zones.

## GCE MACHINE TYPES, PRICING & REGIONS

At amazon there are a lot of [instance types](https://aws.amazon.com/ec2/instance-types/) and
ways to pay for those instances.

The following will be a very very high level look.

### INSTANCE TYPES

There are 5 basic instance types,

* GENERAL PURPOSE
* COMPUTE OPTIMIZED
* MEMORY OPTIMIZED
* ACCELERATED COMPUTING
* STORAGE OPTIMIZED

For the purpose of this cheat sheet we will focus on the general types,

The format is CPU and RAM.

* GENERAL PURPOSE
  * A1 - For ARM based processing (64-bit ARM cores)
    * a1.medium 1, 2
    * a1.large 2, 4
    * a1.4xlarge 16, 32
  * T3 - 2.5 GHz Intel Scalable Processor
    * t3.nano 1, .5
    * t3.medium 2, 4
    * t3.2xlarge 8,32
  * T2 - High frequency Intel Xeon processors
    * t2.nano 1, .5
    * t2.micro 1, 1 (FREE TIER)
    * t2.medium 2, 4
    * t2.2xlarge 8, 32
  * M5 -  3.1 GHz Intel Xeon® Platinum 8175 processors
    * m5.large 2, 8
    * m5.metal 96, 384
    * m5d.metal 96, 384
  * M5a - AMD EPYC 7000 series processors
    * m5a.large 2, 8
    * m5ad.large 2, 8
    * m5ad.24xlarge 96, 384
  * M4 - 2.3 GHz Intel Xeon® E5-2686 v4 (Broadwell) processors
    * m4.large 2, 8
    * m4.4xlarge 16, 64
    * m4.16xlarge 64, 256
  * T3a - 2.5 GHz AMD EPYC 7000 series processors
    * t3a.nano 2, 0.5
    * t3a.medium 2, 4
    * t3a.2xlarge 8, 32

A note that accelerated computing offers F1 instances which are 
customizable hardware acceleration with field programmable gate arrays
(FPGAs).

* f1.2xlarge 1 FPGA, 8 CPU, 122GB ($1.65/hour)
* f1.4xlarge 2 FPGA, 16 CPU, 255GB ($3.3/hour)
* f1.16xlarge 8 FPGA, 64 CPU, 976GB ($13.20/hour)

### PRICING FOR GENERAL PURPOSE TYPE

There are four ways to pay for Amazon EC2 instances

* `On-Demand Instances` - pay for compute capacity by per hour or
  per second depending on which instances you run.
* `Reserved Instances` - Reserved Instances provide
  you with a significant discount (up to 75%).
  compared to On-Demand instance pricing. But you need to
  purchase a 1-3 years.
* `Spot Instances` - Amazon EC2 Spot instances allow
  you to request spare Amazon EC2 computing.  They
  can be shut down at any moment.
  capacity for up to 90% off the On-Demand price.
* `Dedicated Hosts` - A physical EC2 server
  dedicated for your use.

To get an idea on pricing, here are some predefined standard machine
types. These prices also vary by region.

ON-DEMAND & SPOT INSTANCES

| MACHINE TYPE    | V. CPUs |   MEMORY |    ON-DEMAND |        SPOT  |
|:----------------|--------:|---------:|-------------:|-------------:|
| a1.medium       |       1 |      2GB |          ~20 |          N/A |
| t3.nano         |       2 |    0.5GB |          ~$4 |        ~$1.5 |
| t2.micro (FREE) |       1 |      1GB |        ~$8.7 |          ~$3 |
| t2.medium       |       2 |      4GB |         ~$35 |       ~$12.5 |
| f1.2xlarge      | 1FPGA 8 |    122GB |1.65/hr 1,240 |          N/A |

There are many other models depending on what you need. Check out the
latest machines and [pricing](https://aws.amazon.com/ec2/pricing/).

## INTERACTING WITH EC2

There are a few ways to interact with `ec2`,

* Using the gui/console.
* Using the aws cli (e.g. `aws`). See below.
* Using aws SDK Client Libraries / API
  (e.g. [go](https://docs.aws.amazon.com/sdk-for-go/api/)). 
  My [go example](https://github.com/JeffDeCola/my-go-examples/tree/master/cloud-services/amazon-aws).

This cheat sheet will focus on `aws`.

## ELASTIC COMPUTE CLOUD (EC2) MAIN SECTIONS

There are ??  main section of gce:

*
*
*
*

The goal is to deploy an `image` that creates a VM instance(s)
(That contains your App/service).

The following illustration shows a high level view on how an
App/service may be running on `ec2`.  AS you can see in this example,
the VM instances contain services.  It also shows
`????` control the show (They deploy and scale VM instances).

![IMAGE -  amazon elastic compute cloud app service view - IMAGE](../../../../../docs/pics/ec2-app-service-view.jpg)

### IMAGES

### ???

### INSTANCES

## INSTANCES - SSH - KEY PAIR

Create you keys from the Amazon EC2 console.  I called mine
`jeffs-aws-ssh-keys.pem`.

Place this file in ~/.ssh.= folder.

Use this anytime you connect to your instance.

## AWS BASIC COMMANDS









