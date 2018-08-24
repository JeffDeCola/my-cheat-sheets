# INSTALL CONCOURSE USING DOCKER-COMPOSE CHEAT SHEET

Concourse is distributed as a single concourse binary,
making it easy to run just about anywhere, especially with Docker.

These are two ways to install concourse using docker.

* Using `docker run IMAGENAME` (the traditional way)
* Using `docker-compose up`

We will use the later, `docker-compose` since its easier
and has a configuration/provisioning file docker-compose.yml.

## INSTALL CONCOURSE USING DOCKER-COMPOSE

Get the `docker-compose.yml` file above.

```bash
wget https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet/docker-compose.yml
```

Generate the keys needed,

```bash
wget https://github.com/JeffDeCola/my-cheat-sheets/tree/master/operations-tools/continuous-integration-continuous-deployment/concourse-ci-cheat-sheet/generate-keys.sh
sudo sh generate-keys.sh
```

Run docker-compose up,

```bash
docker-compose up
```

Keep the output minimal,

```bash
docker-compose up -d
```

Note that this will pull two docker images,

```bash
docker images
```

* postgres
* concourse/concourse

I like to keep specific versions, so it doesn't
just keep pulling and storing the latest.

It will then create three docker containers:

```bash
docker ps
```

* postgres
* concourse/concourse (web)
* concourse/concourse (worker)

That's it, check that its working,

[192.168.100.5:8080](http://192.168.100.5:8080)

I could not get login as user
`jeff` and password `test`.

I just used `test` `test`.

To stop all docker containers is simple,

```bash
docker-compose down
```

## USER A STATIC IP TO ACCESS OUTSIDE YOUR MACHINE

If you want to access this from outside you box,
first setup a static IP to a network device.
I have a cheat sheet on that called
[network device configuration](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/development/operating-systems/linux/network-device-configuration-cheat-sheet)

Remember to match the static IP to what's in your
`docker-compose.yml` file,

```text
    - CONCOURSE_EXTERNAL_URL=http://192.168.100.5:8080
```

## IF YOU GET AN IMAGE ERROR

Start fresh, so prune the containers,

```bash
sh prune-containers.sh
```

or

```bash
docker container prune
```

And make sure to stop all running containers,

```bash
docker-compose down
```
