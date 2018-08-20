# GIT CHEAT SHEET

`git` _is a free and open source DVCS (Distributed Version Control System)._

## INSTALL

For Linux,

```bash
sudo apt install git
```

For Windows,

There are two main options,

Option 1. Stand alone, git for windows app (completely independent from
bash on ubuntu on windows)

```bash
https://git-scm.com/downloads
```

Option 2. Open bash on ubuntu on windows and install like linux above.
I'm not a huge fan of bash on windows.

## CONFIGURE FOR LINUX

I also like to add the hostname/machine name so I know where it came from,

```bash
git config --global user.name "Jeff DeCola (HOSTNAME/MACHINE NAME)"
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

## HTTPS ACCESS

Generate a personal access token for your machines .netrc
file at github.com.

Creating a .netrc file and be placing this token in
login field as follows,

```bash
nano ~/.netrc
```

Add,

```bash
machine github.com
login {TOKEN FROM GITHUB}
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

Copy your public ssh key,

```bash
cat ~/.ssh/id_rsa.pub
```

Goto GitHub.com and paste your public ssh key (settings -> ssh keys)

Check the fingerprint at github against your local public key.
You may not need the md5 option,

```bash
ssh-keygen -E md5 -lf ~/.ssh/id_rsa.pub
```

To connect to GitHub use,

```bash
ssh -T git@github.com
```

If you want to use a particular public key file,

```bash
ssh -i ~/.ssh/id_rsa.pub git@github.com
```

Force git to use ssh on your local repo, rather the https,

```bash
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
```

Check ssh settings,

```bash
git config --list
```

Hence,

```bash
remote.origin.url=https://github.com/JeffDeCola/my-cheat-sheets.git
```

Was changed to ssh method,

```bash
remote.origin.url=git@github.com:JeffDeCola/my-cheat-sheets.git
```

## WORKFLOW

The Git Workflow, showing how everything fits together,

![IMAGE - GIT Repo Workflow Diagram - IMAGE](../../../docs/pics/GIT-Repo-Workflow-Diagram.jpg)

A simpler view,

![IMAGE - Simple GIT Repo Workflow Diagram - IMAGE](../../../docs/pics/Simple-GIT-Repo-Workflow-Diagram.jpg)

### CLONE

Via http,

```bash
git clone https://github.com/JeffDeCola/my-cheat-sheets.git
```

Via ssh,

```bash
git clone git@github.com:JeffDeCola/my-cheat-sheets.git
```

I would set the repo to always use ssh,

```bash
git remote set-url origin git@github.com:JeffDeCola/my-cheat-sheets.git
```

### ADD/COMMIT/PUSH

Your first push,

```bash
git add .
git commit -m "initial"
git push origin master
```

Normal flow just use `git push`.

### PULL (fetch and Merge)

```bash
git pull
```

### CHECKOUT - CREATE A BRANCH

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

## REMOVING A COMMIT ON GITHUB

Note: You will destroy the commit history and not at all
nice for collaboration.

```bash
git reset —hard <SHA>
git push —force
```

## TO GET ALL CAUGHT UP

If you just want to have all your local branches
and master up to date, the best way to do this is branch-by-branch.

Note `pull = fetch + merge`.

```bash
git checkout master
git pull

git checkout branch1
git pull

git checkout branch2
git pull
```

## INTERGRATE GIT WITH BASH PROPMPT

I use [git-aware-promp](https://github.com/jimeh/git-aware-prompt).

## MAC-OS git chekcout autocomplete

```bash
curl https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash \
-o ~/.git-completion.bash
```

Add to `~/.bash_profile`,

```bash
if [ -f ~/.git-completion.bash ]; then
  . ~/.git-completion.bash
fi
```

## HOW I CREATE A NEW REPO

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
