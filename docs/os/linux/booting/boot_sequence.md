---
slug: os-linux-booting-boot-sequence
title: "Boot Sequence"
authors: [kbbgl]
tags: [os, linux, booting, boot_sequence]
---

# Boot Sequence

Steps in the boot sequence:

1) **BIOS** (Basic Input Output System)/**UEFI** runs **POST** (Power On Self Test) to locate and execute the boot program/loader (typically the boot service can be found in the device Master Boot Record/MBR). Control of the computer is then transferred to the boot program/loader.

2) Boot prorgam/loader loads the kernel into memory and executes it. It then check HW.

3) Kernel starts init process (`pid=1`).

4) `init` manages system initialization using `systemd`.

## BIOS

The BIOS contains all the code required to gain initial access to the keyboard, monitor, serial ports, disks.

It is placed in the ROM chip that comes with the computer. This ensures that the BIOS will never be damaged by disk failures.

## Bootloaders

There are different types of bootloaders:

- **GRUB**: Grand Unified Bootloader, most popular bootloader for non-embedded Linux distros.
- `efibootmgr`: boot manager used with GRUB.
- **LILO**: Old Linux Loader, obsolete.
- Das U-Boot: Most popular bootloader for embedded systems.

## System Configuration

System configuration files are stored in different places in different distros:

```bash
# RHEL
/etc/sysconf

# Debian
/etc/default
```

The files contained within these folders provide extra options when starting a service and contain code to set env vars (.e.g `/etc/default/useradd` sets default user properties when a user is added).

## Shutdown and Reboot

The `shutdown` command brings down the system gracefully, optionally notifying all logged in users it's being taken down.

```bash
# Halt machine in 1 minute, send message to all logged-in users
sudo shutdown -h +1 "Power failure imminent"

# Halt machine now
sudo shutdown -h now

# Reboot machine now
sudo shutdown -r now

# poweroff now
sudo shutdown now
```
