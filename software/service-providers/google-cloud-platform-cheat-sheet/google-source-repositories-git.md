# GOOGLE SOURCE REPOSITORIES (GIT) CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Google source repositories (git), which is part of
[gcp](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet),
is a private git repositories hosted on Google Cloud Platform
that has a lot of limitations._

Table of Contents

* [OVERVIEW (NOT FOR DEVELOPMENT USE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-source-repositories-git.md#overview-not-for-development-use)
* [FREE RESOURCE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-source-repositories-git.md#free-resource)
* [MIRROR A REPO](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-source-repositories-git.md#mirror-a-repo)
* [CREATE/CLONE A REPO ON GOOGLE GIT](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-providers/google-cloud-platform-cheat-sheet/google-source-repositories-git.md#createclone-a-repo-on-google-git)

Documentation and Reference

* [your repos](https://source.cloud.google.com/repos)
* [google source repositories documentation](https://cloud.google.com/source-repositories/docs/)
* [quickstart](https://cloud.google.com/source-repositories/docs/quickstart)
* [google source repositories sdk reference (gcloud source)](https://cloud.google.com/sdk/gcloud/reference/source/)

## OVERVIEW (NOT FOR DEVELOPMENT USE)

You can do two things,

* Mirror a repo from github (no push)
* Create a repo

Google git is extremely limited. Its basically just to push
code to a repo. There are no pull requests or ability to
merge branches at google.  So you must do everything local.

It is also extremely slow.  It takes time (10 minutes)
for you to see any changes at google.

## FREE RESOURCE

As of my last update, the free resources are,

* Up to 5 users per month
* 50GB storage
* 50GB egress

Full list of [free gcp services](https://cloud.google.com/free/docs/gcp-free-tier).

## MIRROR A REPO

You can mirror a repo from github at google.

Use the console to connect to github and pick the repo you want to mirror.

For example, this repo will be named `github_jeffdecola_my-cheat-sheets`

Its just a copy of a repo from github and you can't push to it.

## CREATE/CLONE A REPO ON GOOGLE GIT

You can create a repo using the console or use the command line.

```bash
gcloud source repos create <your-repo-name>
```

Check your repo [here](https://source.cloud.google.com/repos)

Clone your new repo,

```bash
gcloud source repos clone <your-repo-name>
```

Now you can push/pull to this repo and that's about it.

To list your repos,

```bash
gcloud source repos list
```

As stated above, google git is extremely limited. Its basically just to push
code to a repo. There are no pull requests or ability to
merge branches at google.  So you must do everything local.
