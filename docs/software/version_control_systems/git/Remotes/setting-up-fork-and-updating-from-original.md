# Setting up fork and updating from original"

After we create a fork, we need to set up a remote so we can pull from it.
https://stackoverflow.com/questions/7244321/how-do-i-update-or-sync-a-forked-repository-on-github/7244456#7244456

## 1) Configuring remote for a fork

```bash
# Check remotes
$ git remote -v
> origin  https://github.com/app/content.git (fetch)
> origin  https://github.com/app/content.git (push)

# add new remote upstream repo
$ git remote add upstream https://github.com/app/content.git

$ git remote -v
> origin    https://github.com/app/content.git (fetch)
> origin    https://github.com/app/content.git (push)
> upstream  https://github.com/app/content.git (fetch)
> upstream  https://github.com/app/content.git (push)
```

## 2) Syncing a fork

```bash
# fetch commits and branches from original
git fetch upstream

# checkout master to prepare to merge
git checkout master

# merge the changes without losing local changes
git merge upstream/master
```

## 3) Creating PR from a fork

1. go to [app Content repo](https://github.com/app.content.git)

1. click on Pull Request:

  ![pr](https://docs.github.com/assets/images/help/pull_requests/pull-request-start-review-button.png)

1. Select `master` base repository.

## [Another way](https://github.com/app/content/pull/17125/files)

```bash
#!/usr/bin/env bash

#Be aware, only contributors should run this script.

CONTENT_URL='https://github.com/app/content.git'

if [ -z "$1" ]
then
  CURRENT=$(git branch --show-current)
else
  CURRENT=$1
fi

(
  git remote add upstream_content $CONTENT_URL ||
  git remote set-url upstream_content $CONTENT_URL
) &&
git fetch upstream_content &&
git checkout master &&
git rebase upstream_content/master &&
git push -f origin master &&
git checkout $CURRENT &&
git pull origin master
```
