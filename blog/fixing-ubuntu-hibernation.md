---
slug: fixing-ubuntu-hibernation
title: Fixing Ubuntu Hibernation
authors: [kbbgl]
tags: [ubuntu,boot,hibernation,raspberry_pi,swap,fs]
date: 2025-05-14
---

## Introduction

One day, seemingly out of the blue, my (using [Docusaurus](https://docusaurus.io/) blog site, which is set up as an Ubuntu `systemd` service on the Raspberry Pi 4, stopped being accessible.

<!-- truncate -->

I tried pinging it and checking whether is it was connected to my local network but it was nowhere to be found. This meant that I needed to physically access the Raspberry Pi so I connected it to a monitor and reset it.

## Boot No Good

Upon reboot, I was greeted by a console with the following worrying "Kernel panic" message:

```
/init: conf/conf.d/zz-resume-auto: line 1: syntax error: unexpected "{" 
[ 1.344964] Kernel panic - not syncing: Attempted to kill mitt exitcode=0x00000200 
[ 1.352769] CPU: 0 PID: 1 Comm: init Not tainted 5.19.0-1011-raspi #18-Ubuntu
[ 1.360029] Hardware name: Raspberry Pi 4 Model B Rev 1.4 (DT)
[ 1.365959] Call trace:
[ 1.368442]  dump_backtrace+Oxbc/Ox124
[ 1.372260]  show_stackoN20/0x5c
[ 1.375625]  dump stack_lul4x88/Oxb4
[ 1.379349]  dump stack4x18/004
[ 1.382714]  panic•Ox1b4/00ac
[ 1.385816]  do exit•ex518/0x520
[ 1.389095]  do_group_exit4x3c/OxbO
[ 1.392726]  __Nake_gp_parent+Ox0/0x40
[ 1.396534]  inuoke_syscall4x50/0x120
[ 1.400343]  el6_suc_common.constprop.04.0x6c/Ox1a0
[ 1.405211]  do ele_suc+004/0x44
[ 1.400577]  elesuc+040/0x180
[ 1.411768]  elOt_64_sync_handler+Oxf4/0x120
[ 1.416106]  elOt_64_sync+Ox1a0/0x1A4
[ 1.419026] SNP: stopping secondary CPUs
[ 1.423813] Kernel Offset: Ox4b0c05400000 from Oxffff800008000000
[ 1.430004] PHYS_OFFSET: Oxffffd932c0000000 
[ 1.434250] CPU features: Ox2000,04027810,00001086
[ 1.439117] Memory Limit: none
[ 1.442221] ---[ end Kernel panic - not syncing: Attempted to kill init! exitcode=0x00000200 ]--- 
```

The first line was the one that stood out. Seems that the file `zz-resume-auto` has a problem with the curly brackets within that context. At this time, I had to figure out a couple of things:

1. What is this file `zz-resume-auto`?
1. Who writes to this file?

### Who are you `zz-resume-auto`?

Since I couldn't access the Raspberry Pi remotely, I had to somehow be able to access it's filesystem without it booting. Since the Raspberry Pi's storage comes in a form of an SD card, I could just pull out the SD card from the slot on the Raspberry Pi, mount it onto an adapter and connect it to another computer.

After inserting the SD card adapter to one of the USB slots on my other machine, I ran the following command to create a directory for the Raspberry Pi filesystem and mounted the SD card onto it:

```bash
# find the disk
> fdisk -l

Disk /dev/sda: 59.48 GiB, 63864569856 bytes, 124735488 sectors
Disk model: Card  Reader
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0x44eba349

Device     Boot  Start       End   Sectors  Size Id Type
/dev/sda1  *      2048    526335    524288  256M  c W95 FAT32 (LBA)
/dev/sda2       526336 124735454 124209119 59.2G 83 Linux


> mkdir /media/rpi/

> mount /dev/sda1/ /media/rpi

> ls /media/rpi
bin   dev  home  lost+found  mnt  proc  run   snap  sys  usr
boot  etc  lib   media       opt  root  sbin  srv   tmp  var
```
