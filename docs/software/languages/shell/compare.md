---
slug: comparing-text-between-2-files
title: Comparing Text Between 2 Files
authors: [kgal-akl]
tags: [tools, bash, shell, helper, compare, diff]
---

```bash
# First, sort both files (comm requires sorted input)
sort /tmp/assigned > /tmp/assigned_sorted
sort /tmp/all > /tmp/all_sorted

# Find groups in 'all' but NOT in 'assigned' (unassigned groups)
comm -23 /tmp/all_sorted /tmp/assigned_sorted > /tmp/unassigned
```