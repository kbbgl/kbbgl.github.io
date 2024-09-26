# What is `origin`?

`origin` is the default name of the remote repository after cloning a remote repo to local:

```bash
❯ git clone https://github.com/mikegithubber/my-first-github-repository.git\n\nCloning into 'my-first-github-repository'...

remote: Enumerating objects: 52, done.
remote: Total 52 (delta 0), reused 0 (delta 0), pack-reused 52
Receiving objects: 100% (52/52), 8.36 KiB | 2.09 MiB/s, done.
Resolving deltas: 100% (9/9), done.

❯ cd my-first-github-repository
❯ git remote
origin

# git push and fetch operations will use the same remote
❯ git remote -v
origin  https://github.com/mikegithubber/my-first-github-repository (fetch)
origin  https://github.com/mikegithubber/my-first-github-repository (push)
```
