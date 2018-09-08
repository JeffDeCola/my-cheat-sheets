# CONCOURSE CHEAT SHEET

`concourse` _is a pipelined CI (Continuous Integration) for
Software Builds. Concourse's main goal is to run tasks_.

[My Concourse task examples](https://jeffdecola.github.io/my-concourse-ci-tasks/).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL CONCOURSE SERVER

There are many ways to install concourse, here are a few,

* [On VirtualBox using vagrant](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-on-virtualbox-using-vagrant.md)
* [Binary Install](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-binary.md)
* [Docker Repository using docker-compose](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose)
* [Ansible Roles on Google Compute Engine](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-ansible-google-compute-engine.md)

## INSTALL AND CONNECT FLY TO CONCOURSE

Now we need a way to connect to your concourse server.
Fly allows you to uploads/updates your files.

* [Install Fly on Windows](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-fly-on-windows.md)
* [Install Fly on Linux](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-fly-on-linux.md)

## CREATE A TEAM

You are now logged in to concourse with team `main`.
You can check `~/.flyrc` to see this.

To creat a new team,

```bash
fly -t ci set-team --team-name team-jeff --local-user test
fly -t ci login -n team-jeff
```

check `~/.flyrc` to see your new team.

## BASIC STRUCTURE OF CONCOURSE

See a more detailed example at
[my-concourse-ci-tasks](https://jeffdecola.github.io/my-concourse-ci-tasks/).

The following diagram illustrates compase running a task called seach-and-replace.

* `pipeline.yml` A pipeline of resources and jobs.
* `config.yml` Configures task
  * Grabs docker image
  * Sets up inputs/outputs into task container
* `task.sh` does the task

`IMPORTANT - THE JOBS ARE COMPLETELY IDENPENDENT OF EACH OTHER`

If you want to store 'state', use a resource to send it offsite.

![IMAGE - concourse cheat sheet structure - IMAGE](../../../docs/pics/Concourse-structure.jpg)

## CREATING YOUR FIRST PIPELINE

There are plenty of better explanations on the web then I could do here.

I have a few examples in my repo
[my-concourse-ci-tasks](https://github.com/JeffDeCola/my-concourse-ci-tasks).

## LOAD/REMOVE PIPELINE TO/FROM CONCOURSE

Load your pipeline to concourse,

```bash
fly -t ci set-pipeline -p NAME -c pipeline.yml --load-vars-from .credentials.yml
```

To remove,

```bash
fly -t ci destroy-pipeline --pipeline NAME
```

## RESTART A BAD WORKER (STALLED)

Check your workers if they are stalled,

```bash
fly -t ci workers
```

If stalled, make sure you login using this full login with -c,

```bash
fly -t ci login -c http://NAME:8080/
```

And then prune-workers,

```bash
fly -t ci prune-worker -w WORKERNAME
```

## CREATE A NEW TEAM WITH NEW USER

Create a new team,

```bash
fly -t ci set-team --team-name ciusers \
    --basic-auth-username admin \
    --basic-auth-password admin
```

Like above, login to new team,

```bash
fly -t ci login -n ciusers -c http://<ip>:8080
```

Show all targets,

```bash
fly targets
```

Show all teams,

```bash
fly -t ci teams
```

## PASSING SECRETS TO USE IN YOUR SCRIPT (via env)

To pass secrets (files and variables) to your concourse script,
first use fly to upload the secret to concourse,

```bash
fly -t ci set-pipeline -p NAME \
    -c pipeline.yml --load-vars-from .credentials.yml \
    --var "private-key-file=$(cat private-key-file.txt | base64)" \
    --var "private-key=$(echo jeffdecola)"
```

Use base64 on files so you can unpack them properly on the other side.

Update the task in your pipeline with params,

```yml
- task: mybuild
  file: task.yml
  params:
    PRIVATE_KEY-FILE: {{private-key-file}}
    PRIVATE_KEY: {{private-key}}
```

Then in your `task.yml` file, also use params with the same name
(the value will be overwritten).

```yml
params:
  PRIVITE_KEY-FILE: "this will be overwritten"
  PRIVITE_KEY: "this will be overwritten"
```

This will create an env variable `PRIVATE_KEY-FILE` and `PRIVATE_KEY` you
may use in your concourse script.

```bash
echo $PRIVATE_KEY-FILE | base64 -d > private-key-file.txt
echo $PRIVATE_KEY
```
