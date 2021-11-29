#!/bin/bash
# my-cheat-sheets destroy-pipeline.sh

echo " "
echo "Destroy pipeline on target jeffs-ci-target which is team jeffs-ci-team"
fly --target jeffs-ci-target \
    destroy-pipeline \
    --pipeline my-cheat-sheets
echo " "
