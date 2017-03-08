# DOCKER CHEAT SHEET

`docker` _is a Virtual Linux Container._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## VIRTUAL MACHINE VERSUS DOCKER CONTAINER

![IMAGE - Virtual-Machine-vs-Docker - IMAGE](docs/pics/Virtual-Machine-vs-Docker.jpg)

Wit a Virtual Machine you must use a Hypervisor emulated Virtual Hardware.
A Virtual Machine needs a guest OS and takes a lot of system resources.

Containers (Like Docker Containers) use a shared host OS.  But you must use that
Operating System.

## INSTALL

Goto this website to install.  Must be 64-bit machine.

`https://docs.docker.com/engine/installation/linux/ubuntulinux`

Check version,

```bash
docker version
```

## IMAGES

### List Images on Machine

```bash
docker images
```

DELETE IMAGES

```bash
docker rmi IMAGE-ID
```

DELETE ALL IMAGES

```bash
docker rmi $(docker images -q)	
```

CONTAINERS

RUN  A CONTAINER FROM DOCKERHUB

```bash
docker run jeffdecola/hello-go
docker run -t -i jeffdecola/hello-go
```

```bash
docker run docker/whalesay cowsay boo
```

```bash
docker run ubuntu /bin/echo 'Hello World'
```

The docker command looks for it on your local system. If the image isn’t there,
then docker gets it from the hub.

LIST RUNNING CONTAINERS

```bash
docker ps
```

STOP A RUNNING CONTIANER

```bash
docker stop IMAGE-ID
```

LIST OLD CONTAINERS YOU HAVE LYING AROUND (CACHED)

```bash
docker ps -a
```

DELETE A CONTAINER

```bash
docker rm IMAGE-ID
```

DELETE ALL CONTAINERS

```bash
docker rm $(docker ps -a -q)
```

RUN INTERACTIVE CONTAINER

This is cool, gives you a tty terminal,

```bash
docker run -t -i ubuntu /bin/bash
docker run ubuntu:latest echo "Hello, world"
```

RUN A CONTAINER AS A DEAEMON,

```bash
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
```

Now see what the docker container daemon is doing,

```bash
docker logs NAME 
```

Stop a container

```bash
docker stop NAME 
```

BUILD YOUR OWN IMAGE (METHOD 1) - Modify existing image

Get an existing Docker Image and add to it.

```bash
docker run -t -i training/sinatra /bin/bash
```

Now add something to it

```bash
root@ IMAGE-ID:/# apt-get install -y ruby2.0-dev
```

exit

```bash
docker commit -m "Added ruby" -a "Jeff DeCola" IMAGE-ID
jeffdecola/sinatra:jeffver
```

BUILD YOUR OWN IMAGE (METHOD 2) - Use Dockerfile

Dockerfile

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

ANOTHER EXAMPLE of Dockerfile

```bash
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay
```

ANOTHER EXAMPLE of Dockerfile - Copy files into docker container

```bash
#Test
FROM ubuntu:14.04
MAINTAINER Jeff DeCola
COPY whatever /
CMD echo "Hi Jeff"
```

BUILD

```bash
docker build -t jeffdecola/NAME .
```

RUN

```bash
docker run -t -i jeffdecola/NAME
docker run -t -i jeffdecola/NAME /bin/bash
```



