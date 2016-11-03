# GIT CHEAT SHEET - HOW TO CREATE A NEW REPO

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

git checkout -b "develop" master
git push --set-upstream origin develop

## SSH KEY

Once setup just need to run this,

```bash
git remote set-url origin git@github.com:JeffDeCola/REPONAME.git
```

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

Goto github and add your key.

To connect use,

```bash
ssh -T git@github.com
```