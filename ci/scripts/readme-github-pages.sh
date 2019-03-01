#!/bin/bash
# my-cheat-sheets readme-github-pages.sh

set -e -x

# The code is located in /my-cheat-sheets
echo "pwd is: $PWD"
echo "List whats in the current directory"
ls -lat 

# Note: my-cheat-sheets-updated already created becasue of .yml file
git clone my-cheat-sheets my-cheat-sheets-updated

cd my-cheat-sheets-updated
ls -lat 

# FOR GITHUB WEBPAGES
# THE GOAL IS TO COPY README.md to /docs/_includes/README.md

# Remove everything before the second heading.
sed '0,/GitHub Webpage/d' README.md > temp-README.md
# Change the first heading ## to #
sed -i '0,/##/{s/##/#/}' temp-README.md
# update the image links (remove docs/)
sed -i 's#IMAGE](docs/#IMAGE](#g' temp-README.md

commit="yes"

# CHECK IF THE FILE EXISTS
if (test -f docs/_includes/README.md)
then
    echo "docs/_includes/README.md exists"
    # CHECK IF THERE IS A DIFF
    if (cmp -s temp-README.md docs/_includes/README.md)
    then
        commit="no"
        echo "No Changes, Do not need to commit"
    fi
else
    echo "docs/_includes/README.md does not exist"
    echo "Create the _includes directory"
    mkdir docs/_includes
fi

if [ "$commit" = "yes" ]
then
    echo "cp updated README.md to docs/_includes/README.md"
    cp temp-README.md docs/_includes/README.md
    rm temp-README.md
    
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
