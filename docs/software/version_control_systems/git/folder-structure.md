# Folder Structure

Git includes the following files and folders:

```bash
tree -L 1 .git
.git
├── COMMIT_EDITMSG
├── FETCH_HEAD
├── HEAD
├── ORIG_HEAD
├── config
├── description
├── hooks
├── index
├── info
├── logs
├── objects
├── packed-refs
└── refs

5 directories, 8 files
```

## `config` - file that includes configuration of the `git` repo, has info such as

- `core` - default configs
- `remote` - configured remotes
- `branch`- the different branches

## `HEAD` - file includes the reference to the head

```bash
cat .git/HEAD
ref: refs/heads/master
```

---

# `git` Objects\n\nGit has its own file system and objects

The type of objects:

- **Blob**: represents a file of any type.
- **Tree**: directories, can contain blobs or other trees.
- **Commit**: Store versions of the project.
- **Annotated Tag**: persistent text pointer to specific Commit..

```
commit -> tree -> N blobs/trees
```

Low-level commands to interact with `git` objects:

```bash
# Step 1) Create a blob and tree inside repo
# create a git object
# git uses SHA1 hash function
git hash-object



# creates a file in `.git/objects/b7/aec520dec0a7516c18eb4c68b64ae1eb9b5a5e`
echo "Hello, Git" | git hash-object --stdin -w\nb7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e


# creates a git object based on file 
git hash-object /path/to/file -w


# read a git object
git cat-file

# contents of the object
git cat-file -p $HASH
git cat-file -p b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e\nHello, Git


# size of object
git cat-file -s b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e
11

# type of object
git cat-file -t b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e
blob


# Git Objects store the size, type and content within it
# and have the following structure:
$TYPE $SIZE\\0$CONTENT

echo "blob 11\\0Hello, Git" | shasum
b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e


# create a tree
echo "Hello, Git" | git hash-object --stdin -w
b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e
echo 'Hello, Git!' | git hash-object --stdin -w
670a245535fe6316eb2316c1103b1a88bb519334

# trees contain the file permissions, the type of pointer to the child blob/tree, the hash of the blob/tree and the filename with extension

cat temp-tree.txt
100644 blob b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e    file1.txt
100644 blob 670a245535fe6316eb2316c1103b1a88bb519334    file2.txt

cat temp-tree.txt | git mkdir
9d2ce41b82297aad442e3187d87ce6ee9232f657

# Step 2) Move tree to staging
git read-tree $HASH_TREE
git read-tree 9d2c[41b82297aad442e3187d87ce6ee9232f657]

# list all files in staging area
git ls-files -s
100644 b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e 0       file1.txt
100644 670a245535fe6316eb2316c1103b1a88bb519334 0       file2.txt

# Step 2) Move tree to working directory
git checkout-index -a

```

## What is a Commit

It's a `git` object type (1 of 4).

It has the same structure as other `git` objects:

```
Content + Object Type + Object Length = Hash
```

Every commit has the following information in its contents:

- Author name and email
- Commit description
- Parent (optional)
- Pointer hash to `tree`

The `commit` object is a wrapper to the `tree` object that has a pointer (hash) of the `tree`.

When committing changes that are in the staging area, we will see that hash of the commit:

```bash
# list files in staging
git ls-files -s
100644 b7aec520dec0a7516c18eb4c68b64ae1eb9b5a5e 0       file1.txt
100644 670a245535fe6316eb2316c1103b1a88bb519334 0       file2.txt


# cc83ae1 is the hash of this commit
git commit -m "very first commit"
[master (root-commit) cc83ae1] first commit
2 files changed, 2 insertions(+)
create mode 100644 file1.txt
create mode 100644 file2.txt
```

We can see the contents of the commit:

```bash
git cat-file -p cc83

# the hash that points back to the tree being wrapped by commit
tree 9d2ce41b82297aad442e3187d87ce6ee9232f657
author $NAME $EMAIL 1633710080 +0300
committer $NAME $EMAIL 1633710080 +0300

very first commit
```

If there is another commit based on the first one (`cc83ae1`), it will be seen in `parent`:

```bash
git --no-pager log
commit 4e7b9a4b17d4748f79e033264f89abb537e3047d (HEAD -> master)
Author: Me <email>
Date:   Sat Oct 9 01:06:08 2021 +0300
second commit
commit cc83ae14f363a62b08df4afb719bf87fc57ddd95
Author: Me <email>
Date:   Fri Oct 8 19:21:20 2021 +0300

first commit


git cat-file -p 4e7b9a
tree a9d0efa22cfef4b6b6805e35c9f4ef61d71f9d19
parent cc83ae14f363a62b08df4afb719bf87fc57ddd95
author Me <email> 1633730768 +0300
committer Me <email> 1633730768 +0300

second commit
```
