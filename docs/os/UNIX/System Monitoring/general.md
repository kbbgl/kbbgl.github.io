# Tools

## Network

```bash
iptraf
```

## I/O

```bash
iostat
sar
vmstat
```

## Memory

```bash
free
pmap
```

## Process and Load

```bash
pstree
mpstat # multiprocessor statistics
numastat
strace # process system calls
```

We can create stress on system by running:

[stress](https://wiki.ubuntu.com/Kernel/Reference/stress-ng)

```bash
# 8 cpu-intensive processes
# 4 I/O intensive processes 
# 6 memory-intensive processes 
# allocating 256 MB by default.  
# The size can bechanged as in--vm-bytes 128M.•  
# Run the stress test for 20 seconds.

stress-ng -c 8 -i 4 -m 6 -t 20s
```

## Logs

System logs located in `/var/log`.

`rsyslogd` and `journalctl` are used to interact with system logs.

To view system logs:

```bash

tail -f /var/log/messages
tail -f /var/log/syslog

# OR
dmesg -w
```

Important logs:

- `boot.log` - system boot messages
- `dmesg` - kernel messages after boot
- `messages.log`/`syslog.log` - all important system messages.
- `secure` - security related messages.

Files are rotated using `logrotate` (configuration in `/etc/logrotate.conf`.

## Systems Activity Report (`sar`)

Backend of `sar` is System Activity Data Collector (`sadc`). `sadc` collects and accumulates info  and stores it in `/var/log/sa`.

```bash

sudo sar [options] [interval] [count]

# CPU information, every 3 seconds for 3 times
sudo sar 3 3 

# All information every 2 seconds
sudo sar -A 2
```

All commands can be found [here](https://cmdref.net/os/linux/command/sar.html)

We can use the [`ksar`](https://sourceforge.net/projects/ksar/) to generate graphs from `sar` data.

## Process Monitoring

- `ps`

    ```bash
    ps auxf # will show ancestry
    ps -elf
    ps -eL
    ps -C "bash"
    
    # Choose columns
    ps -o pid,cputime,command
    ```

    Commands which are in `[]` are threads. and exist only within the kernel.

    `VSZ` - process virtual size in KB

    `RSS` - resident set size. Non-swapped memory a task is using in KB.

    `STAT` - State of the process.

- `pstree`

    ```bash
    pstree -aAp 2408
    ```
