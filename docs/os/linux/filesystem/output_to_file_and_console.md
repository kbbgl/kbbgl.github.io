---
slug: os-linux-filesystem-output-to-file-and-console
title: "Output to File and STDOUT"
authors: [kbbgl]
tags: [os, linux, filesystem, output_to_file_and_console]
---

# Output to File and STDOUT

```bash
mongo --eval "var cleanup=true" /path/to/sanity/tools/scripts/scripts_loader.js 2>&1 | tee sanity.log
```
