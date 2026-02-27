---
slug: os-linux-disk-lvm
title: "LVM (Logical Volume Management)"
authors: [kbbgl]
tags: [os, linux, disk, lvm]
---

# LVM (Logical Volume Management)

Virtual devices may be easier to manage than physical devices, and can  have  capabilities  beyond  what  the  physical devices provide themselves.  A Volume Group (**VG**) is a collection of one or more physical devices, each called a Physical Volume (**PV**).  A Logical Volume (**LV**) is a virtual block device that can be used by the system or applications.   Each  block  of  data in an LV is stored on one or more PV in the VG, according to algorithms implemented by Device Mapper (**DM**) in the kernel.

![lvm](https://i2.wp.com/manjaro.site/wp-content/uploads/2017/08/lvm-on-ubuntu.png?resize=678%2C381&ssl=1)

To manage VGs:

```bash
vgdisplay
vgcreate
vgextend
vgreduce
vgs # list vgs
```

Manage PVs:

```bash
pvcreate
pvdisplay
pvmove
pvremove
pvs # see all physical volumes
```

All tools to manage LVs in:

```bash
ls -lF /sbin/lv*
lvcreate
lvremove
lvs # see all lvs
```

All LVM utilities are listed in:

```bash
man lvm
```

PV -> VG -> LV path:

```bash
/dev/$vg_name/$lv_name
```

## Creating LVs

1. Create partitions on disk drives (using `fdisk /dev/sda` `t` `8e` option to change type to LVM)
2. Create physical volumes from the partitions.

    ```bash
    sudo pvcreate /dev/sda4
    pvdisplay
    ```

3. Create the VG.

    ```bash
    sudo vgcreate -s 16 vg /dev/sda4
    vgdisplay -v # provides vg, pv and lv
    ```

4. Allocate LV from VG.

    ```bash
    # if needed
    sudo vgextend vg /dev/sda4
    sudo lvcreate -L 50G -n mylvm vg
    lvgdisplay
    ```

5. Format LV

    ```bash
    sudo mkfs -t ext4 /dev/vg/mylvm
    ```

6. Mount LV (and update `/etc/fstab` if needed)

    ```bash
    sudo mkdir /mylvm
    sudo mount /dev/vg/mylvm /mylvm

    echo "/dev/vg/mylvm /mylvm ext4 defaults 1 2" >> /etc/fstab
    ```

To remove:

```bash
sudo umount /mnt
sudo lvremove /dev/vg/mylvm
```

### Resizing LVs

When expanding the size of an LV with a filesystem, we first need to expand the LV then the fs. When shriking an LV with a filesystem, we first need to shrink the fs and then the LV.

```bash
# the -r options resizes fs at the same time as the LV is changed
sudo lvresize -r -L 20 GB /dev/vg/mylvm

# or
sudo lvreduce -L 10G /dev/$vg/$lv
```

Reduce a VG:

```bash
sudo pvmove /dev/sda4
sudo vgreduce vg /dev/sda4
```

### LVM Snapshots

Useful for backups, deploying VMs, app testings as the create an exact copy of the VG.

To create a snapshot:

```bash
sudo lvcreate -l 128 -s -n mysnap /dev/vg/mylvm
```

To make mount:

```bash
mkdir /mysnap
mount -o ro /dev/vg/mysnap /mysnap
```

To use and remove snapshot:

```bash
sudo umount /mysnap
sudo lvremove /dev/vg/mysnap
```
