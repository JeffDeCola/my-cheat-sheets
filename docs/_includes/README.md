
# CHEAT SHEETS

* [git](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/git-cheat-sheet)

* [docker](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/docker-cheat-sheet)

* [concourse ci](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/concourse-ci-cheat-sheet)

* [google compute engine](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/google-compute-engine-cheat-sheet)

* [packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/packer-cheat-sheet)

* [grafana](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/grafana-cheat-sheet)

* [terraform](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/terraform-cheat-sheet)

* [consul](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/consul-cheat-sheet)

* [vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/vagrant-cheat-sheet)

* [kubernetes](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/kubernetes-cheat-sheet)

* [stackdriver](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/stackdriver-cheat-sheet)

## GITHUB WEBPAGE UPDATED USING CONCOURSE

A Concourse Pipeline will automatically update the GitHub WebPage.

![IMAGE - my-cheat-sheets concourse ci piepline - IMAGE](pics/my-cheat-sheets-pipeline.jpg)

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
