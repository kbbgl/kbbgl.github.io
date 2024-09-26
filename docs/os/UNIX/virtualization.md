# Virtualizations

Virtualization is the process of running a virtual instance of a computer system in a layer abstracted from the actual hardware. Most commonly, it refers to running multiple operating systems on a computer system simultaneously.

To check if the guest machine supports hardware virtualization:

```bash

# For Intel IVT processors
grep "vmx" /proc/cpuinfo

# For AMD
grep "svm" /proc/cpuinfo

```

A **hypervisor**, also known as a virtual machine monitor (**VMM**), is software that creates and runs virtual machines (VMs). A hypervisor allows one host computer to support multiple guest VMs by virtually sharing its resources, such as memory and processing. The hypervisor can be run externally (e.g. VMWare) or internally (KVM).

## `KVM` - Kernel Virtual Machine

The KVM project added hypervisor capabilities into the Linux kernel.

`libvirt` is a library that interacts with the VMs, virtual networks, storage. All tools can be found in:

```bash
ls /usr/bin/virt*
```

The most popular tools is `virt-manager`.
