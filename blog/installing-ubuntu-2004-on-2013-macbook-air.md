---
slug: installing-ubuntu-2004-on-2013-macbook-air
title: Installing Ubuntu 20.04 on 2013 MacBook Air
description: Fill me up!
authors: [kbbgl]
tags: [mac,ubuntu]
date: 2021-02-17
---

## Introduction

The other day when visiting my family, under a large pile of torn up binders and laminated documents, I found my sibling’s old 2013 MacBook Air. I thought it would be wasteful to just leave it there so I picked it up and took it to the lab, AKA home.
I discovered that the laptop was password locked with my sibling’s user and password. Since it must’ve been laying at my family’s house for a few years at least, and we as human’s have a tendency to forget our credentials, I decided to save the attempts and just format it and start with a clean OS.

<!-- truncate -->

## Formatting and Reinstalling MacOS

1. In the [menu bar](https://support.apple.com/en-il/guide/mac-help/aside/mchl385f1eea/10.13/mac/10.13), choose Apple menu > Restart. As your Mac restarts, hold down the `Command + R` keys until the macOS Utilities window appears.
1. Select _Disk Utility_, then click Continue.
1. Select your startup disk on the left, then click _Erase_.
1. Click the Format pop-up menu, choose [Mac OS Extended format](https://support.apple.com/en-il/guide/mac-help/aside/dsku61bc26a5/10.13/mac/10.13), enter a name, then click _Erase_.
1. After the disk is erased, choose _Disk Utility > Quit Disk Utility_.
1. Select [Reinstall macOS](https://support.apple.com/en-il/guide/mac-help/reinstall-macos-mchlp1599/10.13/mac/10.13), click _Continue_, then follow the onscreen instructions.

All of the steps ran smoothly and I had a working machine! The battery was still in really good shape as well, the screen was scratchless and the keyboard still had that bounce to it.
So I began downloading all the applications I usually use like Google Chrome, iTerm, Visual Studio Code. But when launching Google Chrome, I discovered that when navigating to GMail or Google Drive, I received an error message indicating that the applications do not support this browser and operating system.
I wasn’t buying it that the hardware would not be able to run the browser and applications and I had a strong belief that it’s some Apple-imposed limitation as part of the ‘[Lightbulb Conspiracy](https://en.wikipedia.org/wiki/Planned_obsolescence).’ So I decided to pursue a different option: installing the latest Ubuntu on this 2013 MacBook Air to free the software restrictions of the machine manufacturer and unleash the power of the hardware.
I decided the best approach would be to initially dual boot MacOS and Ubuntu, ensure that Ubuntu is stable and then remove the partition where MacOS lives.

## Create a Bootable Ubuntu USB Stick

The first step was to generate a bootable USB stick which would hold the Ubuntu image.
I logged into MacOS and downloaded the latest Ubuntu image from [Ubuntu Desktop](https://ubuntu.com/download/desktop) and saved it in `~/Downloads/ubuntu-20.04.2.0-desktop-amd.iso`.

I then opened the Terminal and ran the following command:

```bash
hdiutil convert -format UDRW -o /tmp/ubuntu.img ~/Downloads/ubuntu-20.04.2.0-desktop-amd.iso
```

The command utilizes the MacOS `hdutil` tool to convert an image between two data types: `.iso` (disk image in ISO-9660 standard) to `.img.dmg` (disk image specifically by MacOS). The `-format UDRW` argument specifies that we want to read and write the image with the Apple `UDIF` format. To read more about the hdutil and UDIF format, I suggest reading the [hdutil man page](https://ss64.com/osx/hdiutil.html) and [disk image formats](http://disktype.sourceforge.net/doc/ch03s13.html).

The command above created a new file in `/tmp/ubuntu.img.dmg`. But since we don’t want to run the MacOS installer (which is what the dmg image and extension do by default) in this case but want to create a bootable USB stick, we’ll need to convert the image from `.img.dmg` to `.img`:

```bash
mv /tmp/ubuntu.img.dmg /tmp/ubuntu.img
```

At this point we can insert the USB stick which will likely be mounted by default to `/dev/disk1`.
Next, we need to unmounting the disk representing the USB stick, copying all files from the image to the unmounted disk using the dd binary and then ejecting the disk. We can perform all these steps using the MacOS `diskutil` binary within the Terminal:

```bash
diskutil unmountDisk /dev/disk1
 
sudo dd if=/tmp/ubuntu.img of=/dev/rdisk1 bs=1m
 
diskutil eject /dev/disk1
```

We’re done with this step! We now have a USB stick containing the Ubuntu OS. We can eject it for now.
Next, we need to prepare some disk space on the hard disk where we intend to install Ubuntu.

## Creating a Partition

A partition is basically a separation of the hard disk into individual sections (also known as containers).

![cont](https://www.maketecheasier.com/assets/uploads/2012/05/partitions-partition-diagram.png)

To be able to run both Ubuntu and MacOS on the same machine, we require to create a partition in the hard disk to hold both operating systems.
This step is rather straightforward with the use of the MacOS Disk Utility. We need to launch it, select the first disk we see on the left navver (called _Macintosh HD_ in the image below):

Then click on _Partition_, choose the size of the disk you want to allocate to the partition where we’ll store Ubuntu, make sure the selected format is _Mac OS Extended (Case-sensitive, Journaled)_ and click Apply. My hard disk size was 120GB so I allocated 80GB to Ubuntu and left the rest for MacOS.

## Replacing MacOS Default Boot Manager

Let’s recap what we have done so far:

What would be the next logical step? We would need to tell the machine to let us decide which operating system we want to boot into. By default, the Apple bootloader will load MacOS. To change this behavior, we need to install a different boot manager. We need [rEFInd Boot Manager](https://www.rodsbooks.com/refind/) to do this.

[Download the binary](http://sourceforge.net/projects/refind/files/0.13.2/refind-bin-0.13.2.zip/download) from here and extract it:

```bash
unzip ~/Downloads/refind-bin-0.13.2.zip
```

Then reboot the machine and hold `Command + R`. This should bring you into Recovery Mode.

We need to run the `rEFInd` installation script. From the main menu bar, choose Utilities > Terminal and type the following command:

```bash
cd /Volumes/Macintosh\ HD/Users/YOUR_USERNAME/Downloads/refined
 
./refind-install
```

This will install the rEFInd boot manager. After it completes, shut down the machine, insert the USB stick we’ve prepared earlier and turn on the machine.

## Launching and Installing Ubuntu

When we turn on the machine the next time, we will be greeted with the rEFInd Boot Manager:

![refind](https://www.rodsbooks.com/refind/refind.png)

Select Tux (the penguin, Linux mascot) and use the arrows to navigate to the _Try Ubuntu without install_ option. Then press _e_ which will expose a configuration file where we will make some changes. We will need to add the following commands between ‘splash‘ and ‘---‘:

```bash
nomodeset radeon.audio=1
```

This is because by default, the Radeon HDMI audio driver is disabled in the Linux kernel. After making this change, press F10 to save and exit.
After a minute or two, you should be the Ubuntu menu:

![ubuntu](https://ubuntucommunity.s3-us-east-2.amazonaws.com/original/2X/a/ad5e454a9fd45fd56d90da951702c2f2224cd32a.png)

Click on _Install Ubuntu_, follow the installation wizard you should be set!

## Addendum: Fixing Camera Detection

One of the applications I use most nowadays is Zoom so I was surprised when I joined a call and saw that my camera was not detected by Zoom. I also downloaded [Cheese](https://wiki.gnome.org/Apps/Cheese) just to confirm sure that the issue wasn’t specific to Zoom. I saw the same greeting: ‘no device detected’.
I did some research and found that there was some firmware missing for the Facetime HD (Broadcom 1570) PCIe webcam which prevented the kernel from detecting the driver. I opened up the terminal and ran the following commands to get the driver to work (you will need `git` and `curl` installed):

```bash
> cd /usr/local/src
> sudo git clone https://github.com/patjak/bcwc_pcie.git
 
Cloning into 'bcwc_pcie'...
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 1057 (delta 2), reused 4 (delta 0), pack-reused 1049
Receiving objects: 100% (1057/1057), 352.48 KiB | 537.00 KiB/s, done.
Resolving deltas: 100% (709/709), done.
 
 
> cd /usr/local/src/bcwc_pcie
> sudo git clone https://github.com/patjak/facetimehd-firmware
Cloning into 'facetimehd-firmware'...
remote: Enumerating objects: 1, done.
remote: Counting objects: 100% (1/1), done.
remote: Total 886 (delta 0), reused 0 (delta 0), pack-reused 885
Receiving objects: 100% (886/886), 290.76 KiB | 294.00 KiB/s, done.
Resolving deltas: 100% (585/585), done.
 
> cd /usr/local/src/bcwc_pcie/facetimehd-firmware
> sudo make
 
Checking dependencies for driver download...
/usr/bin/curl
/usr/bin/xzcat
/bin/cpio
 
Downloading the driver, please wait...
 
 
Found matching hash from OS X, El Capitan 10.11.5
==> Extracting firmware...
 --> Decompressing the firmware using gzip...
 --> Deleting temporary files...
 --> Extracted firmware version 1.43.0
 
> sudo make install
Copying firmware into '//lib/firmware/facetimehd'
 
> cd /usr/local/src/bcwc_pcie
> sudo make
make -C /lib/modules/5.0.0-23-generic/build M=/usr/local/src/bcwc_pcie modules
make[1]: Entering directory '/usr/src/linux-headers-5.0.0-23-generic'
  CC [M]  /usr/local/src/bcwc_pcie/fthd_ddr.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_hw.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_drv.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_ringbuf.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_isp.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_v4l2.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_buffer.o
  CC [M]  /usr/local/src/bcwc_pcie/fthd_debugfs.o
  LD [M]  /usr/local/src/bcwc_pcie/facetimehd.o
  Building modules, stage 2.
  MODPOST 1 modules
  CC      /usr/local/src/bcwc_pcie/facetimehd.mod.o
  LD [M]  /usr/local/src/bcwc_pcie/facetimehd.ko
make[1]: Leaving directory '/usr/src/linux-headers-5.0.0-23-generic'
 
> sudo make install
make -C /lib/modules/5.0.0-23-generic/build M=/usr/local/src/bcwc_pcie modules_install
make[1]: Entering directory '/usr/src/linux-headers-5.0.0-23-generic'
  INSTALL /usr/local/src/bcwc_pcie/facetimehd.ko
At main.c:160:
- SSL error:02001002:system library:fopen:No such file or directory: ../crypto/bio/bss_file.c:72
- SSL error:2006D080:BIO routines:BIO_new_file:no such file: ../crypto/bio/bss_file.c:79
sign-file: certs/signing_key.pem: No such file or directory
  DEPMOD  5.0.0-23-generic
Warning: modules_install: missing 'System.map' file. Skipping depmod.
make[1]: Leaving directory '/usr/src/linux-headers-5.0.0-23-generic'
```

Don’t worry too much about the SSL errors above, they are red herrings.
I then needed to load the drivers into the kernel:

```bash
> sudo depmod
> sudo modprobe -r bdc_pci
> sudo modprobe facetimehd
```

I relaunched Cheese and saw my messy face on the camera 🙂
But once I restarted the computer, I saw that the changes were reverted. I needed to persist them somehow. After some more research, I found that I needed to add `facetimehd` to the kernel modules that are loaded during boot time:

```bash
> sudo echo facetimehd >> /etc/modules
```

After restarting, I saw that I was able to detect my camera in Zoom and Cheese!
