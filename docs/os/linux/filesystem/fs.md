---
slug: os-linux-filesystem-fs
title: "Filesystems"
authors: [kbbgl]
tags: [os, linux, filesystem, fs]
---

# Filesystems

Applications write and read files. The files are abstractions to the actual physical/hardware layer. The filesystem is an interface between the applications and the I/O layer.

Multiple filesystems are usually merged together to form a unified tree structure.

## VFS

Linux uses a Virtual File System (**VFS**) to communcate with the filesystem software. When an application needs access to a file, it interacts the the VFS abstraction layer which translates the I/O system calls to specific code relevant to the particular filesystem. This way, the application developer doesn't need to worry about interacting differently with each particular filesystem.

Popular supported Linux filesystems:

- `ext2/3/4` - Linux default. Newest filesystem and default for majority of distros. Features:

  - Journaling
  - Extent (contiguous blocks which increase performance of large files)
  - Backwards compatible with earlier `ext2/3`.
  - Max file size is 16TB, max filesystem size is 1EB.
  - Unlimited amount of subdirectories.
  - The `ext4` layout:

    - Disk blocks are partitioned into block groups which contain adjacent inode and data blocks resulting in better performance.
    - Fields are written to disk in little-endian order (The least significant byte (the "little end") of the data is placed at the byte with the lowest address).

  - the `ext4` superblock contains global information about the filesystem such as mount count, block size, blocks per group, free block count, free inode count and OS system id. The superblock is stored on several block groups.

  - Block sizes are 512/1024/2048/4096 bytes.

  - To see the list of superblocks of a specific :

    ```bash
    sudo dumpe2fs /dev/sda1 | grep superblock
    ```

    To modify a speific filesystem:

    ```bash
    # change number of max mounts
    sudo tune2fs -c 25 /dev/sda1

    # change interval maximal between filesystem chceks
    sudo tune2fs -i 10 /dev/sda1

    # list contents of superblocl
    sudo tunefs -l /dev/sda1
    ```

- `xfs` - RHEL. Engineered to deal with large datasets and effectively handle parallel I/O tasks.

  - Up to 16 EB for total filesystem.
  - Up to 8 EB for an individual file.
  - Employs Direct Memory Access (DMA) I/O.
  - Guranteeing an I/O rate.
  - Can journal quota information.
  - Supports `xattr`.
  - Maintenance can be done while the filesystem is mounted.
  - Can use `xfsdump` and `xfsrestore` to dump/restore respectively.

- `jfs`
- `FAT32/64` - Windows
- `NTFS` - Windows
- `nfs`

A list of supported filesystem exists:

```bash
cat /proc/filesystem
```

Filesystems can be stored on:

- a physical partition on a disk.
- a logical partition controlled by the Logical Volume Manager (**LVM**).
- a network where it is hidden from the local system. The most popular is Network Shares (`nfs`).

Data distinction:

- Shareable vs non-shareable

    user home directories are shareable, device lock files are not.

- Variable vs static

    static data include binaries, docs. variable data includes process information files, etc.

---

The standard for filesystem tree is referred to as FHS (File Hierarchy Standard).

![dir-struct](https://www.tecmint.com/wp-content/uploads/2013/09/Linux-Directory-Structure.jpeg)

## `/boot`

Holds essential files for booting the system.
Some important ones are:

`vmlinuz` - Compressed Linux kernel.

`initramfs/initrd` - The initial RAM filesystem/disk which is mounted before the root filesystem becomes available.

`config` - Used to configure the kernel compilation.

`System.map` - Kernel symbol table. Used for debugging.

## `/dev`

Pseudo-filesystem.
Contains I/O devices files/nodes built into the system which are used to read/write to/from a peripheral hardware device. This directory does not include network interfaces.

## `/etc`

Contains configuration files and startup scripts. Different distributions often add their own configuration files and directories under it.

Important files:

`/etc/systemd` - Contains configuration scripts for interacting with system services.

`/etc/initd` - Constains startup and shutdown scripts

## `/lib` and `/lib64`

These directories should contain only those libraries needed to execute the binaries in /bin and /sbin. These libraries are particularly important for booting the system and executing commands within the root filesystem.

Kernel modules (often device or filesystem drivers) are located under `/lib/modules/<kernel-version-number>`.

PAM (Pluggable Authentication Modules) files are stored in `/lib/security`.

## `/media`

This directory was typically used to mount filesystems on removable media. These include CDs, DVDs, and USB drives, and even Paleolithic era floppies.

## `/mnt`

This directory is provided so that the system administrator can temporarily mount a filesystem when needed.

## `/opt`

This directory is designed for software packages that wish to keep all or most of their files in one isoated place.

For example, if `dolphy_app` were the name of a package which resided under `/opt`, then all of its files should reside in directories under `/opt/dolphy_app`, including `/opt/dolphy_app/bin` for binaries and `/opt/dolphy_app/man` for any `man` pages.

## `/proc`

A [synthetic/pseudo-filesystem](https://en.wikipedia.org/wiki/Synthetic_file_system) which contains directories with virtual files full of processes and system information. All information written to this directory is in memory.

## `/sys`

Also a pseudo-filesystem that contains information about devices, drivers, kernel modules, system configuration. It is used both to gather and modify the behavior of the system.

## `/root`

Home directory of `root` user.

## `/tmp`

Directory for temporary files that are usually cleared automatically (depending on the distro). This directory should not be used to store very large files as it is stored in RAM.

## Journaling

Filesystems that are journaled have the ability to recover from system crashes and ungraceful shutdowns. This is because all operations on the filesystem are grouped into transactions which must each atomically succeed. If the transaction fails, no changes are made on the filesystem. A log file is maintained of all transactions.

### inodes

Inodes are data structures on a disk that contain metadata about files. Each file is represented by one inode. All I/O activity moves both the file and the inode of the file.

Inodes contain:

- Permissions
- User/group ownership
- Size
- Last update time ( in `ns`)
- Last modification time ( in `ns`)

The file names are stored on the directory file and associates the file name and the inode. The association can be:

- **Hard link** - pointer to the inode. Hard link files need to be on the same filesystem.
- **Soft/Symbolic link** -  pointer to a file name which has an associated inode.

We can use `ln` to create links:

```bash
ln -s file file-soft
ln file file-hard
```

## Usage

```bash
# Prints all filesystems and their capacity/type
df -hT

# disk usage 
du -h /var/log/some_app
```

## Creating Filesystems

```bash
sudo fdisk /dev/sda1

-> n
-> p/e
```
