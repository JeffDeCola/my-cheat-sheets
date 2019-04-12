# DOCKER CHEAT SHEET

`docker` _is a virtual linux container. DockerHub is an online
resource to keep docker images (that you can pull)._

[Jeff's images at DockerHub](https://hub.docker.com/u/jeffdecola/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## VIRTUAL MACHINE (VM) vs DOCKER CONTAINER

The following diagram shows the difference between a Virtual Machine
and a Docker Container.

![IMAGE - virtual-machine-vs-docker-container - IMAGE](../../../../../docs/pics/virtual-machine-vs-docker-container.jpg)

Virtual Machine:

* Must use a Hypervisor emulated Virtual Hardware.
* Needs a guest OS.
* Takes a lot of system resources.
* Takes up a lot of memory.

Container:

* Uses a shared host OS.
* You must use that OS.
* Less Resources and lightweight.

## INSTALL

Goto [docs.docker.com/install](https://docs.docker.com/install)
to install.  Must be 64-bit machine.

When you're done, check version,

```bash
docker version
```

To mange docker as non-root, add yourself to the docker group,

```bash
sudo usermod -aG docker $USER
```

## DOCKER RUN (RUN A SINGLE CONTAINER)

To run a simple docker container just use `docker run`,

```bash
docker run jeffdecola/hello-go
docker run --name hello-go -dit jeffdecola/hello-go
```

* `-d` Run container in background.
* `-t` or `-tty` terminal. 
* `-i` interactive container, Keep STDIN open even if
  not attached.

## DOCKER-COMPOSE (RUN MULTIPLE CONTAINERS)

Compose is a tool for defining and running
multi-container Docker applications.

With compose, you use a .yml file to configure your application’s
services. Then, with a single command, you create and
start all the services from your configuration.

Here is an example of using docker-compose to
[start up concourse services using two docker containers](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet/install-concourse-using-docker-compose/README.md).

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

### PULL IMAGES

Get an image from [hub.docker.com](https://hub.docker.com/explore/),

```bash
docker pull concourse/concourse
```

Check you got it,

```bash
docker images
```

### CHECK THE HISTORY OF AN IMAGE

```bash
docker history jeffdecola/hello-go-deploy-gce
```

### CREATE IMAGE USING DOCKERFILE

Create a Dockerfile,

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
LABEL Jeff DeCola
COPY whatever /
CMD echo "Hi Jeff"
```

Build the image,

```bash
docker build -t jeffdecola/NAME .
```

`-t` is tag name.

Check your build,

```bash
docker images
```

### COMPILE YOUR CODE INSIDE CREATE IMAGE

When creating a docker image, your image must match the base docker image.

Hence, you should compile your code in the base image during build.

Here is an illustration why,

![IMAGE - compile-code-in-docker-base-image - IMAGE](../../../../../docs/pics/compile-code-in-docker-base-image.jpg)

For an example, refer to my repo
[here](https://github.com/JeffDeCola/hello-go-deploy-gce/tree/master/example-01/build-push).

## PUSH IMAGE TO DOKERHUB

To push an image to dockerhub,

```bash
docker push jeffdecola/hello-go
```

## CONTAINERS

A container is a running image.

List of Running Containers

```bash
docker ps
```

### START/STOP A CONTAINER

Run an image from Docker Hub,

```bash
docker run jeffdecola/hello-go
docker run --name hello-go jeffdecola/hello-go
docker run docker/whalesay cowsay boo
docker run ubuntu /bin/echo 'Hello World'
```

I like to use,

```bash
docker run --name hello-go -dit jeffdecola/hello-go
```

The docker command looks for it on your local system.
If the image isn’t there, docker gets it.

Stop a running container,

```bash
docker stop IMAGE-ID
```

Find the switches you can use with a container,

```bash
docker run IMAGENAME --help
```

List old containers you have lying around (cached),

```bash
docker ps -a
```

### DELETE A CONTAINER

Delete a container,

```bash
docker rm IMAGE-ID
```

Delete all containers,

```bash
docker rm $(docker ps -a -q)
```

### RUN INTERACTIVE CONTAINER

Run an interactive container, This is cool, it gives you a tty terminal,

* -t tty
* -i interactive

```bash
docker run -t -i ubuntu /bin/bash
docker run -t -i jeffdecola/hello-go
```

Run a container as a daemon,

```bash
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
```

Now see what the docker container daemon is doing,

```bash
docker logs NAME
```

### GET A SHELL PROMPT INSIDE A RUNNING CONTAINER

To get inside a running container,

```bash
docker exec -t -i <container ID> /bin/bash
```

Use `docker ps` to get the container ID.

### CHECK THE STDOUT OF A RUNNING CONTAINER

```bash
docker logs -f <container_name>
```

`-f` switch is check forever.
