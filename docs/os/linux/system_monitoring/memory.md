---
slug: os-linux-system-monitoring-memory
title: "Memory Monitoring and Usage"
authors: [kbbgl]
tags: [os, linux, system_monitoring, memory]
---

# Memory Monitoring and Usage

```bash
pmap
cat /proc/meminfo
cat /proc/sys/vm/*
free
vmstat -S m
vmstat -d
vmstat -p /dev/sdb1
```

## OOM Killer

Ways to deal with memory pressure:

- Permit memory allocations as long as there's free memory.
- Use swap space to push some resident memory out of core. We can turn on/off the swap by `/sbin/swapo[n/ff] -a`
- In Linux, the system overcommits memory (using COW) for user space processes (kernel space are allocated during request time).

We can change the overcommitting [policy](https://www.kernel.org/doc/Documentation/vm/overcommit-accounting) by:

```bash
# $X can be:
# 0: default, permit for certain cases.
# 1: all.
# 2: disable, memory requests will fail when memory commit reaches the size of swap space + configurable (`/proc/sys/vm/overcommit_ratio`) RAM percentage.
echo "vm.overcommit_memory=X"
```

When the system reaches memory exhaustion, it invokes the OOM-Killer which decides to kill process by the `/proc/[pid]/oom_score`. We can modify this value by adjusting `/prov/[pid]/oom_adj` and `/prov/[pid]/oom_adj_score`.
