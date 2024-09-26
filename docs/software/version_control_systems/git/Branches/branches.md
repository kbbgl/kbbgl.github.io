# Branches

A `branch` is just a text reference to a `commit`.
The pointers for all `branch`es are located in:

```bash
cat .git/refs/heads
```

The current `branch` head always points to the latest `commit`:

```bash
git log --no-page
commit 4e7b9a4b17d4748f79e033264f89abb537e3047d (HEAD -> master)
Date:   Sat Oct 9 01:06:08 2021 +0300
second commit
commit cc83ae14f363a62b08df4afb719bf87fc57ddd95
Date:   Fri Oct 8 19:21:20 2021 +0300

first commit

cat .git/refs/heads/master\n4e7b9a4b17d4748f79e033264f89abb537e3047d
```

## `HEAD`

`HEAD` is the pointer to a currently checked-out `branch`. The currently checked-out `branch` points to the latest `commit`. `HEAD` can also point to a specific `commit`

There is only one `HEAD`.

The pointer is located in:

```bash
cat .git/HEAD
ref: refs/heads/master
```

To change the reference to a specific `branch`:

```bash
git checkout $BRANCH_NAME
```

To change the reference to a specific `commit`:

```bash
git checkout $COMMIT_HASH
cat .git/HEAD
$COMMIT_HASH
```

Checking out a specific `commit` will trigger `detached HEAD` since the `HEAD` is not pointing to a `branch`.

```bash
git --no-pager log
# `HEAD` not pointing to `master` (e.g. `HEAD` -> `master`)
commit 4e7b9a4b17d4748f79e033264f89abb537e3047d (HEAD)\nDate:   Sat Oct 9 01:06:08 2021 +0300

second commit

commit cc83ae14f363a62b08df4afb719bf87fc57ddd95\nDate:   Fri Oct 8 19:21:20 2021 +0300\n\n    first commit\n```\n\n\n## Management
```bash
# create new branch
git branch temp

# create branch and checking it out
git branch -b temp

# delete branch
git branch -d $BRANCH_NAME

# rename
git branch -m $BRANCH_OLD_NAME $BRANCH_NEW_NAME
```

New `branch`es will be based on the current one:

```bash
git --no-pager branch
  master

# create a new branch based on currently checked-out branch (master)
git branch temp
git --no-pager branch
master
temp

cat .git/refs/heads/temp\n1ed1cdd2f5d87602a0582db36d9a1a73ac93bb34

cat .git/refs/heads/master\n1ed1cdd2f5d87602a0582db36d9a1a73ac93bb34

git checkout temp
Switched to branch 'temp'

cat .git/HEAD
ref: refs/heads/temp
```
