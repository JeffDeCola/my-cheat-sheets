#!/bin/bash
# my-cheat-sheets set-pipeline.sh

fly -t ci set-pipeline -p my-cheat-sheets -c pipeline.yml --load-vars-from ../../../.credentials.yml
