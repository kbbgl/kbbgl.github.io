---
title: git Cheat Sheet
slug: git-cheatsheet
tags: [git,cheatsheet,vcs]
authors: [kgal-pan]
---

# `git` Cheat Sheet

## Create a Local Branch from An Upstream Branch

```bash
git fetch upstream $REMOTE_BRANCH:$REMOTE_BRANCH
```
## Set the origin URL

```bash
git remote set-url origin git@github.com:nikhilbhardwaj/abc.git
```

## Remove file from cache

```bash
git rm --cached file1.txt
```

## Creating a feature branch\task

```bash
git checkout -b new-feature
git status
git add <some-file>
git commit
git push -u origin new-feature
```

## Merging feature branch with `master`

```bash
git checkout master
git pull origin master
git merge new-feature
git push origin master
```

## Stage any changes to tracked files and commit them in one step

```bash
git commit -a -m "COMMIT MSG"
```

## Updating from remote

### retrieve latest metadata from remote

```bash
git fetch
```

### retrieve latest metadata AND transfer the files

```bash
git pull
```
