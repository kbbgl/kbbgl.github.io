# File Permissions and Ownership

Each file has access rights associated to it. There are 3 groups of permissions:

- **owner/user**: the owner of the file .
- **group**: the group of users who have access.
- **world/other**: rest of the world.

In the example below:

```bash
ls -l /usr/bin/vi
-rw-rw-r-- 1 coop aproject 1601 Mar 9 15:04 a_file
```

The user is `coop` and the group is `aproject`.

Each file also has access right:

- `r`: read access is allowed
- `w`: write access
- `x`: execute.
- `-`: means that the permission is not allowed on the specific access right.

## `chmod`

A user that owns a file can change it's permissions using `chmod` (unless root).

For example:

```bash
$ ls -l a_file
-rw-rw-r-- 1 coop coop 1601 Mar 9 15:04 a_file

# give user and world execution permission
# remove group write permission

$ chmod uo+x,g-w a_file
$ ls -l a_file
-rwxr--r-x 1 coop coop 1601 Mar 9 15:04 a_file
```

We can also set permissions using octal digits. The octal number representation is the sum of each digit:

- `4` if the read permission is desired
- `2` if the write permission is desired
- `1` if the executre permission is desired

```bash
# 7 means rwx for user
# 5 means rx for group
# 5 means rx for world
chmod 755 a_file
```

We can also see the list of permissions in octal:

```bash
stat -c "%a %n" a_file
664 file
```

![chmod](https://lwstatic-a.akamaihd.net/kb/wp-content/uploads/2019/11/fig_permissions_chmod-command.jpg)

### `chown` and `chgrp`

`chown` is used to change file ownership. `chgrp` is used to change the group file ownership.

```bash
chgrp $some_group /path/to/file
```

```bash
chown $some_user /path/to/file
```

We can change both the group and user file ownership in one:

```bash
chown $some_user:$some_group /path/to/file

# Recursive
chown -R $some_user:$some_group /path/to/dir
```

### `umask`

`umask`, or the user file-creation mode, is a Linux command that is used to assign the default file permission sets for newly created folders and files.

When a new file is created, it is created by default with:

```bash
touch file
stat -c "%a %n" file
664
```

When a new directory is created:

```bash
mkdir dir
stat -c "%a %n" dir
775 dir
```

The default value is:

```bash
umask
0002
```

Which means that the created file default permission for world will drop by 2 (`666 -> 664`) and the directory by 2 (`777 -> 775`).

We can set it by:

```bash
umask 0022
```

### Filesystem Access Control Lists (ACLs)

Linux extends the user/group/world and r/w/x permission model with the full POSIX ACLs.

The ACLs must be implemented in the particular filesystem for them to be available.

Some commands to manage ACLs:

```bash
getfacl file|directory
setfacl -m u:$user:rx /home/$user2/file1

# Remove ACL
setfacl -x u:$user /home/$user2/file

```
