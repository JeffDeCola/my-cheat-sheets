# MARATHON CHEAT SHEET

`marathon` _is a framework that runs on top of Mesos
which consists of a scheduler and executor._

Documentation and reference,

* [Marathon Documentation](https://mesosphere.github.io/marathon/)
* My
  [mesos cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/orchestration/cluster-managers-resource-management-scheduling/mesos-cheat-sheet)

My Repo example is
[hello-go-deploy-marathon](https://github.com/JeffDeCola/hello-go-deploy-marathon).

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

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
curl -X PUT http://10.141.141.10:8080/v2/apps/jeffAPP -d @app.json -H "Content-type: application/json"
```


