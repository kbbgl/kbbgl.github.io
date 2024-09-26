---
title: Getting Started with the Kernel
slug: getting-started-kernel
authors: kbbgl
tags: [linux,os_dev]
---

## Source

```bash
gh repo clone torvalds/linux
```

## Kernel Source Tree

| Directory       | Description                                   |
|-----------------|-----------------------------------------------|
| `arch`          | Architecture-specific source                  |
| `block`         | Block I/O Layer                               |
| `crypto`        | Crypto API                                    |
| `Documentation` | Kernel source documentation                   |
| `drivers`       | Device drivers                                |
| `firmware`      | Device firmware needed to use certain drivers |
| `fs`            | the VFS and the individual filesystems        |
| `include`       | Kernel headers                                |
| `init`          | Kernel boot and initialization                |
| `ipc`           | Interprocess communication code               |
| `kernel`        | Core subsystems, such as the scheduler        |
| `lib`           | Helper routines                               |
| `mm`            | Memory management subsystem and the VM        |
| `net`           | Networking subsystem                          |
| `samples`       | Sample, demonstrative code                    |
| `scripts`       | Scripts used to build the kernel              |
| `security`      | Linux Security Module                         |
| `sound`         | Sound subsystem                               |
| `usr`           | Early user-space code (called `initramfs`     |
| `tools`         | Tools helpful for developing Linux            |
| `virt`          | Virtualization infrastructure                 |

## Building the Kernel

### Configuring the Kernel

Kernel configuration is controlled by `CONFIG_$FEATURE`. For example, symmetrical multiprocessing (SMP) is controlled by `CONFIG_SMP`.

Configuration options values can be `yes,no,module`. Kernel features are usually `yes` or `no` whereas drivers are can be either `yes,no,module`.
Configuration options can also be strings or integers.

We can use:

```bash
# Asks for options 1-by-1
make config

# Better tools
make menuconfig
make gconfig

# Creates configuration based on your architecture
make defconfig
```

to facilitate configuration.

Once the configuration is done, the selected options will be stored in `.config`.

After making changes in `.config`, you can validate and update the configuration using:
```bash
make oldconfig
```

Another helpful option is `CONFIG_IKCONFIG_PROC` which places the complete, compressed kernel configuration file at `/proc/config.gz`. You can then run the following command to compy the config:
```bash
zcat /proc/config.gz > .config
make oldconfig
```

To *build* the kernel using:
```bash

# No need for all the noise, just print errors/warnings if encountered
make > /dev/null
```

We can speed up the build process by running `n` concurrent jobs;
```bash
make -jn

# For 16 concurrent jobs
make -j16 > /dev/null
```

[`ccache`](https://ccache.dev/) and [`distcc`](https://www.distcc.org/) can improve kernel build time as well.

## Installing the Kernel

The installation is architecture and bootloader dependent. We need to find the directions for the bootloader on where it's expecting the kernel image and how to set it up to boot.

On x86 system using `grub`, you would:
```bash
cp arch/i386/boot/bzImage /boot
mv /boot/vmlinuz-version
```

And then edit `/boot/grub/grub.conf` and adding a new entry for the new kernel.

## Linux Kernel Attributes

### No `libc` or Standard Headers

This means that the kernel is not linked to any library because of speed and size. The full C library is too large and inefficient for the kernel.

Many usual `libc` functions are implemented inside the kernel such as `lib/string.c`. All we need to do is to include `<linux/string.h>`.


### No Memory Protection

Make sure to not dereference a `NULL` pointer which will cause a major kernel error (`oops`). 
The kernel memory is not pageable so every byte consumed in memory is one less available physical memory.

### No Floating Point

Using a floating point inside the kernel requires manually saving and restoring the floating point registers. Don't do this!

### Stack

The user-space has a large stack that can grow dynamically. The kernel stack is small and fixed-size depending on architecture. On x86, the stack is configurable at compile-time and can be either 4KB or 8KB. Each process receives its own stack.

### Sync and Concurrency

The kernel is susceptible to race conditions since it's a multitasking operating system, supports SMP, interrupts occur asynchronously with respect to the currently executing code.

