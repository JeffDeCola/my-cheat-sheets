# my-cheat-sheets

[![Code Climate](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/badges/gpa.svg)](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

`my-cheat-sheets` _is a place to keep my cheat sheets with
the number one gaol of creating services.

With this in mind, I organized my cheat sheets into three
main sections, 

* DEVELOPMENT
* INFRASTRUCTURE AS A SERVICE
* OPERATIONS TOOLS

[See my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## MY CHEAT SHEETS

_Sections in alphabetical order_

### DEVELOPMENT

* DEVELOPMENT ENVIRONMENT

  * [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environment/vagrant-cheat-sheet)
  * [visual studio code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/development-environment/visual-studio-code-cheat-sheet)

* LANGUAGES

  * [go](https://github.com/JeffDeCola/my-go-examples)
  * [php](https://github.com/JeffDeCola/my-php-containers)
  * [python](https://github.com/JeffDeCola/my-python-examples)

* OPERATING SYSTEMS

  * LINUX

    * [dns](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/operating-systems/linux/dns-cheat-sheet)

  * WINDOWS

    * [bash on ubuntu on windows](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/operating-systems/windows/bash-on-ubuntu-on-windows-cheat-sheet)

* SOFTWARE ARCHITECTURES

  * API

    * [RESTful](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/api/RESTful-cheat-sheet)
    * [youtube content id api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/api/youtube-content-id-api-cheat-sheet)
    * [youtube data api v3](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/api/youtube-data-api-v3-cheat-sheet)
    * [youtube player api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/api/youtube-player-api-cheat-sheet)

  * AUTHORIZATION

    * [OAuth 2.0 authorization](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/authorization/OAuth-2.0-authorization-cheat-sheet)

  * MESSINGING

    * [NATS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/messinging/NATS-cheat-sheet)

  * MICROSERVICES

    * [microservices](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/software-architectures/microservices/microservices-cheat-sheet)

* SOURCE / VERSION CONTROL

  * [git](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/source-version-control/git-cheat-sheet)

### INFRASTRUCTURE AS A SERVICE

* COMPUTE

  * [aws (amazon web services)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/infrastructure-as-a-service/compute/amazon-web-services-cheat-sheet)
  * [gce (google compute engine)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/infrastructure-as-a-service/compute/google-compute-engine-cheat-sheet)

* DATABASE

  * [google cloud spanner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/infrastructure-as-a-service/database/google-cloud-spanner-cheat-sheet)
  * [postgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/infrastructure-as-a-service/database/postgreSQL-cheat-sheet)

### OPERATIONS TOOLS

* CONFIGURATION MANAGEMENT

  * [ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/configuration-management/ansible-cheat-sheet)

* CONTINUOUS INTEGRATION / CONTINUOUS DEPLOYMENT

  * [concourse ci](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet)

* ORCHESTRATION

  * BUILDS / DEPLOYMENT

    * [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/builds-deployment/docker-cheat-sheet)
    * [packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/builds-deployment/packer-cheat-sheet)
    * [terraform](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/builds-deployment/terraform-cheat-sheet)

  * DISCOVERY / CONFIGURATION  

    * [consul](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/discovery-configuration/consul-cheat-sheet)

  * RESOURCE MANAGEMENT / SCHEDULING

    * [kubernetes](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/resource-management-scheduling/kubernetes-cheat-sheet)
    * [marathon](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/resource-management-scheduling/marathon-cheat-sheet-sheet)
    * [mesos](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/orchestration/resource-management-scheduling/mesos-cheat-sheet)

* TELEMETRY

  * [grafana](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/telemetry/grafana-cheat-sheet)
  * [stackdriver](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/telemetry/stackdriver-cheat-sheet)

_A lot of these cheat sheets has info I gathered from other sources._

![IMAGE - Creating Services Environment Overview - IMAGE](docs/pics/Creating-Services-Environment-Overview.jpg)

## GITHUB WEBPAGE UPDATED USING CONCOURSE

A concourse pipeline will automatically update
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

![IMAGE - my-cheat-sheets concourse ci piepline - IMAGE](docs/pics/my-cheat-sheets-pipeline.jpg)

A _ci/.credentials.yml_ file needs to be created for your _slack_url_ and _repo_github_token_.

Use fly to upload the the pipeline file _ci/pipline.yml_ to Concourse:

```bash
fly -t ci set-pipeline -p my-cheat-sheets -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
```

`my-cheat-sheets` also contains a few extra concourse resources:

* A resource (_resource-slack-alert_) uses a [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* A resource (_resource-repo-status_) use a [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.
