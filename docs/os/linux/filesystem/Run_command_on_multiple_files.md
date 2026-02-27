---
slug: os-linux-filesystem-run-command-on-multiple-files
title: "Run Command on Multiple Files"
authors: [kbbgl]
tags: [os, linux, filesystem, run_command_on_multiple_files]
---

# Run Command on Multiple Files

```bash
for file in ~/Downloads/146622/ip/*;do cat $file | wc -l;done

   16138
   16392
   16686
   16272
   16276
   16028
```
