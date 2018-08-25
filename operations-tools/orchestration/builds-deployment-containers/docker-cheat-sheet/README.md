# DOCKER CHEAT SHEET

`docker` _is a Virtual Linux Container._ Docker Hub is an online
resource to place docker containers._

[Jeff's Docker Hub Images](https://hub.docker.com/u/jeffdecola/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## A VIRTUAL MACHINE VS DOCKER CONTAINER

The following diagram shows the difference betwenn a Virtual Machine
and a Docker Container.

![IMAGE - Virtual-Machine-vs-Docker - IMAGE](../../../../docs/pics/Virtual-Machine-vs-Docker.jpg)

Virtual Machine:

* Must use a Hypervisor emulated Virtual Hardware
* Needs a guest OS
* Takes a lot of system resources

Container:

* Use a shared host OS
* You must use that OS
* Less Recources and lightweight

## INSTALL

Goto this website to install.  Must be 64-bit machine.

[docs.docker.com/install](https://docs.docker.com/install)

When you're done, check version,

```bash
docker version
```

## DOCKER RUN (RUN A SINGLE CONTAINER)

To run a simple docker comtainer just use `docker run`,

```bash
docker run jeffdecola/hello-go
```

## DOCKER-COMPOSE (RUN MULTIPLE CONTAINERS)

Compose is a tool for defining and running
multi-container Docker applications.

With compose, you use a .yml file to configure your application’s
services. Then, with a single command, you create and
start all the services from your configuration.

Here is an [example](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet/install-concourse-using-docker.md)
of using docker-compose to start up concourse services using two
docker containers.

## IMAGES

An image is ready to run.

List images,

```bash
docker images
```

Delete an image,

```bash
docker rmi IMAGE-ID
```

Delete all images (-f is force),

```bash
docker rmi $(docker images -q)
docker rmi -f $(docker images -q)
```

Get an image from [hub.docker.com](https://hub.docker.com/explore/),

```bash
docker pull concourse/concourse
```

Check you got it,

```bash
docker images
```

## CONTAINERS - RUNNING IMAGES

A container is a running image.

List of Running Containers

```bash
docker ps
```

Run an image from Docker Hub,

```bash
docker run jeffdecola/hello-go
```

```bash
docker run docker/whalesay cowsay boo
```

```bash
docker run ubuntu /bin/echo 'Hello World'
```

The docker command looks for it on your local system.
If the image isn’t there, docker gets it.

Stop a runing container,

```bash
docker stop IMAGE-ID
```

Find the switches you can use with a container,

```bash
docker run IMAGENAME --help
```

List old containers you have lying around (casched),

```bash
docker ps -a
```

Delete a container,

```bash
docker rm IMAGE-ID
```

Delete all containers,

```bash
docker rm $(docker ps -a -q)
```

Run an interactive container, This is cool, it gives you a tty terminal,

```bash
docker run -t -i ubuntu /bin/bash
docker run -t -i jeffdecola/hello-go
```

-t tty
-i interactive

Run a container as a daemon,

```bash
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
```

Now see what the docker container daemon is doing,

```bash
docker logs NAME
```

Stop a container,

```bash
docker stop NAME
```

## BUILD YOUR OWN IMAGE (METHOD 1) - Modify existing image

Get an existing docker image and add to it,

```bash
docker run -t -i training/sinatra /bin/bash
```

Now add something to it,

```bash
root@ IMAGE-ID:/# apt-get install -y ruby2.0-dev
```

Exit,

```bash
docker commit -m "Added ruby" -a "Jeff DeCola" IMAGE-ID
jeffdecola/sinatra:jeffver
```

## BUILD YOUR OWN IMAGE (METHOD 2) - Create Dockerfile

create a Dockerfile,

```bash
## Test
FROM ubuntu:14.04
MAINTAINER Jeff DeCola
CMD echo "Hi Jeff"
```

Another example,

```bash
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay
```

And another one,

```bash
#Test
FROM ubuntu:14.04
MAINTAINER Jeff DeCola
COPY whatever /
CMD echo "Hi Jeff"
```

Build the image,

```bash
docker build -t jeffdecola/NAME .
```

-t is tag name.

Check your build,

```bash
docker images
```
