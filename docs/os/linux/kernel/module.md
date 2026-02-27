---
slug: os-linux-kernel-module
title: "Modules"
authors: [kbbgl]
tags: [os, linux, kernel, module]
---

# Modules

Modules enable the operating system to extend capabilities of network, disk, peripheral devices and others. They are pieces of software that can be loaded/unloaded into the kernel upon demand and without the need to restart the OS.

## Listing Modules

This command will list all loaded modules:

```bash
lsmod


# Used by column: number indicates how many processes are using it 
# and name of module (1 process, module named bridge)
Module                Size       Used by
stp                   16384      1 bridge

```

We can get more information about a module (loaded or unloaded) using:

```bash
modinfo e1000e

...
# means that the module e1000e depends on module mii
depends:        mii
# what HW it works on
alias:          pci....
# parameters that can be configured
parm:           debug:3c59x debug level (0-6) (int)
...
```

Module information can also be gathered from:

```bash
cat /sys/module/$module_name

# and its parameters
cat /sys/module/$module_name/parameters/$parameter_name
```

## Unloading Modules

We can unload modules by running:

```bash
sudo /sbin/rmmod $module_name
```

Better alternative is to use `modprobe` to remove a module:

```bash
modprobe -r e1000e
```

It is impossible to unload a module while it's being used by other modules/processes. Need to use `lsmod` to check which other modules are using the module attempted to be unloaded.

## Loading Modules

We can load a module using:

```bash
sudo /sbin/insmod /path/to/module_name.ko

# Example
sudo insmod /lib/modules/$(uname -r)/kernel/drivers/net/ethernet/intel/e1000e.ko key1=value1 ...
```

We can also use to load a module using:

```bash
modprobe e1000e
```

To rebuild a module dependency database (needed for `modprobe` and `modinfo`):

```bash
depmod
```

The modules are kernel version-specific and will not be able to load otherwise. They must be compiled when the kernel itself is compiled.
When loading/unloading with `modprobe`, the system will automatically load/unload all dependent modules first.

The location for kernel modules is under:

```text
/lib/modules/$kernel_version/
```

## `modprobe.d`

When loading/unloading modules using `modprobe`, files `/etc/modprobe.d/*.conf` are scanned. These configuration files include automatically supplied options and blacklist specific modules to avoid them from being loaded.

The configuration files are one command per line
