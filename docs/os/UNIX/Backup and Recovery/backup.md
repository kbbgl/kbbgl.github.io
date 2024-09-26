# Backup

What data should be backed up?

- Business-related data.
- System configuration files.
- User files (under `/home`)

## Utilities

```bash
# create/extract file archives
tar
cpio
# Create an archive, use -o or --create:
ls | cpio --create -O /dev/st0

# Extract from an archive, use -i or --extract:
cpio -i somefile -I /dev/st0

# List contents of an archive, use -t or --list:
cpio -t -I /dev/st0

#You can specify the input (-I device) or use redirection on the command line.



# compress archives
gzip
bzip2
xz

# used to transfer raw data between media
# copy entire partitions/disks
dd

# Create a 10 MB file filled with zeros:
dd if=/dev/zero of=outfile bs=1M count=10

# Back up an entire hard drive to another (raw copy):
dd if=/dev/sda of=/dev/sdb

# Create an image of a hard disk (which could later be transferred to another hard disk):
dd if=/dev/sda of=sdadisk.img

# Back up a partition:
dd if=/dev/sda1 of=partition1.img

# Back up a CD ROM:
dd if=/dev/cdrom of=tgsservice.iso bs=2048

# Use dd in a pipeline:
dd if=ndata conv=swab count=1024 | uniq > ofile

# synchronize directory subtrees, filesystems accross networks
rsync [options] $source_file $destination_file
rsync file.tar someone@backup.mydomain:/usr/local
rsync -r a-machine:/usr/local b-machine:/usr/
rsync -r --dry-run /usr/local /BACKUP/usr
rsync -r project-X archive-machine:archives/project-X

# obsolete ways to export/import on specific machine
dump
restore
```

More popular tools:

```bash
# http://www.amanda.org/
amanda

# https://www.bacula.org/7.0.x-manuals/en/main/Main_Reference.html
bacula

# https://clonezilla.org/aman
clonezilla
```
