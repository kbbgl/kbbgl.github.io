# Group Management

Groups are collection of users that have a commonality.
They are defined in:

```plaintext
/etc/group
```

And attributes are:

```plaintext
groupname:password:GID:user1,user2,...
```

## Managing Groups

Some examples:

```bash
# Create system group named staff with GID 215
sudo groupadd -r -g 215 staff

# Change the GID to 101 of group named blah
sudo groupmod -g 101 blah

# Delete a group named newgroup
sudo groupdel newgroup

# Add user student to 3 groups
sudo usermod -G student,group1,group2 student

# Add user rocky to group friends
sudo usermod -G friends rocky
```

### User Private Groups (UPG)

UPG means that each user has their own (primary) group. Upon user creation using `useradd`, the user will get created and added to a group by the same name with GID = UID.

### Group Memberships

Group memberships can be identified by:

```bash
groups [$username]
id -Gn [$username]
```
