---
slug: os-linux-filesystem-inode
title: "inodes"
authors: [kbbgl]
tags: [os, linux, filesystem, inode]
---

# inodes

An `inode` is an index node. It serves as a unique identifier for a specific piece of metadata on a given filesystem.

the maximum number of `inode`s is `2^32`. The number of `inode`s on a system is `1:16KB` of system capacity.

To check the number of `inode`s on the system, run:

```bash
df -i /dev/sdb

Filesystem       Inodes IUsed    IFree IUse% Mounted on
/dev/sdb       78643200  8407 78634793    1% /opt
```

We can also see the number of `inode`s for a specific file:

```bash
ls -i discovery.js

7749258 discovery.js
```

We can check the number of `inode`s per directory:

```bash
ls -idl middlewares

7749233 drwxrwxr-x 7 ansible ansible 4096 Aug 13 10:52 middlewares
```
