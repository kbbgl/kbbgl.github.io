---
slug: loading-coredump
title: Loading Coredump into GDB
tags: [gdb, debug, reverse_engineer, cheatsheet]
last_update:
  date: 12/31/2022
  author: kbbgl
---

In Ubuntu:

```bash
> CRASHED_APP=fmt
> mv /var/crash/$CRASHED_APP.crash ./
> apport-unpack $CRASHED_APP.crash $CRASHED_APP.crash.gdb
> gdb $CRASHED_APP $CRASHED_APP.crash.gdb/CoreDump
```