# Storage Devices

## `/dev`

Each device on system is represented by a file in the `/dev` directory.

Notable devices are:

* Hard drive for newer disk technologies such as SATA and SCSI are named with `sd` prefix. Older Linux machines used `hd` for hard drives. The file system treats each hard drive as a different device with an appended letter:

 `sda` - First SATA hard drive.
 `sda` - Second SATA hard drive.

* Drives are sometimes split up into partitions and are represented by numbers in the end of the name:

```text
    sda1 - First partition of first SATA hard drive
    
    sda2 - Second partition of first SATA hard drive
```

To view more information about block devices how much capacity each drive can hold:

```bash
fdisk -l
lsblk
```

## Character and Block devices

Devices can transfer data in and out in two ways.

```bash
ls -latr /dev

crw-rw-rw-   1 root   wheel           24,  22 Jul 22 20:04 dtrace
brw-r-----   1 root   operator         1,   8 Jul 22 20:04 disk1s6
```

The `c` in the permission list row indicates that the `dtrace` device is a **chatacter** device which sends and receives data character by character such as mice and keyboards.

The `b` indicates the device sends and receives data in blocks of bytes. Hard drives and USB sticks are block devices which have higher data throughput.

## Mounting Devices

A storage device must be first physically connected then logically connected in order to be available for the filesystem. The _mount point_ is the location where devices are attached. External USB drives are usually `/media` while internal hard drives are mounted in `/mnt`.

To mount a device:

```bash
mount /dev/sdb1 /mnt
```

Mount point should be empty!

The filesystems that are mounted are kept in `/etc/fstab`

To unmount:

```bash
umount /dev/sdb1
```

We can run `df` to retrieve information about free and used space of any devices/hard drives.

## Repairing a device

First need to unmount a device and then run:

```bash
# -p automatically repairs device
fsck -p /dev/sdb1
```
