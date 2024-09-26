# User Management

## Attributes

Each user registered in the system has a line entry in `/etc/passwd` which has all attributes of the user are listed. The attributes are:

- Username.
- UID
- GID
- Home directory
- Default Shell

```bash
# beav - username
# 'x' means password is set and used in `/etc/shadow`.
# 1000 - UID
# 1000 - GID
# Theodore Cleaver - Comment
beav:x:1000:1000:Theodore Cleaver:/home/beav:/bin/bash
```

There are some system users (such as `bin`, `daemon`, `sys`) which are created for specific purposes and cannot be used to log in (this is why their default shell is `/sbin/nologin`).

### Creating User Accounts

The command will create a user with the default options:

```bash
sudo adduser $username
```

The following steps will occur when adding a new user (according to defaults set in `/etc/defaults/useradd`):

1) The next available UID value (based on `/etc/login.defs`) will be assigned.

2) A group called `$username` will be created and UID will be set to GID.

3) A home directory `/home/$username` will be created.

4) The contents of `/etc/skel` will be copied to `/home/$username`.  

5) An entry of either `!!` or `!` is placed in the `password` field of the `/etc/shadow` file for `$username`'s entry, thus requiring the administrator to assign a password for the account to be active.

### Modifying User Accounts

This command will remove the user from `/etc/passwd`, `/etc/shadow` and `/etc/group`.

```bash
sudo userdel $username

# Will also remove the home directory for this user
sudo userdel -r $username 
```

This command will modify the user attributes.

```bash
sudo usermod $username

# Will lock account so the user cannot log in
sudo usermod -L $username
# OR
sudo chage -E 1970-01-01 $username

# Will unlock an account
sudo usermod -U $username
```

### Password Management

Use of `/etc/shadow` enables password aging on a per user basis. At the same time, it also allows for maintaining greater security of hashed passwords.

Only user can change its own password (unless `root` which can change for all users).
To change passwords:

```bash
passwd
```

To set password aging (using `root`):

```bash
chage

# Force user to change password on next login
sudo chage -d 0 $username
```
