# CLONE A REPO AND MAKE IT YOUR OWN CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_How I clone a repo from someone else and make it my own.
Remember to always give credit to the original author and
you must keep the license requirements._

## STEPS

```text
#CLONE AND RENAME
cd clones
git clone git@github.com:<USER>/<REPO>.git
mv <REPO> <REPO>-clone
cd <REPO>-clone

# REMOVE OLD GIT
rm -rf .git

# CREATE REPO ON GITHUB "<REPO>-clone"
# ADD DESCRIPTION "A clone of <USER>/<REPO> with my own changes"

#INITIALIZE NEW GIT REPO
git init -b main

# ADD A README
echo "# <REPO>-clone" > README.md

# ADD COMMIT AND PUSH
git add .
git commit -m "first commit"
git remote add origin git@github.com:JeffDeCola/<REPO>-clone.git
git push -u origin main

# MAKE DEVELOP BRANCH AND PUSH
git checkout -b "develop"
git push -u origin develop
```
