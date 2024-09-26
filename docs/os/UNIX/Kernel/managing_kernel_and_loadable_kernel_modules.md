The Linux® kernel is the main component of a Linux operating system (OS) and is the core interface between a computer’s hardware and its processes. It communicates between the 2, managing resources as efficiently as possible.

The kernel has 4 jobs:

* Memory management: Keep track of how much memory is used to store what, and where

* Process management: Determine which processes can use the central processing unit (CPU), when, and for how long

* Device drivers: Act as mediator/interpreter between the hardware and processes

* System calls and security: Receive requests for service from the processes


Where the kernel fits within the OS
To put the kernel in context, you can think of a Linux machine as having 3 layers:

* The hardware: The physical machine—the bottom or base of the system, made up of memory (RAM) and the processor or central processing unit (CPU), as well as input/output (I/O) devices such as storage, networking, and graphics. The CPU performs computations and reads from, and writes to, memory.

* The Linux kernel: The core of the OS. It’s software residing in memory that tells the CPU what to do.

* User processes: These are the running programs that the kernel manages. User processes are what collectively make up user space. User processes are also known as just processes. The kernel also allows these processes and servers to communicate with each other (known as inter-process communication, or IPC).


Memory get's divided into two distinct areas:

* The user space, which is a set of locations where normal user processes run (i.e everything other than the kernel). The role of the kernel is to manage applications running in this space from messing with each other, and the machine.

* The kernel space, which is the location where the code of the kernel is stored, and executes under.

Processes running under the user space have access only to a limited part of memory, whereas the kernel has access to all of the memory. 

Processes running in user space also don't have access to the kernel space (except for `root`). User space processes can only access a small part of the kernel via an interface exposed by the kernel - the system calls. If a process performs a system call, a software interrupt is sent to the kernel, which then dispatches the appropriate interrupt handler and continues its work after the handler has finished.

![](https://www.differencebetween.com/wp-content/uploads/2017/12/Difference-Between-User-Mode-and-Kernel-Mode-fig-2.png)


Kernel space code has the property to run in "kernel mode", which (in your typical desktop -x86- computer) is what you call code that executes under ring 0. Typically in x86 architecture, there are 4 rings of protection. Ring 0 (kernel mode), Ring 1 (may be used by virtual machine hypervisors or drivers), Ring 2 (may be used by drivers, I am not so sure about that though). Ring 3 is what typical applications run under. It is the least privileged ring, and applications running on it have access to a subset of the processor's instructions. Ring 0 (kernel space) is the most privileged ring, and has access to all of the machine's instructions.

![rings](https://upload.wikimedia.org/wikipedia/commons/2/2f/Priv_rings.svg)


### Kernel Module

Loadable kernel modules (**LKMs**) are modules that can be added to the Linux kernel. These modules are usually used for:

* Device drivers - Used to communicate and manage external hardware with the OS.
* Filesystem drivers
* System calls - User space programs use system calls to get services from the kernel. For example, there are system calls to read a file, to create a new process, and to shut down the system. Most system calls are integral to the system and very standard, so are always built into the base kernel (no LKM option). But you can invent a system call of your own and install it as an LKM. Or you can decide you don't like the way Linux does something and override an existing system call with an LKM of your own.
* Network drivers
* TTY line disciplines
* Executable interpreters


Rootkits are usually written as LKMs which allow hackers to be able to control what the target system is reporting in terms of processes, ports, services, hard drive space.

#### Checking the Kernel Version

```bash
uname -a
```

Or:
```bash
cat /proc/version
```

#### Kernel Tuning with `sysctl`

We can modify kernel behavior using `sysctl`. Any changes made will be effective until reboot. To make the changes permanent, we need to edit the `/etc/sysctl.conf` file.

To see the available options for tuning:
```bash
sudo systcl -a
```

To enable packet forwarding (for MITM):
```bash
sudo sysctl -w net.ipv4.ip_forward=1
```

To make this change permanent by editing `/etc/sysctl.conf` and uncommenting the same line.

To make system unpingable, add to `/etc/sysctl.conf`:
```
net.ipv4.icmp_echo_ignore_all=1
```

Then run to load in `sysctl` settings from `/etc/sysctl.conf`:
```bash
sysctl -p
```


#### Managing Kernel Modules

Linux has 2 ways to manage modules:

* `insmod` suite (insert module) or `rmmod` (remove module). Need to be used with caution because they do not take into account module dependencies.
* `modprobe` command - preferable and modern way to interact with modules.

To list installed modules, their size and which modules are using it:
```bash
lsmod
```

To gather more info on a module:
```bash
modinfo $MODULE_NAME
```
It will also print the `depends` property which lists the module dependencies.


To insert a module:
```bash
modprobe -a $NAME_OF_MODULE
```

To print the kernel message buffer:
```bash
dmesg
```

To remove a module:
```bash
modprobe -r $NAME_OF_MODULE
```

### Write kernel module

A kernel module is a piece of code, usually written in `c` that can interact with Ring 0.

Steps are:
1) Create `lkm_example.c` with sample:

```c
#include <linux/kernel.h>
#include <linux/module.h>


/*
 * Init function of our module
 */
static int __init hellokernelmod_init(void)
{
 printk(KERN_INFO "Hello Kernel!\n");
 return 0;
}

/*
 * Exit function of our module.
 */
static void __exit hellokernelmod_exit(void)
{
 printk(KERN_INFO "Hasta la vista, Kernel!\n");
}

MODULE_AUTHOR("Your name here");
MODULE_DESCRIPTION("Simple Kernel Module to display messages on init and exit.");
MODULE_LICENSE("MIT");


module_init(hellokernelmod_init);
module_exit(hellokernelmod_exit);
```

2) Create `Makefile`

```
obj-m += lkm_example.o

all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules
clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
```

3) Compile:
```bash
make
```
This will create a `lkm_example.ko` file.
We can check the module description by:
```bash
modinfo lkm_example.ko
```

4) Insert module into kernel:
```bash
sudo insmod lkm_example.ko
```

5) Verify the insertion:
```bash
lsmod | grep lkm_example
```

6) See messages:
```bash
dmesg

#[11144.986911] lkm_example: loading out-of-tree module taints kernel.
#[11144.986914] lkm_example: module license 'MIT' taints kernel.
#[11144.986915] Disabling lock debugging due to kernel taint
#[11144.987001] lkm_example: module verification failed: signature and/or required key missing - tainting kernel
#[11144.989369] module initialized
```

7) Remove module:
```bash
sudo rmmod lkm_example
```

```bash
dmesg
#[11404.347066] module exiting
```