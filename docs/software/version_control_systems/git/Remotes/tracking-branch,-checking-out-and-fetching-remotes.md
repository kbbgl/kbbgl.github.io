# What is a Tracking Branch?

A tracking branch is a local branch that is directly connected to a remote branch.

If we're cloning a remote repository that has 2 branches (`master (default)` and `br-1`) to local, `git` by will create only the default branch locally.

```bash
❯ git clone https://github.com/mikegithubber/my-first-github-repository
cd my-first-github-repository

# will show the configured default branch
❯ git remote
  *master
  
  # remote branches
  # we can see that the default is `master`
  # also we see there are other remote branches (feature-1/2) that don't exist locally
❯ git --no-pager branch -r
  origin/HEAD -> origin/master
  origin/feature-1
  origin/feature-2
  origin/master
```

To check which branches are tracking and which are not:

```bash
# our local master branch is tracking origin master remote branch
❯ git --no-pager branch -vv
  * master f38cf54 [origin/master] Create hello-github.txt
```

## Checkout Remote Branch

```bash
❯ git checkout feature-1
Branch 'feature-1' set up to track remote branch 'feature-1' from 'origin'.
Switched to a new branch 'feature-1'

❯ git --no-pager branch -vv
  * feature-1 3efc27d [origin/feature-1] Another local change
  master    f38cf54 [origin/master] Create hello-github.txt
  
# shows full info about remote
❯ git remote show origin
  * remote origin
  Fetch URL: https://github.com/mikegithubber/my-first-github-repository
  Push  URL: https://github.com/mikegithubber/my-first-github-repository
  HEAD branch: master
  Remote branches:
  feature-1 tracked
  feature-2 tracked
  master    tracked
  ...
  ```
