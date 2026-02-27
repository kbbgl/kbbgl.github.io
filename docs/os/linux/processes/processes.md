---
slug: os-linux-processes-processes
title: "Processes"
authors: [kbbgl]
tags: [os, linux, processes]
---

# Processes

In Linux, threads are treated as standalone processes.

Every process has its own:

* `pid`
* `ppid` - parent process id
* `pgid` - process group id
* program code, data, variables - **static libraries** are loaded at compile time. **shared libraries** (or DLLs) are loaded during run time.
* file descriptors
* environment

Process attributes:

* Program being executed
* Context/state
* Permissions - (`setuid` programs are run with the permissions of the owner)
* Associated resources

Process states:

* Running
* Sleeping
* Stopped
* Zombie

`init` is the first process that runs on a system and is the parent of all other processes except for ones originating from the kernel.

If a parent process dies before a child, the `ppid` of the child is set to 1 and the process is adopted by `init`. On newer systems, the `pid` is actually set to 2 which corresponds to an internal kernel thread (`kthresadd`).

Processes that have terminated but their parent process has not requested their exit code are called a **zombie processes**. The `init` process ensures processes die gracefully.

Processes are started by `fork and exec` where the parent process creates a child process which inherits its `pid` and then the parent process terminates.

All processes are **scheduled** by the system. This is important because it enables the kernel to perform context switching when a program is waiting for a response, the resources can be transferred to another program in need of the resources until the other first program receives the data.

Processes interact with the kernel using **system calls** which act as interfaces to request/release access to the hardware. They run in their own user space for added security where no other process can access it (**Process Resource Isolation**).

## Controlling Resources

We can use `ulimit` to control resource limits for processes. There are 2 types of limits:

* Hard, the maximum value, can be set by `root`:

    ```bash
    ulimit -H -n
    1048576
    ```

* Soft, the current limiting value, can be set by any user:

   ```bash
   ulimit -S -n
   1024
   ```

We can set the limit for the current session by using:

```bash
ulimit -n 1600
```

To modify system-wide, we can modify the file `/etc/security/limits.conf` and reboot.

We can also use `nice` to set the `niceness` level of a process:

```bash
nice -n 5 cat
nice -n -5 cat
renice +5 -p 20003
```

## Execution Modes

A process can be run in two modes:

* System/Kernel mode - Ring 0
* User mode - Ring 3
![rings](https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Priv_rings.svg/1920px-Priv_rings.svg.png)
A mode is enforced on the hardware level as it is the state of the processor.

When a process needs access to the hardware, it issues a system call which changes the context switch from user mode to kernel mode.

## Daemons

Processes that run in the background, usually started at boot time and provide security.

Scripts in `/etc/init.d` run daemons on startup.

## Shared Libraries

Libraries that can be linked to any program at run-time. They provide a means to use code that can be loaded anywhere in the memory. Once loaded, the shared library code can be used by any number of programs.

Use `ldd` to list the shared libraries of an executable. It displays a list of `sonames` the name of the object (e.g. `linux-vdso.so.1`)

```bash
ldd /usr/bin/apt
 linux-vdso.so.1 (0x00007fff9339b000)
    ...
```

The `ldconfig` binary uses `/etc/ld.so.conf` to configure where to look for shared libraries. We can also use the `LD_LIBRARY_PATH` to set up the shared libraries.
