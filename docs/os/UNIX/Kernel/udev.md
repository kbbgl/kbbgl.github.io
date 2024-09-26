# `udev`

`udev` is a device manager for the Linux kernel. It manages device nodes in `/dev`. It also handles all user space events raised when HW devices are added or removed into the system.

## Device Nodes

The `/dev` directory includes different device nodes such as network and block devices.

![udev](https://us.v-cdn.net/6030874/uploads/editor/vi/n36syvedw95x.png)

Device nodes can be created:

```bash
sudo mknod [-m mode] /dev/name

# e.g.
sudo mknod -m 666 /dev/mycdrv c 254 1
```

Device drivers can manage multiple devices nodes. To identify the driver associated with a device, there are **major** and **minor** numbers. Device nodes of the same type (e.g. block) with the same major number use the same driver.

We can see the major/minor numbers:

```bash
ll /dev/sda*

# major is 8, minor is 0-3
brw-rw---- 1 root disk 8, 0 Dec 31 11:10 /dev/sda
brw-rw---- 1 root disk 8, 1 Dec 31 11:10 /dev/sda1
brw-rw---- 1 root disk 8, 2 Dec 31 11:10 /dev/sda2
brw-rw---- 1 root disk 8, 3 Dec 31 11:10 /dev/sda3
```

## `udev` Components

`udev` is built upon 3 components:

- `libudev` library which allows access to devices' informations.
- `udevd` or `systemd-udevd` which manages the devices in `/dev` directory.
- `udevadm`, a tool for diagnostics and control.

## Configuration

The main configuration file can be found in `/etc/udev/udev.conf` which contains information about where to place device nodes, permissions, etc. The rules for device naming are kept in `/etc/udev/rules.d` and `/usr/lib/udev/rules.d`.

## How `udev` Loads Devices

`udev` runs a daemon (either `udevd` or `systemd-udevd`) and monitors a netlink socket (communication between kernel and user space).

The hotplug kernel subsystem dynamically handles the addition and removal of devices, by loading the appropriate drivers and by creating the corresponding device files (with the help of `udevd`). When `udev` is notified by the kernel of the appearance of a new device, it collects various information on the given device by consulting the corresponding entries in `/sys/`. Armed with all of this information, `udev` then consults all of the rules contained in `/etc/udev/rules.d/` and `/lib/udev/rules.d/`. In this process it decides how to name the device, what symbolic links to create (to give it alternative names), and what commands to execute. All of these files are consulted, and the rules are all evaluated sequentially.

## Rule Files

All rule files are located in `/etc/udev/rules.d/$rulename.rules`.
