---
slug: os-linux-disk-raid
title: "RAID"
authors: [kbbgl]
tags: [os, linux, disk, raid]
---

# RAID

**RAID** (Redundant Array of Independent Disks) is a data storage virtualization technology.

RAID allows:

- To spread I/O across multiple disks to increase performance and fault tolerance.
- Implemetation on the either HW/SW level.
- Mirroring: writing the same data to more than one disk.
- Striping: splitting data to more than one disk.
- Parity: extra data is stored to allow problem detection and repair.

RAID creates filesystems which spans more than one disk. RAID devices are created by combining partitions from several disks together.

To manage RAID devices:

```bash
mdadm
```

Raid device can then be used like any filesystem (`/dev/mdX`)

## Levels

- **RAID 0**: only striping so gives performance boost. no redundancy so any disk fail will cause total lost.
- **RAID 1**: only mirroring. each disk has a duplicate so at least two disks required. good for recovery.
- **RAID 5**: rotating parity stripe. single disk failure will cause no data loss, only performance drop. must have 3 disks.
- **RAID 6**: striping with dual parity. can handle 2 disk failures and requires 4 disk. RAID 6 is preferred to RAID 5.
- **RAID 10**: mirrored and striped. at least 4 disks needed.

## Configuring RAID

1) Create partitions using `fdisk` ((t)ype `fd`)

    ```bash
    sudo fdisk /dev/sdb
    sudo fdisk /dev/sdc
    ```

2) Create RAID device using `mmdadm`.

    ```bash
    sudo mdadm --create /dev/md0 --level=1 --raid-disks=2 /dev/sdb1 /dev/dev/sdc1
    ```

3) Format RAID device.

    ```bash
    sudo mkfs.ext4 /dev/md0
    ```

4) Add device to `/etc/fstab` and capture configuration to file:

    ```bash
    echo "/dev/md0 /myraid ext4 defaults 0 2" >> /etc/fstab
    sudo bash -c "mdadm --detail --scan >> /etc/mdadm.conf"
    ```

5) Mount RAID device.

    ```bash
    sudo mkdir /myraid
    sudo mount /dev/md0 /myraid
    ```

6) Check RAID device

    ```bash
    cat /proc/mdstat
    ```

To stop RAID device:

```bash
sudo mdadm -S /dev/md0
```

## Monitoring

```bash
sudo mdadm --detail /dev/md0

# OR 
cat /proc/mdstat

# OR (requires configuring /etc/mdadm.conf, can set MAILADDR me@domain.com to receive email.
# need to turn on 
sudo systemctl start mdmonitor
sudo systemctl enable mdmonitor
mdmonitor 
```

## Hot Spares

Can use hot spares for redundancy to fix an issue ASAP.

To initialize a RAID device with hot spare:

```bash
# -x 1 tells mdadm to configure one psare device.
sudo mdadm --create /dev/md0 -l 5 -n3 -x 1 /dev/sda8 /dev/sda9 /dev/sda10 /dev/sd11
```

To add it later:

```bash
sudo mdadm --fail /dev/md0 /dev/sdb2
```

To replace faulty with spare:

```bash
sudo mdadm --remove /dev/md0 /dev/sdb2
sudo mdadm --add /dev/md0 /dev/sdc2
```
