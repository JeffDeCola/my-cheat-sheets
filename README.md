# my-cheat-sheets

[![Code Climate](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/badges/gpa.svg)](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/my-cheat-sheets/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

`my-cheat-sheets` _is a place to keep my cheat sheets for various Apps,
Operating Systems and Languages._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## CHEAT SHEETS

* OPERATIONS TOOLS

  * CONFIGURATION MANAGEMENT

    * [ansible](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/ansible-cheat-sheet)

  * CONTINUOUS INTEGRATION

    * [concourse ci](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/concourse-ci-cheat-sheet)

  * DISCOVERY / CONFIGURATION

    * [consul](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/consul-cheat-sheet)

  * OCHESTRATION (DEPLOYMENT)

    * [kubernetes](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/kubernetes-cheat-sheet)

    * [packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/packer-cheat-sheet)

    * [terraform](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/terraform-cheat-sheet)

  * TELEMETRY

    * [grafana](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/grafana-cheat-sheet)

    * [stackdriver](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/stackdriver-cheat-sheet)

  * VIRTUAL MACHINES / CONTAINERS

    * [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/docker-cheat-sheet)

    * [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/vagrant-cheat-sheet)

* COMPUTE

  * [google compute engine](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/google-compute-engine-cheat-sheet)

* DATABASE

  * [google cloud spanner](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/google-cloud-spanner-cheat-sheet)

  * [postgreSQL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/postgreSQL-cheat-sheet)

* API

  * [RESTful](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/RESTful-cheat-sheet)

  * [youtube content id api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/youtube-content-id-api-cheat-sheet)

  * [youtube data api v3](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/youtube-data-api-v3-cheat-sheet)

  * [youtube player api](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/youtube-player-api-cheat-sheet)

* AUTHORIZATION

  * [OAuth 2.0 authorization](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/OAuth-2.0-authorization-cheat-sheet)

* MESSINGING

  * [NATS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/NATS-cheat-sheet)

* SOURCE / VERSION CONTROL

  * [git](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/git-cheat-sheet)

* OPERATING SYSTEMS

  * LINUX

    * [dns](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/Operating-Systems/Linux/dns-cheat-sheet)

  * WINDOWS

* LANGUAGES

  * [go](https://github.com/JeffDeCola/my-go-examples)

  * [php](https://github.com/JeffDeCola/my-php-containers)

  * [python](https://github.com/JeffDeCola/my-python-examples)

A lot of these cheat sheets have info I gathered from other sources.

## GITHUB WEBPAGE UPDATED USING CONCOURSE

A Concourse Pipeline will automatically update the GitHub WebPage.

![IMAGE - my-cheat-sheets concourse ci piepline - IMAGE](docs/pics/my-cheat-sheets-pipeline.jpg)

A _ci/.credentials.yml_ file needs to be created for your _slack_url_ and _repo_github_token_.

Use fly to upload the the pipeline file _ci/pipline.yml_ to Concourse:

```bash
fly -t ci set-pipeline -p my-cheat-sheets -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
```

## CONCOURSE RESOURCES IN PIPELINE

`my-cheat-sheets` also contains a few extra concourse resources:

* A resource (_resource-slack-alert_) uses a [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* A resource (_resource-repo-status_) use a [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

The above resources can be removed from the pipeline.
