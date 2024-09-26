# Disk Partitioning

Popular hard disk types:

- Small Computer Systems Interface (**SCSI**): Older type, very flexible and still used. Has different capacities, speeds.

- Serial Attached SCSI (**SAS**): Newer and have better performance than SCSI.

- Integrated Drive Electronics (**IDE**): Obsolete.

- Serial Advanced Technology Attachment (**SATA**): Replaced IDEs, seen as SCSI devices. Have built-in hotswap and fast data transfer rates.

- Universal Serial Bus (**USB**): Used by flash drives, seen as SCSI devices.

- Solid State Drives (**SSD**): Use less power than rotational disks, very fast data transfer.

## Disk Geometry

A hard drives is composed of a number of 'platters'. The platters store data. The geometry of the drive defines how the data is organized on these platters.

![cylinder](https://upload.wikimedia.org/wikipedia/commons/thumb/0/02/Cylinder_Head_Sector.svg/720px-Cylinder_Head_Sector.svg.png)

Heads: are the needles that read/write data while the platter spins.

Cylinders: Heads are attached to the cylinders which are separators between each platter. It's a group which consists of the same track on all platters.

Tracks: Circular regions where the head needle reads/writes data onto.

Sectors: Are data blocks (of 512/4096 bytes) which are slices of a the tracks.

We can see the sectors by:

```bash
sudo fdisk -l /dev/sda
```

We can see the partitions and the type by:

```bash
blkid /dev/sda*
/dev/sda: PTUUID="cac1e54c-3ea8-4e5e-8486-cdb1c02055f3" PTTYPE="gpt"
/dev/sda1: UUID="2A46-DB7A" TYPE="vfat" PARTUUID="c197e0f8-56d2-4726-a424-0de20f1dc9ae"
/dev/sda2: UUID="7a140733-3f99-4878-a4e5-e738b2614353" TYPE="ext4" PARTUUID="6955ca22-d490-437d-9840-e571d743a624"
/dev/sda3: UUID="fciQ3l-PQWL-c78G-1wZz-x1aw-Sdp6-3fo7l1" TYPE="LVM2_member" PARTUUID="383fe0d3-f69c-43b4-9576-e5e516b5f773"
```

To see the block device information in tree format:

```bash
lsblk
```

### Partition Organization

Disks are divided into partitions. A partition is a physically contiguous groups of sectors on the disk.

There are two main partitions layouts:

- Master Boot Record (**MBR**): Older scheme, limited to 2TB, 4 primary partitions, 1 can be extensible partition and may be subdivided into 16 other logical partitions.

- GUID Partition Table(**GPT**): Modern, based on UEFI, up to 128 primary partitions and no need for extended partitions.

#### Reasons for Partition

- Separatation of user/app data from OS files.
- Sharing between machines/OSs.
- Security, permissions.
- Keeping volatile storage isolated from stable.
- Keeping most important/used files on faster storage.

Most systems use at least two partitions:

- `/`
- Swap, which is used as an extension of physical memory.

### Disk Devices and Nodes

The Linux kernel interacts with the hardware through the `/dev` directory. The devices are should be interacted with only through the VFS.

### Create a partition

```bash
sudo fdisk /dev/sda
# [-> inside fdisk prompt]
-> n
+1000M

# for LVM partition
-> t
-> 8e

# list new partition
-> p

# refresh partition table. Doesn't always work, recommended to reboot.
sudo partprobe -s
```

We can use this command to format a partition:

```bash
sudo mkfs.ext4 /dev/sda9
```

SATA/SCSI hard disks have the following format:

```bash
# xx is the device type (usually sd)
# y is the letter for the drive number
# z is the partition number
xxy[z]

sda - first disk.
sdb - second disk.
sda1 - first disk, first partition
...
```
