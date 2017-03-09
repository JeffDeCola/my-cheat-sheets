# GIT CHEAT SHEET

`git` _is a free and open source DVCS (Distributed Version Control System)._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL

For linux,

```bash
sudo apt install git
```

For Winodws,

```bash
https://git-scm.com/downloads
```

## CONFIGURE FOR LINUX

```bash
git config --global user.name "Jeff DeCola (HOSTNAME)"
git config --global user.email YOUREMAIL
git config --global core.editor nano
```

Check configuration,

```bash
git config --list
```

## CONFIGURE FOR WINDOWS

Same as above but may also have to add permissions to ci scripts.

e.g.

```bash
git update-index --chmod=+x readme-github-pages.sh
```

## SSH KEY

Generate public and private key for Ubuntu,

```bash
ssh-keygen -t rsa -b 4096 -C "phrase"
```

Now add key,

```bash
ssh-add ~/.ssh/id_rsa
```

Copy this file,

```bash
cat ~/.ssh/id_rsa.pub
```

Goto GitHub and add your key.

To connect ti GitHub use,

```bash
ssh -T git@github.com
```

Force git to use ssh rather the http,

```bash
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
```

Check Settings,

```bash
git config --list
```

## CLONE A REPO

```bash
git clone https://github.com/JeffDeCola/REPONAME.git
```

## INITIAL PUSH

```bash
git add .
git commit -m "initial"
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
git push origin master
```

## CREATE A BRANCH

```bash
git checkout -b "develop" master
git push --set-upstream origin develop
```

## LOST DATA

If you lost local data (maybe a snapshot screwed up), 
just grab the lastest code from the git master. 

```bash
git fetch origin master
git reset --hard FETCH_HEAD
git clean -df
```

## INTERGRATE GIT WITH BASH PROPMPT

I use [git-aware-promp] https://github.com/jimeh/git-aware-prompt

# HOW I CREATE A NEW REPO

First setup a repo on github.

Clone the Repo,

```bash
git clone https://github.com/JeffDeCola/REPONAME.git
cd REPONAME
git status
```

Add Files,

```bash
README.md
LICENSE
.gitignore
.codeclimate
update_concourse.sh
-R /ci
-R /docs
```

Update all the above files with your REPONAME.

Initial Push,

```bash
git add .
git commit -m "initial"
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
git push origin master
```

Create a "develop" Branch,

```bash
git checkout -b "develop" master
git push --set-upstream origin develop
```

At GitHub add GitHub Webpage on /docs.

Add REPONAME to codeclimate.

Login to codeclimate and add new REPONAME.

Configure Concourse CI,

```bash
fly -t ci set-pipeline -p REPONAME -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
```

Snap a Concourse picture for README.md and place in
`/docs/pics/REPONAME-pipeline.jpg`.
