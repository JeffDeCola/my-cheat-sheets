#!/bin/bash
# concourse-cheat-sheet search-and-replace.sh

set -e -x

echo "List whats in the current directory"
ls -lat 

pwd

cd /my-cheat-sheets
ls -lat

cd ..
touch result.txt
mv result.txt result
cd result
ls -lat
pwd


