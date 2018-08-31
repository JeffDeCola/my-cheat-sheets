#!/bin/bash
# my-cheat-sheets update_concourse.sh

fly -t ci set-pipeline -p my-cheat-sheets -c ci/pipeline.yml --load-vars-from ../.credentials.yml
