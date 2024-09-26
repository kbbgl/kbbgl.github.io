# Kernel

The kernel is the core component of the OS. It:

- Connects the HW/SW
- Manages system resources (RAM/CPU)
- Handles connected device drivers and makes devices available to OS.
- Is responsible for **system initialization and boot up**.
- Process scheduling.
- Memory management.
- Control access to HW.
- I/O between applications and storage devices.
- Implementation of local/network fs.
- local and remove security control(e.g. fs permissions)
- Networking control.

## Boot

Parameters are passed to the system at boot using the kernel command line. An example:

```bash
# root: root filesystem
# ro: mounts root device read-only on boot
# crashkernel: how much memory to set aside for kernel crashdumps

linux boot/vmlinuz-4.19.0 root=UUID=3829711-91487194-173541957 ro quiet crashkernel=384M-:128M
```

The parameters are `key=value` after `boot/vmlinuz-4.19.0`.

To see what command the system booted with:

```bash
cat /proc/cmdline
```

To view all possible parameters:

```bash
man bootparam
```

## `sysctl`

Interface that can be used to read and tune kernel parameters at runtime.

We can see all keys and values:

```bash
sudo systctl -a

abi.cp15_barrier = 2
...
```

Each key corresponds to an entry in `/proc/sys/`, i.e. `/proc/sys/abi/cp15_barrier`.

We can set a value:

```bash
# Either
sudo sh -c 'echo 1 > /proc/sys/abi/cp15_barrier'

# Or
sudo sysctl abi.cp15_barrier=1
```

Settings can be fixed at boot time by modifying:

```bash
sudo nano /etc/sysctl.conf

# makes change effective immediately
sudo sysctl -p
```
