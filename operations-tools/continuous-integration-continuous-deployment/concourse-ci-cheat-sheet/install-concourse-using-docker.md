# DOCKER CHEAT SHEET

Concourse is distributed as a single concourse binary,
making it easy to run just about anywhere, especially with Docker.

These are two ways to install concourse using docker.

* Using `docker run IMAGENAME` (the traditional way)
* Using `docker-compose up`

We will use `docker-compose` since its easier and has a
configuration/provisioning file.

## INSTALL CONCOURSE USING DOCKER-COMPOSE

Get the `docker-compose.yml` file above.

```bash
wget https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet/docker-compose.yml
```

Gernerate the keys needed

```bash
wget https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet/generate-keys.sh
sudo sh generate-keys.sh
```

Note that this file will use two docker containers.

To run
```bash
$ docker-compose up
```

That's it, check that its working,

[http://127.0.0.1:8080](http://127.0.0.1:8080)

You can also check the images you have,

```bash
docker images
```

And see which containers are running,

```bash
docker ps
```

## IF YOU GET AN IMAGE ERROR

Just prune the containers,

```bash
docker container prune
```
