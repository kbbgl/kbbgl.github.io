# List Remote and Local Branches

After cloning a remote repo, `git` will not automatically create remote branches in local, except for `master`.

To see all branches (remote + local):

```bash
❯ git --no-pager branch -a
* master
  remotes/origin/HEAD -> origin/master
  remotes/origin/feature-1
  remotes/origin/feature-2
  remotes/origin/master
```

Remote branches have the `remotes/` path.

To see all remote branches:

```bash
❯ git --no-pager branch -r
  origin/HEAD -> origin/master
  origin/feature-1
  origin/feature-2
  origin/master
```
