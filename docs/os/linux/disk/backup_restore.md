---
slug: os-linux-disk-backup-restore
title: "Backing up and Restoring Partition Tables"
authors: [kbbgl]
tags: [os, linux, disk, backup_restore]
---

# Backing up and Restoring Partition Tables

```bash
# Backup
sudo dd if=/dev/sda of=mbrbackup bs=512 count=1

# Restore
sudo dd if=mbrbackup of=/dev/sda bs=512 count=1
```

For GPT systems, we can use:

```bash
sudo sgdisk --backup=/tmp/sda_backup /dev/sda
```

Other tools that can be used:

- `fdisk`
- `sfdisk`
- `parted`
- `gdisk`

We can see what partitions:

```bash
cat /proc/partitions 
```
