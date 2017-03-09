# DOCKER CHEAT SHEET

`docker` _is a Virtual Linux Container._ Docker Hub is an online 
resource to place docker containers.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

[Jeff's Docker Hub Images](https://hub.docker.com/u/jeffdecola/)

## A VIRTUAL MACHINE VS DOCKER CONTAINER

The followng disagram shows the difference betwenn a Virtual Machine and a Docker Container.

![IMAGE - Virtual-Machine-vs-Docker - IMAGE](../docs/pics/Virtual-Machine-vs-Docker.jpg)

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

`https://docs.docker.com/engine/installation/linux/ubuntulinux`

Check version,

```bash
docker version
```

## IMAGES

An image is a ready to run.

List Images on Machine

```bash
docker images
```

Delete a docker image,

```bash
docker rmi IMAGE-ID
```

Delete all images,

```bash
docker rmi $(docker images -q)
```

## CONTAINERS

A Container is a running image.

List of Running Containers

```bash
docker ps
```

Run an image from Docker Hub

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

Run an interactive container,

This is cool, gives you a tty terminal,

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

```bash
mkdir hello-world-test
```

nano  Dockerfile

```bash
## Test
FROM ubuntu:14.04
MAINTAINER Jeff DeCola
CMD echo "Hi Jeff"
```

Another example of a Dockerfile

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

### Build ???

```bash
docker build -t jeffdecola/NAME .
```

### Run Image,

```bash
docker run -t -i jeffdecola/NAME
docker run -t -i jeffdecola/NAME /bin/bash
```



