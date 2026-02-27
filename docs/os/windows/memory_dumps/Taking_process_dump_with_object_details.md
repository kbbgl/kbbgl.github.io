---
slug: os-windows-memory-dumps-taking-process-dump-with-object-details
title: "Create Process Dump"
authors: [kbbgl]
tags: [os, windows, memory_dumps, taking_process_dump_with_object_details]
---

# Create Process Dump

## load symbols

```batch
.loadby sos clr
```

## List all objects, how many of each type, and how much memory each is using

```batch
!dumpheap –stat
```
