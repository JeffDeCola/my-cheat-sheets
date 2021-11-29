#!/bin/sh
# my-cheat-sheets readme-github-pages.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "readme-github-pages.sh -debug (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status. Needed for concourse.
    # set -x enables a mode of the shell where all executed commands are printed to the terminal.
    set -e -x
    echo " "
else
    echo "readme-github-pages.sh (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status.  Needed for concourse.
    set -e
    echo " "
fi

echo "GOAL ----------------------------------------------------------------------------------"
echo " "

echo "The goal is to git clone /my-cheat-sheets to /my-cheat-sheets-updated"
echo "Then script will edit the /docs/_includes/README.md for GITHUB WEBPAGES"
echo "Finally push the changes in /docs/_includes/README.md to github"
echo " "

echo "CHECK THINGS --------------------------------------------------------------------------"
echo " "

echo "At start, you should be in a /tmp/build/xxxxx directory with two folders:"
echo "   /my-cheat-sheets"
echo "   /my-cheat-sheets-updated (created in task-build-push.yml task file)"
echo " "

echo "pwd is: $PWD"
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "GIT CLONE -----------------------------------------------------------------------------"
echo " "

echo "git clone my-cheat-sheets to my-cheat-sheets-updated"
git clone my-cheat-sheets my-cheat-sheets-updated
echo " "

echo "cd my-cheat-sheets-updated"
cd my-cheat-sheets-updated
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "EDIT README FOR GITHUB WEBPAGES -------------------------------------------------------"
echo " "

echo "Copy README.md to /docs/_includes/README.md and edit"
echo "    Remove everything before the second heading in README.md.  Place in temp-README.md"
echo "    sed '0,/github webpage/d' README.md > temp-README.md"
sed '0,/github webpage/d' README.md > temp-README.md
echo "    Change the first heading ## to #"
echo "    sed -i '0,/##/{s/##/#/}' temp-README.md"
sed -i '0,/##/{s/##/#/}' temp-README.md
echo "    Update the image links (remove docs/)"
echo "    sed -i 's#IMAGE](docs/#IMAGE](#g' temp-README.md"
sed -i 's#IMAGE](docs/#IMAGE](#g' temp-README.md
echo "    Update the image links for svgs (if you have them)"
echo "    Add \"https://raw.githubusercontent.com/JeffDeCola/REPONAME/master/svgs/\" to \"svgs/\""
echo "    sed -i 's/svgs\//https:\/\/raw.githubusercontent.com\/JeffDeCola\/my-cheat-sheets\/master\/svgs\//g' temp-README.md"
sed -i 's/svgs\//https:\/\/raw.githubusercontent.com\/JeffDeCola\/my-cheat-sheets\/master\/svgs\//g' temp-README.md
echo " "

echo "GIT COMMIT OR NOT ---------------------------------------------------------------------"
echo " "

commit="yes"

echo "Does docs/_includes/README.md exist?"
if test -f docs/_includes/README.md
then
    echo "    Yes, it exists."
    # CHECK IF THERE IS A DIFF
    if (cmp -s temp-README.md docs/_includes/README.md)
    then
        commit="no"
        echo "    No changes are needed, Do not need to git commit and push"
    else
        echo "    Updates are needed"
    fi
    echo " "
else
    echo "    No, it does not exist"
    echo "    Creating the _includes directory if it doesn't exist"
    mkdir -p docs/_includes
    echo " "
fi

if [ "$commit" = "yes" ]
then

    echo "GIT SETUP -------------------------------------------------------------------------"
    echo " "

    echo "cp temp-README.md docs/_includes/README.md"
    cp temp-README.md docs/_includes/README.md
    echo " "

    echo "Update some global git variables"
    echo "git config --global user.email \"jeffdecola@gmail.com\""
    echo "git config --global user.name \"Jeff DeCola (Concourse)\""
    git config --global user.email "jeffdecola@gmail.com"
    git config --global user.name "Jeff DeCola (Concourse)"
    echo " "
    git config --list
    echo " "

    echo "GIT PUSH MASTER BRANCH ------------------------------------------------------------"
    echo " "

    echo "git add and commit what is needed to protect from unforseen issues"
    echo "git add docs/_includes/README.md"
    git add docs/_includes/README.md
    echo " "

    echo " git commit -m \"Update docs/_includes/README.md for GitHub WebPage\""
    git commit -m "Update docs/_includes/README.md for GitHub WebPage"
    echo " "

    echo "git status"
    git status
    echo " "
    
    echo "git push  - Not needed here since its done in pipeline"
    echo " "
fi

echo "CLEAN UP ------------------------------------------------------------------------------"
echo " "

echo "rm temp-README.md"
rm temp-README.md
echo " "

echo "readme-github-pages.sh (END)"
echo " "
