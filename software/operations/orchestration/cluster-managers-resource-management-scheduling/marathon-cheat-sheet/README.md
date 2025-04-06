# MARATHON CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Marathon is a framework that runs on top of Mesos
which consists of a scheduler and executor._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet#install)
* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet#overview)
* [MARATHON .JSON FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet#marathon-json-file)
* [DEPLOY DOCKER CONTAINER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/marathon-cheat-sheet#deploy-docker-container)

Documentation and Reference

* [marathon documentation](https://mesosphere.github.io/marathon/)
* [mesos cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet)
* [hello-go-deploy-marathon](https://github.com/JeffDeCola/hello-go-deploy-marathon)

## INSTALL

This is out of the scope of this cheat sheet.
I used a vagrant file.

## OVERVIEW

Deploy, run and scale docker containers easily
with native support.

Mesos provides resource management and scheduling
while a framework such as marathon provides
scheduling and execution.

There are many types of frameworks that a designed to do
different things such as data storage, or processing.
Marathon is design to keep an app running forever.

## MARATHON .JSON FILE

This is sent to marathon to get the docker container.

As an example,

```json
{
    "id": "hello-go-deploy-marathon-long-run",
    "cpus": 0.1,
    "mem": 4.0,
    "container": {
        "type": "DOCKER",
        "docker": {
            "forcePullImage": true,
            "image": "jeffdecola/hello-go-deploy-marathon:latest",
            "network": "BRIDGE",
            "portMappings": [{
                "containerPort": 8080,
                "hostPort": 0
            }]
        }
    }
}
```

## DEPLOY DOCKER CONTAINER

To deploy your app, send the json file to marathon,

```bash
curl -X PUT http://10.141.141.10:8080/v2/apps/jeffAPP \
    -d @app.json -H "Content-type: application/json"
```
