---
slug: os-linux-account-management-controlling-file-directory-permissions
title: "Controlling File and Directory Permissions"
authors: [kbbgl]
tags: [os, linux, account_management, controlling_file_directory_permissions]
---

# Controlling File and Directory Permissions

## Granting ownership to specific user

```bash
chmod $USER $FILE
chmod kg /etc/hosts
```

## Granting ownership to specific group

```bash
chgrp $GROUP $FILE
```

## Checking permissions

```bash
ls -l /path/to/dir
```

Taking an example value of `drwxrwxrwx+`, the meaning of each character is explained in the following tables:

`d`
The file type, technically not part of its permissions. See info ls -n "What information is listed" for an explanation of the possible values.

`rwx`
The permissions that the owner has over the file.

`rwx`
The permissions that the group has over the file.

`rwx`
The permissions that all the other users have over the file.

`+` A single character that specifies whether an alternate access method applies to the file. When this character is a space, there is no alternate access method. A . character indicates a file with a security context, but no other alternate access method. A file with any other combination of alternate access methods is marked with a + character, for example in the case of **Access Control Lists**.

### Changing Permissions

Only the file owner or `root` can change permissions.

Can be changed using digits:
![alt](https://danielmiessler.com/images/permissions.png)

Or using `UGO` (user, group, other):

`+` - add permission
`-` - removes permission
`=` - sets permission

Remove permission for user `u` for file:

```bash
chmod u-w /etc/hosts
```

Remove permission for group `g` for file:

```bash
chmod g-w /etc/hosts
```

### Changing Default Permissions

 `unmask` takes the Linux base permissions and _subtracts_ the `unmask` values to set the default permissions.

![](https://www.computernetworkingnotes.org/images/linux/rhce-study-guide/ls22-c-1-calculate-default-permisison.png)

Each user can set a personal default umask value for the files and directories in their personal `~/.profile` file. To see the current value when logged on as the user, simply enter the command `umask` and note what is returned.

### SUID/SGID bits

We can grant temporary `root` access to a user or group by setting the appropriate bit.

To set the SUID bit, add a `4` to the beginning of the permission:

```bash
chmod 4644 /path/to/file
```

To set the SGID bit, add a `2`:

```bash
chmod 2644 /path/to/file
```

We can look for all files that have these bits set:

```bash
find / -user root -perm 4000
```

```bash
find / -user root -perm 2000
```

When the SUID bit is set, we will see it as an `s` instead of `x` when checking the permission of the file:

```
-rwxr-xr-x 1 root root 26696 Mar 17 2020 sucrack
-rwsr-xr-x 1 root root 140944 Jul 5 2020 sudo
```
