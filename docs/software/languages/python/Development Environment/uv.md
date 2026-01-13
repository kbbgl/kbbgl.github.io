---
slug: execute-python-using-uv
title: Execute Python using uv
authors: [kgal-akl]
tags: [uv, python, package_managers, executable, bash, shell]
---


```python
#!/usr/bin/env -S uv run --script  
# /// script  
# requires-python = ">=3.11"  
# dependencies = [ "modules", "here" ]  
# ///  
```
  
  
The script now works like a standalone executable, and `uv` will magically install and use the specified modules.