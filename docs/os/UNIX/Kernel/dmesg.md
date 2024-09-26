# `dmesg`

Provides vital information that can be used for troubleshooting production performance problems.

`dmesg` command prints all the kernel-related log messages in the console. It includes messages related to the device driver, OS patching, memory, disk drives, network, etc.

## Verbosity

```bash
dmesg --level=emerg,alert,crit,err,warn,notice,info,debug
```

To output the verbosity level and the facility (`kernel`, `user`, etc.):

```bash
dmesg -x
```

## Human-Readable Timestamp

```bash
dmesg -T
```
