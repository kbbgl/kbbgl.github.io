---
slug: os-linux-filesystem-features
title: "Filesystem Features"
authors: [kbbgl]
tags: [os, linux, filesystem, features]
---

# Filesystem Features

## Extended Attributes

Filesystems support extended attributes (`xattr`). Extended attributes associate metadata not interpreted directly by the filesystem with files.

There are four flags that can be set on files/directories to modify the extended attributes:

- `i`: Immutable, cannot be modified/deleted/renamed and no hardlink.
- `a`: Append-only. Can only be opened for writing and appending.
- `d`: no dump. Files with this flag will be ignored when `dump` program is run.
- `A`: access time (`atime`) of the file will not be updated.

Only the superuser can modify these flags.

We can use `lsattr` to list and `chattr` to modify file/directory `xattr`.

There are four namespaces:

- `user`
- `trusted`
- `security` - used by SELinux.
- `system` - Used to access Access Control Lists (ACLs)

## Quota

Linux allows to control filesystem usage for users/groups.

```bash
sudo apt install quota

quotacheck
quotaon
quotaoff
edquota
quota
```

Quota operations require the existence of the files `aquota.user` and `aquota.group` in the root directory of the filesystem using quotas.
