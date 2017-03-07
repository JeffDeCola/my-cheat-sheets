#!/bin/bash
# my-cheat-sheets readme-github-pages.sh

set -e -x

# The code is located in /my-cheat-sheets
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Note: my-cheat-sheets-updated already created becasue of .yml file
git clone my-cheat-sheets my-cheat-sheets-updated

cd my-cheat-sheets-updated
ls -lat 

# FOR GITHUB WEBPAGES
# BASICALLY COPY README.md to /docs/_includes/README.md
# Remove everything before the second heading.
sed '0,/GitHub Webpage/d' README.md > temp-README.md
# update the image links (remove docs/)
sed -i 's#IMAGE](docs/#IMAGE](#g' temp-README.md

# CHECK IF THEERE IS A DIFF, IF THERE IS COMMIT, IF NOT DON'T
if !(cmp -s temp-README.md docs/_includes/README.md)
then
    cp temp-README.md docs/_includes/README.md
    rm temp-README.md
    echo "README.md and docs/_includes/README.md differ"
    #ADD AND COMMIT
    git config --global user.email "jeff@keeperlabs.com"
    git config --global user.name "Jeff DeCola (Concourse)"

    git status
    # ONLY add what is needed to protect from unforseen issues.
    git add docs/_includes/README.md
    git commit -m "cp README.md docs/_includes/README.md for GitHub Page"
    git status
else
    rm temp-README.md
fi
echo "complete"
