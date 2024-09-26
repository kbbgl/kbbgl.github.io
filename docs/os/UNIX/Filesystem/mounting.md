# Mounting Filesystems

To be able to use a filesystem after creation, we need to mount it somewhere within the tree structure in a child directory within the `/` directory.

To be able to mount a filesystem, we first need to create a directory.

Then we can mount it.

```bash
mkdir /home/mount_point
sudo mount -t ext /dev/sdb4 /home/mount_point
```

This will mount a `ext4` filesystem on the partition `/dev/sdb4` into mount point `/home/mount_point`. Any files residing in `/home/mount_point` will be hidden until `umount`.

We can unmount the filesystem:

```bash
sudo umount /home/mount_point
# OR
sudo umount /dev/sdb4
```

This needs to be done only on filesystems not currently running/in use (`device is busy`).

We can create and mount a filesystem:

```bash
# Create a fs full of zeros
# read/write 1M at a time
# and copy 512 input blocks
dd if=/dev/zero of=/tmp/zeros bs=1M count=512

# Load XFS filesystem
sudo /sbin/mkfs.xfs /tmp/zeros

# OR
sudo mkfs -t xfs /tmp/zeros

# Mount the filesystem 
sudo mount /tmp/zeros /mnt
```

We can see that the filesystem is mounted:

```bash
lsmod | grep xfs

Module                  Size  Used by
xfs                  1204224  1
```

## Mounting at Boot and `/etc/fstab`

During system initialization, the following command is executed:

```bash
mount -a
```

This mounts all filesystems listed in `/etc/fstab` configuration file.
It is also popular to use `autofs` or `automount` tools to automatically mount filesystems.

## Checking Filesystem for Errors

Every filesyste has a utility designed to check and fix errors. The generic name is `fsck`.

```bash
fsck [-t fstype] [options] [device-file] [-a] [-r]

# using -r will prompt fixing issue one by one
# using -a the issues will be fixed automatically
```

`fsck` should only be run on an unmounted filesystem.

We can use the following to force filesystem check of all mounted filesystems upon reboot:

```bash
sudo touch /forcefsck #file will dissappear after a successful check.
sudo reboot
```
