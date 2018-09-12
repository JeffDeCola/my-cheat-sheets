#!/usr/bin/env bash

docker container prune
docker volume prune

#docker-compose rm -fv
#/bin/bash -c 'docker volume ls -qf "name=concourse_" | xargs docker volume rm'
#/bin/bash -c 'docker network ls -qf "name=concourse_" | xargs docker network rm'
#/bin/bash -c 'docker ps -aqf "name=concourse_*" | xargs docker rm'

