#!/usr/bin/env bash

docker container prune
docker stop $(docker ps -aq)
#docker rm $(docker ps -aq)

