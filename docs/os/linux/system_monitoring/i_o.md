---
slug: os-linux-system-monitoring-i-o
title: "I/O Monitoring"
authors: [kbbgl]
tags: [os, linux, system_monitoring, i_o]
---

# I/O Monitoring

1) The main tool used is `iostat`.

    ```bash
    # Show in human readable
    # And extended information
    iostat --human -x

    # Example
    # check for particular devices 
    # refresh every 2 seconds
    # run for 200 times 
    iostat --human /dev/sda /dev/sdb 2 200
    ```

    if `%util` is close to 100%, the system is I/O-bound.

2) Also used is:

    ```bash
    iotop -o
    ```

    Which gives a list of processes that utilize the I/O. `be` stands for best effort, `rt` stands for real time in the `prio` column.

3) `ionice` we can schedule I/O and priority for processes.

    ```bash
    # -c arg:
    # 0 default 
    # 1 real-time 
    # 2 best-effort
    # 3 no access unless no other program askes for it

    # -n arg:
    # 0 highest priority

    # -p
    # pid
    ionice -c [0-4] -n [0-7]

    # Example
    ionice -c 2 -n 3 -p 30078
    ```

## Load Testing Tools

We can use `bonnie++`, `bon_csv2html` and `bon_csv2txt` to benchmarking

```bash
time sudo bonnie++ -n 0 -u 0 -r 100 -f -b -d /mnt
```

Another option is [fsmark](http://sourceforge.net/projects/fsmark/).
