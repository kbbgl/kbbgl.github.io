---
slug: os-linux-awk-cheat-sheet
title: "AWK Cheat Sheet"
authors: [kbbgl]
tags: [os, linux, awk_cheat_sheet]
---

# AWK Cheat Sheet

### Add line numbers at start of each line

```bash
awk -v count=0 '{print ++count " " $0}' input.txt
```

### Print conditional (`$2` is an integer)

```bash
awk  '{if ($2 > 30) print $0}' input.txt
```
