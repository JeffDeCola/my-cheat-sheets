# DOCKER CHEAT SHEET

Concourse is distributed as a single concourse binary,
making it easy to run just about anywhere, especially with Docker.

## INSTALL CONCOURSE

Get the `docker-compose` yml file from concourse and run docker,

```bash
$ wget https://concourse-ci.org/docker-compose.yml
$ docker-compose up
```

Which will show something like,

```bash
Creating docs_concourse-db_1 ... done
Creating docs_concourse_1    ... done
Concourse will be running at 127.0.0.1:8080.
```

That's it, check that its working,

[http://127.0.0.1:8080](http://127.0.0.1:8080)

You can also check the images you have in your docker ???,

```bash
docker images
```

postgres and concourse/concourse

And see which containers are running,

```bash
docker ps
```

## DOCKER-COMPOSE .yml FILE

For refererence, the docker compose .yml fiels looks like,

```yml
version: '3'

services:
  concourse-db:
    image: postgres
    environment:
    - POSTGRES_DB=concourse
    - POSTGRES_PASSWORD=concourse_pass
    - POSTGRES_USER=concourse_user
    - PGDATA=/database

  concourse:
    image: concourse/concourse
    command: quickstart
    privileged: true
    depends_on: [concourse-db]
    ports: ["8080:8080"]
    environment:
    - CONCOURSE_POSTGRES_HOST=concourse-db
    - CONCOURSE_POSTGRES_USER=concourse_user
    - CONCOURSE_POSTGRES_PASSWORD=concourse_pass
    - CONCOURSE_POSTGRES_DATABASE=concourse
    - CONCOURSE_EXTERNAL_URL
    - CONCOURSE_ADD_LOCAL_USER=test:$$2a$$10$$0W9/ilCpYXY/yCPpaOD.6eCrGda/fnH3D4lhsw1Mze0WTID5BuiTW
    - CONCOURSE_MAIN_TEAM_ALLOW_ALL_USERS=true
    - CONCOURSE_WORKER_GARDEN_NETWORK
```