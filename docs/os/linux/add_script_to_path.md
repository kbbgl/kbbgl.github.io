---
slug: os-linux-add-script-to-path
title: "Add Script to Path"
authors: [kbbgl]
tags: [os, linux, add_script_to_path]
---

# Add Script to Path

```bash
nano ~/.zshrc
```

Then add script path:

```bash
# ~/.zshrc
export PATH="~/scripts:$PATH"
```

```bash
ls ~/scripts/
# bastion.sh
```

```bash
source ~/.zshrc
```

```bash
bastion.sh
```
