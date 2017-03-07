# GIT CHEAT SHEET

`git` _is a free and open source Distributed Version Control System._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL

```bash
sudo apt install git
```

## CONFIGURE

```bash
git config --global user.name "Jeff DeCola"
git config --global user.email EMAIL
git config --global core.editor nano
```

Check,

```bash
git config --list
```

## SSH KEY

Generate public and private key for ubuntu on virtual box and ubuntu on windows,

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

To connect use,

```bash
ssh -T git@github.com
```

Force git to use ssh rather the http (Get this from GitHub.com),

```bash
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
```

Check Settings,

```bash
git config --list
```

## CLONE

```bash
git clone https://github.com/JeffDeCola/REPONAME.git
```

## FIRST PUSH

```bash
git add .
git commit
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
git push origin master
```

## CREATE BRANCH

```bash
git checkout -b "develop" master
git push --set-upstream origin develop
```

## HOW I CREATE A NEW REPO FOR THE FIRST TIME

### Setup repo on GitHub

Do this first.

### Create local repo

````bash
git clone https://github.com/JeffDeCola/REPONAME.git
cd REPONAME
git status
```

### Add Files

```bash
README.md
LICENSE
.gitignore
.codeclimate
-R /ci
-R /docs
```

Update all the above files with new REPONAME.

### Initial Push

```bash
git add .
git commit (initial)
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
git push origin master
```

### Create "develop" Branch

```bash
git checkout -b "develop" master
git push --set-upstream origin develop
```

### GitHub WebPage Setting

In Settings add GitHub Webpage on /docs

### Add REPONAME to codeclimate

Login to codeclimate and add new REPONAME.

### Configure Concourse

```bash
fly -t ci set-pipeline -p REPONAME -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
```

### Snap a Concourse Picture for README.md

Place in `/docs/pics/REPONAME-pipeline.jpg`
