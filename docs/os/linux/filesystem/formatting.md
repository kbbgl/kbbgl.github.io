---
slug: os-linux-filesystem-formatting
title: "Making/Formatting"
authors: [kbbgl]
tags: [os, linux, filesystem, formatting]
---

# Making/Formatting

Every filesystem has a utility for formatting/making a filesystem on a partition. Usually it's prefix is `mkfs` and the binary resides in `/sbin/mkfs`.
The format is:

```bash
mkfs [-t fstype] [options] [device-file]

# Example
sudo mkfs -t ext4 /dev/sda10
```
