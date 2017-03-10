#!/bin/bash
# concourse-cheat-sheet search-and-replace.sh

set -e -x

echo "List whats in the current directory"
ls -lat 

pwd

docker images
docker ps

cd /my-cheat-sheets
ls -lat

