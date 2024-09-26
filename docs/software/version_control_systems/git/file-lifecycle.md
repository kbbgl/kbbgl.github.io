# Basic `git` Operations

Move between:

```plain
working dir <-> staging <-> repo
```

## `git` Tracking Statuses

### Untracked

new file in working directory. when using it will move the file to staging area:

```bash
git add $file
```

A file that was staged could be moved back here after using:

```bash
git rm $file
```

### Modified

edited file. can move it to staging using

```bash
git add $file
```

### Staged

file in staging, got here after:

```bash
git add $file
```

### Unmodified

happens after

```bash
git commit
```

## Git File Lifecycle

![lifecycle](https://git-scm.com/book/en/v2/images/lifecycle.png)

```bash
echo \"third file\" >> file3.txt
# move file from untracked to staged

git add file3.txt

# check staging
git ls-files -s
100644 b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e 0       file1.txt
100644 670a245535fe6316eb2316c1103b1a88bb519334 0       file2.txt
100644 667bb3858a056cc96e79c0c3b1edfb60135c2359 0       file3.txt

# unstage file but don't remove it from working dir

git rm --cached file3.txt
rm 'file3.txt'

# check staging
git ls-files -s
100644 b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e 0       file1.txt
100644 670a245535fe6316eb2316c1103b1a88bb519334 0       file2.txt

# read file
git add file3.txt

# check staging
git ls-files -s
100644 b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e 0       file1.txt
100644 670a245535fe6316eb2316c1103b1a88bb519334 0       file2.txt
100644 667bb3858a056cc96e79c0c3b1edfb60135c2359 0       file3.txt

# go to unmodified
git commit -m "second commit"
[master 4e7b9a4] second commit
1 file changed, 1 insertion(+)
create mode 100644 file3.txt
```
